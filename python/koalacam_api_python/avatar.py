#! encoding=utf-8

import json
import traceback
from common import HOST, SESSION


def upload_avatar(subject_id, path):
    url = HOST + '/subject/avatar'
    img = open(path, 'rb')
    ret = None
    try:
        ret = SESSION.post(url, {'subject_id': subject_id}, files={'avatar': ('filename.jpg', img)}).content
        data = json.loads(ret)
        if data['code'] !=0:
            print data['code'], data['desc']
        elif data['code'] == 0:
            return data['data']['url']
    except:
        print ret
        print traceback.format_exc()
    return None



if __name__ == '__main__':
    #上传识别头像
    imageUrl = upload_avatar(126526, '/Users/Desktop/subjects/megvii.png')
    print imageUrl
