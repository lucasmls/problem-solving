package fizzbuzz

import "fmt"

/**
 * Write a program that console logs the numbers
 * from 1 to n. But for multiples of three print
 * “fizz” instead of the number and for the multiples
 * of five print “buzz”. For numbers which are multiples
 * of both three and five print “fizzbuzz”.

 *   fizzBuzz(5);
 *   1
 *   2
 *   fizz
 *   4
 *   buzz
 */

// FizzBuzz ...
func FizzBuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}
