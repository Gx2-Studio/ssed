# ssed Implementation Guide

**A Step-by-Step Guide to Implementing ssed from Scratch**

This guide provides practical tips, hints, and guidance for implementing all the TODOs in this project. It's designed to help you learn by doing, not by copying code.

---

## 📋 Project Overview

**What you're building**: A tool that converts plain English (e.g., "replace all foo with bar in file.txt") into sed operations, making text processing accessible without memorizing cryptic syntax.

**Current state**:
- Documentation complete ✅
- Specifications complete ✅
- Implementation not started ⚠️

**Technology**: Go (excellent choice for text processing and CLI tools!)

---

## 🎯 Understanding the TODOs

Your `IMPLEMENTATION_CHECKLIST.md` is organized into **6 phases**:

### **Phase 1: Core Functionality (MVP)** - START HERE!
This is your foundation. Includes:
- Basic infrastructure (CLI, file I/O, stream processing)
- Essential command-line options (`-n`, `-e`, `-f`, `--help`)
- Core addressing (line numbers, ranges)
- Essential commands (`s///`, `p`, `d`, `q`)
- Basic regex support

### **Phase 2: Extended Core Features**
Expands on Phase 1 with more options and regex features.

### **Phase 3: Advanced Features**
Hold space, multi-line commands, file I/O commands.

### **Phase 4: Advanced Options & Modes**
Polish and advanced features.

### **Phase 5: Polish & Optimization**
Error handling, performance, compatibility.

### **Phase 6: Super Features**
Modern improvements that make ssed better than regular sed.

---

## 🚀 Step-by-Step Implementation Guide

### **Step 1: Set Up Your Go Project Structure**

**What to do:**
```bash
# Initialize Go module (if not done)
go mod init github.com/yourusername/ssed

# Create the directory structure
mkdir -p cmd/ssed
mkdir -p pkg/{lexer,parser,ast,patterns,operations,stream}
mkdir -p internal/{executor,fileops}
mkdir -p test/{lexer,parser,integration}
```

**Why this structure?**
- `cmd/ssed/`: Entry point (main.go) - the CLI application
- `pkg/`: Reusable packages (can be imported by others)
  - `lexer/`: Breaks input into tokens
  - `parser/`: Builds Abstract Syntax Tree from tokens
  - `ast/`: Defines AST node types
  - `patterns/`: Pattern matching (literal, regex, predefined)
  - `operations/`: Text transformations (replace, delete, etc.)
  - `stream/`: Stream processor (line-by-line processing)
