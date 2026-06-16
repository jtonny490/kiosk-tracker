package sales

import "fmt"

type Sale struct {
	ItemName  string
	Quantity  int
	UnitPrice float64
}

// TotalAmount returns the total revenue for this sale.
func (s Sale) TotalAmount() float64 {
	return float64(s.Quantity) * s.UnitPrice
}

// store holds all recorded sales in memory (replaced by DB later).
var store []Sale

// RecordSale prompts the user for sale details and stores the transaction.
func RecordSale(itemName string, quantity int, unitPrice float64) error {
	if itemName == "" {
		return fmt.Errorf("item name cannot be empty")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if unitPrice <= 0 {
		return fmt.Errorf("unit price must be greater than zero")
	}

	sale := Sale{
		ItemName:  itemName,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}
	store = append(store, sale)
	return nil
}

// GetAllSales returns a copy of all recorded sales.
func GetAllSales() []Sale {
	result := make([]Sale, len(store))
	copy(result, store)
	return result
}

// TotalRevenue calculates the sum of all sale amounts.
func TotalRevenue() float64 {
	var total float64
	for _, s := range store {
		total += s.TotalAmount()
	}
	return total
}
