/*
This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; version 2.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see <http://www.gnu.org/licenses/>.

Copyright (C) Mohammed Saed, 2022
*/
package boid

import (
	"math/rand"
	"time"
)

type Boid struct {
	ID       int
	Position vector2D
	Velocity vector2D
}

func CreateBoid(bid int, boidMap [][]int) *Boid {
	screenWidth := len(boidMap[0])
	screenHeight := len(boidMap)

	boid := &Boid{
		ID: bid,
		Position: vector2D{
			X: rand.Float64() * float64(screenWidth),
			Y: rand.Float64() * float64(screenHeight),
		},
		Velocity: vector2D{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
	}

	boidMap[int(boid.Position.Y)][int(boid.Position.X)] = boid.ID
	return boid
}

func (boid *Boid) calculateAcceleration() vector2D {
	accel := vector2D{X: 0, Y: 0}
	return accel
}

func (boid *Boid) moveOne(boidMap [][]int) {
	screenWidth := len(boidMap[0])
	screenHeight := len(boidMap)

	// Calculate acceleration
	boid.Velocity = boid.Velocity.add(boid.calculateAcceleration()).limit(-1, 1)

	// Determine next move and flip velocity if at edge of screen
	next := boid.Position.add(boid.Velocity)
	if next.X >= float64(screenWidth) || next.X <= 0 {
		boid.Velocity = vector2D{X: -boid.Velocity.X, Y: boid.Velocity.Y}
		next = boid.Position.add(boid.Velocity)
	}
	if next.Y >= float64(screenHeight) || next.Y <= 0 {
		boid.Velocity = vector2D{X: boid.Velocity.X, Y: -boid.Velocity.Y}
		next = boid.Position.add(boid.Velocity)
	}

	// Reset boID map to new boID position
	boidMap[int(boid.Position.Y)][int(boid.Position.X)] = -1
	boid.Position = next
	boidMap[int(boid.Position.Y)][int(boid.Position.X)] = boid.ID
}

func (boid *Boid) Start(boidMap [][]int) {
	for {
		boid.moveOne(boidMap)
		time.Sleep(5 * time.Millisecond)
	}
}