- `internal/`: Private implementation (can't be imported externally)
  - `executor/`: Executes operations
  - `fileops/`: File operations
- `test/`: Test files organized by component

**Learning resources:**
- [Go project layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go)

---

### **Step 2: Start with the Lexer (Tokenizer)**

**What it does**:
Breaks input string into tokens (words, symbols, operators). This is the first stage of parsing.

**Example transformation:**
```
Input:  "replace all foo with bar"
Output: [REPLACE, ALL, IDENTIFIER("foo"), WITH, IDENTIFIER("bar")]
```

**How to implement:**

1. **Create `pkg/lexer/token.go`**:
   - Define a `Token` struct with `Type` and `Value` fields
   - Define token types as constants (enum pattern in Go)

   Example token types you'll need:
   - Keywords: `REPLACE`, `DELETE`, `INSERT`, `SHOW`, `WITH`, `IN`, `FROM`
   - Literals: `STRING`, `NUMBER`, `IDENTIFIER`
   - Special: `EOF`, `ILLEGAL`, `NEWLINE`

2. **Create `pkg/lexer/lexer.go`**:
   - Implement a `Lexer` struct that holds:
     - Input string
     - Current position
     - Current character
   - Implement `NextToken()` method that:
     - Reads the current character
     - Determines token type
     - Returns the token
     - Advances to next position
   - Helper methods:
     - `readChar()`: Move to next character
     - `peekChar()`: Look at next character without advancing
     - `skipWhitespace()`: Skip spaces, tabs
     - `readIdentifier()`: Read word characters
     - `readString()`: Read quoted strings
     - `readNumber()`: Read numeric literals

**Key concepts to learn:**
- **State machines**: The lexer is a state machine that transitions based on input
- **Character-by-character parsing**: Process input one character at a time
- **Lookahead**: Peek at next character to decide current token type
- **String vs rune**: Go uses runes (int32) for characters, strings are UTF-8

**Go packages to use:**
- `strings`: String manipulation (`strings.ContainsRune`, etc.)
- `unicode`: Character classification (`unicode.IsLetter`, `unicode.IsDigit`, `unicode.IsSpace`)

**Testing approach:**
Write tests for each token type and combination:
```go
func TestLexer_ReplaceCommand(t *testing.T) {
    input := "replace foo with bar"
    lexer := NewLexer(input)

    tests := []struct {
        expectedType  TokenType
        expectedValue string
    }{
        {REPLACE, "replace"},
        {IDENTIFIER, "foo"},
        {WITH, "with"},
        {IDENTIFIER, "bar"},
        {EOF, ""},
    }

    for i, tt := range tests {
        tok := lexer.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("test[%d] - wrong token type. expected=%q, got=%q",
                i, tt.expectedType, tok.Type)
        }
        if tok.Value != tt.expectedValue {
            t.Fatalf("test[%d] - wrong literal. expected=%q, got=%q",
                i, tt.expectedValue, tok.Value)
        }
    }
}
```

**Common pitfalls:**
- Not handling whitespace properly
- Forgetting to advance position after reading
- Not handling EOF correctly
- Not supporting quoted strings with spaces

---

### **Step 3: Build the Parser**

**What it does**:
Takes tokens from lexer and builds an Abstract Syntax Tree (AST). The AST represents the structure of the command in a way that's easy to execute.

**Your grammar** (from GRAMMAR.md):
```
<query> ::= <action> <target> <context>? <option>*
```

**How to implement:**

1. **Create `pkg/ast/ast.go`**:
   Define node types for your AST. Each node represents a part of the command.

   Example node types:
   ```go
   type Node interface {
       TokenLiteral() string
   }

   type Query struct {
       Action      Action
       Target      Target
       Replacement string      // For replace operations
       Context     Context     // Where to apply
       Options     []Option    // How to apply
   }

   type Action interface {
       Node
       actionNode()
   }

   type ReplaceAction struct {
       Token Token
       Scope string  // "all", "first", "last", etc.
   }

   type DeleteAction struct {
       Token Token
   }

   type Target interface {
       Node
       targetNode()
   }

   type LiteralTarget struct {
       Token Token
       Value string
   }

   type LineTarget struct {
       Token Token
       Line  int
   }

   type Context interface {
       Node
       contextNode()
   }

   type FileContext struct {
       Token Token
       Path  string
   }
   ```

2. **Create `pkg/parser/parser.go`**:
   Implement recursive descent parser. This is a top-down parser that works by calling functions for each grammar rule.

   Structure:
   ```go
   type Parser struct {
       lexer        *lexer.Lexer
       currentToken Token
       peekToken    Token
       errors       []string
   }

   func NewParser(l *lexer.Lexer) *Parser {
       p := &Parser{lexer: l}
       // Read two tokens to initialize current and peek
       p.nextToken()
       p.nextToken()
       return p
   }

   func (p *Parser) nextToken() {
       p.currentToken = p.peekToken
       p.peekToken = p.lexer.NextToken()
   }

   func (p *Parser) Parse() (*ast.Query, error) {
       query := &ast.Query{}

       // Parse action (required)
       query.Action = p.parseAction()
       if query.Action == nil {
           return nil, errors.New("expected action")
       }

       // Parse target (required)
       query.Target = p.parseTarget()
       if query.Target == nil {
           return nil, errors.New("expected target")
       }

       // Parse "with X" for replace operations
       if _, ok := query.Action.(*ast.ReplaceAction); ok {
           query.Replacement = p.parseReplacement()
       }

       // Parse context (optional)
       if p.currentTokenIs(IN) || p.currentTokenIs(FROM) {
           query.Context = p.parseContext()
       }

       // Parse options (optional, multiple)
       for p.currentToken.Type != EOF {
           opt := p.parseOption()
           if opt != nil {
               query.Options = append(query.Options, opt)
           }
           p.nextToken()
       }

       return query, nil
   }

   func (p *Parser) parseAction() ast.Action {
       switch p.currentToken.Type {
       case REPLACE:
           return p.parseReplaceAction()
       case DELETE:
           return p.parseDeleteAction()
       // ... other actions
       default:
           return nil
       }
   }

   func (p *Parser) parseReplaceAction() *ast.ReplaceAction {
       action := &ast.ReplaceAction{Token: p.currentToken}
       p.nextToken()

       // Check for scope modifier (all, first, last)
       if p.currentTokenIs(ALL) || p.currentTokenIs(FIRST) || p.currentTokenIs(LAST) {
           action.Scope = p.currentToken.Value
           p.nextToken()
       } else {
           action.Scope = "first" // default
       }

       return action
   }

   func (p *Parser) parseTarget() ast.Target {
       // Determine target type based on current token
       if p.currentTokenIs(STRING) {
           return &ast.LiteralTarget{
               Token: p.currentToken,
               Value: p.currentToken.Value,
           }
       } else if p.currentTokenIs(NUMBER) {
           lineNum, _ := strconv.Atoi(p.currentToken.Value)
           return &ast.LineTarget{
               Token: p.currentToken,
               Line:  lineNum,
           }
       }
       // ... handle other target types
       return nil
   }

   func (p *Parser) currentTokenIs(t TokenType) bool {
       return p.currentToken.Type == t
   }

   func (p *Parser) expectPeek(t TokenType) bool {
       if p.peekToken.Type == t {
           p.nextToken()
           return true
       }
       p.peekError(t)
       return false
   }
   ```

**Example AST for "replace all foo with bar in file.txt":**
```go
&ast.Query{
    Action: &ast.ReplaceAction{
        Scope: "all",
    },
    Target: &ast.LiteralTarget{
        Value: "foo",
    },
    Replacement: "bar",
    Context: &ast.FileContext{
        Path: "file.txt",
    },
    Options: []ast.Option{},
}
```

**Key concepts to learn:**
- **Abstract Syntax Trees**: Tree representation of syntactic structure
- **Recursive descent parsing**: Parse by recursively calling functions for each grammar rule
- **Lookahead**: Use `peekToken` to decide how to parse current token
- **Error recovery**: Collect parsing errors instead of failing immediately

**Testing approach:**
Parse queries and verify AST structure:
```go
func TestParser_BasicReplace(t *testing.T) {
    input := "replace foo with bar"
    l := lexer.NewLexer(input)
    p := NewParser(l)

    query := p.Parse()
    if query == nil {
        t.Fatal("Parse() returned nil")
    }

    if _, ok := query.Action.(*ast.ReplaceAction); !ok {
        t.Fatalf("Action is not ReplaceAction. got=%T", query.Action)
    }

    target, ok := query.Target.(*ast.LiteralTarget)
    if !ok {
        t.Fatalf("Target is not LiteralTarget. got=%T", query.Target)
    }

    if target.Value != "foo" {
        t.Errorf("Target value wrong. expected='foo', got='%s'", target.Value)
    }

    if query.Replacement != "bar" {
        t.Errorf("Replacement wrong. expected='bar', got='%s'", query.Replacement)
    }
}
```

**Common pitfalls:**
- Not handling optional parts of grammar
- Not advancing tokens properly (infinite loops)
- Confusing `currentToken` and `peekToken`
- Not handling EOF properly

---

### **Step 4: Implement Pattern Matching**

**What it does**:
Matches patterns in text. Supports multiple types: literal strings, regular expressions, and predefined patterns (email, URL, etc.).

**How to implement:**

1. **Create `pkg/patterns/matcher.go`**:
   Define interface and implementations for different pattern types.

   ```go
   type Matcher interface {
       Match(line string) bool
       FindAll(line string) []string
       Replace(line string, replacement string) string
   }

   // Literal matching (simplest - start here)
   type LiteralMatcher struct {
       pattern string
   }

   func NewLiteralMatcher(pattern string) *LiteralMatcher {
       return &LiteralMatcher{pattern: pattern}
   }

   func (m *LiteralMatcher) Match(line string) bool {
       return strings.Contains(line, m.pattern)
   }

   func (m *LiteralMatcher) FindAll(line string) []string {
       if !m.Match(line) {
           return nil
       }
       // Return all occurrences
       var matches []string
       remaining := line
       for strings.Contains(remaining, m.pattern) {
           idx := strings.Index(remaining, m.pattern)
           matches = append(matches, m.pattern)
           remaining = remaining[idx+len(m.pattern):]
       }
       return matches
   }

   func (m *LiteralMatcher) Replace(line string, replacement string) string {
       return strings.ReplaceAll(line, m.pattern, replacement)
   }
   ```

2. **Add regex matching**:
   ```go
   type RegexMatcher struct {
       pattern *regexp.Regexp
   }

   func NewRegexMatcher(pattern string) (*RegexMatcher, error) {
       re, err := regexp.Compile(pattern)
       if err != nil {
           return nil, err
       }
       return &RegexMatcher{pattern: re}, nil
   }

   func (m *RegexMatcher) Match(line string) bool {
       return m.pattern.MatchString(line)
   }

   func (m *RegexMatcher) FindAll(line string) []string {
       return m.pattern.FindAllString(line, -1)
   }

   func (m *RegexMatcher) Replace(line string, replacement string) string {
       return m.pattern.ReplaceAllString(line, replacement)
   }
   ```

3. **Add predefined patterns** (Phase 2):
   ```go
   var PredefinedPatterns = map[string]string{
       "email": `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
       "url":   `https?://[^\s]+`,
       "phone": `\b\d{3}[-.]?\d{3}[-.]?\d{4}\b`,
       "ip":    `\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`,
       // Add more as needed
   }

   func NewPredefinedMatcher(name string) (*RegexMatcher, error) {
       pattern, ok := PredefinedPatterns[name]
       if !ok {
           return nil, fmt.Errorf("unknown predefined pattern: %s", name)
       }
       return NewRegexMatcher(pattern)
   }
   ```

**Key concepts to learn:**
- **Regular expressions**: Pattern matching language
- **Go's regexp package**: How to compile and use regex
- **Escaping**: Special characters in regex
- **Greedy vs non-greedy matching**

**Go packages to use:**
- `regexp`: Regular expression support
- `strings`: String operations

**Testing approach:**
Test each pattern type independently:
```go
func TestLiteralMatcher_Match(t *testing.T) {
    matcher := NewLiteralMatcher("foo")

    tests := []struct {
        input    string
        expected bool
    }{
        {"foo", true},
        {"foobar", true},
        {"barfoo", true},
        {"bar", false},
        {"", false},
    }

    for _, tt := range tests {
        result := matcher.Match(tt.input)
        if result != tt.expected {
            t.Errorf("Match(%q) = %v, want %v", tt.input, result, tt.expected)
        }
    }
}

