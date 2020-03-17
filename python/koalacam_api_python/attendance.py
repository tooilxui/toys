#! encoding=utf-8

import json
import time
import traceback
from common import HOST, SESSION


##获取考勤记录
def get_records(user_name):
    url = HOST + '/attendance/records'
    ret = None
    try:
        ret = SESSION.get(url + '?user_name=' + user_name)
        data = json.loads(ret.text)
        if data['code'] != 0:
            print data['code'], data['desc']
        elif data['code'] == 0:
            for record in data['data']:
                print record['subject']['id'], time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(record['date'])), record['worktime']
    except:
        print ret
        print traceback.format_exc()


if __name__ == '__main__':
    get_records('xx')
