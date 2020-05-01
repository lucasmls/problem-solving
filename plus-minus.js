function plusMinus(arr) {
  const countings = arr.reduce((countings, item) => {
    if (item >= 1) {
      return { ...countings, positives: countings.positives += 1 }
    }

    if (item === 0) {
      return { ...countings, zeros: countings.zeros += 1 }
    }

    if (item <= -1) {
      return { ...countings, negatives: countings.negatives += 1 }
    }

    return countings
  }, {
    positives: 0,
    negatives: 0,
    zeros: 0,
  })

  const itemsCount = arr.length

  console.log((countings.positives / itemsCount).toFixed(6))
  console.log((countings.negatives / itemsCount).toFixed(6))
  console.log((countings.zeros / itemsCount).toFixed(6))
}

plusMinus([-4, 3, -9, 0, 4, 1])