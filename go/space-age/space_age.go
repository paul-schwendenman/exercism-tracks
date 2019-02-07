package space

// Planet is a planet
type Planet string

// Age in years based on orbit time of a given planet
func Age(age float64, planet Planet) float64 {
	secondsInEarthYear := 60 * 60 * 24 * 365.25

	earthYears := age / secondsInEarthYear
	years := float64(0)

	if planet == "Earth" {
		years = earthYears
	} else if planet == "Mercury" {
		years = earthYears / 0.2408467
	} else if planet == "Venus" {
		years = earthYears / 0.61519726
	} else if planet == "Mars" {
		years = earthYears / 1.8808158
	} else if planet == "Jupiter" {
		years = earthYears / 11.862615
	} else if planet == "Saturn" {
		years = earthYears / 29.447498
	} else if planet == "Uranus" {
		years = earthYears / 84.016846
	} else if planet == "Neptune" {
		years = earthYears / 164.79132
	}
	return years
}
