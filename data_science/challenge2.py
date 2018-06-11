import pandas as pd

# from . import make_answer, submit
from data_science import make_answer, submit


def get_math_grades():
    df = pd.read_csv('test2.csv')
    media = df['NU_NOTA_CH']
    media += df['NU_NOTA_CN']
    media += df['NU_NOTA_LC']
    media += df['NU_NOTA_REDACAO']
    media /= 4
    answer = []
    for index in df.index:
        answer.append({
            'NU_INSCRICAO': df.at[index, 'NU_INSCRICAO'],
            'NU_NOTA_MT': float(f'{media.at[index]:.1f}'),
        })
    return answer


if __name__ == '__main__':
    url = 'https://api.codenation.com.br/v1/user/acceleration/data-science/challenge/enem-2/submit'
    answer = get_math_grades()
    data = make_answer(answer)
    submit(url, data)
