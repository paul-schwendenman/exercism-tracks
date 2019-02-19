import re


def abbreviate(words):
    word_list = re.findall(r'([A-Za-z])[A-Za-z\']*', words)
    return ''.join(word_list).upper()
