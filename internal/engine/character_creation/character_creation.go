package character_creation

import (
	"errors"

	"github.com/reonardoleis/adventurer/internal/entities"
	"github.com/reonardoleis/adventurer/internal/tables"
)

var attributeList = []string{
	"Strength",
	"Dexterity",
	"Constitution",
	"Intelligence",
	"Wisdom",
	"Charisma",
}

type CharacterCreator struct {
	Character        *entities.Character
	currentAttribute int
	buyPoints        int
}

func NewCharacterCreator(character *entities.Character) *CharacterCreator {
	return &CharacterCreator{Character: character, buyPoints: 27}
}

func (cc *CharacterCreator) SetName(name string) {
	cc.Character.Name = name
}

func (cc *CharacterCreator) SetStory(story string) {
	cc.Character.Story = story
}

func (cc *CharacterCreator) SetAttributeTo(val int) error {
	if cc.currentAttribute >= len(attributeList) {
		return errors.New("all attributes are already set")
	}

	if cc.buyPoints-tables.AttributeCost(val) < 0 {
		return errors.New("not enough points to buy this attribute")
	}

	switch attributeList[cc.currentAttribute] {
	case "Strength":
		cc.Character.Attributes.Strength = &entities.Attribute{Value: val}
	case "Dexterity":
		cc.Character.Attributes.Dexterity = &entities.Attribute{Value: val}
	case "Constitution":
		cc.Character.Attributes.Constitution = &entities.Attribute{Value: val}
	case "Intelligence":
		cc.Character.Attributes.Intelligence = &entities.Attribute{Value: val}
	case "Wisdom":
		cc.Character.Attributes.Wisdom = &entities.Attribute{Value: val}
	case "Charisma":
		cc.Character.Attributes.Charisma = &entities.Attribute{Value: val}
	}

	cc.Character.Attributes.SetModifiers()

	cc.currentAttribute++
	cc.buyPoints -= tables.AttributeCost(val)

	return nil
}

func (cc *CharacterCreator) GetPointsLeft() int {
	return cc.buyPoints
}

func (cc *CharacterCreator) GetCurrentAttributeName() string {
	if cc.currentAttribute >= len(attributeList) {
		return ""
	}

	return attributeList[cc.currentAttribute]
}
