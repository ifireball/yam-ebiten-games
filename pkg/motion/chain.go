package motion

type chain []Motion

func (c chain) Run(out chan<- Step) {
	defer lastStep(out)
	in := make(chan Step)
	defer close(in)

	for _, m := range c {
		go m.Run(in)
		for s := range in {
			if s.Stop {
				break
			}
			out <- s
		}
	}
}

func Chain(m... Motion) chain {
	return chain(m)
}
