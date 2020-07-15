/*
 * Check to see if two provided strings are anagrams of eachother.
 * One string is an anagram of another if it uses the same characters
 * in the same quantity. Only consider characters, not spaces
 * or punctuation.  Consider capital letters to be the same as lower case

 *   anagrams('rail safety', 'fairy tales') --> True
 *   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
 *   anagrams('Hi there', 'Bye there') --> False
 */

function replaceSpecialCharacters(string) {
  return string.replace(/[`~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]/gi, "");
}

function getStringCharsCountingDictionary(string) {
  return string.split("").reduce((acc, letter) => {
    return {
      ...acc,
      [letter]: (acc[letter] || 0) + 1
    };
  }, {});
}

function anagrams(stringA, stringB) {
  const firstString = replaceSpecialCharacters(stringA).toLowerCase();
  const secondString = replaceSpecialCharacters(stringB).toLowerCase();

  const firstStringCharsCountingDictionary = getStringCharsCountingDictionary(
    firstString
  );

  const secondStringCharsCountingDictionary = getStringCharsCountingDictionary(
    secondString
  );

  if (
    Object.keys(firstStringCharsCountingDictionary).length !==
    Object.keys(secondStringCharsCountingDictionary).length
  ) {
    return false;
  }

  for (letter in firstStringCharsCountingDictionary) {
    const firstStringLetterCount = firstStringCharsCountingDictionary[letter];
    const secondStringLetterCount = secondStringCharsCountingDictionary[letter];

    if (firstStringLetterCount !== secondStringLetterCount) {
      return false;
    }
  }

  return true;
}

// function anagrams(stringA, stringB) {
//   const firstString = replaceSpecialCharacters(stringA)
//     .toLowerCase()
//     .split("")
//     .sort()
//     .join("")

//   const secondString = replaceSpecialCharacters(stringB)
//     .toLowerCase()
//     .split("")
//     .sort()
//     .join("")

//   return firstString === secondString
// }

module.exports = anagrams
