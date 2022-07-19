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
package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position vector2D
	velocity vector2D
}

func CreateBoid(bid int) *Boid {
	boid := &Boid{
		id: bid,
		position: vector2D{
			x: rand.Float64() * float64(screenWidth),
			y: rand.Float64() * float64(screenHeight),
		},
		velocity: vector2D{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
	}

	boidMap[int(boid.position.x)][int(boid.position.y)] = boid.id
	return boid
}

func (boid *Boid) calculateAcceleration() vector2D {
	upper, lower := boid.position.addV(viewRadius), boid.position.addV(-viewRadius)
	avgPosition, avgVelocity := vector2D{x: 0, y: 0}, vector2D{x: 0, y: 0}
	count := 0.0

	lock.Lock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != boid.id {
				if dist := boids[otherBoidID].position.distance(boid.position); dist <= viewRadius {
					count++
					avgPosition = avgPosition.add(boids[otherBoidID].position)
					avgVelocity = avgVelocity.add(boids[otherBoidID].velocity)
				}
			}
		}
	}
	lock.Unlock()

	accel := vector2D{0, 0}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.divideV(count), avgVelocity.divideV(count)
		accelAlignment := avgVelocity.subtract(boid.velocity).multiplyV(adjustRate)
		accelCohesion := avgPosition.subtract(boid.position).multiplyV(adjustRate)
		accel = accel.add(accelAlignment).add(accelCohesion)
	}
	return accel
}

func (boid *Boid) moveOne() {
	// Calculate acceleration
	boid.velocity = boid.velocity.add(boid.calculateAcceleration()).limit(-1, 1)

	lock.Lock()
	// Determine next move and flip velocity if at edge of screen
	next := boid.position.add(boid.velocity)
	if next.x >= float64(screenWidth) || next.x <= 0 {
		boid.velocity = vector2D{x: -boid.velocity.x, y: boid.velocity.y}
		next = boid.position.add(boid.velocity)
	}
	if next.y >= float64(screenHeight) || next.y <= 0 {
		boid.velocity = vector2D{x: boid.velocity.x, y: -boid.velocity.y}
		next = boid.position.add(boid.velocity)
	}

	// Reset boID map to new boID position
	boidMap[int(boid.position.x)][int(boid.position.y)] = -1
	boid.position = next
	boidMap[int(boid.position.x)][int(boid.position.y)] = boid.id
	lock.Unlock()
}

func (boid *Boid) Start() {
	for {
		boid.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}
