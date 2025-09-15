package systems

func CanAddItem(c *Character, item Item) bool {
	if len(c.Inventory) >= c.MaxSlots {
		return false
	} else {
		return true
	}
}
