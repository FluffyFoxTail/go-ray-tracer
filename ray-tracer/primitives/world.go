package primitives

type World struct {
	Elems []Hittable
}

func (w *World) Add(h Hittable) {
	w.Elems = append(w.Elems, h)
}

func (w *World) AddAll(hittables ...Hittable) {
	for _, h := range hittables {
		w.Add(h)
	}
}

func (w *World) Count() int {
	return len(w.Elems)
}

func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	hitAny := false
	closest := tMax
	record := Hit{}

	for _, elem := range w.Elems {
		isHit, tempRecord := elem.Hit(r, tMin, closest)

		if isHit {
			hitAny = true
			closest = tempRecord.T
			record = tempRecord
		}
	}
	return hitAny, record
}
