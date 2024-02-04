package entities

type Attribute struct {
	Value    int
	Modifier int
}

type Attributes struct {
	Strength     *Attribute
	Dexterity    *Attribute
	Constitution *Attribute
	Intelligence *Attribute
	Wisdom       *Attribute
	Charisma     *Attribute
}

func NewAttributes() *Attributes {
	return &Attributes{
		Strength:     &Attribute{Value: 0},
		Dexterity:    &Attribute{Value: 0},
		Constitution: &Attribute{Value: 0},
		Intelligence: &Attribute{Value: 0},
		Wisdom:       &Attribute{Value: 0},
		Charisma:     &Attribute{Value: 0},
	}
}

func (a *Attribute) SetModifier() {
	a.Modifier = (a.Value - 11) / 2
}

func (a Attributes) SumAll() int {
	return a.Strength.Value + a.Dexterity.Value + a.Constitution.Value + a.Intelligence.Value + a.Wisdom.Value + a.Charisma.Value
}

func (a *Attributes) SetModifiers() {
	a.Strength.SetModifier()
	a.Dexterity.SetModifier()
	a.Constitution.SetModifier()
	a.Intelligence.SetModifier()
	a.Wisdom.SetModifier()
	a.Charisma.SetModifier()
}

func (a *Attributes) GetNames() []string {
	return []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
}

func (a *Attributes) SetByName(name string, value int) {
	var attr *Attribute
	switch name {
	case "Strength":
		attr = a.Strength
	case "Dexterity":
		attr = a.Dexterity
	case "Constitution":
		attr = a.Constitution
	case "Intelligence":
		attr = a.Intelligence
	case "Wisdom":
		attr = a.Wisdom
	case "Charisma":
		attr = a.Charisma
	}

	attr.Value = value
	attr.SetModifier()
}
