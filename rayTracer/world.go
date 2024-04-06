package rayTracer

type World struct {
	Elems []Hitable
}

func (w *World) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAny := false
	closest := tMax
	record := HitRecord{}

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
