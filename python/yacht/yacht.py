from collections import Counter

YACHT = 0
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = 7
FOUR_OF_A_KIND = 8
LITTLE_STRAIGHT = 10
BIG_STRAIGHT = 11
CHOICE = 12


def score(dice, category):
    if category in (ONES, TWOS, THREES, FOURS, FIVES, SIXES):
        return score_numbers(dice, category)
    elif category == FULL_HOUSE and valid_full_house(dice):
        return sum(dice)
    elif category == FOUR_OF_A_KIND:
        return score_four_of_a_kind(dice)
    elif category == LITTLE_STRAIGHT and valid_little_straight(dice):
        return 30
    elif category == BIG_STRAIGHT and valid_big_straight(dice):
        return 30
    elif category == YACHT and valid_yacht(dice):
        return 50
    elif category == CHOICE:
        return sum(dice)
    else:
        return 0


def score_numbers(dice, number):
    return sum(number for die in dice if die == number)


def valid_full_house(dice):
    common = Counter(dice).most_common(2)
    return common[0][1] == 3 and common[1][1] == 2


def score_four_of_a_kind(dice):
    common = Counter(dice).most_common(1)
    return common[0][0] * 4 if common[0][1] >= 4 else 0


def valid_yacht(dice):
    common = Counter(dice).most_common(1)
    return common[0][1] == 5


def valid_little_straight(dice):
    return set(dice) == set((1, 2, 3, 4, 5))


def valid_big_straight(dice):
    return set(dice) == set((2, 3, 4, 5, 6))

