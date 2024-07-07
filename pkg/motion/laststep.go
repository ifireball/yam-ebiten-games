package motion

func lastStep(out chan<- Step) {
	out <- Step{Stop: true}
}