func TestRegexMatcher_Email(t *testing.T) {
    matcher, _ := NewPredefinedMatcher("email")

    line := "Contact: john@example.com or jane@test.org"
    matches := matcher.FindAll(line)

    if len(matches) != 2 {
        t.Errorf("Expected 2 matches, got %d", len(matches))
    }
}
```

**Common pitfalls:**
- Not escaping special regex characters when needed
- Forgetting to compile regex before using
- Not handling regex compilation errors
- Case sensitivity issues

---

### **Step 5: Build the Stream Processor**

**What it does**:
Reads input line-by-line and applies operations. This is the heart of sed-like functionality.

**How to implement:**

1. **Create `pkg/stream/processor.go`**:

   ```go
   type Processor struct {
       query       *ast.Query
       input       io.Reader
       output      io.Writer
       lineNumber  int
       patternSpace string  // sed's working buffer
       holdSpace    string  // sed's hold buffer (Phase 3)
   }

   func NewProcessor(query *ast.Query, input io.Reader, output io.Writer) *Processor {
       return &Processor{
           query:  query,
           input:  input,
           output: output,
       }
   }

   func (p *Processor) Process() error {
       scanner := bufio.NewScanner(p.input)

       for scanner.Scan() {
           p.lineNumber++
           p.patternSpace = scanner.Text()

           // Check if this line should be processed
           if p.shouldProcess() {
               // Apply the operation
               result, err := p.processLine()
               if err != nil {
                   return err
               }

               // Output the result (unless suppressed)
               if result != "" && !p.isSuppressed() {
                   fmt.Fprintln(p.output, result)
               }
           } else {
               // Line doesn't match, output unchanged (unless -n flag)
               if !p.isSuppressed() {
                   fmt.Fprintln(p.output, p.patternSpace)
               }
           }
       }

       return scanner.Err()
   }

   func (p *Processor) shouldProcess() bool {
       // Check if addressing matches current line
       // For now, process all lines (implement addressing later)
       return true
   }

   func (p *Processor) processLine() (string, error) {
       // Dispatch to appropriate operation based on action type
       switch action := p.query.Action.(type) {
       case *ast.ReplaceAction:
           return p.handleReplace(action)
       case *ast.DeleteAction:
           return p.handleDelete(action)
       case *ast.ShowAction:
           return p.handleShow(action)
       // ... other actions
       default:
           return "", fmt.Errorf("unknown action type: %T", action)
       }
   }

   func (p *Processor) handleReplace(action *ast.ReplaceAction) (string, error) {
       // Get matcher based on target type
       matcher, err := p.getMatcher()
       if err != nil {
           return "", err
       }

       // Apply replacement based on scope
       result := p.patternSpace
       switch action.Scope {
       case "all":
           result = matcher.Replace(result, p.query.Replacement)
       case "first":
           result = strings.Replace(result,
               p.query.Target.(*ast.LiteralTarget).Value,
               p.query.Replacement, 1)
       case "last":
           // More complex - find last occurrence and replace
           // Hint: strings.LastIndex
       }

       return result, nil
   }

   func (p *Processor) handleDelete(action *ast.DeleteAction) (string, error) {
       // Return empty string to delete the line
       return "", nil
   }

   func (p *Processor) handleShow(action *ast.ShowAction) (string, error) {
       // Check if pattern matches
       matcher, err := p.getMatcher()
       if err != nil {
           return "", err
       }

       if matcher.Match(p.patternSpace) {
           return p.patternSpace, nil
       }

       return "", nil  // Don't show this line
   }

   func (p *Processor) getMatcher() (patterns.Matcher, error) {
       switch target := p.query.Target.(type) {
       case *ast.LiteralTarget:
           return patterns.NewLiteralMatcher(target.Value), nil
       case *ast.RegexTarget:
           return patterns.NewRegexMatcher(target.Pattern)
       // ... other target types
       default:
           return nil, fmt.Errorf("unknown target type: %T", target)
       }
   }

   func (p *Processor) isSuppressed() bool {
       // Check if output is suppressed (-n flag)
       // For now, return false (implement in Phase 1)
       return false
   }
   ```

**Key concepts to learn:**
- **Buffered I/O**: `bufio` package for efficient line reading
- **Line-by-line processing**: Core sed model
- **Pattern space**: sed's working buffer for current line
- **Hold space**: sed's secondary buffer (Phase 3)
- **Addressing**: Selecting which lines to process

**Go packages to use:**
- `bufio`: Buffered I/O (`bufio.Scanner` for line reading)
- `io`: I/O interfaces (`io.Reader`, `io.Writer`)
- `fmt`: Formatted I/O

**Testing approach:**
Test with various inputs and operations:
```go
func TestProcessor_Replace(t *testing.T) {
    input := strings.NewReader("foo bar\nfoo baz")
    output := &strings.Builder{}

    query := &ast.Query{
        Action: &ast.ReplaceAction{Scope: "all"},
        Target: &ast.LiteralTarget{Value: "foo"},
        Replacement: "qux",
    }

    proc := NewProcessor(query, input, output)
    err := proc.Process()

    if err != nil {
        t.Fatalf("Process() error: %v", err)
    }

    expected := "qux bar\nqux baz\n"
    if output.String() != expected {
        t.Errorf("Expected %q, got %q", expected, output.String())
    }
}
```

**Common pitfalls:**
- Not handling EOF properly
- Forgetting to output unchanged lines (when not using -n)
- Not handling empty lines correctly
- Buffer/memory issues with large files

---

### **Step 6: Implement Operations**

**What it does**:
Actual text transformations. Each operation is a separate, testable unit.

**How to implement:**

1. **Create `pkg/operations/replace.go`**:
   ```go
   package operations

   import "strings"

   type ReplaceOp struct {
       Pattern     string
       Replacement string
       Scope       string  // "all", "first", "last", or number
   }

   func (op *ReplaceOp) Apply(line string) string {
       switch op.Scope {
       case "all":
           return strings.ReplaceAll(line, op.Pattern, op.Replacement)
       case "first":
           return strings.Replace(line, op.Pattern, op.Replacement, 1)
       case "last":
           return op.replaceLast(line)
       default:
           // Handle numeric scope (Nth occurrence)
           return line
       }
   }

   func (op *ReplaceOp) replaceLast(line string) string {
       // Find last occurrence and replace it
       lastIdx := strings.LastIndex(line, op.Pattern)
       if lastIdx == -1 {
           return line
       }
       return line[:lastIdx] + op.Replacement + line[lastIdx+len(op.Pattern):]
   }
   ```

2. **Create `pkg/operations/delete.go`**:
   ```go
   package operations

   type DeleteOp struct {
       Condition func(string) bool
   }

   func (op *DeleteOp) ShouldDelete(line string) bool {
       if op.Condition != nil {
           return op.Condition(line)
       }
       return true  // Delete all if no condition
   }

   // Helper constructors
   func DeleteEmptyLines() *DeleteOp {
       return &DeleteOp{
           Condition: func(line string) bool {
               return strings.TrimSpace(line) == ""
           },
       }
   }

   func DeleteMatching(pattern string) *DeleteOp {
       return &DeleteOp{
           Condition: func(line string) bool {
               return strings.Contains(line, pattern)
           },
       }
   }
   ```

3. **Create `pkg/operations/insert.go`**:
   ```go
   package operations

   type InsertOp struct {
       Text     string
       Position string  // "before", "after", "beginning", "end"
   }

   func (op *InsertOp) Apply(line string) string {
       switch op.Position {
       case "before":
           return op.Text + "\n" + line
       case "after":
           return line + "\n" + op.Text
       default:
           return line
       }
   }
   ```

4. **Create `pkg/operations/show.go`**:
   ```go
   package operations

   type ShowOp struct {
       Matcher func(string) bool
   }

   func (op *ShowOp) ShouldShow(line string) bool {
       if op.Matcher != nil {
           return op.Matcher(line)
       }
       return true  // Show all if no matcher
   }
   ```

**Key concepts to learn:**
- **Single responsibility**: Each operation does one thing
- **Composition**: Operations can be combined
- **Functional programming**: Use functions as data

**Testing approach:**
Test each operation in isolation:
```go
func TestReplaceOp_All(t *testing.T) {
    op := &ReplaceOp{
        Pattern:     "foo",
        Replacement: "bar",
        Scope:       "all",
    }

    tests := []struct {
        input    string
        expected string
    }{
        {"foo", "bar"},
        {"foo foo", "bar bar"},
        {"foobar", "barbar"},
        {"baz", "baz"},
    }

    for _, tt := range tests {
        result := op.Apply(tt.input)
        if result != tt.expected {
            t.Errorf("Apply(%q) = %q, want %q", tt.input, result, tt.expected)
        }
    }
}
```

**Common pitfalls:**
- Not handling edge cases (empty strings, no matches)
- Mutating state when you shouldn't
- Not being clear about what each operation does

---

### **Step 7: Build the CLI (Main Entry Point)**

**What it does**:
Command-line interface that ties everything together.

**How to implement:**

1. **Create `cmd/ssed/main.go`**:

   **Option A: Simple (start here)**
   ```go
   package main

   import (
       "fmt"
       "os"

       "github.com/yourusername/ssed/pkg/lexer"
       "github.com/yourusername/ssed/pkg/parser"
       "github.com/yourusername/ssed/pkg/stream"
   )

   func main() {
       if len(os.Args) < 2 {
           fmt.Fprintf(os.Stderr, "Usage: ssed <query> [file...]\n")
           os.Exit(1)
       }

       query := os.Args[1]

       // Parse query
       l := lexer.NewLexer(query)
       p := parser.NewParser(l)
       ast, err := p.Parse()
       if err != nil {
           fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
           os.Exit(1)
       }

       // Determine input source
       var input io.Reader = os.Stdin
       if len(os.Args) > 2 {
           // Read from file
           filename := os.Args[2]
           file, err := os.Open(filename)
           if err != nil {
               fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
               os.Exit(1)
           }
           defer file.Close()
           input = file
       }

       // Execute
       processor := stream.NewProcessor(ast, input, os.Stdout)
       if err := processor.Process(); err != nil {
           fmt.Fprintf(os.Stderr, "Execution error: %v\n", err)
           os.Exit(1)
       }
   }
   ```

   **Option B: Using cobra (more features)**
   ```go
   package main

   import (
       "github.com/spf13/cobra"
       "github.com/yourusername/ssed/cmd/ssed/commands"
   )

   func main() {
       rootCmd := &cobra.Command{
           Use:   "ssed",
           Short: "Super Simple sed - text transformation in plain English",
           Long: `ssed is a natural language interface for text transformation.

   Instead of learning sed syntax, just describe what you want in plain English.

   Examples:
     ssed "replace all foo with bar" < input.txt
     ssed "delete empty lines" file.txt
     ssed "show lines containing error" app.log`,
           Run: commands.RunQuery,
       }

       // Add flags
       rootCmd.Flags().BoolP("in-place", "i", false, "Edit files in-place")
       rootCmd.Flags().StringP("backup", "b", "", "Backup suffix for in-place editing")
       rootCmd.Flags().BoolP("quiet", "n", false, "Suppress automatic printing")
       rootCmd.Flags().Bool("preview", false, "Preview changes without applying")

       if err := rootCmd.Execute(); err != nil {
           os.Exit(1)
       }
   }
   ```

2. **Add version command**:
   ```go
   var versionCmd = &cobra.Command{
       Use:   "version",
       Short: "Print version information",
       Run: func(cmd *cobra.Command, args []string) {
           fmt.Println("ssed version 0.1.0")
       },
   }

   func init() {
       rootCmd.AddCommand(versionCmd)
   }
   ```

3. **Add examples command**:
   ```go
   var examplesCmd = &cobra.Command{
       Use:   "examples",
       Short: "Show usage examples",
       Run: func(cmd *cobra.Command, args []string) {
           fmt.Println("ssed Examples:")
           fmt.Println()
           fmt.Println("  Replace text:")
           fmt.Println("    ssed \"replace hello with hi\" greeting.txt")
           fmt.Println()
           fmt.Println("  Delete lines:")
           fmt.Println("    ssed \"delete empty lines\" document.txt")
           // ... more examples
       },
   }
   ```

**Key concepts to learn:**
- **CLI design**: Flags, arguments, subcommands
- **Error handling**: Exit codes, error messages
- **I/O redirection**: stdin, stdout, files
- **User experience**: Help messages, examples

**Go packages to consider:**
- `flag`: Built-in flag parsing (simple)
- `github.com/spf13/cobra`: Popular CLI framework (powerful)
- `github.com/spf13/pflag`: POSIX-compatible flags
- `os`: OS interaction (args, exit codes, file operations)

**Testing approach:**
Integration tests with different arguments:
```go
func TestCLI_ReplaceWithFile(t *testing.T) {
    // Create temp file
    tmpfile, _ := ioutil.TempFile("", "test")
    tmpfile.WriteString("foo bar\n")
    tmpfile.Close()
    defer os.Remove(tmpfile.Name())

    // Run command
    cmd := exec.Command("./ssed", "replace foo with baz", tmpfile.Name())
    output, err := cmd.Output()

    if err != nil {
        t.Fatalf("Command failed: %v", err)
    }

    expected := "baz bar\n"
    if string(output) != expected {
        t.Errorf("Expected %q, got %q", expected, string(output))
    }
}
```

**Common pitfalls:**
- Not handling stdin vs file input properly
- Poor error messages
- Not checking file permissions
- Not handling interrupt signals (Ctrl+C)

---

### **Step 8: Write Comprehensive Tests**

**Testing strategy:**

1. **Unit tests**: Test each component independently
2. **Integration tests**: Test end-to-end workflows
3. **Compatibility tests**: Compare with actual sed (Phase 5)

**Unit test structure:**

```
test/
├── lexer/
│   └── lexer_test.go        # Test tokenization
├── parser/
│   └── parser_test.go       # Test AST generation
├── patterns/
│   └── matcher_test.go      # Test pattern matching
├── operations/
│   └── operations_test.go   # Test each operation
└── integration/
    └── e2e_test.go          # End-to-end tests
