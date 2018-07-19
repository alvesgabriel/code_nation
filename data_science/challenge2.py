import json

from sklearn.linear_model import LinearRegression
from sklearn import svm
import pandas as pd
import requests

# from . import make_answer, submit
# from data_science import make_answer, submit


def make_answer(answer):
    return {
        'token': '5a66cb603690853e851f162caa295cc7c7dc721c',
        'email': 'gabriel.alves.monteiro1@gmail.com',
        'answer': answer,
    }


def submit(url, data):
    print(type(data))
    resp = requests.post(url, data=json.dumps(data))
    print(resp, resp.json())


def get_math_grades():
    df = pd.read_csv('test2.csv')
    df['NU_NOTA_CH'] = df['NU_NOTA_CH'].fillna(0)
    df['NU_NOTA_CN'] = df['NU_NOTA_CN'].fillna(0)
    df['NU_NOTA_LC'] = df['NU_NOTA_LC'].fillna(0)
    df['NU_NOTA_REDACAO'] = df['NU_NOTA_REDACAO'].fillna(0)
    grades = pd.concat([df['NU_NOTA_CH'], df['NU_NOTA_CN'], df['NU_NOTA_LC'], df['NU_NOTA_REDACAO']], axis=1)

    media = df['NU_NOTA_CH']
    media += df['NU_NOTA_CN']
    media += df['NU_NOTA_LC']
    media += df['NU_NOTA_REDACAO']
    media /= 4

    lm = LinearRegression()
    lm.fit(grades, media)
    nu_nota_mt = lm.predict(grades)

    answer = []
    for index in df.index:
        answer.append({
            'NU_INSCRICAO': df.at[index, 'NU_INSCRICAO'],
            'NU_NOTA_MT': float(f'{nu_nota_mt[index]:.1f}'),
        })
    print(answer[0], media[0])
    return answer


if __name__ == '__main__':
    url = 'https://api.codenation.com.br/v1/user/acceleration/data-science/challenge/enem-2/submit'
    answer = get_math_grades()
    data = make_answer(answer)
    # submit(url, data)
