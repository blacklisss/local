import argparse

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
    parser.add_argument('filename1', type=str, help='filename of the first text')
    parser.add_argument('filename2', type=str, help='filename of the second text')
    parser.add_argument('--n', type=int, default=2, help='value of n for the n-grams')
    args = parser.parse_args()

    text1 = load_text(args.filename1)
    text2 = load_text(args.filename2)
    similarity = compare_texts(text1, text2, args.n)
    print(f'The similarity score between the texts is: {similarity}')
