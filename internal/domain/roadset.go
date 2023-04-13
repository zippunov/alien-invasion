package domain

type void struct{}
type RoadSet map[road]void

func (cs RoadSet) add(r road) {
	cs[r] = void{}
}

func (cs RoadSet) remove(r road) {
	delete(cs, r)
}
