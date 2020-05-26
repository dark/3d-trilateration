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
	"fmt"
	"testing"
)

// Parametrized function to test the random selection of a point in a
// cubic range around a center point.
func testCubicRange(t *testing.T, center Point, distance float64) {
	fmt.Print("Testing cubic range around point:", center)
	fmt.Println(" with distance:", distance)
	random_point := SelectRandomPointInCubicRange(center, distance)
	fmt.Println("Selected random point:", random_point)

	if random_point.X > center.X+distance || random_point.X < center.X-distance {
		t.Errorf("X-coordinate of random point is out of range")
	}
	if random_point.Y > center.Y+distance || random_point.Y < center.Y-distance {
		t.Errorf("Y-coordinate of random point is out of range")
	}
	if random_point.Z > center.Z+distance || random_point.Z < center.Z-distance {
		t.Errorf("Z-coordinate of random point is out of range")
	}
}

func TestSmallCubicRangeAroundOrigin(t *testing.T) {
	testCubicRange(t, Point{}, 10)
}

func TestLargeCubicRangeAroundOrigin(t *testing.T) {
	testCubicRange(t, Point{}, 1000)
}

func TestHugeCubicRangeAroundOrigin(t *testing.T) {
	testCubicRange(t, Point{}, 1000000)
}

func TestSmallCubicRangeAroundPoint(t *testing.T) {
	testCubicRange(t, Point{X: 123, Y: -456, Z: 7890}, 10)
}

func TestLargeCubicRangeAroundPoint(t *testing.T) {
	testCubicRange(t, Point{X: 123, Y: -456, Z: 7890}, 1000)
}

func TestHugeCubicRangeAroundPoint(t *testing.T) {
	testCubicRange(t, Point{X: 123, Y: -456, Z: 7890}, 1000000)
}

// Parametrized function to test the random selection of a point on a
// sphere.
func testPointOnSphere(t *testing.T, center Point, radius float64) {
	fmt.Print("Testing sphere around point:", center)
	fmt.Println(" with radius:", radius)
	random_point := SelectRandomPointOnSphere(center, radius)
	fmt.Println("Selected random point:", random_point)
	distance := Distance(center, random_point)
	fmt.Println("Distance of the random point from the center:", distance)

	const precision_factor = 0.0001
	if distance > radius*(1.+precision_factor) || distance < radius*(1.-precision_factor) {
		t.Errorf("The distance is outside of the permissible range")
	}
}

func TestSmallSphereAroundOrigin(t *testing.T) {
	testPointOnSphere(t, Point{}, 10)
}

func TestLargeSphereAroundOrigin(t *testing.T) {
	testPointOnSphere(t, Point{}, 1000)
}

func TestHugeSphereAroundOrigin(t *testing.T) {
	testPointOnSphere(t, Point{}, 1000000)
}

func TestSmallSphereAroundPoint(t *testing.T) {
	testPointOnSphere(t, Point{X: 123, Y: -456, Z: 7890}, 10)
}

func TestLargeSphereAroundPoint(t *testing.T) {
	testPointOnSphere(t, Point{X: 123, Y: -456, Z: 7890}, 1000)
}

func TestHugeSphereAroundPoint(t *testing.T) {
	testPointOnSphere(t, Point{X: 123, Y: -456, Z: 7890}, 1000000)
}
