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

// The package 'trilateration' provides utilities to perform
// trilateration (also known as true range multilateration), to
// determine the location in a 3D space of a target point, given the
// distances from "station" points whose location is already known.
package trilateration

import (
	"gonum.org/v1/gonum/mat"
)

// Returns the Jacobian matrix of the residuals function. The matrix
// has size [i x 3], where 'i' is the number of observations.
func residualsJacobian(observations []Range, current_guess Point) *mat.Dense {
	j_r := mat.NewDense(len(observations), 3, nil)
	for i, observation := range observations {
		// Cache the distance from the station to the current
		// guess, since it's a common denominator.
		distance := Distance(observation.Station, current_guess)

		j_r.Set(i, 0, (current_guess.X-observation.Station.X)/distance)
		j_r.Set(i, 1, (current_guess.Y-observation.Station.Y)/distance)
		j_r.Set(i, 2, (current_guess.Z-observation.Station.Z)/distance)
	}

	return j_r
}

// Returns the column vector with the residuals. It has size [i x 1].
func residuals(observations []Range, current_guess Point) *mat.VecDense {
	r := mat.NewVecDense(len(observations), nil)
	for i, observation := range observations {
		distance := Distance(observation.Station, current_guess)
		r.SetVec(i, observation.Distance-distance)
	}

	return r
}

// Implements one iteration of the Gauss–Newton algorithm
// (https://en.wikipedia.org/wiki/Gauss%E2%80%93Newton_algorithm), as
// applied to true range multilateration. Returns the new guess.
func gaussNetwonIteration(observations []Range, current_guess Point) Point {
	// The variable names in this function mirror those from the Wikipedia article.

	// Current guess of the algorithm
	beta_s := mat.NewVecDense(3, nil)
	beta_s.SetVec(0, current_guess.X)
	beta_s.SetVec(1, current_guess.Y)
	beta_s.SetVec(2, current_guess.Z)

	// Jacobian matrix of the residuals, and its transpose.
	j_r := residualsJacobian(observations, current_guess)
	j_r_t := j_r.T()

	// Left pseudoinverse of j_r.
	var product mat.Dense
	product.Mul(j_r_t, j_r)
	var left_pseudoinverse mat.Dense
	left_pseudoinverse.Solve(&product, j_r_t)

	// Matrix of residuals.
	r := residuals(observations, current_guess)

	// Compute and return the new guess.
	var subtrahend mat.VecDense
	subtrahend.MulVec(&left_pseudoinverse, r)
	beta_s.SubVec(beta_s, &subtrahend)

	return Point{X: beta_s.AtVec(0), Y: beta_s.AtVec(1), Z: beta_s.AtVec(2)}
}
