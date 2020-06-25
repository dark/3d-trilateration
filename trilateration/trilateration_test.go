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

// A set of stations to be used for testing.
var s1 = Point{
	X: -9529.96875,
	Y: -41.71875,
	Z: -10613.03125,
}
var s2 = Point{
	X: -9570.0625,
	Y: -60.28125,
	Z: -10585.375,
}
var s3 = Point{
	X: -9617.125,
	Y: -76.59375,
	Z: -10570.6875,
}
var s4 = Point{
	X: -9662.1875,
	Y: -80.34375,
	Z: -10544.375,
}
var s5 = Point{
	X: -9674.40625,
	Y: -81.4375,
	Z: -10528.3437,
}
var s6 = Point{
	X: -9780.125,
	Y: 9.75,
	Z: -10414.21875,
}
var s7 = Point{
	X: -9898.96875,
	Y: 49.09375,
	Z: -10440.8125,
}

// An arbitrary set of observations that converges close to the
// origin.
var observations_set_origin = []Range{
	{s1, 14263.89},
	{s2, 14270.25},
	{s3, 14291.06},
	{s4, 14302.03},
	{s5, 14298.49},
	{s6, 14286.60},
	{s7, 14387.58},
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
	solution, err := Trilaterate(observations_set_origin, initial_guess,
		100 /*max_iterations*/, 1 /*min_sum_of_residual_squares*/)
	if err != nil {
		t.Errorf("Error should be nil, found instead: %s", err)
	}
	final_sum_of_squares := SumOfResidualSquares(observations_set_origin, solution)
	fmt.Println("Solution:", solution)
	fmt.Println("Sum of squares of the residuals:", final_sum_of_squares)
	if final_sum_of_squares > 1 {
		t.Errorf("Final sum of squares is too high: %f", final_sum_of_squares)
	}
}

// Store all observations and expected results in a single variable to
// be used by all following tests.
type testObservation struct {
	measurements      []Range
	expected_solution Point
}

var test_observations = []testObservation{
	{observations_set_origin, Point{0, 0, 0}},
	{
		measurements: []Range{
			{s1, 30433.55},
			{s2, 30405.41},
			{s3, 30390.37},
			{s4, 30364.13},
			{s5, 30348.13},
			{s6, 30237.38},
			{s7, 30266.39},
		},
		expected_solution: Point{-9530.5, -910.28125, 19808.125},
	},
	{
		measurements: []Range{
			{s1, 14263.78},
			{s2, 14270.17},
			{s3, 14291.01},
			{s4, 14302.01},
			{s5, 14298.48},
			{s6, 14286.69},
			{s7, 14387.72},
		},
		expected_solution: Point{6.25, -1.28125, -5.75},
	},
	{
		measurements: []Range{
			{s1, 3184.68},
			{s2, 3157.20},
			{s3, 3143.50},
			{s4, 3118.80},
			{s5, 3103.34},
			{s6, 2997.21},
			{s7, 3037.08},
		},
		expected_solution: Point{-9529.437, -64.5, -7428.4375},
	},
	{
		measurements: []Range{
			{s1, 3911.65},
			{s2, 3874.52},
			{s3, 3832.39},
			{s4, 3787.03},
			{s5, 3773.53},
			{s6, 3632.56},
			{s7, 3511.50},
		},
		expected_solution: Point{-13243.15625, 1026.5625, -10003.09375},
	},
}

// Parametric function to test one observation, given some parameters.
func testOneObservation(t *testing.T, observation testObservation, initial_guess Point, min_sum_of_residual_squares float64) {
	// Solve
	solution, err := Trilaterate(observation.measurements, initial_guess,
		100 /*max_iterations*/, min_sum_of_residual_squares)
	if err != nil {
		t.Errorf("Error should be nil, found instead: %s", err)
	}

	fmt.Println("  Actual solution:", solution)

	// Check the solution
	distance := Distance(observation.expected_solution, solution)
	fmt.Println("  Distance from expected solution:", distance)
	if distance > 1.0 {
		t.Errorf("Distance from expected solution is too high")
	}

	// Check the sum of squares of the residuals
	final_sum_of_squares := SumOfResidualSquares(observation.measurements, observation.expected_solution)
	fmt.Println("  Sum of squares of the residuals:", final_sum_of_squares)
	if final_sum_of_squares > min_sum_of_residual_squares {
		t.Errorf("Final sum of squares is too high")
	}
}

// Test all test observations listed above, using a random point in
// cubic range from the first station as initial guess.
func TestAllObservationsInitialGuessInCubicRange(t *testing.T) {
	for _, observation := range test_observations {
		fmt.Println("Testing observations that should converge at:", observation.expected_solution)

		random_guess := SelectRandomPointInCubicRange(observation.measurements[0].Station, observation.measurements[0].Distance)
		fmt.Println("  Using initial random guess in cubic range:", random_guess)

		testOneObservation(t, observation, random_guess, 1 /*min_sum_of_residual_squares*/)
	}
}

// Test all test observations listed above, using a random point on a
// sphere identified by the first measurement as initial guess.
func TestAllObservationsInitialGuessOnSphere(t *testing.T) {
	for _, observation := range test_observations {
		fmt.Println("Testing observations that should converge at:", observation.expected_solution)

		random_guess := SelectRandomPointOnSphere(observation.measurements[0].Station, observation.measurements[0].Distance)
		fmt.Println("  Using initial random guess over sphere:", random_guess)

		testOneObservation(t, observation, random_guess, 1 /*min_sum_of_residual_squares*/)
	}
}
