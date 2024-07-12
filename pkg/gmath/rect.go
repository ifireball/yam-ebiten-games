package gmath

type Rect struct {
	TopLeft, BottomRight Vec2
}

func (r *Rect) Top() float64    { return r.TopLeft.Y }
func (r *Rect) Bottom() float64 { return r.BottomRight.Y }
func (r *Rect) Left() float64   { return r.TopLeft.X }
func (r *Rect) Right() float64  { return r.BottomRight.X }

func (r *Rect) OutSide(or *Rect) bool {
	return r.Top() > or.Bottom() ||
		r.Bottom() < or.Top() ||
		r.Left() > or.Right() ||
		r.Right() < or.Left()
}

func (r *Rect) Overlap(or *Rect) bool {
	return !r.OutSide(or)
}

func (r *Rect) Zero() {
	r.TopLeft.Zero()
	r.BottomRight.Zero()
}
