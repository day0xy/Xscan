from scapy.all import *


def tcp_syn_scan(target_ip, port):
    src_port = RandShort()
    syn_packet = IP(dst=target_ip) / TCP(sport=src_port, dport=port, flags="S")
    response = sr1(syn_packet, timeout=1, verbose=False)

    if response is None:
        print(f"Port {port} is filtered (no response)")
    elif response.haslayer(TCP):
        if response.getlayer(TCP).flags == 0x12:  # SYN+ACK
            # Send RST to close the connection
            rst_packet = IP(dst=target_ip) / TCP(sport=src_port, dport=port, flags="R")
            send(rst_packet, verbose=False)
            print(f"Port {port} is open")
        elif response.getlayer(TCP).flags == 0x14:  # RST
            print(f"Port {port} is closed")
    elif response.haslayer(ICMP):
        if int(response.getlayer(ICMP).type) == 3 and int(
                response.getlayer(ICMP).code
        ) in [1, 2, 3, 9, 10, 13]:
            print(f"Port {port} is filtered (ICMP response)")
