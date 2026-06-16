package reports

import (
	"fmt"
	"strings"

	"kiosk-tracker/internal/purchases"
	"kiosk-tracker/internal/sales"
)

type Report struct {
	TotalRevenue     float64
	TotalExpenditure float64
	Profit           float64
}

func GenerateReport() Report {
	revenue := sales.TotalRevenue()
	expenditure := purchases.TotalExpenditure()
	return Report{
		TotalRevenue:     revenue,
		TotalExpenditure: expenditure,
		Profit:           revenue - expenditure,
	}
}

func FormatReport() string {
	r := GenerateReport()

	var sb strings.Builder
	divider := strings.Repeat("─", 36)

	sb.WriteString("\n")
	sb.WriteString(divider + "\n")
	sb.WriteString("        KIOSK FINANCIAL REPORT\n")
	sb.WriteString(divider + "\n")
	sb.WriteString(fmt.Sprintf("  Total Sales Revenue : KES %10.2f\n", r.TotalRevenue))
	sb.WriteString(fmt.Sprintf("  Total Purchases Cost: KES %10.2f\n", r.TotalExpenditure))
	sb.WriteString(divider + "\n")

	if r.Profit >= 0 {
		sb.WriteString(fmt.Sprintf("  Net Profit          : KES %10.2f\n", r.Profit))
	} else {
		sb.WriteString(fmt.Sprintf("  Net Loss            : KES %10.2f\n", r.Profit))
	}

	sb.WriteString(divider + "\n")
	return sb.String()
}
