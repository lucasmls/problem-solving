/*
 * Given an array and chunk size, divide the array into many subarrays
 * where each subarray is of length size

 * chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
 * chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
 * chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
 * chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
 * chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]
 */

function chunk(array, size) {
  let chunks = []

  for (let i = 0; i < array.length; i += size) {
  const chunk = array.slice(i, i + size)
  chunks.push(chunk)
  }

  return chunks
}

// function chunk(array, size) {
//   let chunks = [];

//   for (let element of array) {
//     const lastChunk = chunks[chunks.length - 1];

//     if (!lastChunk || lastChunk.length === size) {
//       chunks.push([element]);
//       continue;
//     }

//     lastChunk.push(element);
//   }

//   return chunks;
// }

module.exports = chunk