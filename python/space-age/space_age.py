from __future__ import division


EARTH_YEAR_IN_SECONDS = 31557600
MERCURY_RATIO = 0.2408467
VENUS_RATIO = 0.61519726
MARS_RATIO = 1.8808158
JUPITER_RATIO = 11.862615
SATURN_RATIO = 29.447498
URANUS_RATIO = 84.016849
NEPTUNE_RATIO = 164.79132


class SpaceAge(object):
    def __init__(self, seconds):
        self.seconds = seconds

    def on_earth_raw(self):
        return self.seconds / EARTH_YEAR_IN_SECONDS

    def on_earth(self):
        return round(self.on_earth_raw(), 2)

    def on_mercury(self):
        return round(self.on_earth_raw() / MERCURY_RATIO, 2)

    def on_venus(self):
        return round(self.on_earth_raw() / VENUS_RATIO, 2)

    def on_mars(self):
        return round(self.on_earth_raw() / MARS_RATIO, 2)

    def on_jupiter(self):
        return round(self.on_earth_raw() / JUPITER_RATIO, 2)

    def on_saturn(self):
        return round(self.on_earth_raw() / SATURN_RATIO, 2)

    def on_uranus(self):
        return round(self.on_earth_raw() / URANUS_RATIO, 2)

    def on_neptune(self):
        return round(self.on_earth_raw() / NEPTUNE_RATIO, 2)
