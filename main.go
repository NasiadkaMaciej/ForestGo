package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Parse command-line flags
	sdlDisplayFlag := flag.Bool("sdl", false, "Display forest in SDL window")
	widthFlag := flag.Int("width", 120, "Width of the forest")
	heightFlag := flag.Int("height", 80, "Height of the forest")
	strictDensityFlag := flag.Float64("density", -1.0, "Run simulation with specified density only (0.0-1.0)")
	strictHumidityFlag := flag.Float64("humidity", -1.0, "Run simulation with specified humidity only (0.0-1.0)")
	flag.Parse()

	// Validate density if provided
	if *strictDensityFlag > 1.0 || (*strictDensityFlag < 0.0 && *strictDensityFlag != -1.0) {
		fmt.Println("Error: Density must be between 0.0 and 1.0")
		os.Exit(1)
	}

	// Validate humidity if provided
	if *strictHumidityFlag > 1.0 || (*strictHumidityFlag < 0.0 && *strictHumidityFlag != -1.0) {
		fmt.Println("Error: Humidity must be between 0.0 and 1.0")
		os.Exit(1)
	}

	// Set up simulation environment
	rand.Seed(time.Now().UnixNano())
	simConfig := SimConfig{
		ForestWidth:    *widthFlag,
		ForestHeight:   *heightFlag,
		OutputDir:      "simulation_results",
		UseSDL:         *sdlDisplayFlag,
		Iterations:     10,
		StrictDensity:  *strictDensityFlag >= 0.0,
		DensityValue:   *strictDensityFlag,
		StrictHumidity: *strictHumidityFlag >= 0.0,
		HumidityValue:  *strictHumidityFlag,
	}

	// Create output directory
	os.MkdirAll(simConfig.OutputDir, 0755)

	// Initialize SDL if needed
	if simConfig.UseSDL {
		forest := &Forest{Width: simConfig.ForestWidth, Height: simConfig.ForestHeight}
		if err := InitSDL(forest, "Forest Fire Simulation"); err != nil {
			simConfig.UseSDL = false
		} else {
			defer CleanupSDL()
		}
	}

	// Run simulations with progress tracking
	RunAllSimulations(simConfig)
}

type SimConfig struct {
	ForestWidth    int
	ForestHeight   int
	OutputDir      string
	UseSDL         bool
	Iterations     int
	StrictDensity  bool
	DensityValue   float64
	StrictHumidity bool
	HumidityValue  float64
}

func RunAllSimulations(config SimConfig) {
	statsFile := filepath.Join(config.OutputDir, "forest_fire_stats.csv")

	var allStats []SimulationStats

	if config.UseSDL {
		// By default, run simulations enlessly with random parameters
		// User can specify strict values for density and humidity
		continueRunning := true

		for continueRunning {
			var density float64
			if config.StrictDensity {
				density = config.DensityValue
			} else {
				density = rand.Float64()
			}

			var humidity float64
			if config.StrictHumidity {
				humidity = config.HumidityValue
			} else {
				humidity = rand.Float64()
			}

			forest := CreateForest(config.ForestWidth, config.ForestHeight, density, humidity)

			stats := RunSingleSimulation(forest, config.UseSDL, density, humidity)

			continueRunning = window != nil

			ReportProgress(density, humidity, 0, stats)

			if continueRunning {
				time.Sleep(time.Second)
			}
		}
	} else {
		// By default, run with all parameter combinations
		// User can specify strict values for density and humidity
		densityStart := 0.1
		densityEnd := 1.0
		densityStep := 0.01

		if config.StrictDensity {
			densityStart = config.DensityValue
			densityEnd = config.DensityValue
			densityStep = 1.0 // Only run once
		}

		humidityStart := 0.0
		humidityEnd := 1.0
		humidityStep := 0.01

		if config.StrictHumidity {
			humidityStart = config.HumidityValue
			humidityEnd = config.HumidityValue
			humidityStep = 1.0 // Only run once
		}

		for density := densityStart; density <= densityEnd; density += densityStep {
			for humidity := humidityStart; humidity <= humidityEnd; humidity += humidityStep {
				for iter := 0; iter < config.Iterations; iter++ {
					forest := CreateForest(config.ForestWidth, config.ForestHeight, density, humidity)

					stats := RunSingleSimulation(forest, config.UseSDL, density, humidity)
					ReportProgress(density, humidity, iter, stats)

					allStats = append(allStats, stats)
				}
			}
		}

		// Save all collected stats
		if len(allStats) > 0 {
			SaveAllStats(statsFile, allStats)
		}

		fmt.Printf("\nAll simulations completed\n")
		fmt.Printf("Statistics saved to %s\n", statsFile)
		fmt.Println("Run 'gnuplot charts.gnuplot' to generate charts.")
	}
}

func RunSingleSimulation(forest *Forest, useSDL bool, density, humidity float64) SimulationStats {
	title := fmt.Sprintf("Forest Fire Simulation (Density: %.2f, Humidity: %.2f)", density, humidity)
	continueSimulation := true

	if useSDL && window != nil {
		window.SetTitle(title)
		continueSimulation = RenderSimulation(forest, 0, title+" (Initial State)")
		time.Sleep(500 * time.Millisecond)
	}

	strikeX, strikeY := forest.StrikeLightning()

	if useSDL {
		continueSimulation = RenderSimulation(forest, 0, title+" (Lightning Strike!)")
		time.Sleep(500 * time.Millisecond)
	}

	forest.IgniteTree(strikeX, strikeY)

	steps := 0
	if useSDL {
		continueSimulation = RenderSimulation(forest, steps, title+" (Fire Started)")
		time.Sleep(500 * time.Millisecond)
	}

	for forest.SimulateStep() && continueSimulation {
		steps++
		if useSDL {
			continueSimulation = RenderSimulation(forest, steps, title)
		}
	}

	if useSDL {
		RenderSimulation(forest, steps, title+" (Finished)")
		time.Sleep(1 * time.Second)
	}

	stats := forest.CalculateStats(strikeX, strikeY)
	stats.SimulationSteps = steps
	stats.StrikeX = strikeX
	stats.StrikeY = strikeY

	return stats
}

func ReportProgress(density, humidity float64, iter int, stats SimulationStats) {
	fmt.Printf("Completed simulation: density=%.2f, humidity=%.2f, iter=%d\n",
		density, humidity, iter)
	fmt.Printf("Strike location: (%d,%d), Simulation steps: %d\n",
		stats.StrikeX, stats.StrikeY, stats.SimulationSteps)
	fmt.Printf("Total trees: %d, Burned: %d (%.2f%%)\n",
		stats.TotalTrees, stats.BurnedTrees,
		float64(stats.BurnedTrees)/float64(stats.TotalTrees)*100)
}
