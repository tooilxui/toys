#! encoding=utf-8

import json
import traceback
from common import HOST, SESSION


##年龄性别统计
def overview(start_time, end_time):
    url = HOST + '/statistics/overview'
    ret = None
    try:
        ret = SESSION.get(url + '?start_time=' + start_time + '&end_time=' + end_time)
        data = json.loads(ret.text)
        if data['code'] != 0:
            print(data['code'], data['desc'])
        elif data['code'] == 0:
            for age in data['data']['ages']:
                print(age)
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    overview('2015-09-07 00:00:00', '2016-09-10 23:59:59')
