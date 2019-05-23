import json
import os
import requests
from hashlib import sha1


def open_json(filename):
    with open(filename) as fl:
        json_obj = json.loads(fl.read())
    fl.close()
    return json_obj


def decrypt(msg, step=3):
    decrypt_msg = ''
    for char in msg:
        ascii_num = ord(char)
        if ascii_num in range(102,123):
            decrypt_msg += chr(ascii_num - step)
        elif ascii_num in range(97,103):
            decrypt_msg += chr(122 - (step + (96 - ascii_num)))
        else:
            decrypt_msg += char
    return decrypt_msg


def to_sha1(msg):
    b = bytes(msg, 'utf-8')
    m = sha1(b)
    return m.hexdigest()

if __name__ == "__main__":
    j = open_json('answer.json')
    j['decifrado'] =  decrypt(j['cifrado'], j['numero_casas'])
    j['resumo_criptografico'] = to_sha1(j['decifrado'])
    url = 'https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token='
    files = {'answer': open('answer.json', 'rb')}
    resp = requests.post(url + j['token'], files=files)
    print(j)
    print(resp.status_code, resp.headers)
    print(resp.content)