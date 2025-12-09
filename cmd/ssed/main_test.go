package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func runSsed(args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	err := Run(args, strings.NewReader(""), &stdout, &stderr)

	return stdout.String(), stderr.String(), err
}

func runSsedWithStdin(stdin string, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	err := Run(args, strings.NewReader(stdin), &stdout, &stderr)

	return stdout.String(), stderr.String(), err
}

func TestCLI_Version(t *testing.T) {
	stdout, _, err := runSsed("version")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stdout, "ssed version") {
		t.Errorf("expected version output, got: %s", stdout)
	}
}

func TestCLI_Help(t *testing.T) {
	stdout, _, err := runSsed("--help")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stdout, "natural language interface") {
		t.Errorf("expected help text, got: %s", stdout)
	}

	if !strings.Contains(stdout, "--preview") {
		t.Errorf("expected --preview flag in help, got: %s", stdout)
	}
}

func TestCLI_Examples(t *testing.T) {
	stdout, _, err := runSsed("examples")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stdout, "Replace text") {
		t.Errorf("expected examples output, got: %s", stdout)
	}
}

func TestCLI_ReplaceWithStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		query    string
		expected string
	}{
		{
			name:     "simple replace",
			input:    "hello world\n",
			query:    "replace hello with hi",
			expected: "hi world\n",
		},
		{
			name:     "replace multiple lines",
			input:    "foo bar\nfoo baz\n",
			query:    "replace foo with qux",
			expected: "qux bar\nqux baz\n",
		},
		{
			name:     "no match",
			input:    "hello world\n",
			query:    "replace xyz with abc",
			expected: "hello world\n",
		},
		{
			name:     "replace with empty",
			input:    "hello world\n",
			query:    "replace world with ",
			expected: "hello \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _, err := runSsedWithStdin(tt.input, tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if stdout != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, stdout)
			}
		})
	}
}

func TestCLI_DeleteWithStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		query    string
		expected string
	}{
		{
			name:     "delete matching line",
			input:    "hello\nerror line\nworld\n",
			query:    "delete error",
			expected: "hello\nworld\n",
		},
		{
			name:     "delete multiple matches",
			input:    "keep\ndelete this\nkeep too\ndelete this too\n",
			query:    "delete delete",
			expected: "keep\nkeep too\n",
		},
		{
			name:     "no match keeps all",
			input:    "hello\nworld\n",
			query:    "delete xyz",
			expected: "hello\nworld\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _, err := runSsedWithStdin(tt.input, tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if stdout != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, stdout)
			}
		})
	}
}

func TestCLI_ShowWithStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		query    string
		expected string
	}{
		{
			name:     "show matching line",
			input:    "hello\nerror line\nworld\n",
			query:    "show error",
			expected: "error line\n",
		},
		{
			name:     "show multiple matches",
			input:    "match this\nno dice\nmatch that\n",
			query:    "show match",
			expected: "match this\nmatch that\n",
		},
		{
			name:     "no match shows nothing",
			input:    "hello\nworld\n",
			query:    "show xyz",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _, err := runSsedWithStdin(tt.input, tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if stdout != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, stdout)
			}
		})
	}
}

func TestCLI_InsertWithStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		query    string
		expected string
	}{
		{
			name:     "insert before",
			input:    "line one\ntarget line\nline three\n",
			query:    "insert INSERTED before target",
			expected: "line one\nINSERTED\ntarget line\nline three\n",
		},
		{
			name:     "insert after",
			input:    "line one\ntarget line\nline three\n",
			query:    "insert INSERTED after target",
			expected: "line one\ntarget line\nINSERTED\nline three\n",
		},
		{
			name:     "insert first (prepend)",
			input:    "line one\nline two\n",
			query:    "insert HEADER first",
			expected: "HEADER\nline one\nline two\n",
		},
		{
			name:     "insert last (append)",
			input:    "line one\nline two\n",
			query:    "insert FOOTER last",
			expected: "line one\nline two\nFOOTER\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _, err := runSsedWithStdin(tt.input, tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if stdout != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, stdout)
			}
		})
	}
}

func TestCLI_FileInput(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "foo bar\nfoo baz\nhello world\n"

	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	stdout, _, err := runSsed("replace foo with qux", tmpFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "qux bar\nqux baz\nhello world\n"
	if stdout != expected {
		t.Errorf("expected %q, got %q", expected, stdout)
	}
}

func TestCLI_PreviewMode(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "foo bar\n"

	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	stdout, _, err := runSsed("replace foo with qux", tmpFile, "--preview")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stdout, "Preview") {
		t.Errorf("expected preview header, got: %s", stdout)
	}

	if !strings.Contains(stdout, "qux bar") {
		t.Errorf("expected transformed content in preview, got: %s", stdout)
	}

	if !strings.Contains(stdout, "no changes made") {
		t.Errorf("expected 'no changes made' message, got: %s", stdout)
	}

	afterContent, _ := os.ReadFile(tmpFile)
	if string(afterContent) != content {
		t.Errorf("file should not be modified in preview mode")
	}
}

func TestCLI_InPlaceEdit(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "foo bar\n"

	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, stderr, err := runSsed("replace foo with qux", tmpFile, "-i")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stderr, "Modified") {
		t.Errorf("expected 'Modified' message, got stderr: %s", stderr)
	}

	afterContent, _ := os.ReadFile(tmpFile)
	expected := "qux bar\n"

	if string(afterContent) != expected {
		t.Errorf("expected file content %q, got %q", expected, string(afterContent))
	}
}

func TestCLI_InPlaceWithBackup(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "foo bar\n"

	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, stderr, err := runSsed("replace foo with qux", tmpFile, "-i", "--backup", ".bak")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(stderr, "Backup created") {
		t.Errorf("expected 'Backup created' message, got stderr: %s", stderr)
	}

	backupFile := tmpFile + ".bak"

	backupContent, err := os.ReadFile(backupFile)
	if err != nil {
		t.Fatalf("backup file should exist: %v", err)
	}

	if string(backupContent) != content {
		t.Errorf(
			"backup should contain original content %q, got %q",
			content,
			string(backupContent),
		)
	}

	afterContent, _ := os.ReadFile(tmpFile)
	expected := "qux bar\n"

	if string(afterContent) != expected {
		t.Errorf("expected file content %q, got %q", expected, string(afterContent))
	}
}

func TestCLI_QuietMode(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "foo bar\n"

	err := os.WriteFile(tmpFile, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, stderr, err := runSsed("replace foo with qux", tmpFile, "-i", "-q")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stderr != "" {
		t.Errorf("expected no stderr in quiet mode, got: %s", stderr)
	}
}

func TestCLI_InvalidQuery(t *testing.T) {
	_, _, err := runSsedWithStdin("hello\n", "invalid command")
	if err == nil {
		t.Error("expected error for invalid query")
	}
}

func TestCLI_MissingQuery(t *testing.T) {
	_, _, err := runSsed()
	if err == nil {
		t.Error("expected error when no query provided")
	}
}

func TestCLI_FileNotFound(t *testing.T) {
	_, _, err := runSsed("replace foo with bar", "/nonexistent/file.txt")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}