```

**Example integration test:**
```go
func TestE2E_BasicReplace(t *testing.T) {
    tests := []struct {
        name     string
        query    string
        input    string
        expected string
    }{
        {
            name:     "simple replace",
            query:    "replace foo with bar",
            input:    "foo",
            expected: "bar",
        },
        {
            name:     "replace all",
            query:    "replace all foo with bar",
            input:    "foo foo foo",
            expected: "bar bar bar",
        },
        {
            name:     "no match",
            query:    "replace foo with bar",
            input:    "baz",
            expected: "baz",
        },
        {
            name:     "multiple lines",
            query:    "replace foo with bar",
            input:    "foo\nbaz\nfoo",
            expected: "bar\nbaz\nbar",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output := runSsed(tt.query, tt.input)
            if output != tt.expected {
                t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, output)
            }
        })
    }
}

func runSsed(query, input string) string {
    l := lexer.NewLexer(query)
    p := parser.NewParser(l)
    ast, _ := p.Parse()

    inputReader := strings.NewReader(input)
    outputWriter := &strings.Builder{}

    proc := stream.NewProcessor(ast, inputReader, outputWriter)
    proc.Process()

    return strings.TrimSpace(outputWriter.String())
}
```

**Test coverage goals:**
- Lexer: 100% (it's small and critical)
- Parser: 95%+ (cover all grammar rules)
- Operations: 100% (each operation is isolated)
- Integration: Cover all examples from EXAMPLES.md

**Running tests:**
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test
go test -run TestLexer_ReplaceCommand ./test/lexer

# Verbose output
go test -v ./...
```

