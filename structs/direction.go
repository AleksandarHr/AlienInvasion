package structs

type Direction int64

const (
	North Direction = iota
	East
	South
	West
	Invalid
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

func (dir Direction) OppositeDirection() Direction {
	switch dir {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}
	return Invalid
}
