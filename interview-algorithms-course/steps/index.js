/** 
 * Write a function that accepts a positive number N.
 * The function should console log a step shape
 * with N levels using the # character.  Make sure the
 * step has spaces on the right hand side!

 *   steps(2)
 *       '# '
 *       '##'
 *   steps(3)
 *       '#  '
 *       '## '
 *       '###'
 *   steps(4)
 *       '#   '
 *       '##  '
 *       '### '
 *       '####'
 */

// function steps(n) {
//   for (let i = 1; i <= n; i++) {
//     // Spaces 3 - 1 = 2  || Hashes 3 - 2 = 1
//     // Spaces 3 - 2 = 1  || Hashes 3 - 1 - 2
//     // Spaces 3 - 3 = 0  || Hashes 3 - 3 = 0

//     const spacesToFill = n - i;
//     const hashesToFill = n - spacesToFill;

//     const hashes = Array(hashesToFill).fill("#");
//     const spaces = Array(spacesToFill).fill(" ");

//     console.log(`${hashes.join("")}${spaces.join("")}`);
//   }
// }

// function steps(n) {
//   for (let row = 0; row < n; row++) {
//     let strRow = "";

//     for (let col = 0; col < n; col++) {
//       if (col <= row) {
//         strRow += "#";
//       } else {
//         strRow += " ";
//       }
//     }

//     console.log(strRow);
//   }
// }

// function steps(n, i = 1) {
//   if (i > n) {
//     return;
//   }

//   // 5 - 1 = 4
//   // 5 - 2 = 3
//   const spacesToFill = n - i;
//   // 5 - 4 == 1
//   // 5 - 3 == 2
//   const hashesToFill = n - spacesToFill;

//   const spaces = Array(spacesToFill).fill(" ");
//   const hashes = Array(hashesToFill).fill("#");
//   console.log(`${hashes.join("")}${spaces.join("")}`);

//   steps(n, i + 1);
// }

function steps(n, row = 0, stair = "") {
  if (n === row) {
    return;
  }

  if (n === stair.length) {
    console.log(stair);
    return steps(n, row + 1);
  }

  stair += stair.length <= row ? "#" : " ";
  steps(n, row, stair);
}

module.exports = steps;