**Common pitfalls:**
- Not testing edge cases
- Tests that depend on each other
- Not cleaning up temp files
- Skipping error cases

---

## 📚 Learning Resources

### **Go Fundamentals**
- [Tour of Go](https://tour.golang.org/) - Official interactive tutorial (START HERE)
- [Effective Go](https://golang.org/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com/) - Practical code examples
- [Go Documentation](https://pkg.go.dev/) - Package documentation

### **Parsing & Compilers**
- [Crafting Interpreters](https://craftinginterpreters.com/) - Excellent book (free online)
- [Writing An Interpreter In Go](https://interpreterbook.com/) - Go-specific (paid)
- [Let's Build a Simple Interpreter](https://ruslanspivak.com/lsbasi-part1/) - Blog series
- [Parsing Techniques](https://dickgrune.com/Books/PTAPG_2nd_Edition/) - Comprehensive reference

### **Text Processing**
- [GNU sed Manual](https://www.gnu.org/software/sed/manual/sed.html) - sed reference
- [sed One-Liners](http://sed.sourceforge.net/sed1line.txt) - Common patterns
- [Regular Expressions](https://regexone.com/) - Interactive regex tutorial
- [Regex101](https://regex101.com/) - Regex testing tool

### **Go CLI Development**
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [CLI Guidelines](https://clig.dev/) - Best practices for CLI tools
- [12 Factor CLI Apps](https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46)

### **Testing in Go**
- [Testing Package](https://golang.org/pkg/testing/) - Official docs
- [Table-Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) - Best practice
- [Advanced Testing](https://about.sourcegraph.com/blog/go/advanced-testing-in-go) - Patterns

---

## 🎯 Recommended Implementation Timeline

### **Week 1-2: Foundation**
**Goal**: Get comfortable with Go and set up infrastructure

- [ ] Complete [Tour of Go](https://tour.golang.org/)
- [ ] Set up project structure
- [ ] Implement basic token types
- [ ] Implement lexer for simple queries
- [ ] Write lexer tests

**Success criteria**: Can tokenize "replace foo with bar"

### **Week 3-4: Parsing**
**Goal**: Build parser and AST

- [ ] Design AST node types
- [ ] Implement parser for simple queries
- [ ] Write parser tests
- [ ] Test end-to-end: lexer + parser

**Success criteria**: Can parse "replace foo with bar" into AST

### **Week 5-6: Core Operations**
**Goal**: Get basic operations working

- [ ] Implement literal pattern matching
- [ ] Implement replace operation
- [ ] Build basic stream processor
- [ ] Write operation tests
- [ ] Test with real inputs

**Success criteria**: Can execute "replace foo with bar" on actual text

### **Week 7-8: CLI & File I/O**
**Goal**: Build usable CLI tool

- [ ] Build main.go with argument parsing
- [ ] Add file input/output
- [ ] Add stdin/stdout support
- [ ] Write integration tests
- [ ] Test with example files

**Success criteria**: Can run `./ssed "replace foo with bar" file.txt`

### **Week 9-10: Expand Operations**
**Goal**: Add more operations

- [ ] Implement delete operation
- [ ] Implement insert operation
- [ ] Implement show/print operation
- [ ] Test each operation thoroughly

**Success criteria**: Can execute delete, insert, show commands

### **Week 11-12: Regex & Patterns**
**Goal**: Add pattern matching beyond literals

- [ ] Implement regex matching
- [ ] Add predefined patterns (email, URL)
- [ ] Add natural language patterns
- [ ] Write pattern tests

**Success criteria**: Can match patterns like "email addresses" or "lines starting with #"

### **Weeks 13+: Phase 2 and Beyond**
**Goal**: Expand to full feature set

- [ ] Implement line addressing
- [ ] Add command-line options
- [ ] Implement range operations
- [ ] Add advanced features from checklist
- [ ] Performance optimization
- [ ] Polish and release

---

## 💡 Implementation Tips & Best Practices

### **1. Start REALLY Small**
Don't try to implement the full grammar on day one. Start with the absolute minimum:

**Day 1**: Make "replace foo with bar" work
**Day 2**: Add "replace all foo with bar"
**Day 3**: Add file input
**Week 2**: Add delete operation
**Month 2**: Add regex

Build incrementally. Each feature should build on working code.

### **2. Test-Driven Development (TDD)**
Write tests FIRST, then implement:

```go
// 1. Write the test (it will fail)
func TestLexer_ReplaceToken(t *testing.T) {
    input := "replace"
    l := NewLexer(input)
    tok := l.NextToken()

    if tok.Type != REPLACE {
        t.Fatalf("wrong token type. expected=REPLACE, got=%q", tok.Type)
    }
}

// 2. Run test: go test
// Result: FAIL (good!)

// 3. Implement just enough to pass
func (l *Lexer) NextToken() Token {
    if l.input == "replace" {
        return Token{Type: REPLACE, Value: "replace"}
    }
    return Token{Type: ILLEGAL}
}

// 4. Run test: go test
// Result: PASS (good!)

// 5. Refactor if needed

// 6. Repeat
```

This ensures you always have tests, and you know when something breaks.

### **3. Use Table-Driven Tests**
Go idiom for testing multiple cases:

```go
func TestLexer_Keywords(t *testing.T) {
    tests := []struct {
        input    string
        expected TokenType
    }{
        {"replace", REPLACE},
        {"delete", DELETE},
        {"insert", INSERT},
        {"show", SHOW},
        {"with", WITH},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            l := NewLexer(tt.input)
            tok := l.NextToken()
            if tok.Type != tt.expected {
                t.Errorf("Expected %q, got %q", tt.expected, tok.Type)
            }
        })
    }
}
```

### **4. Debug with Print Statements**
When things aren't working, add debug output:

```go
func (p *Parser) Parse() *ast.Query {
    fmt.Printf("DEBUG: Starting parse, currentToken=%+v\n", p.currentToken)

    action := p.parseAction()
    fmt.Printf("DEBUG: Parsed action: %+v\n", action)

    // ... more parsing

    return query
}
```

Remove or comment out before committing.

### **5. Use Go's Built-in Tools**
```bash
# Format code (do this often!)
go fmt ./...

# Check for common mistakes
go vet ./...

# Run tests
go test ./...

# Check test coverage
go test -cover ./...

# Build binary
go build -o ssed cmd/ssed/main.go

# Install to $GOPATH/bin
go install cmd/ssed/main.go
```

### **6. Handle Errors Properly**
Go requires explicit error handling:

```go
// BAD: Ignoring errors
result, _ := parser.Parse()

// GOOD: Check and handle errors
result, err := parser.Parse()
if err != nil {
    return nil, fmt.Errorf("parsing failed: %w", err)
}

// Use %w to wrap errors (Go 1.13+)
// This preserves the error chain
```

### **7. Document Your Code**
Write package-level documentation:

```go
// Package lexer implements tokenization for the ssed language.
//
// The lexer breaks input text into tokens, which are the smallest
// meaningful units of the language (keywords, identifiers, strings, etc.).
//
// Example usage:
//
//     l := lexer.NewLexer("replace foo with bar")
//     for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
//         fmt.Printf("%+v\n", tok)
//     }
//
package lexer

// Token represents a single lexical token in the ssed language.
type Token struct {
    Type  TokenType  // The type of token (REPLACE, IDENTIFIER, etc.)
    Value string     // The actual text of the token
}
```

### **8. Commit Often**
Make small, focused commits:

```bash
git add pkg/lexer/
git commit -m "Add basic lexer with keyword recognition"

git add pkg/lexer/lexer_test.go
git commit -m "Add lexer tests for keywords"

git add pkg/parser/
git commit -m "Implement basic parser for replace actions"
```

This makes it easy to roll back if something breaks.

### **9. Read Other People's Code**
Study well-written Go projects:
- [Hugo](https://github.com/gohugoio/hugo) - Static site generator
- [Docker](https://github.com/moby/moby) - Container platform
- [Cobra](https://github.com/spf13/cobra) - CLI library
- [Go standard library](https://golang.org/src/) - Best practices

### **10. Don't Optimize Prematurely**
**First**: Make it work
**Second**: Make it right
**Third**: Make it fast

Don't worry about performance until you have working code and tests.

---

## 🤔 Common Pitfalls & How to Avoid Them

### **Pitfall 1: Trying to Do Everything at Once**
**Problem**: Attempting to implement the entire grammar on day one.
**Solution**: Start with one command: "replace foo with bar". Get that working end-to-end. Then add features one at a time.

### **Pitfall 2: Not Writing Tests**
**Problem**: Writing lots of code without tests, then discovering bugs later.
**Solution**: Write tests alongside code. Use TDD: write test → run test (fails) → implement → test passes.

### **Pitfall 3: Poor Error Messages**
**Problem**: Cryptic errors like "parse error" with no context.
**Solution**: Include helpful context:
```go
// BAD
return nil, errors.New("parse error")

// GOOD
return nil, fmt.Errorf("expected 'with' after replacement target, got '%s' at position %d",
    p.currentToken.Value, p.currentToken.Position)
```

### **Pitfall 4: Forgetting to Advance Position**
**Problem**: Lexer or parser gets stuck in infinite loop.
**Solution**: Always advance position/token in loops:
```go
// BAD
for p.currentToken.Type != EOF {
    if p.currentToken.Type == REPLACE {
        // ... process replace
        // FORGOT TO ADVANCE!
    }
}

// GOOD
for p.currentToken.Type != EOF {
    if p.currentToken.Type == REPLACE {
        // ... process replace
        p.nextToken()  // ADVANCE!
    }
}
```

### **Pitfall 5: Not Handling EOF**
**Problem**: Code crashes when reaching end of input.
**Solution**: Always check for EOF:
```go
func (l *Lexer) NextToken() Token {
    if l.position >= len(l.input) {
        return Token{Type: EOF, Value: ""}
    }
    // ... rest of implementation
}
```

### **Pitfall 6: Ignoring Edge Cases**
**Problem**: Code works for normal cases but fails on edge cases.
**Solution**: Test edge cases:
- Empty input
- Single character
- Very long input
- Special characters
- Whitespace variations
- No matches
- Multiple matches

### **Pitfall 7: Mutating Shared State**
**Problem**: Operations modify shared state, causing unexpected behavior.
**Solution**: Keep operations pure (no side effects):
```go
// BAD: Mutates input
func (op *ReplaceOp) Apply(line *string) {
    *line = strings.ReplaceAll(*line, op.Pattern, op.Replacement)
}

// GOOD: Returns new string
func (op *ReplaceOp) Apply(line string) string {
    return strings.ReplaceAll(line, op.Pattern, op.Replacement)
}
```

### **Pitfall 8: Over-Engineering**
**Problem**: Creating complex abstractions before knowing requirements.
**Solution**: Start simple, refactor later:
```go
// Start simple
func Replace(line, pattern, replacement string) string {
    return strings.ReplaceAll(line, pattern, replacement)
}

// Refactor later when you need more features
type ReplaceOp struct {
    Pattern     string
    Replacement string
    Scope       string
}
```

### **Pitfall 9: Not Reading Go Idioms**
**Problem**: Writing Go like it's Python/Java/C++.
**Solution**: Learn Go idioms:
- Error handling: explicit, not exceptions
- Interfaces: small, implicit
- Composition over inheritance
- Channels for concurrency
- Defer for cleanup

### **Pitfall 10: Giving Up Too Soon**
**Problem**: Getting stuck and giving up.
**Solution**:
- Break problem into smaller pieces
- Ask for help (Go community is friendly!)
- Take breaks
- Review working examples
- Remember: everyone struggles with parsing at first

---

## 🎓 What You'll Learn

By completing this project, you'll gain deep understanding of:

### **Compiler/Interpreter Concepts**
- Lexical analysis and tokenization
- Parsing and Abstract Syntax Trees
- Grammar design and BNF notation
- Recursive descent parsing
- Pattern matching and recognition

### **Go Programming**
- Project structure and organization
- Interfaces and composition
- Error handling patterns
- Testing strategies
- CLI development
- File I/O and stream processing

### **Software Engineering**
- Test-driven development
- Incremental development
- Code organization
- API design
- Documentation

### **Text Processing**
- Regular expressions
- Stream processing
- Line-oriented processing
- Pattern matching algorithms

---

## 🎬 Quick Start Commands

```bash
# 1. Initialize project (if not done)
cd /home/user/ssed
go mod init github.com/yourusername/ssed

# 2. Create directory structure
mkdir -p cmd/ssed
mkdir -p pkg/{lexer,parser,ast,patterns,operations,stream}
mkdir -p internal/{executor,fileops}
mkdir -p test/{lexer,parser,integration}

# 3. Create initial main.go
cat > cmd/ssed/main.go << 'EOF'
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("ssed v0.1.0 - Development")

    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: ssed <query>\n")
        os.Exit(1)
    }

    query := os.Args[1]
    fmt.Printf("Query: %s\n", query)
    fmt.Println("TODO: Implement parsing and execution")
}
EOF

# 4. Create lexer package
mkdir -p pkg/lexer
cat > pkg/lexer/token.go << 'EOF'
package lexer

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
    Type  TokenType
    Value string
}

// Token types
const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Keywords
    REPLACE = "REPLACE"
    DELETE  = "DELETE"
    INSERT  = "INSERT"
    SHOW    = "SHOW"
    WITH    = "WITH"
    IN      = "IN"
    FROM    = "FROM"
    ALL     = "ALL"
    FIRST   = "FIRST"
    LAST    = "LAST"

    // Literals
    IDENTIFIER = "IDENTIFIER"
    STRING     = "STRING"
    NUMBER     = "NUMBER"
)

var keywords = map[string]TokenType{
    "replace": REPLACE,
    "delete":  DELETE,
    "insert":  INSERT,
    "show":    SHOW,
    "with":    WITH,
    "in":      IN,
    "from":    FROM,
    "all":     ALL,
    "first":   FIRST,
    "last":    LAST,
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENTIFIER
}
EOF

# 5. Create first test
cat > test/lexer/lexer_test.go << 'EOF'
package lexer_test

import (
    "testing"
)

func TestLexer_Keywords(t *testing.T) {
    t.Skip("TODO: Implement lexer first, then unskip this test")

    // This test will check keyword recognition
    // Implement after creating lexer.go
}
EOF

# 6. Build and run
go build -o ssed cmd/ssed/main.go
./ssed "replace foo with bar"

# 7. Run tests
go test ./...

# 8. Check everything is set up
go fmt ./...
go vet ./...

echo "Setup complete! Ready to start implementing."
```

---

## 📝 Your Action Plan

### **This Week**
1. ✅ Set up project structure (use commands above)
2. ✅ Create basic main.go
3. ✅ Create token types
4. 📖 Read about lexing/tokenization
5. 📖 Complete [Tour of Go](https://tour.golang.org/) if new to Go

### **Next Week**
6. 🔨 Implement lexer (start with keywords)
7. 🧪 Write lexer tests
8. ✅ Get "replace foo with bar" tokenizing correctly
9. 📖 Read about parsing and ASTs

### **Week 3**
10. 🔨 Design AST node types
11. 🔨 Implement basic parser
12. 🧪 Write parser tests
13. ✅ Get "replace foo with bar" parsing to AST

### **Week 4**
14. 🔨 Implement literal pattern matching
15. 🔨 Implement basic replace operation
16. 🔨 Create stream processor
17. 🧪 Write integration test
18. ✅ Get end-to-end execution working

### **Month 2**
19. 🔨 Add more operations (delete, insert, show)
20. 🔨 Add file I/O
21. 🔨 Add command-line options
22. 🧪 Comprehensive testing

### **Month 3+**
23. 🔨 Phase 2: Extended features
24. 🔨 Phase 3: Advanced features
25. 🚀 Polish and release!

---

## 🎯 Success Milestones

Track your progress with these milestones:

- [ ] **Milestone 1**: Can tokenize "replace foo with bar"
- [ ] **Milestone 2**: Can parse "replace foo with bar" into AST
- [ ] **Milestone 3**: Can execute basic replace on stdin
- [ ] **Milestone 4**: Can execute from command line
- [ ] **Milestone 5**: Can process files
- [ ] **Milestone 6**: All Phase 1 features working
- [ ] **Milestone 7**: Test coverage > 80%
- [ ] **Milestone 8**: Can handle all examples from EXAMPLES.md
- [ ] **Milestone 9**: Phase 2 features complete
- [ ] **Milestone 10**: Full sed compatibility

---

## 💭 Final Thoughts

**This is a substantial project**, but it's absolutely doable. The key is:

1. **Start small** - One feature at a time
2. **Test everything** - Catch bugs early
3. **Iterate** - Build → Test → Refine → Repeat
4. **Learn** - You'll make mistakes, that's how you learn
5. **Be patient** - Parsing is hard, but you'll get it
6. **Have fun** - This is a cool project!

You have **excellent documentation** already. The specs (GRAMMAR.md, LANGUAGE_SPEC.md) are well thought out. Now it's time to bring them to life with code.

**Remember**: Every expert was once a beginner. Every complex project started with a simple "Hello World". You've got this!

---

## 🆘 When You Get Stuck

1. **Read the error message carefully** - Go's error messages are usually helpful
2. **Add debug print statements** - See what's actually happening
3. **Write a test** - Isolate the problem
4. **Simplify** - Reduce to smallest failing case
5. **Take a break** - Fresh eyes help
6. **Search online** - Someone has probably solved this before
7. **Read the code** - Step through line by line
8. **Ask for help** - Go community is friendly

**Common Go resources**:
- [Go Forum](https://forum.golangbridge.org/)
- [Go Subreddit](https://reddit.com/r/golang)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/go)
- [Gophers Slack](https://gophers.slack.com/)

---

## 📚 Additional Resources

### **Books** (optional but helpful)
- "The Go Programming Language" by Donovan & Kernighan
- "Writing An Interpreter In Go" by Thorsten Ball
- "Crafting Interpreters" by Robert Nystrom (free online)

### **Videos**
- [Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE) by Rob Pike
- [GopherCon talks](https://www.youtube.com/c/GopherAcademy) - Many great talks

### **Documentation to Bookmark**
- [Go standard library](https://pkg.go.dev/std)
- [strings package](https://pkg.go.dev/strings)
- [regexp package](https://pkg.go.dev/regexp)
- [bufio package](https://pkg.go.dev/bufio)
- [io package](https://pkg.go.dev/io)

---

**Good luck, and enjoy the journey! 🚀**

Remember: The goal isn't to write perfect code on the first try. The goal is to learn by building. Every bug you fix, every test you write, every refactoring you do - you're learning.

You've got this! 💪
