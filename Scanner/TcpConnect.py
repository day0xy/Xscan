from scapy.all import *
import socket


def tcp_connect_scan(target_ip, ports):
    scan_results = {}  # 存储扫描结果的字典
    for port in ports:
        try:
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.settimeout(1)
            result = sock.connect_ex((target_ip, port))
            if result == 0:
                scan_results[port] = "open"
            else:
                scan_results[port] = "closed"
            sock.close()
        except socket.error as err:
            print(f"Socket error on port {port}: {err}")

    # 打印扫描结果
    print("Xscan Tcp-Connect scan result:")
    print("PORT\tSTATE")
    for port, status in scan_results.items():
        print(f"{port}\t{status}")
