/**
 * Given the array candies and the integer extraCandies, where candies[i] represents the number
 * of candies that the ith kid has.
 * For each kid check if there is a way to distribute extraCandies among the kids such that he or
 * she can have the greatest number of candies among them. Notice that multiple kids can have the
 * greatest number of candies.
 * 
 * @link https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
 */

/**
 * @param {number[]} candies
 * @param {number} extraCandies
 * @return {boolean[]}
 */
function kidsWithCandies(candies, extraCandies) {
  const maxCandyValue = Math.max(...candies)  
  const result = candies.reduce((acc, candy) => {
    if ((candy + extraCandies) >= maxCandyValue) {
      return [...acc, true]
    }

    return [...acc, false]
  }, [])

  return result
};

console.log(kidsWithCandies([12, 1, 12], 10))