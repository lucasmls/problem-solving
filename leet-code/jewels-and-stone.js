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