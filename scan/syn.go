package scan

import (
	"context"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"sync"
	"time"
)

type SynScanner struct {
	timeout int
	thread  int
	ctx     context.Context
	mutex   sync.Mutex
}

// NewSynScanner 创建SynScanner
func NewSynScanner(timeout int, thread int) *SynScanner {
	return &SynScanner{
		timeout: timeout,
		thread:  thread,
	}
}

func (s *SynScanner) Start(ctx context.Context, ip []string, port []int) (<-chan Result, <-chan error) {
	jobChan := make(chan PortJob)
	resultChan := make(chan Result)
	errChan := make(chan error)

	results := make(map[string]map[int]PortState)

	var wg sync.WaitGroup

	// 创建并启动协程
	for i := 0; i < s.thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Scan(ctx, jobChan, results, errChan)
		}()
	}

	// 分发端口扫描任务
	go func() {
		defer close(jobChan)
		for _, ip := range ip {
			for _, port := range port {
				select {
				case <-ctx.Done():
					return
				case jobChan <- PortJob{ip: ip, port: port}:
				}
			}
		}
	}()

	// 等待所有的协程完成
	go func() {
		wg.Wait()
		for ip, ports := range results {
			resultChan <- Result{Host: ip, Ports: ports}
		}
		close(resultChan)
	}()

	return resultChan, errChan
}

func (s *SynScanner) Scan(ctx context.Context, jobChan <-chan PortJob, results map[string]map[int]PortState, errChan chan<- error) {
	for job := range jobChan {
		state, err := s.ScanPort(ctx, job.ip, job.port)
		if err != nil {
			errChan <- err
			continue
		}

		s.mutex.Lock()
		if _, ok := results[job.ip]; !ok {
			results[job.ip] = make(map[int]PortState)
		}
		results[job.ip][job.port] = state
		s.mutex.Unlock()
	}
}

func (s *SynScanner) ScanPort(ctx context.Context, ip string, port int) (PortState, error) {
	localIP, err := GetLocalIP()
	if err != nil {
		return PortClosed, err
	}

	// Create a new packet buffer
	buffer := gopacket.NewSerializeBuffer()

	// Create a new packet options
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	// Create a new ethernet layer
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC:       net.HardwareAddr{0xBD, 0xBC, 0xB5, 0xBC, 0xB5, 0xBD},
		EthernetType: layers.EthernetTypeIPv4,
	}

	// Create a new IP layer
	ipLayer := &layers.IPv4{
		SrcIP:    localIP,
		DstIP:    net.ParseIP(ip),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}

	// Create a new TCP layer
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(1234), // You can change this to a random port
		DstPort: layers.TCPPort(port),
		SYN:     true,
		Window:  1024,
	}

	// Update the checksum for the TCP layer
	tcpLayer.SetNetworkLayerForChecksum(ipLayer)

	// Serialize the packet
	err = gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
	)
	if err != nil {
		return PortClosed, err
	}

	// Open a network interface
	handle, err := pcap.OpenLive("wlo1", 1024, false, 30*time.Second) // Change "eth0" to your network interface
	if err != nil {
		return PortClosed, err
	}
	defer handle.Close()

	// Write the packet to the network interface
	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		return PortClosed, err
	}

	// Set a filter for the network interface to only receive SYN-ACK packets
	filter := fmt.Sprintf("tcp and src host %s and dst host %s and tcp[13] == 18", ip, localIP.String())
	err = handle.SetBPFFilter(filter)
	if err != nil {
		return PortClosed, err
	}

	// Read packets from the network interface
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()

	// Wait for a SYN-ACK packet
	for {
		select {
		case <-ctx.Done():
			return PortClosed, nil
		case packet := <-packets:
			// Check if the packet is a TCP packet
			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				// The packet is a TCP packet, return that the port is open
				return PortOpen, nil
			}
		case <-time.After(time.Duration(s.timeout) * time.Second):
			// Timeout, return that the port is closed
			return PortClosed, nil
		}
	}
}

func GetLocalIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil

}
