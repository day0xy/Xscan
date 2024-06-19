import socket


def udp_scan(target_ip, ports):
    results = {}

    for port in ports:
        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        sock.settimeout(1.0)
        is_open = False

        try:
            # Send a specific request data
            message = "ping"
            sock.sendto(message.encode(), (target_ip, port))

            # Retry mechanism, send data multiple times
            for _ in range(3):
                try:
                    data, _ = sock.recvfrom(1024)
                    if data.decode().startswith("Received:"):
                        is_open = True
                        break
                except socket.timeout:
                    pass
        except Exception as e:
            # Other exceptional cases
            print(f"Error on port {port}: {e}")
        finally:
            sock.close()

        results[port] = is_open

    # Print scan results
    print("UDP scan results:\n")
    print("PORT    STATE")

    for port, is_open in sorted(results.items()):
        state = "open" if is_open else "closed"
        print(f"{port}/udp  {state}")
