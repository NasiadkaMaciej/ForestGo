# ForestGo: Forest Fire Simulation

ForestGo is a Go program that simulates the spread of fire in a randomly generated forest, considering environmental factors and tree-specific parameters.

| Fire progress | Lightning strike |
| ------------------------------------------------- | ------------------------------------------------- |
| ![](https://nasiadka.pl/project/forest-go/fire.png) | ![](https://nasiadka.pl/project/forest-go/lightning.png) |

## Features

- Random forest generation with configurable density and humidity
- Four tree types (Pine, Oak, Birch, Maple) with different burning properties
- Lightning strike simulation with realistic fire propagation
- Real-time visualization using SDL
- Statistics collection and CSV output for analysis

## Building

```sh
go build .
```

## Running the Simulation

```sh
./ForestGo [flags]
```

### Available Flags

- `-sdl` - Display the forest in an SDL window (without this flag, runs in batch mode)
- `-width` (default: 120) - Width of the forest grid
- `-height` (default: 80) - Height of the forest grid
- `-density` (default: varies 0.1-1.0) - Run with specific forest density (0.0-1.0)
- `-humidity` (default: varies 0.0-1.0) - Run with specific humidity value (0.0-1.0)

### Examples

```sh
# Run batch simulations with default parameters
./ForestGo
# It will run simulation for all possible parameters, 10 times each

# Run with visualization
./ForestGo -sdl
# It will run random parameter simulation till program is topped

# Run at specific density and humidity
./ForestGo -density 0.21 -humidity 0.37
# It will run this simulation 10 times

# Visualize a specific scenario
./ForestGo -sdl -width 100 -height 80 -density 0.6 -humidity 0.2
# It will run this specific scenario till program is stopped
```
