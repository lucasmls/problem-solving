package steps

import (
	"fmt"
	"strings"
)

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

// Steps ...
func Steps(n int) {
	for i := 1; i <= n; i++ {
		spacesToFill := n - i
		hashesToFill := n - spacesToFill

		var spacesStr strings.Builder
		var hashesStr strings.Builder

		for i := 0; i < spacesToFill; i++ {
			fmt.Fprintf(&spacesStr, " ")
		}

		for i := 0; i < hashesToFill; i++ {
			fmt.Fprintf(&hashesStr, "#")
		}

		var row strings.Builder
		fmt.Fprintf(&row, "%s%s", hashesStr.String(), spacesStr.String())

		fmt.Println(row.String())
	}
}
