#! encoding=utf-8


import json
import traceback

# from pip._vendor import requests
import requests
from common import HOST, USER, PASSWD


##pad
def login(pad_id):
    url = HOST + '/pad/login'
    ret = None
    try:
        data = {'username': USER, 'password': PASSWD, 'device_type': 2, 'pad_id': pad_id}
        ret = requests.session().post(url, json=data).content
        data = json.loads(ret)
        if data['code'] != 0:
            print(data['code'], data['desc'])
        elif data['code'] == 0:
            for box in data['data']['boxes']:
                print(box['box_address'], box['box_token'])
            print(data['data']['screen_token'])
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    login('pad_id')
