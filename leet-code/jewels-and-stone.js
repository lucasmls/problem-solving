/**
 * You're given strings J representing the types of stones that are jewels, and S representing the
 * stones you have. Each character in S is a type of stone you have. You want to know how many of
 * the stones you have are also jewels.
 * 
 * The letters in J are guaranteed distinct, and all characters in J and S are letters.
 * Letters are case sensitive, so "a" is considered a different type of stone from "A".
 * 
 * @link https://leetcode.com/problems/jewels-and-stones
 */

/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
function numJewelsInStones(J, S) {
  const jewelsMap = J.split("").reduce((acc, letter) => {
    return  { ...acc, [letter]: true }
  }, {})

  const result = S.split("").reduce((acc, stone) => {
    if (jewelsMap[stone] === true) {
      return acc + 1
    }

    return acc
  }, 0)

  return result
};

console.log(numJewelsInStones("z", "ZZ"))