package comparethetriplets

/*
 * Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from 1 to 100 for three categories: problem clarity, originality, and difficulty.
 * The rating for Alice's challenge is the triplet a = (a[0], a[1], a[2]), and the rating for Bob's challenge is the triplet b = (b[0], b[1], b[2]).
 * The task is to find their comparison points by comparing a[0] with b[0], a[1] with b[1], and a[2] with b[2].

 * If a[i] > b[i], then Alice is awarded 1 point.
 * If a[i] < b[i], then Bob is awarded 1 point.
 * If a[i] = b[i], then neither person receives a point.

 * Given a and b, determine their respective comparison points.

 * @link https://www.hackerrank.com/challenges/compare-the-triplets/problem
 */

// CompareTriplets ...
func CompareTriplets(a []int32, b []int32) []int32 {
	aScore := int32(0)
	bScore := int32(0)

	for i := 0; i < len(a); i++ {
		a := a[i]
		b := b[i]

		if a > b {
			aScore++
			continue
		}

		if a == b {
			continue
		}

		bScore++
	}

	return []int32{aScore, bScore}
}
