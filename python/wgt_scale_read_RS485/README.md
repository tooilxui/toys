# testing hardware info. :
- data acquisition module : ADAM-4561
- weighing Scales : AD-4402
- pc : windows 10

# require software:
- [ADAM-4561 driver](http://downloadt.advantech.com/download/downloadsr.aspx?File_Id=1-19MVPYF)
- PySerail Package : install by following command  ```$  pip install pyserial```

# code ref:
- https://pythonhosted.org/pyserial/pyserial_api.html
- https://pythonhosted.org/pyserial/shortintro.html#testing-ports
- https://www.jianshu.com/p/767fd1fbcaae

# ADAM-4561 port info. check steps (Windows) :
1. [我的電腦] 右鍵 [內容] 
2. [裝置管理員] > 展開[連接埠(COM和LPT)] 
3. [Silicon Labs CP210x USB to UART Bridge (COM3)] 右鍵 [內容] 
4. [連接埠設定]

# AD-4402 Comunication Mode Setting :
- 「 + Enter -> Function -> Set Function -> Serial -> RS485 -> r5 f- 2 (Comunication mode) -> Set to Stream Mode
- ref :  [AD-4402 User manual ](http://esp.andweighing.com/uploads/documents/AD4402_IM.pdf) Pg. 124