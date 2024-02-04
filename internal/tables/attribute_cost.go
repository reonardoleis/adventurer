package tables

var (
	attributeCost = map[int]int{
		8:  0,
		9:  1,
		10: 2,
		11: 3,
		12: 4,
		13: 5,
		14: 7,
		15: 9,
	}
)

func AttributeCost(val int) int {
	return attributeCost[val]
}
