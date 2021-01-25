package defanginganipaddress

import "strings"

/**
 * Given a valid (IPv4) IP address, return a defanged version of that IP address.
 *
 * Input: address = "1.1.1.1"
 * Output: "1[.]1[.]1[.]1"
 *
 * Input: address = "255.100.50.0"
 * Output: "255[.]100[.]50[.]0"
 *
 * @link https://leetcode.com/problems/defanging-an-ip-address/
 */

func defangIPaddr(address string) string {
	result := ""
	parts := strings.Split(address, ".")

	for i := 0; i < len(parts)-1; i++ {
		result += parts[i] + "[.]"
	}

	result += parts[len(parts)-1]

	return result
}
