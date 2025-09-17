package systems

import (
	"TheVengefulBlade/internal/models"
	"errors"
	"strings"
)

func CanAddItem(c *models.Character, item models.Item) bool {
	return len(c.Inventory) < c.MaxSlots
}

func AddItem(c *models.Character, item models.Item) error {
	if !CanAddItem(c, item) {
		return errors.New("Inventaire plein...")
	}
	c.Inventory = append(c.Inventory, item)
	return nil
}

func RemoveItem(c *models.Character, itemName string, qty int) error {
	itemName = strings.ToLower(itemName)
	for i, item := range c.Inventory {
		if strings.ToLower(item.Name) == itemName {
			if item.Qty > qty {
				c.Inventory[i].Qty -= qty
			} else if item.Qty == qty {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			} else {
				return errors.New("Quantité insuffisante...")
			}
			return nil
		}
	}
	return errors.New("Objet introuvable...")
}

func CountItem(c *models.Character, itemName string) int {
	count := 0
	itemName = strings.ToLower(itemName)
	for _, item := range c.Inventory {
		if strings.ToLower(item.Name) == itemName {
			count += item.Qty
		}
	}
	return count
}

func UpgradeInventorySlot(c *models.Character) error {
	if c.InventoryUpgradesUsed >= 3 {
		return errors.New("Inventaire déjà amélioré au maximum...")
	}
	c.MaxSlots += 10
	c.InventoryUpgradesUsed++
	return nil
}
