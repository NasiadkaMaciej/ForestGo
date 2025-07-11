package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func SaveAllStats(filename string, allStats []SimulationStats) error {
	os.MkdirAll(filepath.Dir(filename), 0755)

	fileExists := false
	if _, err := os.Stat(filename); err == nil {
		fileExists = true
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExists {
		writer.Write([]string{
			"ForestDensity", "Humidity", "TotalTrees", "BurnedTrees",
			"BurnPercentage", "SimulationSteps", "StrikeLocation",
		})
	}

	for _, stats := range allStats {
		// Calculate burn percentage
		burnPercentage := float64(0)
		if stats.TotalTrees > 0 {
			burnPercentage = float64(stats.BurnedTrees) / float64(stats.TotalTrees) * 100
		}

		// Write data row
		writer.Write([]string{
			strconv.FormatFloat(stats.Parameters["ForestDensity"], 'f', 4, 64),
			strconv.FormatFloat(stats.Parameters["Humidity"], 'f', 4, 64),
			strconv.Itoa(stats.TotalTrees),
			strconv.Itoa(stats.BurnedTrees),
			strconv.FormatFloat(burnPercentage, 'f', 2, 64),
			strconv.Itoa(stats.SimulationSteps),
			fmt.Sprintf("%d,%d", stats.StrikeX, stats.StrikeY),
		})
	}

	return nil
}
