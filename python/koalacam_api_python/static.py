#! encoding=utf-8
import request
import json
import traceback

# from pip._vendor import requests
import requests
from common import LOCAL_HOST_IP


# 1:N识别
def recognize(screen_token, path):
    url = 'http://' + LOCAL_HOST_IP + ':8866/recognize'
    ret = None
    try:
        ret = requests.session().post(url, {'screen_token': screen_token},
                                      files={'image': ('filename.jpg', open(path))}).content
        data = json.loads(ret)
        print(data)
    except:
        print(ret)
        print(traceback.format_exc())


# 1:1比对
def checkin(person_id, path):
    url = 'http://' + LOCAL_HOST_IP + ':8866/checkin'
    ret = None
    try:
        ret = requests.session().post(url, {'person_id': person_id},
                                      files={'image': ('filename.jpg', open(path))}).content
        data = json.loads(ret)
        print(data)
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    # 1:N识别
    recognize('screen_token', '/Users/Desktop/subjects/Pad_B.jpg')
    # 1:1比对
    checkin(135, '/Users/Desktop/subjects/Pad_B.jpg')
