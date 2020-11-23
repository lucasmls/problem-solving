package findwordsthatcanbeformedbycharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	tt := []struct {
		firstInput  []string
		secondInput string
		output      int
	}{
		{
			firstInput: []string{
				"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin",
				"ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb",
				"ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl",
				"boygirdlggnh",
				"xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx",
				"nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop",
				"hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx",
				"juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr",
				"lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo",
				"oxgaskztzroxuntiwlfyufddl",
				"tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp",
				"qnagrpfzlyrouolqquytwnwnsqnmuzphne",
				"eeilfdaookieawrrbvtnqfzcricvhpiv",
				"sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz",
				"yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue",
				"hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv",
				"cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo",
				"teyygdmmyadppuopvqdodaczob",
				"qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs",
				"qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs",
			},
			secondInput: "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp",
			output:      0,
		},
		{
			firstInput:  []string{"cat", "bt", "hat", "tree"},
			secondInput: "atach",
			output:      6,
		},
		{
			firstInput:  []string{"hello", "world", "leetcode"},
			secondInput: "welldonehoneyr",
			output:      10,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, CountCharacters(tc.firstInput, tc.secondInput))
	}
}
