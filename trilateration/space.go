/*
 *  A library for three-dimensional space trilateration
 *  Copyright (C) 2020  Marco Leogrande
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package trilateration

import (
	"math"
	"math/rand"
	"time"
)

// Cartesian coordinates of a Point in a 3D space.
type Point struct {
	X float64
	Y float64
	Z float64
}

// Representation of a range, as distance from a given starting point
// (called 'station').
type Range struct {
	Station  Point
	Distance float64
}

// Computes the distance between two points in a 3D space.
func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2) + math.Pow(p1.Z-p2.Z, 2))
}

// Select a random point from the cubic space identified by these properties:
// * its center is the provided Point
// * its edges are parallel to the 3d coordinate axes;
// * its edges are 2*distance units long.
//
// Please note that this function re-seeds the global random pool at
// each invocation.
func SelectRandomPointInCubicRange(center Point, distance float64) Point {
	rand.Seed(time.Now().Unix())
	return Point {
		X: center.X + (rand.Float64() - 0.5) * 2 * distance,
		Y: center.Y + (rand.Float64() - 0.5) * 2 * distance,
		Z: center.Z + (rand.Float64() - 0.5) * 2 * distance,
	}
}
