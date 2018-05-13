// Package space contains helpers and values for space age computations.
package space

// Planet describes a planet object.
type Planet string

// PlanetYears describes the mapping of planets and their earth years.
type PlanetYears map[Planet]float64

// lookup is the lookup table for planet years.
var lookup PlanetYears

// Register all planets.
func init() {
	lookup = make(PlanetYears)
	lookup["Earth"] = 31557600
	lookup["Mercury"] = 0.2408467
	lookup["Venus"] = 0.61519726
	lookup["Mars"] = 1.8808158
	lookup["Jupiter"] = 11.862615
	lookup["Saturn"] = 29.447498
	lookup["Uranus"] = 84.016846
	lookup["Neptune"] = 164.79132
}

// Age computes for the Earth-years old given seconds olds.
func Age(s float64, p Planet) float64 {
	if p == "Earth" {
		return s / lookup["Earth"]
	}

	return (s / lookup[p]) / lookup["Earth"]
}
