function birthdayCakeCandles(ar) {
  const maxCandle = Math.max(...ar)
  const maxCandleList = ar.filter(candle => candle === maxCandle)
  return maxCandleList.length
}

console.log(birthdayCakeCandles())