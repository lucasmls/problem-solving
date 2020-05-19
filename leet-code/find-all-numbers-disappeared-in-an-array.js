/**
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 * 
 * @link https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
 */

/**
 * @param {number[]} nums
 * @return {number[]}
 */
function findDisappearedNumbers (nums) {
  const numbersMap = {}

  for (const n of nums) {
    numbersMap[n] = true
  }

  let result = []
  for (let i = 1; i < nums.length +1; i++) {
    if (!numbersMap[i]) {
      result.push(i)
    }
  }

  return result
}

const res = findDisappearedNumbers([4,3,2,7,8,2,3,1])
console.log(res)