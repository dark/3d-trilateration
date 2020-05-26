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

// An arbitrary set of observations that converges close to the
// origin.
var observations_set_origin = []Range{
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

// Test manual iterations over the low-level function, with a specific
// observations set and a static guess.
func TestIterationsWithStaticGuess(t *testing.T) {
	guess := Point{X: 20000, Y: -30000, Z: 90000}

	fmt.Println("Initial guess:", guess)
	fmt.Println("Sum of squares:", SumOfResidualSquares(observations_set_origin, guess))
	for i := 0; i < 20; i += 1 {
		fmt.Println("Iteration:", i)
		guess = GaussNetwonIteration(observations_set_origin, guess)
		fmt.Println("   New guess:", guess)
		fmt.Println("   Sum of squares of the residuals:",
			SumOfResidualSquares(observations_set_origin, guess))
	}

	final_sum_of_squares := SumOfResidualSquares(observations_set_origin, guess)
	if final_sum_of_squares > 1 {
		t.Errorf("Final sum of squares is too high: %f", final_sum_of_squares)
	}
}

// Test the high-level function to iterate internally based on the
// provided parameters, including the specific observations set and a
// static guess.
func TestConvergenceWithStaticGuess(t *testing.T) {
	initial_guess := Point{X: 20000, Y: -30000, Z: 90000}
	solution := Trilaterate(observations_set_origin, initial_guess,
		100 /*max_iterations*/, 1 /*min_sum_of_residual_squares*/)
	final_sum_of_squares := SumOfResidualSquares(observations_set_origin, solution)
	fmt.Println("Solution:", solution)
	fmt.Println("Sum of squares of the residuals:", SumOfResidualSquares(observations_set_origin, solution))
	if final_sum_of_squares > 1 {
		t.Errorf("Final sum of squares is too high: %f", final_sum_of_squares)
	}
}
