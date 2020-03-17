from urllib.parse import urlencode

import websocket
import json
from common import LOCAL_HOST_IP, RTSP_URL


def on_message(ws, message):
    result = json.loads(message)
    # print json.dumps(result, indent=4, sort_keys=True)
    print(json.dumps(print(json.dumps(result, indent=4, sort_keys=True))))


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
    ws.run_forever()
