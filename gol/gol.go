package main

const ALIVE = 0xFF
const DEAD = 0x00

// This function creates a new world of the same size of the current world.
func createNewWorld(p golParams) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for i := 0; i < p.imageHeight; i++ {
		newWorld[i] = make([]byte, p.imageWidth)
	}
	return newWorld
}

// This function when given the position of a cell and the world, returns the number of alive neighbors.
func countAliveNeighbors(p golParams, world [][]byte, x int, y int) int {
	aliveNeighbors := 0
	// Iterate over the 8 neighbors of the cell.
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// Skip the cell itself.
			if dx == 0 && dy == 0 {
				continue
			}
			// Calculate the position of the neighbor.
			nx, ny := x+dx, y+dy

			// If the neighbor is out of bounds, wrap around.
			if nx < 0 {
				nx = p.imageWidth - 1
			} else if nx >= p.imageWidth {
				nx = 0
			}

			// If the neighbor is out of bounds, wrap around.
			if ny < 0 {
				ny = p.imageHeight - 1
			} else if ny >= p.imageHeight {
				ny = 0
			}

			// If the neighbor is alive, increment the counter.
			if world[ny][nx] == ALIVE {
				aliveNeighbors++
			}
		}
	}
	return aliveNeighbors
}

// This function when given the world, returns the next state of the world.
func calculateNextState(p golParams, world [][]byte) [][]byte {
	nextWorld := createNewWorld(p)
	// Iterate over the cells of the world.
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			neighbours := countAliveNeighbors(p, world, x, y)
			// Apply the rules of the game.
			switch world[y][x] {
			case ALIVE:
				// If the cell is alive and has 2 or 3 neighbors, it stays alive.
				if neighbours == 2 || neighbours == 3 {
					nextWorld[y][x] = ALIVE
					// Otherwise, it dies.
				} else {
					nextWorld[y][x] = DEAD
				}
			case DEAD:
				// If the cell is dead and has 3 neighbors, it becomes alive.
				if neighbours == 3 {
					nextWorld[y][x] = ALIVE
					// Otherwise, it stays dead.
				} else {
					nextWorld[y][x] = DEAD
				}
			}
		}
	}
	return nextWorld
}

// This function when given the world, returns a slice of all the positions of the alive cells.
func calculateAliveCells(p golParams, world [][]byte) []cell {
	// Create a slice to store the alive cells.
	var aliveCells []cell
	// Iterate over the cells of the world.
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == ALIVE {
				aliveCells = append(aliveCells, cell{x, y})
			}
		}
	}
	return aliveCells
}
