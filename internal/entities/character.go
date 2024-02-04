package entities

type InventoryItem struct {
	Item     *Item
	Quantity int
}

type Character struct {
	Name  string
	Story string

	CreationFinished bool

	Race  string
	Class string

	IsNPC     bool
	IsHostile bool

	Level      int
	Experience int
	NextLevel  int

	MainHand *Item
	OffHand  *Item
	Armor    *Item

	Inventory  []*InventoryItem
	Attributes *Attributes
}

func (c *Character) AddExperience() {
	c.Experience++
	if c.Experience >= c.NextLevel {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.Experience = 0
	c.NextLevel = int(float32(c.NextLevel) * 1.5)
}

func (c *Character) AddItem(item *Item) {
	for i, invItem := range c.Inventory {
		if invItem.Item.ID == item.ID {
			c.Inventory[i].Quantity++
			return
		}
	}

	c.Inventory = append(c.Inventory, &InventoryItem{Item: item, Quantity: 1})
}

func (c *Character) RemoveItem(id int) {
	for i, invItem := range c.Inventory {
		if invItem.Item.ID == id {
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity <= 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			return
		}
	}
}
