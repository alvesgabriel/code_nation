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


def write_json(filename, json_obj):
    with open(filename, 'w') as fl:
        fl.write(json.dumps(json_obj, indent=True))
        fl.close()


if __name__ == "__main__":
    filename = 'answer.json'
    json_obj = open_json(filename)
    json_obj['decifrado'] =  decrypt(json_obj['cifrado'], json_obj['numero_casas'])
    json_obj['resumo_criptografico'] = to_sha1(json_obj['decifrado'])
    write_json(filename, json_obj)
    url = 'https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token='
    with open('answer.json', 'rb') as fl:
        files = {'answer': fl}
        print(files)
        resp = requests.post(url + json_obj['token'], files=files)
        print(resp.status_code, resp.headers)
        print(resp.json())
        fl.close()