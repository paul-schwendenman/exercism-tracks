package space

// Planet is a planet
type Planet string

var secondsInEarthYear = 60 * 60 * 24 * 365.25
var orbitalPeriod = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age in years based on orbit time of a given planet
func Age(age float64, planet Planet) float64 {
	return age / (orbitalPeriod[planet] * secondsInEarthYear)
}
