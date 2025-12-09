package benchmark

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var (
	ssedBinary  string
	testDataDir string
)

func TestMain(m *testing.M) {
	tmpDir, err := os.MkdirTemp("", "ssed-bench")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create temp dir: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	ssedBinary = filepath.Join(tmpDir, "ssed")

	cmd := exec.Command("go", "build", "-o", ssedBinary, "../cmd/ssed")
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build ssed: %v\n", err)
		os.Exit(1)
	}

	testDataDir = tmpDir
	generateTestData(tmpDir)

	os.Exit(m.Run())
}

func generateTestData(dir string) {
	small := strings.Repeat("hello world foo bar\n", 1000)
	_ = os.WriteFile(filepath.Join(dir, "small.txt"), []byte(small), 0o644)

	medium := strings.Repeat("error: something went wrong\ninfo: all good\nwarning: check this\n", 33334)
	_ = os.WriteFile(filepath.Join(dir, "medium.txt"), []byte(medium), 0o644)

	large := strings.Repeat("The quick brown fox jumps over the lazy dog\n", 1000000)
	_ = os.WriteFile(filepath.Join(dir, "large.txt"), []byte(large), 0o644)

	var logBuilder strings.Builder

	for i := range 100000 {
		switch i % 10 {
		case 0:
			logBuilder.WriteString(fmt.Sprintf("[ERROR] 2024-01-01 %06d: Connection failed\n", i))
		case 1, 2:
			logBuilder.WriteString(fmt.Sprintf("[WARN] 2024-01-01 %06d: Retrying operation\n", i))
		default:
			logBuilder.WriteString(fmt.Sprintf("[INFO] 2024-01-01 %06d: Processing request\n", i))
		}
	}

	_ = os.WriteFile(filepath.Join(dir, "log.txt"), []byte(logBuilder.String()), 0o644)

	wideLine := strings.Repeat("abcdefghij", 10000) + "\n"
	wide := strings.Repeat(wideLine, 100)
	_ = os.WriteFile(filepath.Join(dir, "wide.txt"), []byte(wide), 0o644)
}

func BenchmarkReplace_Small_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "small.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "replace foo with baz", file)
	}
}

func BenchmarkReplace_Small_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "small.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "s/foo/baz/g", file)
	}
}

func BenchmarkReplace_Medium_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "medium.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "replace error with ERROR", file)
	}
}

func BenchmarkReplace_Medium_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "medium.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "s/error/ERROR/g", file)
	}
}

func BenchmarkReplace_Large_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "replace fox with cat", file)
	}
}

func BenchmarkReplace_Large_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "s/fox/cat/g", file)
	}
}

func BenchmarkDelete_Log_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "delete ERROR", file)
	}
}

func BenchmarkDelete_Log_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "/ERROR/d", file)
	}
}

func BenchmarkDelete_Log_Grep(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runGrep(b, "-v", "ERROR", file) // inverted match = delete matching lines
	}
}

func BenchmarkDeleteFirst_Medium_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "medium.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "delete first 1000 lines", file)
	}
}

func BenchmarkDeleteFirst_Medium_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "medium.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "1,1000d", file)
	}
}

func BenchmarkDeleteLast_Medium_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "medium.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "delete last 1000 lines", file)
	}
}

func BenchmarkShow_Log_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "show ERROR", file)
	}
}

func BenchmarkShow_Log_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "-n", "/ERROR/p", file)
	}
}

func BenchmarkShow_Log_Grep(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runGrep(b, "ERROR", file)
	}
}

func BenchmarkShowFirst_Large_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "show first 100 lines", file)
	}
}

func BenchmarkShowFirst_Large_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "-n", "1,100p;100q", file) // Add ;100q for early exit
	}
}

func BenchmarkShowFirst_Large_Head(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runHead(b, "-100", file)
	}
}

func BenchmarkShowLast_Large_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "show last 100 lines", file)
	}
}

func BenchmarkShowLast_Large_Tail(b *testing.B) {
	file := filepath.Join(testDataDir, "large.txt")
	for i := 0; i < b.N; i++ {
		runTail(b, "-100", file)
	}
}

func BenchmarkRegexReplace_Log_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, `replace /[0-9]{6}/ with XXXXXX`, file)
	}
}

func BenchmarkRegexReplace_Log_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "-E", `s/[0-9]{6}/XXXXXX/g`, file)
	}
}

func BenchmarkRegexDelete_Log_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, `delete /^\[WARN\]/`, file)
	}
}

func BenchmarkRegexDelete_Log_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, `-E`, `/^\[WARN\]/d`, file)
	}
}

func BenchmarkRegexDelete_Log_Grep(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runGrep(b, "-Ev", `^\[WARN\]`, file) // -E extended regex, -v for inverted
	}
}

func BenchmarkCompound_Log_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "delete INFO then replace ERROR with CRITICAL", file)
	}
}

func BenchmarkCompound_Log_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "log.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "/INFO/d; s/ERROR/CRITICAL/g", file)
	}
}

func BenchmarkWideLines_Ssed(b *testing.B) {
	file := filepath.Join(testDataDir, "wide.txt")
	for i := 0; i < b.N; i++ {
		runSsed(b, "replace abc with xyz", file)
	}
}

func BenchmarkWideLines_Sed(b *testing.B) {
	file := filepath.Join(testDataDir, "wide.txt")
	for i := 0; i < b.N; i++ {
		runSed(b, "s/abc/xyz/g", file)
	}
}

func runSsed(b *testing.B, query, file string) {
	b.Helper()

	var stdout, stderr bytes.Buffer

	cmd := exec.Command(ssedBinary, query, file)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		b.Fatalf("ssed failed: %v\nstderr: %s", err, stderr.String())
	}
}

func runSed(b *testing.B, args ...string) {
	b.Helper()

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("sed", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		b.Fatalf("sed failed: %v\nstderr: %s", err, stderr.String())
	}
}

func runTail(b *testing.B, args ...string) {
	b.Helper()

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("tail", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		b.Fatalf("tail failed: %v\nstderr: %s", err, stderr.String())
	}
}

func runHead(b *testing.B, args ...string) {
	b.Helper()

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("head", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		b.Fatalf("head failed: %v\nstderr: %s", err, stderr.String())
	}
}

func runGrep(b *testing.B, args ...string) {
	b.Helper()

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("grep", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// grep return 1 when no matches found, not an error
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return // no matches ok
		}

		b.Fatalf("grep failed: %v\nstderr: %s", err, stderr.String())
	}
}
