// It handles all menu navigation and delegates logic to internal packages.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"kiosk-tracker/internal/purchases"
	"kiosk-tracker/internal/reports"
	"kiosk-tracker/internal/sales"
	"kiosk-tracker/internal/stock"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║     KIOSK TRACKER  — CLI MVP     ║")
	fmt.Println("╚══════════════════════════════════╝")

	for {
		printMenu()
		choice := promptString("Enter choice")

		switch strings.TrimSpace(choice) {
		case "1":
			handleRecordSale()
		case "2":
			handleRecordPurchase()
		case "3":
			handleViewReport()
		case "4":
			handleViewStock()
		case "5":
			fmt.Println("\nGoodbye! Keep tracking those profits.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please enter 1–5.")
		}
	}
}

// printMenu displays the main navigation menu.
func printMenu() {
	fmt.Println("\n┌─────────────────────────────┐")
	fmt.Println("│          MAIN MENU          │")
	fmt.Println("├─────────────────────────────┤")
	fmt.Println("│  1. Record Sale             │")
	fmt.Println("│  2. Record Purchase         │")
	fmt.Println("│  3. View Report             │")
	fmt.Println("│  4. View Stock              │")
	fmt.Println("│  5. Exit                    │")
	fmt.Println("└─────────────────────────────┘")
}

// handleRecordSale collects sale details from the user and delegates to the sales package.
func handleRecordSale() {
	fmt.Println("\n── Record Sale ──────────────────")

	itemName := promptString("  Item name")
	quantity, err := promptInt("  Quantity sold")
	if err != nil {
		fmt.Printf("Invalid quantity: %v\n", err)
		return
	}
	price, err := promptFloat("  Selling price (KES)")
	if err != nil {
		fmt.Printf("Invalid price: %v\n", err)
		return
	}

	if err := sales.RecordSale(itemName, quantity, price); err != nil {
		fmt.Printf(" Could not record sale: %v\n", err)
		return
	}

	fmt.Printf("Sale recorded — %d × %s @ KES %.2f\n", quantity, itemName, price)
}

// handleRecordPurchase collects purchase details and delegates to the purchases package.
func handleRecordPurchase() {
	fmt.Println("\n── Record Purchase ──────────────")

	itemName := promptString("  Item name")
	quantity, err := promptInt("  Quantity purchased")
	if err != nil {
		fmt.Printf("Invalid quantity: %v\n", err)
		return
	}
	cost, err := promptFloat("  Buying price (KES)")
	if err != nil {
		fmt.Printf("Invalid price: %v\n", err)
		return
	}

	if err := purchases.RecordPurchase(itemName, quantity, cost); err != nil {
		fmt.Printf("Could not record purchase: %v\n", err)
		return
	}

	fmt.Printf("Purchase recorded — %d × %s @ KES %.2f\n", quantity, itemName, cost)
}

// handleViewReport prints the formatted financial report.
func handleViewReport() {
	fmt.Print(reports.FormatReport())
}

// handleViewStock displays the current stock levels for all items.
func handleViewStock() {
	items := stock.CalculateStock()

	fmt.Println("\n── Current Stock ─────────────────────────────────")
	if len(items) == 0 {
		fmt.Println("  No stock data available yet.")
		return
	}

	fmt.Printf("  %-20s %10s %10s %10s\n", "ITEM", "PURCHASED", "SOLD", "REMAINING")
	fmt.Println("  " + strings.Repeat("─", 52))

	for _, item := range items {
		status := ""
		if item.RemainingStock <= 0 {
			status = "OUT"
		}
		fmt.Printf("  %-20s %10d %10d %10d%s\n",
			item.ItemName,
			item.TotalPurchased,
			item.TotalSold,
			item.RemainingStock,
			status,
		)
	}
	fmt.Println("  " + strings.Repeat("─", 52))
}

// ── Input helpers ─────────────────────────────────────────────────────────────

// promptString reads a trimmed string from stdin.
func promptString(label string) string {
	fmt.Printf("%s: ", label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// promptInt reads and parses an integer from stdin.
func promptInt(label string) (int, error) {
	raw := promptString(label)
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid integer", raw)
	}
	return val, nil
}

// promptFloat reads and parses a float64 from stdin.
func promptFloat(label string) (float64, error) {
	raw := promptString(label)
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid number", raw)
	}
	return val, nil
}
