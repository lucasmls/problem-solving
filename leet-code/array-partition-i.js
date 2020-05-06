/**
 * Given an array of 2n integers, your task is to group these integers into n pairs of integer,
 * say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.
 * Example:
 * Input: [1,4,3,2]
 * Output: 4
 * 
 * @link https://leetcode.com/problems/array-partition-i/
 */

/**
 * @param {number[]} nums
 * @return {number}
 */
function arrayPairSum (nums) { 
  const sortedList = Int16Array.from(nums).sort();

  let result = 0
  for(let i = 0; i < sortedList.length -1; i += 2) {
    result += sortedList[i]
  }

  return result
};

console.log(arrayPairSum([1,2,3,2]))