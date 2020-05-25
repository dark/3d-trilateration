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

package main

import (
	"fmt"
	"github.com/dark/3d-trilateration/trilateration"
)

func main() {
	g := trilateration.Point{X: 2, Y: -3, Z: 9}

	obs := []trilateration.Range{
		{
			trilateration.Point{
				X: -9529.96875,
				Y: -41.71875,
				Z: -10613.03125,
			},
			14263.89,
		},
		{
			trilateration.Point{
				X: -9570.0625,
				Y: -60.28125,
				Z: -10585.375,
			},
			14270.25,
		},
		{
			trilateration.Point{
				X: -9617.125,
				Y: -76.59375,
				Z: -10570.6875,
			},
			14291.06,
		},
		{
			trilateration.Point{
				X: -9662.1875,
				Y: -80.34375,
				Z: -10544.375,
			},
			14302.03,
		},
		{
			trilateration.Point{
				X: -9674.40625,
				Y: -81.4375,
				Z: -10528.3437,
			},
			14298.49,
		},
		{
			trilateration.Point{
				X: -9780.125,
				Y: 9.75,
				Z: -10414.21875,
			},
			14286.60,
		},
	}

	fmt.Println("Initial guess:", g)
	fmt.Println("Sum of squares:", trilateration.SumOfResidualSquares(obs, g))
	for i := 0; i < 100; i += 1 {
		fmt.Println("\nIteration:", i)
		g = trilateration.GaussNetwonIteration(obs, g)
		fmt.Println("   New guess:", g)
		fmt.Println("   Sum of squares:", trilateration.SumOfResidualSquares(obs, g))
	}
}
