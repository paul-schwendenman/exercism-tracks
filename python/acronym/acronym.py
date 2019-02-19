import re


def get_letter(word):
    for char in word:
        if char.isalpha():
            break
    return char.upper()


def abbreviate(words):
    word_list = re.split(r'[ -]+', words)
    return ''.join(filter(lambda char: char.isalpha(), map(get_letter, word_list)))
