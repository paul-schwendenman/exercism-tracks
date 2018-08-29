def is_leap_year(year):
    return True if is_divisible_by(year, 4) and (not is_divisible_by(year, 100) or is_divisible_by(year, 400)) else False

def is_divisible_by(number, divisor):
    return (number % divisor) == 0
