package main

import (
	"math/rand"
)

type Coord struct {
	X, Y int
}

func (f *Forest) StrikeLightning() (x, y int) {
	// Try to find a healthy tree by random sampling
	attempts := 0
	maxAttempts := f.Width * f.Height

	for attempts < maxAttempts {
		x = rand.Intn(f.Width)
		y = rand.Intn(f.Height)

		if f.Trees[y][x] != nil && f.Trees[y][x].Status == Healthy {
			f.Trees[y][x].Status = Struck
			return x, y
		}

		attempts++
	}

	// If no healthy tree found, just return random coordinates
	return rand.Intn(f.Width), rand.Intn(f.Height)
}

// IgniteTree ignites a tree at the given location
func (f *Forest) IgniteTree(x, y int) bool {
	if x < 0 || y < 0 || x >= f.Width || y >= f.Height ||
		f.Trees[y][x] == nil ||
		(f.Trees[y][x].Status != Healthy && f.Trees[y][x].Status != Struck) {
		return false
	}

	tree := f.Trees[y][x]
	fireProbability := calculateIgnitionProbability(tree, f.Humidity)

	if fireProbability > rand.Float64() {
		tree.Status = Burning
		tree.BurningTime = calculateBurnTime(tree, f.Humidity)
		return true
	}

	return false
}

func (f *Forest) SimulateStep() bool {
	// Track new fires that will start in this step
	newFires := make(map[Coord]bool)
	fireActive := false

	// Process burning trees
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			if f.Trees[y][x] != nil && f.Trees[y][x].Status == Burning {
				fireActive = true

				// Try to spread fire to neighbors
				spreadFireToNeighbors(f, x, y, newFires)

				// Update burning time and status
				burnTree := f.Trees[y][x]
				burnTree.BurningTime--
				if burnTree.BurningTime <= 0 {
					burnTree.Status = Burned
				}
			}
		}
	}

	// Apply new fires
	for coords := range newFires {
		x, y := coords.X, coords.Y
		if f.Trees[y][x] != nil && f.Trees[y][x].Status == Healthy {
			f.Trees[y][x].Status = Burning
			f.Trees[y][x].BurningTime = calculateBurnTime(f.Trees[y][x], f.Humidity)
		}
	}

	return fireActive || len(newFires) > 0
}

// Helper function to calculate ignition probability
func calculateIgnitionProbability(tree *Tree, humidity float64) float64 {
	baseProbability := 0.8 - (humidity * 0.5)            // Humidity factor
	typeFactor := tree.IgnitionFactor()                  // Type-specific factor
	ageFactor := 0.8 + (float64(tree.Age) / 100.0 * 0.4) // Older trees are more likely to ignite
	return baseProbability * typeFactor * ageFactor
}

func calculateBurnTime(tree *Tree, humidity float64) int {
	baseBurnTime := tree.BaseBurnTime()
	ageFactor := 1.0 + float64(tree.Age)/100.0
	humidityFactor := 1.0 - (humidity * 0.3)
	burnTime := int(float64(baseBurnTime) * ageFactor * humidityFactor)

	return burnTime
}

// Helper function to spread fire to neighboring trees
func spreadFireToNeighbors(f *Forest, x, y int, newFires map[Coord]bool) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx, ny := x+dx, y+dy

			if nx >= 0 && nx < f.Width && ny >= 0 && ny < f.Height &&
				f.Trees[ny][nx] != nil && f.Trees[ny][nx].Status == Healthy {

				baseProbability := 0.4
				if dx != 0 && dy != 0 {
					baseProbability *= 0.7
				}

				typeMultiplier := f.Trees[y][x].IgnitionFactor()
				ageMultiplier := 0.8 + (float64(f.Trees[ny][nx].Age) / 100.0 * 0.4)
				humidityFactor := 1.0 - (f.Humidity * f.Humidity * 0.9)
				fireProb := baseProbability * typeMultiplier * humidityFactor * ageMultiplier

				if fireProb > rand.Float64() {
					newFires[Coord{nx, ny}] = true
				}
			}
		}
	}
}

// CalculateStats computes statistics for the simulation
func (f *Forest) CalculateStats(strikeX, strikeY int) SimulationStats {
	totalTrees := 0
	burnedTrees := 0

	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			if f.Trees[y][x] != nil {
				totalTrees++
				if f.Trees[y][x].Status == Burned {
					burnedTrees++
				}
			}
		}
	}

	return SimulationStats{
		TotalTrees:  totalTrees,
		BurnedTrees: burnedTrees,
		Parameters: map[string]float64{
			"ForestDensity": float64(totalTrees) / float64(f.Width*f.Height),
			"Humidity":      f.Humidity,
		},
	}
}
