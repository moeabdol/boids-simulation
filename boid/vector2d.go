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

type Vector2D struct {
	X float64
	Y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vector2D) MultiplY(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{X: v1.X + d, Y: v1.Y + d}
}

func (v1 Vector2D) MultiplYV(d float64) Vector2D {
	return Vector2D{X: v1.X * d, Y: v1.Y * d}
}

func (v1 Vector2D) DivideV(d float64) Vector2D {
	return Vector2D{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
