/**
 * 
 * Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
 * That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].
 * 
 * @link https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
 */

/**
 * @param {number[]} nums
 * @return {number[]}
 */
function smallerNumbersThanCurrent (nums) {
   const sortedNums = [...nums]
    .sort((n1, n2) => n2 - n1)
    .reduce((acc, item, index) => {
      return {
        ...acc,
        [item]: index
      }
    }, {})

    const result = nums.reduce((acc, item) => {
      const position = sortedNums[item]
      const count = (nums.length - 1) - position
      return [ ...acc, count ]
    }, [])
   
    return result
};

// console.log(smallerNumbersThanCurrent([8,1,2,2,3]))
console.log(smallerNumbersThanCurrent([0, 1, 10]))