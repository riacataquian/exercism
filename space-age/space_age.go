// package space ...
package space

// Planet ...
type Planet string

// Age ...
func Age(s float64, p Planet) (age float64) {
	switch p {
	case "Mercury":
		age = 280.88
	case "Venus":
		age = 9.78
	case "Mars":
		age = 39.25
	case "Jupiter":
		age = 2.41
	case "Saturn":
		age = 3.23
	case "Uranus":
		age = 1.21
	case "Neptune":
		age = 1.58
	default:
		age = 31.69
	}

	return
}
