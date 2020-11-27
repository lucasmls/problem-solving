package carsrace

/**
 * In a race of N cars, numbered between 0 to N-1, we want to determine the arrival
 * ordder of each one of them.
 *
 * For a give race, we have an array A of size N in which the value contained in each
 * position represents the number of the car that ends up immediately ahead the car woth
 * the index of the same position.
 *
 * Generates and array C in which the value contained in each position represents the
 * placing of the car that has the number equals to the index of the same position
 *
 *					0  1  2  3   4  5
 * Input:  [3, 4, 1, 5, -1, 2]
 * Output: [6, 2, 3, 5, 1, 4]
 *
 */

func indexOf(list []int, valueToBeFound int) int {
	for i, value := range list {
		if value == valueToBeFound {
			return i
		}
	}

	return -1
}

// CarsRace ...
func CarsRace(cars []int) []int {
	result := make([]int, len(cars))

	// -1 means that no one ends up the ahead of the cars of number equal to the index of the -1 value
	carNumberToFind := -1
	position := 1
	for position <= len(cars) {
		index := indexOf(cars, carNumberToFind)

		result[index] = position
		position++
		carNumberToFind = index
	}

	return result
}
