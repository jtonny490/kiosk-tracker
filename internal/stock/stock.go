package stock

import (
	"kiosk-tracker/internal/purchases"
	"kiosk-tracker/internal/sales"
)

type StockItem struct {
	ItemName         string
	TotalPurchased   int
	TotalSold        int
	RemainingStock   int
}

func CalculateStock() []StockItem {
	purchased := make(map[string]int)
	for _, p := range purchases.GetAllPurchases() {
		purchased[p.ItemName] += p.Quantity
	}

	sold := make(map[string]int)
	for _, s := range sales.GetAllSales() {
		sold[s.ItemName] += s.Quantity
	}

	itemSet := make(map[string]struct{})
	for name := range purchased {
		itemSet[name] = struct{}{}
	}
	for name := range sold {
		itemSet[name] = struct{}{}
	}

	var result []StockItem
	for name := range itemSet {
		totalBought := purchased[name]
		totalSold := sold[name]
		result = append(result, StockItem{
			ItemName:       name,
			TotalPurchased: totalBought,
			TotalSold:      totalSold,
			RemainingStock: totalBought - totalSold,
		})
	}

	return result
}
