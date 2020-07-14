/**
 * Write a function that accepts a string.  The function should
 * capitalize the first letter of each word in the string then
 * return the capitalized string.

 * capitalize('a short sentence') --> 'A Short Sentence'
 * capitalize('a lazy fox') --> 'A Lazy Fox'
 * capitalize('look, it is working!') --> 'Look, It Is Working!'
 */

// function capitalize(str) {
//   return str
//     .split("")
//     .map((letter, index) =>
//       str[index - 1] === " " || str[index - 1] === undefined
//         ? letter.toUpperCase()
//         : letter
//     )
//     .join("");
// }

function capitalize(str) {
  return str
    .split(" ")
    .map(word => `${word[0].toUpperCase()}${word.slice(1)}`)
    .join(" ")
}

module.exports = capitalize