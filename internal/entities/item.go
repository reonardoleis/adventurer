package entities

type Item struct {
	ID          int
	Name        string
	Description string
	Value       int
	Damage      [3]int // count, sides, bonus
	ArmorClass  int
}
