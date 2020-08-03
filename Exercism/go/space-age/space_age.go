package space

// Planet type
type Planet string

// SecondsEarth constant
const SecondsEarth = 31557600

// MercuryPeriod constant
const MercuryPeriod = 0.2408467

// VenusPeriod constant
const VenusPeriod = 0.61519726

// MarsPeriod constant
const MarsPeriod = 1.8808158

// JupiterPeriod constant
const JupiterPeriod = 11.862615

// SaturnPeriod constant
const SaturnPeriod = 29.447498

// UranusPeriod constant
const UranusPeriod = 84.016846

// NeptunePeriod constant
const NeptunePeriod = 164.79132

// Age method
func Age(seconds float64, planet Planet) (age float64) {

	ageEarch := seconds / SecondsEarth
	age = ageEarch
	switch planet {
	case "Mercury":
		age = ageEarch / MercuryPeriod
	case "Venus":
		age = ageEarch / VenusPeriod
	case "Mars":
		age = ageEarch / MarsPeriod
	case "Jupiter":
		age = ageEarch / JupiterPeriod
	case "Saturn":
		age = ageEarch / SaturnPeriod
	case "Uranus":
		age = ageEarch / UranusPeriod
	case "Neptune":
		age = ageEarch / NeptunePeriod
	}

	return age
}
