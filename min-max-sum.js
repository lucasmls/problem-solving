function minMaxSum(arr) {
  const sortedArr = arr.sort()

  const firstFour = sortedArr.slice(0, 4)
  const lastFour = sortedArr.slice(1, 5)

  const firstSum = firstFour.reduce((acc, num) => acc + num, 0)
  const lastSum =  lastFour.reduce((acc, num) => acc + num, 0)

  console.log(`${firstSum} ${lastSum}`)
}

minMaxSum([5, 4, 3, 2, 1])