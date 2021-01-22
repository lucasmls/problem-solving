package defanginganipaddress

import "strings"

// https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	result := ""
	parts := strings.Split(address, ".")

	for i := 0; i < len(parts)-1; i++ {
		result += parts[i] + "[.]"
	}

	result += parts[len(parts)-1]

	return result
}
