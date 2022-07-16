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
package boid

import "math"

type vector2D struct {
	X float64
	Y float64
}

func (v1 vector2D) add(v2 vector2D) vector2D {
	return vector2D{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 vector2D) subtract(v2 vector2D) vector2D {
	return vector2D{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 vector2D) multiply(v2 vector2D) vector2D {
	return vector2D{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 vector2D) addV(d float64) vector2D {
	return vector2D{X: v1.X + d, Y: v1.Y + d}
}

func (v1 vector2D) multiplyV(d float64) vector2D {
	return vector2D{X: v1.X * d, Y: v1.Y * d}
}

func (v1 vector2D) divideV(d float64) vector2D {
	return vector2D{X: v1.X / d, Y: v1.Y / d}
}

func (v1 vector2D) limit(lower, upper float64) vector2D {
	return vector2D{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

func (v1 vector2D) distance(v2 vector2D) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
