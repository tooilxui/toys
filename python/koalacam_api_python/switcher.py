#! encoding=utf-8

import time
import socket
import traceback


###网络开关

def open_door(ip):
    if ip is None:
        return False
    address = (ip, 5000)
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    try:
        sock.sendto("on1", address)
        time.sleep(0.5);
        sock.sendto("off1", address)
    except:
        print(traceback.format_exc())
        return False
    return True


if __name__ == '__main__':
    print(open_door('y.y.y.y'))
