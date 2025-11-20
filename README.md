# ssed - Super Simple sed

**Text transformation in plain English. No syntax to memorize.**

```bash
# Instead of this:
sed -i.bak 's/foo/bar/g' file.txt

# Just say this:
ssed "replace all foo with bar in file.txt with backup"
```

## Project Status

🚧 **In Development** - This project is in the planning and early development phase.

## What is ssed?

ssed is a natural language interface for text transformation that provides all the power of sed without the cryptic syntax. Just describe what you want in plain English, and ssed figures out how to do it.

### Core Philosophy

**"No syntax to memorize, just describe what you want"**

ssed should be usable by anyone without reading documentation - even on their first try.

### Goals

1. **Zero Learning Curve**: Use plain English instead of cryptic sed syntax
2. **Intuitive**: Natural language queries that read like sentences
3. **Safe by Default**: Preview changes before applying, easy backups
4. **Functionally Complete**: Everything sed can do, but in plain English
5. **Interactive & Helpful**: Guided mode with suggestions and examples
6. **sed Compatible**: Also accepts traditional sed syntax for power users

## Why ssed?

**The Problem with sed:**

Everyone knows sed is powerful, but every time you need it, you have to:
- Look up the syntax (again)
- Remember the difference between BRE and ERE
- Figure out escaping rules
- Debug cryptic error messages
- Hope you don't mess up a file

**The ssed Solution:**

```bash
# Traditional sed - requires expertise
sed -n '/ERROR/,/^$/p' log.txt | sed 's/^/  /' | sed '/DEBUG/d'

# ssed - just describe what you want
ssed "show lines between ERROR and empty line from log.txt, indent them, and remove lines with DEBUG"
```

## Quick Examples

### Basic Operations
```bash
# Replace text
ssed "replace hello with hi in greeting.txt"

# Delete lines
ssed "delete empty lines from document.txt"

# Extract content
ssed "show lines containing error from app.log"

# Insert text
ssed "insert header at beginning of README.md"
```

### Smart Patterns
```bash
# Predefined patterns - no regex needed!
ssed "extract all email addresses from contacts.txt"
ssed "find all URLs in webpage.html"
ssed "remove phone numbers from privacy.txt"

# Natural language patterns
ssed "delete lines starting with # from config.txt"
ssed "show words ending with ing from text.txt"
```

### Safe Editing
```bash
# Preview first
ssed "replace all TODO with DONE in tasks.txt show preview"

# Automatic backups
ssed "delete test data from production.db with backup"

# Dry run mode
ssed "remove all comments from code.js dry run"
```

## How It Works

### Command-Line Mode
```bash
# Direct execution
ssed "your natural language query here"

# Works on multiple files
ssed "remove trailing spaces in all .js files"

# Pipes and streams
cat file.txt | ssed "delete empty lines"
```

### Interactive Mode
```bash
$ ssed

ssed> What would you like to do?
User: replace old with new

ssed> In which file?
User: config.txt

ssed> [Preview shown]
Apply changes? (yes/no)

User: yes
ssed> ✓ Done! Modified config.txt
```

## Documentation

### For Users
- **[EXAMPLES.md](EXAMPLES.md)** - Comprehensive usage examples
  - Real-world scenarios
  - Common patterns and operations
  - Interactive mode examples
  - Quick reference card

### For Developers
- **[LANGUAGE_SPEC.md](LANGUAGE_SPEC.md)** - Complete natural language specification
  - Grammar and syntax definition
  - Pattern matching rules
  - Operation types and modifiers
  - Implementation guidelines

- **[SED_TO_SSED_MAPPING.md](SED_TO_SSED_MAPPING.md)** - sed to ssed translation guide
  - Every sed command in natural language
  - Complete feature parity mapping
  - Migration guide from sed

### Reference Materials
- **[SED_FEATURES.md](SED_FEATURES.md)** - Complete sed feature reference
  - All sed commands and capabilities
  - Used as feature completeness checklist

- **[QUICK_REFERENCE.md](QUICK_REFERENCE.md)** - Traditional sed reference
  - For sed compatibility mode

## Implementation Phases

### Phase 1: Natural Language Parser (MVP)
- Basic query parsing (replace, delete, show, insert)
- Literal text matching
- Line number addressing
- Single file operations
- Preview mode

### Phase 2: Pattern Matching
- Predefined patterns (email, URL, phone, etc.)
- Natural language patterns (starting with, containing, etc.)
- Case-insensitive matching
- Multiple file support

### Phase 3: Advanced Operations
- Range operations (between patterns)
- Transform operations (case conversion, etc.)
- Compound operations (multiple actions)
- Conditional operations

### Phase 4: Interactive Mode
- Guided prompts with suggestions
- Tab completion
- Command history and learning
- Smart error messages and corrections

