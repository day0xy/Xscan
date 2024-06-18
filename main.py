import argparse

from utils.PortParser import *


def main():
    parser = argparse.ArgumentParser(
        description="Xscan_python - A simple port scanner written in Python"
    )
    parser.add_argument("target", help="Target IP address")
    parser.add_argument(
        "-p", "--ports", nargs="+", type=int, help="Ports to scan", required=True
    )
    parser.add_argument(
        "-s",
        "--scan",
        choices=["connect", "syn"],
        default="connect",
        help="Scan type: connect or syn",
    )

    args = parser.parse_args()

    target_ip = args.target
    ports = args.ports
    scan_type = args.scan

    if scan_type == "connect":
        for port in ports:
            tcp_connect_scan(target_ip, port)
    elif scan_type == "syn":
        for port in ports:
            tcp_syn_scan(target_ip, port)
    elif scan_type == "udp":
        for port in ports:
            udp_scan(target_ip, port)
    elif scan_type == "fid":
        fin_scan(target_ip)
    elif scan_type == "ping":
        ping_scan(target_ip)


if __name__ == "__main__":
    main()
