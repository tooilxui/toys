#!/usr/bin/env python3
import time
import serial
import serial.tools.list_ports

coms = serial.tools.list_ports.comports()
for a in coms:
    print(a)

# ADAM-4561 driver download:
# http://downloadt.advantech.com/download/downloadsr.aspx?File_Id=1-19MVPYF
#
# code ref:
# https://pythonhosted.org/pyserial/pyserial_api.html
# https://pythonhosted.org/pyserial/shortintro.html#testing-ports
# https://www.jianshu.com/p/767fd1fbcaae
#
ser = serial.Serial(
    port='COM3',                  # 連接PORT
    baudrate=9600,                # 波特率
    bytesize=serial.EIGHTBITS,    # 資料位元
    parity=serial.PARITY_NONE,    # 同位檢查
    stopbits=serial.STOPBITS_ONE, # 停止位元
    # xonxoff=False,                # 軟體流控 Xon / Xoff
    # rtscts=False,                 # 硬體流控 RTS/CTS
    # dsrdtr=False,                 # 硬體流控 DSR/DTR
    timeout=0.5                   # 讀取超時
)

a = 0
while a < 1000:
    # x = ser.read(18)      # read one byte
    # print(a, ' : ', x)
    # x = ser.read()    # read up to ten bytes (timeout)
    # print(a, ' : ', x)
    x = ser.readline()  # read a '\n' terminated line
    print(a, ' : ', x)
    a = a+1

if ser.is_open:
    ser.close()