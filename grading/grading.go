package grading

/**
 * HackerLand University has the following grading policy:
 * Every student receives a grade in the inclusive range from 0 to 100.
 * 	- Any grade less than 40 is a failing grade.
 * Sam is a professor at the university and likes to round each student's grade according to these rules:
 * 	- If the difference between the grade and the next multiple of 5 is less than 3, round grade up to the next multiple of 5.
 * 	- If the value of grade is less than 38, no rounding occurs as the result will still be a failing grade.
 *
 * @link https://www.hackerrank.com/challenges/grading/problem
 */

// GradingStudents ...
func GradingStudents(grades []int32) []int32 {
	result := []int32{}

	for i := 0; i < len(grades); i++ {
		grade := grades[i]

		if grade < 38 {
			result = append(result, grade)
			continue
		}

		mod := grade % 5
		if mod < 3 {
			result = append(result, grade)
			continue
		}

		result = append(result, grade+(5-mod))
	}

	return result
}
