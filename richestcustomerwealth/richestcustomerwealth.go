package richestcustomerwealth

/**
 * You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank.
 * Return the wealth that the richest customer has.
 * A customer's wealth is the amount of money they have in all their bank accounts.
 * The richest customer is the customer that has the maximum wealth.
 *
 * Input: accounts = [[1,2,3],[3,2,1]]
 * Output: 6
 *
 * Input: accounts = [[1,5],[7,3],[3,5]]
 * Output: 10
 *
 * @link https://leetcode.com/problems/richest-customer-wealth/
 */

func maximumWealth(accounts [][]int) int {
	result := 0

	for i := 0; i < len(accounts); i++ {
		personBanksMoney := accounts[i]

		personMoney := 0

		for j := 0; j < len(personBanksMoney); j++ {
			bankMoney := accounts[i][j]
			personMoney += bankMoney
		}

		if personMoney > result {
			result = personMoney
		}
	}

	return result
}
