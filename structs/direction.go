package structs

type Direction int64

// enum to represent the four directions for neighbour city links
const (
	North Direction = iota
	East
	South
	West
	Invalid
)

// String returns a representation of the given direction
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

// StringToDirection returns a direction provided a relevant string
func StringToDirection(dir string) Direction {
	switch dir {
	case "north":
		return North
	case "east":
		return East
	case "south":
		return South
	case "west":
		return West
	}
	return Invalid
}
