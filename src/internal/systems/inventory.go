package systems

import (
	"TheVengefulBlade/internal/models"
	"errors"
)

func CanAddItem(c *models.Character, item models.Item) bool {
	if len(c.Inventory) >= c.MaxSlots {
		return false
	} else {
		return true
	}
}

func AddItem(c *models.Character, item models.Item) error {
	var err error
	if len(c.Inventory) >= c.InventoryCapacity {
		err = errors.New("Inventaire plein...")
	} else {
		c.Inventory = append(c.Inventory, item)
		err = nil
	}
	return err
}

func RemoveItem(c *models.Character, itemName string, qty int) error {
	err := errors.New("Objet introuvable...")
	for i := 0; i < len(c.Inventory); i++ {
		if itemName == c.Inventory[i].Name {
			err = nil
			break
		}
	}
	return err
}

func CountItem(c *models.Character, itemName string) int {}

func UpgradeInventorySlot(c *models.Character) error {
	var err error
	if c.InventoryUpgradesUsed >= 3 {
		err = errors.New("Inventaire déjà amélioré au maximum...")
	} else {
		c.MaxSlots = c.MaxSlots + 10
		err = nil
	}
	return err
}
