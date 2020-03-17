#! encoding=utf-8

import json
from common import HOST, SESSION
import traceback


##员工管理

def import_subject(subject_type, name, gender, photo_ids):
    url = HOST + '/subject'
    data = {'subject_type': subject_type, 'name': name, 'gender': gender, 'photo_ids': photo_ids}
    ret = None
    try:
        ret = SESSION.post(url, json=data).content
        data = json.loads(ret)
        if data['code'] == 0:
            return data['data']['id']
        else:
            print(data['code'], data['desc'])
    except:
        print(ret)
        print(traceback.format_exc())
    return None


def update_subject(subject_type, name, gender, subject_id, photo_ids):
    url = HOST + '/subject/' + str(subject_id)
    data = {'subject_type': subject_type, 'name': name, 'gender': gender, 'photo_ids': photo_ids}
    ret = None
    try:
        ret = SESSION.put(url, json=data).content
        data = json.loads(ret)
        if data['code'] == 0:
            return data['data']['id']
        else:
            print(data['code'], data['desc'])
    except:
        print(ret)
        print(traceback.format_exc())


def photo(path):
    url = HOST + '/subject/photo'
    img = open(path, 'rb')
    ret = None
    try:
        ret = SESSION.post(url, files={'photo': ('filename.jpg', img)}).content
        data = json.loads(ret)
        if data['code'] == 0:
            return data['data']['id']
        elif data['code'] != 0:
            return data['code'], data['desc']
    except:
        print(ret)
        print(traceback.format_exc())


if __name__ == '__main__':
    ##上传底库照片
    photo_ids = [photo("/Users/Desktop/subjects/T1.jpg")]
    print(photo_ids)
    ##新增一个用户
    subject_id = import_subject(0, '测试员工', 1, photo_ids)
    print(subject_id)
    ##更新底库到用户
    update_subject(0, '测试员工' + str(subject_id), 2, subject_id, photo_ids)
