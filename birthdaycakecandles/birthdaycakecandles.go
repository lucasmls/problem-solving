package birthdaycakecandles

/*
 * You are in charge of the cake for your niece's birthday and have decided the cake will have one candle for each year of her total age. When she blows out the candles, she’ll only be able to blow out the tallest ones. Your task is to find out how many candles she can successfully blow out.
 * For example, if your niece is turning  years old, and the cake will have  candles of height , , , , she will be able to blow out  candles successfully, since the tallest candles are of height  and there are  such candles.
 * Complete the function birthdayCakeCandles in the editor below. It must return an integer representing the number of candles she can blow out.
 *
 * @link https://www.hackerrank.com/challenges/birthday-cake-candles/problem
 */

// BirthdayCakeCandles ...
func BirthdayCakeCandles(ar []int32) int32 {
	maxCandleSize := int32(0)

	for i := 0; i < len(ar); i++ {
		candle := ar[i]

		if candle > maxCandleSize {
			maxCandleSize = candle
		}
	}

	var tallestCandles []int32
	for i := 0; i < len(ar); i++ {
		candle := ar[i]

		if candle == maxCandleSize {
			tallestCandles = append(tallestCandles, candle)
		}
	}

	return int32(len(tallestCandles))
}
