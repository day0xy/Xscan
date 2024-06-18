import argparse

from utils.PortParser import *
from Scanner.TcpConnect import *
from Scanner.TcpSyn import *
from Scanner.Ping import *

# from Scanner.Fin import *
# from Scanner.Udp import *


def main():
    parser = argparse.ArgumentParser(
        description="Xscan_python - A simple port scanner written in Python"
    )
    parser.add_argument("-t", "--target", help="Target Ip")

    parser.add_argument("-p", "--ports", nargs="+", type=str, required=True)
    parser.add_argument(
        "-s",
        "--scan",
        choices=["connect", "syn" "udp", "fin", "ping"],
        default="connect",
    )

    args = parser.parse_args()

    target_ip = args.target
    ports = parse_ports(args.ports)
    scan_type = args.scan

    if scan_type == "connect":
        tcp_connect_scan(target_ip, ports)
    elif scan_type == "syn":
        tcp_syn_scan(target_ip, ports)
    elif scan_type == "udp":
        udp_scan(target_ip, ports)
    elif scan_type == "fin":
        fin_scan(target_ip)
    elif scan_type == "ping":
        ping(target_ip)


if __name__ == "__main__":
    main()
