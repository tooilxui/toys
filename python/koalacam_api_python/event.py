#! encoding=utf-8

import json
import traceback
from common import HOST, SESSION


##获取历史识别记录
def get_events(user_role):
    url = HOST + '/event/events'
    ret = None
    try:
        ret = SESSION.get(url + '?user_role=' + str(user_role))
        data = json.loads(ret.text)
        if data['code'] != 0:
            print(data['code'], data['desc'])
        elif data['code'] == 0:
            for record in data['data']:
                print(record['id'], record['confidence'])
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    print(None)
