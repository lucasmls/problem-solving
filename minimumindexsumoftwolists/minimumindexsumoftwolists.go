package minimumindexsumoftwolists

/**
 * Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.
 * You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.
 *
 * Input:
 *		list1 = ["Shogun","Tapioca Express","Burger King","KFC"],
 *    list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
 * Output: ["Shogun"]
 * Explanation: The only restaurant they both like is "Shogun".
 *
 * @link https://leetcode.com/problems/minimum-index-sum-of-two-lists/
 */

func findRestaurant(list1 []string, list2 []string) []string {
	firstRestaurantsMap := map[string]int{}
	for index, restaurant := range list1 {
		firstRestaurantsMap[restaurant] = index
	}

	secondRestaurantsMap := map[string]int{}
	for index, restaurant := range list2 {
		secondRestaurantsMap[restaurant] = index
	}

	matchesMap := map[string]int{}
	for restaurant, firstIndex := range firstRestaurantsMap {
		secondIndex, ok := secondRestaurantsMap[restaurant]
		if !ok {
			continue
		}

		matchesMap[restaurant] = firstIndex + secondIndex
	}

	result := []string{}

	if len(matchesMap) == 1 {
		for restaurant := range matchesMap {
			result = append(result, restaurant)
		}

		return result
	}

	min := -1
	minRestaurant := ""

	prev := -1
	isTie := true
	for restaurant, indexSum := range matchesMap {
		if min == -1 {
			min = indexSum
			prev = indexSum
			minRestaurant = restaurant
		}

		if indexSum < min {
			min = indexSum
			minRestaurant = restaurant
		}

		if indexSum != prev {
			isTie = false
		}

		prev = indexSum
	}

	if isTie == true {
		for restaurant := range matchesMap {
			result = append(result, restaurant)
		}

		return result
	}

	return []string{minRestaurant}
}
