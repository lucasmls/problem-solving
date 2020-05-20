
/**
 * Given an array A of non-negative integers, return an array consisting of
 * all the even elements of A, followed by all the odd elements of A.
 */

/**
 * @param {number[]} A
 * @return {number[]}
 */
function sortArrayByParity(A) {
  let nextEvenIndex = 0

  for (let i = 0; i < A.length; i++) {
    const actualItem = A[i]
    const itemInEvenPosition = A[nextEvenIndex]

    if (actualItem % 2 === 0) {
      A[nextEvenIndex] = actualItem
      A[i] = itemInEvenPosition 
      nextEvenIndex++
    }
  }

  return A
}

console.log(sortArrayByParity([3,1,2,4]))
console.log(sortArrayByParity([1, 6, 3, 8, 11]))
console.log(sortArrayByParity([2, 1, 2, 3, 4, 7, 8, 10, 123, 54, 32, 33, 6]))
console.log(sortArrayByParity([0, 1, 2]))
