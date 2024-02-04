package character_creation

import (
	"github.com/reonardoleis/adventurer/internal/dice"
	"github.com/reonardoleis/adventurer/internal/entities"
)

type CharacterCreator struct {
	Character     *entities.Character
	AttributeRoll int
}

func NewCharacterCreator(character *entities.Character) *CharacterCreator {
	return &CharacterCreator{Character: character}
}

func (cc *CharacterCreator) CreateCharacter(name, story, race, class string) {
	cc.Character.Name = name
	cc.Character.Story = story
	cc.Character.Race = race
	cc.Character.Class = class
	cc.Character.Level = 1
	cc.Character.Experience = 0
	cc.Character.NextLevel = 100

	cc.Character.Inventory = make([]*entities.InventoryItem, 0)
	cc.Character.Attributes = entities.NewAttributes()
}

func (cc *CharacterCreator) RollAttributes() {
	cc.AttributeRoll = dice.Roll(4, 6).GetGreaterN(3).Sum()
}

func (cc *CharacterCreator) SetAttributes(
	strength, dexterity, constitution, intelligence, wisdom, charisma int,
) bool {
	if strength+dexterity+constitution+intelligence+wisdom+charisma != cc.AttributeRoll {
		return false
	}

	cc.Character.Attributes.Strength.Value = strength
	cc.Character.Attributes.Dexterity.Value = dexterity
	cc.Character.Attributes.Constitution.Value = constitution
	cc.Character.Attributes.Intelligence.Value = intelligence
	cc.Character.Attributes.Wisdom.Value = wisdom
	cc.Character.Attributes.Charisma.Value = charisma

	cc.Character.Attributes.SetModifiers()

	return true
}

func (cc *CharacterCreator) BuyAttributes(
	strength, dexterity, constitution, intelligence, wisdom, charisma int,
) bool {
	if strength < 8 || dexterity < 8 || constitution < 8 || intelligence < 8 || wisdom < 8 || charisma < 8 {
		return false
	}

	if strength > 15 || dexterity > 15 || constitution > 15 || intelligence > 15 || wisdom > 15 || charisma > 15 {
		return false
	}

	cc.Character.Attributes.Strength.Value = strength
	cc.Character.Attributes.Dexterity.Value = dexterity
	cc.Character.Attributes.Constitution.Value = constitution
	cc.Character.Attributes.Intelligence.Value = intelligence
	cc.Character.Attributes.Wisdom.Value = wisdom
	cc.Character.Attributes.Charisma.Value = charisma

	cc.Character.Attributes.SetModifiers()

	cc.Character.CreationFinished = true

	return true
}
