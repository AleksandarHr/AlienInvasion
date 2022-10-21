package structs

type Direction int64

const (
	North Direction = iota
	East
	South
	West
)

func (dir Direction) String() string {
	switch dir {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return "Invalid Direction"
}
