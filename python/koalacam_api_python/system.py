#! encoding=utf-8

import json
import traceback
from common import HOST, SESSION


##门禁管理

def get_screen():
    url = HOST + '/system/screen'
    ret = None
    try:
        ret = SESSION.get(url)
        data = json.loads(ret.text)
        print(data)
        if data['code'] != 0:
            print(data['code'], data['desc'])
        elif data['code'] == 0:
            for screen in data['data']:
                print(screen['id'], screen['camera_position'])
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    get_screen()
