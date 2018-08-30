from __future__ import division
import math


class Rational(object):
    def __init__(self, numer, denom):
        gcd = math.gcd(numer, denom)
        reduced_numer = numer // gcd

        self.numer = -reduced_numer if denom < 0 else reduced_numer
        self.denom = abs(denom // gcd)

    def __eq__(self, other):
        return self.numer == other.numer and self.denom == other.denom

    def __repr__(self):
        return '{}/{}'.format(self.numer, self.denom)

    def __add__(self, other):
        numer = (self.numer * other.denom) + (self.denom * other.numer)
        denom = (self.denom * other.denom)
        return Rational(numer, denom)

    def __sub__(self, other):
        numer = (self.numer * other.denom) - (self.denom * other.numer)
        denom = (self.denom * other.denom)
        return Rational(numer, denom)

    def __mul__(self, other):
        numer = self.numer * other.numer
        denom = self.denom * other.denom
        return Rational(numer, denom)

    def __truediv__(self, other):
        numer = self.numer * other.denom
        denom = self.denom * other.numer
        return Rational(numer, denom)

    def __abs__(self):
        return Rational(abs(self.numer), abs(self.denom))

    def __pow__(self, power):
        if isinstance(power, float):
            return (self.numer ** power) / (self.denom ** power)
        elif power < 0:
            return Rational(self.denom**abs(power), self.numer**abs(power))

        return Rational(self.numer**power, self.denom**power)

    def __rpow__(self, base):
        return base**(self.numer / self.denom)
