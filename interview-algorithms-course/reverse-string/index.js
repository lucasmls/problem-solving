/**
 * Given a string, return a new string with the reversed order of characters
 *  reverse("apple") === "leppa";
 *  reverse("hello") === "olleh";
 *  reverse("Greetings!") === "!sgniteerG";
 */

// function reverse(str) {
//   const splitedString = str.split("");

//   let reversedString = "";

//   for (let index = splitedString.length; index > 0; index--) {
//     reversedString = reversedString + splitedString[index - 1];
//   }

//   return reversedString;
// }

// function reverse(str) {
//   return str.split("").reduceRight((acc, strSlice) => {
//     return acc + strSlice;
//   });
// }

// function reverse(str) {
//   return str.split("").reduce((acc, strSlice) => strSlice + acc);
// }

// function reverse(str) {
//   let reversedString = "";

//   for (let char of str) {
//     reversedString = char + reversedString;
//   }

//   return reversedString;
// }

function reverse(str) {
  return str
    .split("")
    .reverse()
    .join("");
}

module.exports = reverse;
