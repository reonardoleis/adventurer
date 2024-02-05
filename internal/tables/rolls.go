package tables

var (
	decisionRollLow    = [2]int{0, 5}
	decisionRollMedium = [2]int{6, 10}
	decisionRollHigh   = [2]int{11, 15}
	decisionRollVHigh  = [2]int{16, 20}
)

var (
	DecisionRolls = map[string][2]int{
		"low":       decisionRollLow,
		"medium":    decisionRollMedium,
		"high":      decisionRollHigh,
		"very_high": decisionRollVHigh,
	}
)
