#! encoding=utf-8

import request
import json

##登录
# from pip._vendor import requests
import requests


class Login(object):
    def __init__(self, host, user, password):
        self.host = host
        self.user = user
        self.password = password
        self.session = self.init_session()

    def build_url(self, path):
        return self.host + path

    def init_session(self):
        session = requests.session()
        url = self.build_url('/auth/login')
        headers = {'User-Agent': 'Koala Admin'}
        data = {'username': self.user, 'password': self.password}
        ret = session.post(url, json=data, headers=headers).content
        data = json.loads(ret)
        if data['code'] == 0:
            print(data['data']['id'])
        elif data['code'] != 0:
            print(data['code'], data['desc'])
        return session


if __name__ == '__main__':
    Login('host', 'username', 'password')
