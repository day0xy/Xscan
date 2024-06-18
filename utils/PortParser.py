def parse_ports(port_str):
    ports = []
    # 分割字符串以处理多个端口或端口范围
    port_parts = port_str.split(",")
    for part in port_parts:
        if "-" in part:
            # 处理端口范围
            start, end = part.split("-")
            ports.extend(range(int(start), int(end) + 1))
        else:
            # 处理单个端口
            ports.append(int(part))
    return ports
