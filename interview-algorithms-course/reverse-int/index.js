/* 
 * Given an integer, return an integer that is the reverse
 * ordering of numbers.

 *   reverseInt(15) === 51
 *   reverseInt(981) === 189
 *   reverseInt(500) === 5
 *   reverseInt(-15) === -51
 *   reverseInt(-90) === -9
 */

// function reverseInt(n) {
//   return Number(
//     String(n)
//       .split("")
//       .reverse()
//       .reduce((acc, item) => {
//         if (item === "-") {
//           return `${item}${acc}`;
//         }

//         if (item === "0") {
//           return acc;
//         }

//         return (acc += item);
//       }, "")
//   );
// }

// function reverseInt(n) {
//   return Number(
//     String(n)
//       .split("")
//       .reverse()
//       .reduce((acc, item) => {
//         const reservedCharsDictionary = {
//           "-": `${item}${acc}`,
//           "0": acc
//         };

//         return reservedCharsDictionary[item] || (acc += item);
//       }, "")
//   );
// }

function reverseInt(n) {
  const reversed = String(n)
    .split("")
    .reverse()
    .join("");

  // Math.sign(n) will return [1] for positive numbers and [-1] for negative numbers
  return parseInt(reversed) * Math.sign(n);
}

module.exports = reverseInt;
