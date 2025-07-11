package main

import "math/rand"

type TreeType int

const (
	Pine TreeType = iota
	Oak
	Birch
	Maple
)

type TreeStatus int

const (
	Healthy TreeStatus = iota
	Struck
	Burning
	Burned
)

type Tree struct {
	Type        TreeType
	Age         int
	Status      TreeStatus
	BurningTime int
}

type Forest struct {
	Trees    [][]*Tree
	Width    int
	Height   int
	Humidity float64
}

type SimulationStats struct {
	TotalTrees      int
	BurnedTrees     int
	SimulationSteps int
	StrikeX         int
	StrikeY         int
	Parameters      map[string]float64
}

var (
	ignitionFactors = map[TreeType]float64{
		Pine:  1.4,
		Oak:   0.8,
		Birch: 1.2,
		Maple: 1.0,
	}

	burnTimes = map[TreeType]int{
		Pine:  5,
		Oak:   4,
		Birch: 3,
		Maple: 4,
	}
)

func (t *Tree) IgnitionFactor() float64 {
	if factor, ok := ignitionFactors[t.Type]; ok {
		return factor
	}
	return 1.0
}

func (t *Tree) BaseBurnTime() int {
	if time, ok := burnTimes[t.Type]; ok {
		return time
	}
	return 3
}

/*
Better function - makes forests vary slightly more
func CreateForest(width, height int, density, humidity float64) *Forest {
	forest := &Forest{
		Trees:    make([][]*Tree, height),
		Width:    width,
		Height:   height,
		Humidity: humidity,
	}

	for y := 0; y < height; y++ {
		forest.Trees[y] = make([]*Tree, width)
		for x := 0; x < width; x++ {
			if rand.Float64() < density {
				forest.Trees[y][x] = &Tree{
					Type:        TreeType(rand.Intn(4)),
					Age:         rand.Intn(100) + 1,
					Status:      Healthy,
					BurningTime: 0,
				}
			}
		}
	}
	return forest
}
*/

func CreateForest(width, height int, density, humidity float64) *Forest {
	forest := &Forest{
		Trees:    make([][]*Tree, height),
		Width:    width,
		Height:   height,
		Humidity: humidity,
	}

	for y := 0; y < height; y++ {
		forest.Trees[y] = make([]*Tree, width)
	}

	totalCells := width * height
	numberOfTrees := int(float64(totalCells) * density)
	treesPlaced := 0

	for treesPlaced < numberOfTrees {
		x := rand.Intn(width)
		y := rand.Intn(height)

		if forest.Trees[y][x] == nil {
			forest.Trees[y][x] = &Tree{
				Type:        TreeType(rand.Intn(4)),
				Age:         rand.Intn(100) + 1,
				Status:      Healthy,
				BurningTime: 0,
			}
			treesPlaced++
		}
	}

	return forest
}
