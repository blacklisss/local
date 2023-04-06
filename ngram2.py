import argparse
import os

def load_text(filename):
    with open(filename, 'r', encoding='utf-8') as f:
        text = f.read()
    return text

def generate_ngrams(text, n):
    ngrams = []
    for i in range(len(text)-n+1):
        ngrams.append(text[i:i+n])
    return ngrams

def compare_texts(text1, text2, n):
    ngrams1 = generate_ngrams(text1, n)
    ngrams2 = generate_ngrams(text2, n)
    set1 = set(ngrams1)
    set2 = set(ngrams2)
    overlap = set1 & set2
    score = len(overlap) / max(len(set1), len(set2))
    return score

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Compare two texts using n-gram similarity')
    parser.add_argument('mainfile', type=str, help='filename of the main text')
    parser.add_argument('dir', type=str, help='directory containing the other texts')
    parser.add_argument('--n', type=int, default=2, help='value of n for the n-grams')
    args = parser.parse_args()

    text1 = load_text(args.mainfile)

    # iterate over files in the directory
    scores = []
    for filename in os.listdir(args.dir):
        if filename.endswith('.html'):
            filepath = os.path.join(args.dir, filename)
            text2 = load_text(filepath)
            similarity = compare_texts(text1, text2, args.n)
            scores.append((filename, similarity))

    # sort by descending similarity
    scores.sort(key=lambda x: x[1], reverse=True)

    # print the results
    print(f'Similarities between "{args.mainfile}" and texts in "{args.dir}":')
    for filename, similarity in scores:
        print(f'{filename}: {similarity:.2f}')
