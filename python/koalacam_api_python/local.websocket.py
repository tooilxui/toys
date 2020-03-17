#! encoding=utf-8
from urllib.parse import urlencode

import websocket
from common import LOCAL_HOST_IP, RTSP_URL


##视频流识别
#  @param wsUrl websocket接口 例如 ws://192.168.1.50:9000/video
#  @param rtspUrl 视频流地址 门禁管理-门禁设备-视频流地址
#  				  例如 rtsp://192.168.0.100/live1.sdp
#                 或者 rtsp://admin:admin12345@192.168.1.100/live1.sdp
#                 或者 rtsp://192.168.1.100/user=admin&password=&channel=1&stream=0.sdp
#                 或者 rtsp://192.168.1.100/live1.sdp
#                 		?__exper_tuner=lingyun&__exper_tuner_username=admin
#                 		&__exper_tuner_password=admin&__exper_mentor=motion
#                 		&__exper_levels=312,1,625,1,1250,1,2500,1,5000,1,5000,2,10000,2,10000,4,10000,8,10000,10
#                 		&__exper_initlevel=6

def on_message(ws, message):
    print('### message ###', message)


def on_error(ws, error):
    print('### error ###', error)


def on_close(ws):
    print('### closed ###')


def on_open(ws):
    print('### open ###')


if __name__ == "__main__":
    websocket.enableTrace(True)
    param = {'url': RTSP_URL}
    ws = websocket.WebSocketApp("ws://" + LOCAL_HOST_IP + ':9000/video?' + urlencode(param),
                                on_open=on_open,
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    print("After websocket")
    ws.run_forever()
