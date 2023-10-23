package space

type Planet string

var planet_to_earth_years = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

var one_earth_year_second float64 = 31557600

func Age(seconds float64, planet Planet) float64 {
	weight, ok := planet_to_earth_years[planet]
	if ok {
		return seconds / one_earth_year_second / weight
	} else {
		return -1
	}
}
