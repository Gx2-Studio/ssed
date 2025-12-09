//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type result struct {
	name  string
	nsOp  float64
	iters int
}

var testNames = map[string]string{
	"Replace_Small":      "Replace Small (1K lines)",
	"Replace_Medium":     "Replace Medium (100K lines)",
	"Replace_Large":      "Replace Large (1M lines)",
	"Delete_Log":         "Delete Pattern (100K log)",
	"DeleteFirst_Medium": "Delete First N (100K)",
	"DeleteLast_Medium":  "Delete Last N (100K)",
	"Show_Log":           "Show Pattern (100K log)",
	"ShowFirst_Large":    "Show First N (1M)",
	"ShowLast_Large":     "Show Last N (1M)",
	"RegexReplace_Log":   "Regex Replace (100K)",
	"RegexDelete_Log":    "Regex Delete (100K)",
	"Compound_Log":       "Compound (100K)",
	"WideLines":          "Wide Lines (100KB lines)",
}

var testOrder = []string{
	"Replace_Small", "Replace_Medium", "Replace_Large",
	"Delete_Log", "DeleteFirst_Medium", "DeleteLast_Medium",
	"Show_Log", "ShowFirst_Large", "ShowLast_Large",
	"RegexReplace_Log", "RegexDelete_Log",
	"Compound_Log", "WideLines",
}

func main() {
	re := regexp.MustCompile(`^Benchmark(\w+)_(Ssed|Sed)-\d+\s+(\d+)\s+([\d.]+)\s+ns/op`)

	results := make(map[string]map[string]result)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if matches == nil {
			continue
		}

		testName := matches[1]
		tool := strings.ToLower(matches[2])
		iters, _ := strconv.Atoi(matches[3])
		nsOp, _ := strconv.ParseFloat(matches[4], 64)

		if results[testName] == nil {
			results[testName] = make(map[string]result)
		}

		results[testName][tool] = result{name: testName, nsOp: nsOp, iters: iters}
	}

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                              BENCHMARK RESULTS                                     ║")
	fmt.Println("╠═══════════════════════════════╤═══════════╤═══════════╤═══════════╤════════════════╣")
	fmt.Println("║ Test                          │ ssed      │ sed       │ Winner    │ Ratio          ║")
	fmt.Println("╠═══════════════════════════════╪═══════════╪═══════════╪═══════════╪════════════════╣")

	var ssedWins, sedWins, ssedOnly int

	for _, key := range testOrder {
		res, ok := results[key]
		if !ok {
			continue
		}

		displayName := testNames[key]
		if displayName == "" {
			displayName = key
		}

		ssedRes, hasSsed := res["ssed"]
		sedRes, hasSed := res["sed"]

		ssedTime := "—"
		sedTime := "—"
		winner := "—"
		ratio := "—"

		if hasSsed {
			ssedTime = formatDuration(ssedRes.nsOp)
		}

		if hasSed {
			sedTime = formatDuration(sedRes.nsOp)
		}

		if hasSsed && hasSed {
			if ssedRes.nsOp < sedRes.nsOp {
				r := sedRes.nsOp / ssedRes.nsOp
				winner = "ssed"
				ratio = fmt.Sprintf("%.1fx faster", r)
				ssedWins++
			} else if sedRes.nsOp < ssedRes.nsOp {
				r := ssedRes.nsOp / sedRes.nsOp
				if r < 1.1 {
					winner = "~same"
					ratio = "—"
				} else {
					winner = "sed"
					ratio = fmt.Sprintf("%.1fx faster", r)
				}
				sedWins++
			} else {
				winner = "tie"
				ratio = "—"
			}
		} else if hasSsed {
			winner = "ssed only"
			ssedOnly++
		}

		fmt.Printf("║ %-29s │ %9s │ %9s │ %9s │ %14s ║\n",
			truncate(displayName, 29), ssedTime, sedTime, winner, ratio)
	}

	fmt.Println("╚═══════════════════════════════╧═══════════╧═══════════╧═══════════╧════════════════╝")
	fmt.Println()
	fmt.Printf("Summary: sed wins %d, ssed wins %d, ssed-only features: %d\n", sedWins, ssedWins, ssedOnly)
	fmt.Println()
}

func formatDuration(ns float64) string {
	switch {
	case ns >= 1e9:
		return fmt.Sprintf("%.1fs", ns/1e9)
	case ns >= 1e6:
		return fmt.Sprintf("%.0fms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.0fµs", ns/1e3)
	default:
		return fmt.Sprintf("%.0fns", ns)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	return s[:maxLen-1] + "…"
}

