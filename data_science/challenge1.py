import pandas as pd

from data_science import make_answer, submit


def get_better_grades():
    df = pd.read_csv('train.csv')
    total = df['NU_NOTA_MT'] * 3
    total += df['NU_NOTA_CN'] * 2
    total += df['NU_NOTA_LC'] * 1.5
    total += df['NU_NOTA_CH'] * 1
    total += df['NU_NOTA_REDACAO'] * 3
    avg = total / 10.5
    data = avg.sort_values(ascending=False).head(20)
    values = data.to_dict()
    answer = []
    for index in data.index:
        answer.append({
            'NU_INSCRICAO': df.at[index, 'NU_INSCRICAO'],
            'NOTA_FINAL': float(f'{values[index]:.1f}'),
        })
    return answer


if __name__ == '__main__':
    url = 'https://api.codenation.com.br/v1/user/acceleration/data-science/challenge/enem-1/submit'
    answer = get_better_grades()
    data = make_answer(answer)
    submit(url, data)
