from scapy.all import sr1, IP, TCP
import logging

# 关闭 scapy 的冗长输出
logging.getLogger("scapy.runtime").setLevel(logging.ERROR)


def tcp_fin_scan(target_ip, ports):
    # 结果字典
    results = {}

    for port in ports:
        # 构造 FIN 包
        fin_packet = IP(dst=target_ip) / TCP(dport=port, flags="F")

        # 发送包并等待响应
        response = sr1(fin_packet, timeout=2, verbose=0)

        if response is None:
            results[port] = "open"
        elif response.haslayer(TCP) and response.getlayer(TCP).flags == 0x14:  # RST-ACK
            results[port] = "closed"
        else:
            results[port] = "filtered"

    # 打印结果
    print("PORT     STATE")
    for port, state in results.items():
        print(f"{port}/tcp {' ' * (5-len(str(port)))} {state}")
