package motion

type Pause int

func (p Pause) Run(out chan<- Step) {
	defer lastStep(out)
	var step Step
	for i := 0; i < int(p); i++ {
		out <- step
	}
}
