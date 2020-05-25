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

// Test converge with a specific set of observation and a static
// guess.
func TestConvergeWithStaticGuess(t *testing.T) {
	guess := Point{X: 20000, Y: -30000, Z: 90000}
	observations := []Range{
		{
			Point{
				X: -9529.96875,
				Y: -41.71875,
				Z: -10613.03125,
			},
			14263.89,
		},
		{
			Point{
				X: -9570.0625,
				Y: -60.28125,
				Z: -10585.375,
			},
			14270.25,
		},
		{
			Point{
				X: -9617.125,
				Y: -76.59375,
				Z: -10570.6875,
			},
			14291.06,
		},
		{
			Point{
				X: -9662.1875,
				Y: -80.34375,
				Z: -10544.375,
			},
			14302.03,
		},
		{
			Point{
				X: -9674.40625,
				Y: -81.4375,
				Z: -10528.3437,
			},
			14298.49,
		},
		{
			Point{
				X: -9780.125,
				Y: 9.75,
				Z: -10414.21875,
			},
			14286.60,
		},
	}

	fmt.Println("Initial guess:", guess)
	fmt.Println("Sum of squares:", SumOfResidualSquares(observations, guess))
	for i := 0; i < 20; i += 1 {
		fmt.Println("\nIteration:", i)
		guess = GaussNetwonIteration(observations, guess)
		fmt.Println("   New guess:", guess)
		fmt.Println("   Sum of squares:", SumOfResidualSquares(observations, guess))
	}

	final_sum_of_squares := SumOfResidualSquares(observations, guess)
	if final_sum_of_squares > 1 {
		t.Errorf("Final sum of squares is too high: %f", final_sum_of_squares)
	}
}
