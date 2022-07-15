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
	Position Vector2D
	Velocity Vector2D
}

func CreateBoid(bid int, screenWidth, screenHeight float64) *Boid {
	boid := &Boid{
		ID: bid,
		Position: Vector2D{
			X: rand.Float64() * screenWidth,
			Y: rand.Float64() * screenHeight,
		},
		Velocity: Vector2D{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
	}
	return boid
}

func (b *Boid) moveOnePixel(screenWidth, screenHeight float64) {
	next := b.Position.Add(b.Velocity)
	b.Position = next

	if next.X >= screenWidth || next.X <= 0 {
		b.Velocity = Vector2D{X: -b.Velocity.X, Y: b.Velocity.Y}
	}
	if next.Y >= screenHeight || next.Y <= 0 {
		b.Velocity = Vector2D{X: b.Velocity.X, Y: -b.Velocity.Y}
	}
}

func (b *Boid) Start(screenWidth, screenHeight float64) {
	for {
		b.moveOnePixel(screenWidth, screenHeight)
		time.Sleep(5 * time.Millisecond)
	}
}
