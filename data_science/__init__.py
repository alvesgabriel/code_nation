import json

import requests


def make_answer(answer):
    return {
        'token': '5a66cb603690853e851f162caa295cc7c7dc721c',
        'email': 'gabriel.alves.monteiro1@gmail.com',
        'answer': answer,
    }


def submit(url, data):
    resp = requests.post(url, data=json.dumps(data))
    print(resp.request.body)
    print(resp, resp.json())