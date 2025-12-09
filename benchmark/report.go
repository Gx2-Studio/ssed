//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

var toolOrder = []string{"ssed", "sed", "grep", "head", "tail"}

func main() {
	re := regexp.MustCompile(`^Benchmark(\w+)_(Ssed|Sed|Grep|Head|Tail)-\d+\s+(\d+)\s+([\d.]+)\s+ns/op`)

	results := make(map[string]map[string]result)
	toolsSeen := make(map[string]bool)
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
		toolsSeen[tool] = true
	}

	var tools []string
	for _, t := range toolOrder {
		if toolsSeen[t] {
			tools = append(tools, t)
		}
	}

	var sb strings.Builder

	sb.WriteString("SSED BENCHMARK REPORT\n")
	sb.WriteString("=====================\n\n")

	testColWidth := 30
	toolColWidth := 10

	sb.WriteString(fmt.Sprintf("%-*s", testColWidth, "Test"))
	for _, tool := range tools {
		sb.WriteString(fmt.Sprintf(" │ %*s", toolColWidth, tool))
	}
	sb.WriteString(" │ Winner\n")

	sb.WriteString(strings.Repeat("─", testColWidth))
	for range tools {
		sb.WriteString("─┼─")
		sb.WriteString(strings.Repeat("─", toolColWidth))
	}
	sb.WriteString("─┼─")
	sb.WriteString(strings.Repeat("─", 20))
	sb.WriteString("\n")

	wins := make(map[string]int)

	for _, key := range testOrder {
		res, ok := results[key]
		if !ok {
			continue
		}

		displayName := testNames[key]
		if displayName == "" {
			displayName = key
		}

		sb.WriteString(fmt.Sprintf("%-*s", testColWidth, truncate(displayName, testColWidth)))

		var bestTool string
		var bestTime float64 = -1
		var toolTimes []struct {
			tool string
			time float64
		}

		for _, tool := range tools {
			if r, has := res[tool]; has {
				toolTimes = append(toolTimes, struct {
					tool string
					time float64
				}{tool, r.nsOp})
				if bestTime < 0 || r.nsOp < bestTime {
					bestTime = r.nsOp
					bestTool = tool
				}
			}
		}

		for _, tool := range tools {
			if r, has := res[tool]; has {
				sb.WriteString(fmt.Sprintf(" │ %*s", toolColWidth, formatDuration(r.nsOp)))
			} else {
				sb.WriteString(fmt.Sprintf(" │ %*s", toolColWidth, "—"))
			}
		}

		if len(toolTimes) >= 2 {

			sort.Slice(toolTimes, func(i, j int) bool {
				return toolTimes[i].time < toolTimes[j].time
			})
			secondBest := toolTimes[1].time
			ratio := secondBest / bestTime

			if ratio < 1.1 {
				sb.WriteString(fmt.Sprintf(" │ ~same (%s/%s)", toolTimes[0].tool, toolTimes[1].tool))
			} else {
				sb.WriteString(fmt.Sprintf(" │ %s (%.1fx)", bestTool, ratio))
				wins[bestTool]++
			}
		} else if len(toolTimes) == 1 {
			sb.WriteString(fmt.Sprintf(" │ %s only", toolTimes[0].tool))
		} else {
			sb.WriteString(" │ —")
		}

		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString("SUMMARY\n")
	sb.WriteString("-------\n")

	type winCount struct {
		tool  string
		count int
	}
	var winCounts []winCount
	for tool, count := range wins {
		winCounts = append(winCounts, winCount{tool, count})
	}
	sort.Slice(winCounts, func(i, j int) bool {
		return winCounts[i].count > winCounts[j].count
	})

	for _, wc := range winCounts {
		sb.WriteString(fmt.Sprintf("  %s: %d wins\n", wc.tool, wc.count))
	}

	err := os.WriteFile("BENCHMARK_REPORT.txt", []byte(sb.String()), 0o644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Report written to BENCHMARK_REPORT.txt")
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