### Phase 5: sed Compatibility
- Accept traditional sed syntax
- Translate sed to natural language
- Full feature parity with GNU sed
- Compatibility flags

### Phase 6: AI Enhancement (Future)
- Intent recognition for ambiguous queries
- Smart suggestions based on file type
- Auto-correction of queries
- Learning from user behavior

## Technology Stack

- **Language**: Go
  - Fast, compiled, cross-platform
  - Excellent for text processing
  - Easy deployment (single binary)

- **Core Components**:
  - Natural Language Parser
  - Pattern Matching Engine
  - Stream Processor
  - Interactive REPL

- **Target**: Cross-platform (Linux, macOS, Windows)

## Project Structure

```
ssed/
├── README.md                      # This file
├── LANGUAGE_SPEC.md              # Natural language specification
├── SED_TO_SSED_MAPPING.md        # sed feature mapping
├── EXAMPLES.md                   # Usage examples
├── SED_FEATURES.md               # sed reference (for feature parity)
├── QUICK_REFERENCE.md            # sed syntax reference
├── cmd/
│   └── ssed/                     # Main CLI application
├── pkg/
│   ├── nlp/                      # Natural language parser
│   ├── patterns/                 # Pattern matching (email, URL, etc.)
│   ├── operations/               # Core operations (replace, delete, etc.)
│   ├── stream/                   # Stream processor
│   ├── interactive/              # Interactive mode & REPL
│   └── translator/               # sed syntax translator
├── internal/
│   ├── executor/                 # Operation executor
│   └── fileops/                  # File operations
└── test/
    ├── nlp/                      # Parser tests
    ├── integration/              # End-to-end tests
    └── compatibility/            # sed compatibility tests
```

## Feature Comparison

| Feature | sed | ssed |
|---------|-----|------|
| Learning Curve | High | None |
| Syntax | Cryptic | Plain English |
| Error Messages | Minimal | Helpful |
| Preview Mode | No | Yes |
| Interactive Mode | No | Yes |
| Pattern Library | No | Yes (email, URL, etc.) |
| Multiple Files | Manual | Built-in |
| Safety Features | Minimal | Extensive |
| Documentation Needed | Always | Never |
| Power | Full | Full |
| Speed | Fast | Fast |

## Installation (Future)

```bash
# Via Go
go install github.com/yourusername/ssed@latest

# Via Homebrew (macOS/Linux)
brew install ssed

# Via apt (Debian/Ubuntu)
apt install ssed

# Download binary
# Binaries for Linux, macOS, Windows on releases page
```

## Usage

```bash
# Natural language (recommended)
ssed "replace all foo with bar in file.txt"

# Traditional sed syntax (compatibility mode)
ssed 's/foo/bar/g' file.txt

# Interactive mode
ssed

# With preview
ssed "delete empty lines from file.txt" --preview

# Help and examples
ssed --help
ssed --examples
ssed --examples replace
```

## Development Roadmap

- [x] Define natural language specification
- [x] Map all sed features to natural language
- [x] Create comprehensive examples
- [ ] Implement natural language parser
- [ ] Implement basic operations (Phase 1)
- [ ] Add pattern matching (Phase 2)
- [ ] Build interactive mode (Phase 4)
- [ ] Add sed compatibility mode (Phase 5)
- [ ] Optimize performance
- [ ] Release v1.0

## Contributing

This project is in early planning stages. We welcome:
- Feedback on natural language design
- Suggestions for common use cases
- Ideas for predefined patterns
- Testing and bug reports (once implemented)
- Documentation improvements

## FAQ

**Q: Will ssed be slower than sed?**
A: The parser adds minimal overhead. Once parsed, execution speed should be comparable to sed.

**Q: Can I use ssed as a drop-in replacement for sed?**
A: Yes! ssed accepts traditional sed syntax for backwards compatibility.

**Q: What about complex sed scripts?**
A: ssed can either translate them to natural language or execute them directly in compatibility mode.

**Q: Do I need to learn regex?**
A: No! ssed provides predefined patterns (email, URL, etc.) and natural language patterns. But regex is supported if you want it.

**Q: Is this AI-powered?**
A: Not initially. Phase 1-5 use traditional parsing. AI enhancement is planned for Phase 6 for advanced intent recognition.

**Q: Why not just improve sed documentation?**
A: Documentation helps, but syntax is the root problem. Natural language eliminates the need to learn syntax entirely.

## License

TBD (likely MIT or Apache 2.0)

## Resources

- [GNU sed Manual](https://www.gnu.org/software/sed/manual/sed.html) - Reference for feature parity
- [POSIX sed Specification](https://pubs.opengroup.org/onlinepubs/9699919799/) - Standards compliance
- [sed One-Liners](http://sed.sourceforge.net/sed1line.txt) - Will be translated to ssed examples

---

**Note**: This project is in active planning. The specification is being finalized before implementation begins. Star and watch for updates!
