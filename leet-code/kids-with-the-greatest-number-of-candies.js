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