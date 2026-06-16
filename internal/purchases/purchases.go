package purchases

import "fmt"

type Purchase struct {
	ItemName  string
	Quantity  int
	UnitCost  float64
}

func (p Purchase) TotalCost() float64 {
	return float64(p.Quantity) * p.UnitCost
}

var store []Purchase

func RecordPurchase(itemName string, quantity int, unitCost float64) error {
	if itemName == "" {
		return fmt.Errorf("item name cannot be empty")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if unitCost <= 0 {
		return fmt.Errorf("unit cost must be greater than zero")
	}

	purchase := Purchase{
		ItemName: itemName,
		Quantity: quantity,
		UnitCost: unitCost,
	}
	store = append(store, purchase)
	return nil
}

func GetAllPurchases() []Purchase {
	result := make([]Purchase, len(store))
	copy(result, store)
	return result
}

func TotalExpenditure() float64 {
	var total float64
	for _, p := range store {
		total += p.TotalCost()
	}
	return total
}
