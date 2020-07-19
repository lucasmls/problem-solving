/*
 * Given a string, return true if the string is a palindrome
 * or false if it is not.  Palindromes are strings that
 * form the same word if it is reversed. *Do* include spaces
 * and punctuation in determining if the string is a palindrome.
 *  palindrome("abba") === true
 *  palindrome("abcdefg") === false
 */

// const reverseString = str =>
//   str
//     .split("")
//     .reverse()
//     .join("");

// function palindrome(str) {
//   const reversedString = reverseString(str);
//   return reversedString === str;
// }

// function palindrome(str) {
//   return str.split("").every((char, index) => {
//     return char === str[str.length - index - 1];
//   });
// }

function palindrome(str) {
  const middle = Math.floor(str.length / 2)

  for (let i = 0; i < middle; i++) {
    const letter = str[i]
    const counterLetter = str[str.length - (i + 1)]

    if (letter !== counterLetter) {
      return false;
    }
  }

  return true
}

module.exports = palindrome