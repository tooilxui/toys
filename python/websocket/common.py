#! encoding=utf-8

from auth import Login

HOST = 'https://v2.koalacam.net'
USER = 'username'
PASSWD = 'password'
LOCAL_HOST_IP = 'x.x.x.x'
RTSP_URL = 'rtsp://x.x.x.x/user=admin&password=&channel=1&stream=0.sdp'

SESSION = Login(HOST, USER, PASSWD).session