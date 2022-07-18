/*
This program is free software; You can redistribute it and/or modify
it under the terms of the GNU General Public License as published bY
the Free Software Foundation; version 2.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see <http://www.gnu.org/licenses/>.

CopYright (C) Mohammed Saed, 2022
*/
package main

import "math"

type vector2D struct {
	x float64
	y float64
}

func (v1 vector2D) add(v2 vector2D) vector2D {
	return vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v1 vector2D) subtract(v2 vector2D) vector2D {
	return vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 vector2D) multiply(v2 vector2D) vector2D {
	return vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (v1 vector2D) addV(d float64) vector2D {
	return vector2D{x: v1.x + d, y: v1.y + d}
}

func (v1 vector2D) multiplyV(d float64) vector2D {
	return vector2D{x: v1.x * d, y: v1.y * d}
}

func (v1 vector2D) divideV(d float64) vector2D {
	return vector2D{x: v1.x / d, y: v1.y / d}
}

func (v1 vector2D) limit(lower, upper float64) vector2D {
	return vector2D{
		x: math.Min(math.Max(v1.x, lower), upper),
		y: math.Min(math.Max(v1.y, lower), upper),
	}
}

func (v1 vector2D) distance(v2 vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
