package dice

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Rolls []int
}

type Value struct {
	Value int
}

func (v Value) Present() string {
	n := fmt.Sprintf("%d", v.Value)
	return "You rolled " + n
}

func (v Value) Get() int {
	return v.Value
}

func Roll(count, sides int) Dice {
	dice := Dice{}
	for i := 0; i < count; i++ {
		r := rand.Intn(sides) + 1
		dice.Rolls = append(dice.Rolls, r)
	}
	return dice
}

func DecisionRoll(dificulty [2]int) (rolled int, needed int, success bool) {
	lowestBound := dificulty[0]
	highestBound := dificulty[1]
	needed = rand.Intn(highestBound-lowestBound) + lowestBound
	rolled = Roll(1, 20).Value().Value

	return rolled, needed, rolled >= needed
}

func (d Dice) Value() Value {
	if len(d.Rolls) == 0 {
		return Value{Value: 0}
	}

	return Value{Value: d.Rolls[0]}
}

func (d Dice) Sum() int {
	sum := 0
	for _, r := range d.Rolls {
		sum += r
	}
	return sum
}

func (d Dice) Max() int {
	max := 0
	for _, r := range d.Rolls {
		if r > max {
			max = r
		}
	}
	return max
}

func (d Dice) Min() int {
	min := 0
	for i, r := range d.Rolls {
		if i == 0 || r < min {
			min = r
		}
	}
	return min
}

func (d *Dice) Remove(value int) {
	for i, r := range d.Rolls {
		if r == value {
			d.Rolls = append(d.Rolls[:i], d.Rolls[i+1:]...)
			return
		}
	}
}

func (d Dice) GetGreaterN(n int) Dice {
	values := make([]int, n)
	for i := 0; i < n; i++ {
		max := d.Max()
		values[i] = max
		d.Remove(max)
	}

	return Dice{Rolls: values}
}
