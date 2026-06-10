package darts

func Score(x, y float64) int {
	radiusSquared := x*x + y*y

	switch {
	case radiusSquared <= 1:
		return 10
	case radiusSquared <= 25:
		return 5
	case radiusSquared <= 100:
		return 1
	default:
		return 0
	}
}