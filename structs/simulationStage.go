package structs

type SimulationStage int64

const (
	SimulationStart SimulationStage = iota
	InitializingWorld
	SpawningAliens
	MovingAliens
	SimulationEnd
)
