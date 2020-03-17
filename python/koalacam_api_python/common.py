#! encoding=utf-8

from auth import Login

# HOST = 'https://v2.koalacam.net'
# USER = 'username'
# PASSWD = 'password'
# LOCAL_HOST_IP = 'x.x.x.x'
# RTSP_URL = 'rtsp://x.x.x.x/user=admin&password=&channel=1&stream=0.sdp'

HOST = 'http://192.1.15.217'
USER = 'face@aptg.com.tw'
PASSWD = '123456'
LOCAL_HOST_IP = '192.1.15.217'
RTSP_URL = 'rtsp://192.1.15.218/user=admin&password=&channel=1&stream=0.sdp'


SESSION = Login(HOST, USER, PASSWD).session

