package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellSize   = 10  // Size of each tree cell in pixels
	frameDelay = 200 // Milliseconds between animation frames
)

var (
	// Colors for different tree types
	pineColor  = sdl.Color{R: 0, G: 120, B: 50, A: 255}    // Darker green
	oakColor   = sdl.Color{R: 50, G: 150, B: 50, A: 255}   // Medium green
	birchColor = sdl.Color{R: 150, G: 200, B: 100, A: 255} // Light green
	mapleColor = sdl.Color{R: 100, G: 180, B: 70, A: 255}  // Yellow-green

	// Status colors
	struckColor  = sdl.Color{R: 255, G: 255, B: 0, A: 255} // Yellow for lightning strike
	burningColor = sdl.Color{R: 255, G: 100, B: 0, A: 255}
	burnedColor  = sdl.Color{R: 100, G: 100, B: 100, A: 255}

	// Window and renderer
	window   *sdl.Window
	renderer *sdl.Renderer
)

func InitSDL(forest *Forest, title string) error {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}

	width := forest.Width * cellSize
	height := forest.Height * cellSize

	var err error
	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		return err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		sdl.Quit()
		return err
	}

	return nil
}

func CleanupSDL() {
	if renderer != nil {
		renderer.Destroy()
	}
	if window != nil {
		window.Destroy()
	}
	sdl.Quit()
}

func getTreeColor(tree *Tree) sdl.Color {
	switch tree.Status {
	case Burning:
		return burningColor
	case Burned:
		return burnedColor
	case Struck:
		return struckColor
	case Healthy:
		var color sdl.Color
		switch tree.Type {
		case Pine:
			color = pineColor
		case Oak:
			color = oakColor
		case Birch:
			color = birchColor
		case Maple:
			color = mapleColor
		}

		return color
	}
	return sdl.Color{R: 255, G: 0, B: 255, A: 255} // Default purple (should not occur)
}

func DrawForest(forest *Forest) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	// Draw each cell in the grid
	for y := 0; y < forest.Height; y++ {
		for x := 0; x < forest.Width; x++ {
			if forest.Trees[y][x] != nil {
				tree := forest.Trees[y][x]
				color := getTreeColor(tree)

				// Draw the tree as a colored rectangle
				renderer.SetDrawColor(color.R, color.G, color.B, color.A)
				rect := sdl.Rect{
					X: int32(x * cellSize),
					Y: int32(y * cellSize),
					W: int32(cellSize),
					H: int32(cellSize),
				}
				renderer.FillRect(&rect)
			}
		}
	}

	// Present the renderer
	renderer.Present()
}

func RenderSimulation(forest *Forest, stepNumber int, title string) bool {
	window.SetTitle(fmt.Sprintf("%s - Step %d", title, stepNumber))
	DrawForest(forest)
	sdl.Delay(frameDelay)
	return true
}
