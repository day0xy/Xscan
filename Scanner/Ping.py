import os
import socket
import struct
import time
import select


def checksum(source_string):
    """
    计算校验和
    """
    sum = 0
    max_count = (len(source_string) / 2) * 2
    count = 0
    while count < max_count:
        val = source_string[count + 1] * 256 + source_string[count]
        sum = sum + val
        sum = sum & 0xFFFFFFFF
        count = count + 2
    if max_count < len(source_string):
        sum = sum + source_string[len(source_string) - 1]
        sum = sum & 0xFFFFFFFF
    sum = (sum >> 16) + (sum & 0xFFFF)
    sum = sum + (sum >> 16)
    answer = ~sum
    answer = answer & 0xFFFF
    answer = answer >> 8 | (answer << 8 & 0xFF00)
    return answer


def create_packet(id):
    """
    创建ICMP请求包
    """
    header = struct.pack("bbHHh", 8, 0, 0, id, 1)
    data = 192 * b"Q"
    my_checksum = checksum(header + data)
    header = struct.pack("bbHHh", 8, 0, socket.htons(my_checksum), id, 1)
    return header + data


def ping(host):
    """
    发送ping请求
    """
    try:
        dest_addr = socket.gethostbyname(host)
    except socket.gaierror:
        return "Host name could not be resolved. Exiting"

    print(f"Pinging {host} [{dest_addr}] with 32 bytes of data:")

    icmp = socket.getprotobyname("icmp")
    try:
        my_socket = socket.socket(socket.AF_INET, socket.SOCK_RAW, icmp)
    except socket.error as e:
        if e.errno == 1:
            # Not run as root
            return "ICMP messages can only be sent from processes running as root."
        raise

    my_id = os.getpid() & 0xFFFF

    packet = create_packet(my_id)
    my_socket.sendto(packet, (dest_addr, 1))

    start = time.time()
    ready = select.select([my_socket], [], [], 2)
    if ready[0] == []:
        print("Request timed out.")
        return

    rec_packet, addr = my_socket.recvfrom(1024)
    time_received = time.time()
    rtt = (time_received - start) * 1000
    print(f"Reply from {addr[0]}: bytes=32 time={rtt:.2f}ms")
    my_socket.close()
