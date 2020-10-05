package api

import "fmt"

type Ocean interface {
	PutLand(x, y int) error
	Islands() int
}

type defaultOcean struct {
	coordinates [][]bool
	size        int
}

// NewOcean creates a new ocean.
func NewOcean(size int) (Ocean, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size %q must be greater than 0", size)
	}

	c := make([][]bool, size)

	for y := range c {
		c[y] = make([]bool, size)
	}

	return &defaultOcean{
		coordinates: c,
		size:        size,
	}, nil
}

// PutLand puts land at specific coordinates.
func (d *defaultOcean) PutLand(x, y int) error {
	if x >= d.size {
		return fmt.Errorf("x must be less than %d", d.size)
	}

	if y >= d.size {
		return fmt.Errorf("y must be less than %d", d.size)
	}

	if d.coordinates[y] == nil {
		d.coordinates[y] = make([]bool, d.size)
	}

	d.coordinates[y][x] = true

	return nil
}

// Islands returns the amount of islands.
func (d *defaultOcean) Islands() int {
	islands := 0
	visited := make([][]bool, d.size)

	for y := range visited {
		visited[y] = make([]bool, d.size)
	}

	for y := range d.coordinates {
		for x, land := range d.coordinates[y] {
			if !land {
				continue
			}

			if !visited[y][x] {
				islands++

				d.walk(visited, x, y)
			}
		}
	}

	return islands
}

func (d *defaultOcean) walk(visited [][]bool, x, y int) {
	var (
		left   = x - 1
		right  = x + 1
		top    = y - 1
		bottom = y + 1
	)

	visited[y][x] = true

	// left
	if left > 0 && d.coordinates[y][left] && !visited[y][left] {
		d.walk(visited, left, y)
	}

	// right
	if right < d.size && d.coordinates[y][right] && !visited[y][right] {
		d.walk(visited, right, y)
	}

	// top
	if top > 0 && d.coordinates[top][x] && !visited[top][x] {
		d.walk(visited, x, top)
	}

	// bottom
	if bottom < d.size && d.coordinates[bottom][x] && !visited[bottom][x] {
		d.walk(visited, x, bottom)
	}
}
