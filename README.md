# 3d-trilateration: a library for three-dimensional space trilateration [![Workflow status](https://github.com/dark/3d-trilateration/workflows/Go/badge.svg)](https://github.com/dark/3d-trilateration/actions?query=workflow%3AGo) [![Build Status](https://travis-ci.org/dark/3d-trilateration.svg?branch=master)](https://travis-ci.org/dark/3d-trilateration)

This library provides utilities to perform **trilateration** (also
known as **true range multilateration**) in a **three-dimensional
space**. In other words, it determines the location in a 3D space of a
target point, given the distances from "station" points whose
locations are already known. A more thorough explanation is available
at the [Wikipedia page for "True range
multilateration"](https://en.wikipedia.org/wiki/True_range_multilateration).

The implementation uses the [Gauss-Newton
algorithm](https://en.wikipedia.org/wiki/Gauss%E2%80%93Newton_algorithm)
to solve a nonlinear least squares problem. In practice, the algorithm
looks iteratively for estimates of the solution that are closer and
closer to the actual solution. This is better suited to account for
inaccuracies in the measurements of distances and locations.
