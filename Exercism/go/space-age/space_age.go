package space

// Planet type
type Planet string

// SecondsEarth constant
const SecondsEarth = 31557600

// PlanetPeriod map
var PlanetPeriod = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132}

// Age method
func Age(seconds float64, planet Planet) (age float64) {

	ageEarch := seconds / SecondsEarth

	age = ageEarch / PlanetPeriod[planet]

	return age
}
