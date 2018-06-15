package robotname

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	return "robot"
}

func (r *Robot) Reset() {
}
