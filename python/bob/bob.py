import re


def hey(phrase):
    trimmed_phrase = phrase.strip()

    if is_silence(trimmed_phrase):
        return 'Fine. Be that way!'
    elif is_question(trimmed_phrase) and is_yelling(trimmed_phrase):
        return 'Calm down, I know what I\'m doing!'
    elif is_question(trimmed_phrase):
        return 'Sure.'
    elif is_yelling(trimmed_phrase):
        return 'Whoa, chill out!'

    return 'Whatever.'


def is_yelling(phrase):
    return phrase == phrase.upper() and re.search(r'[A-Z]', phrase)


def is_question(phrase):
    return phrase[-1] == '?'


def is_silence(phrase):
    return len(phrase) == 0
