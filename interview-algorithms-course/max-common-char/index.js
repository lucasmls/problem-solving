/*
 *  Given a string, return the character that is most
 *  commonly used in the string.

 *  maxChar("abcccccccd") === "c"
 *  maxChar("apple 1231111") === "1"
 */

function maxChar(str) {
  const charsDictionaryCount = str
    .split("")
    .reduce((map, letter) => {
      return {
        ...map,
        [letter]: (map[letter] || 0) + 1
      }
    }, {})

  let mostCommon = {
    letter: "",
    quantity: 0,
  }

  for (const letter in charsDictionaryCount) {
    const quantity = charsDictionaryCount[letter]

    if (quantity > mostCommon.quantity) {
      mostCommon = { letter, quantity }
    }
  }

  return mostCommon.letter
}

module.exports = maxChar
