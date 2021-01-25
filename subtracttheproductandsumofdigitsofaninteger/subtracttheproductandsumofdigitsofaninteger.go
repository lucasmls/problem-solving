package subtracttheproductandsumofdigitsofaninteger

/**
 * Given an integer number n, return the difference between the product of its digits and the sum of its digits.
 *
 * Input: n = 234
 * Output: 15
 *
 * Input: n = 4421
 * Output: 21
 *
 * @link https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
 */

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		mod := n % 10

		product *= mod
		sum += mod

		n /= 10
	}

	return product - sum
}
