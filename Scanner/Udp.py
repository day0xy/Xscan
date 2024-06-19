import socket


def udp_scan(target_ip, ports):
    results = {}

    for port in ports:
        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        sock.settimeout(1.0)
        is_open = False

        try:
            # 发送一个空数据包到指定端口
            sock.sendto(b"", (target_ip, port))
            data, _ = sock.recvfrom(1024)
            is_open = True
        except socket.timeout:
            # 超时表示端口没有响应，可能是关闭的或被防火墙阻止
            pass
        except Exception as e:
            # 其他异常情况
            print(f"Error on port {port}: {e}")
        finally:
            sock.close()

        results[port] = is_open

    # 打印扫描结果
    print("Xscan UDP connect scan result:\n")
    print("PORT    STATE")

    for port, is_open in sorted(results.items()):
        state = "open" if is_open else "closed"
        print(f"{port}/udp  {state}")
