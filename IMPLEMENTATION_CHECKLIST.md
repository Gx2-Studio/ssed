# ssed Implementation Checklist

This checklist organizes sed features by priority and complexity to guide the implementation of ssed.

## Phase 1: Core Functionality (MVP)

### Basic Infrastructure
- [ ] Command-line argument parsing
- [ ] File input/output handling
- [ ] stdin/stdout processing
- [ ] Line-by-line stream processing
- [ ] Pattern space management
- [ ] Error handling and reporting

### Essential Command-Line Options
- [ ] `-n` (suppress automatic printing)
- [ ] `-e SCRIPT` (add script)
- [ ] `-f FILE` (script from file)
- [ ] `--help` (help message)
- [ ] `--version` (version info)

### Core Addressing
- [ ] Line numbers (e.g., `5`)
- [ ] Line ranges (e.g., `5,10`)
- [ ] Last line (`$`)
- [ ] All lines (no address)
- [ ] Address negation with `!`

### Essential Commands
- [ ] `s///` - Substitution (basic)
  - [ ] Basic replacement
  - [ ] `g` flag (global)
  - [ ] Numeric flag (replace Nth)
  - [ ] `p` flag (print if substituted)
- [ ] `p` - Print pattern space
- [ ] `d` - Delete pattern space
- [ ] `q` - Quit
- [ ] `#` - Comments

### Basic Regular Expressions
- [ ] Literal characters
- [ ] `.` (any character)
- [ ] `*` (zero or more)
- [ ] `^` (start of line)
- [ ] `$` (end of line)
- [ ] `[...]` (character class)
- [ ] `[^...]` (negated class)

### Basic Replacement Features
- [ ] Literal text replacement
- [ ] `&` (matched string)
- [ ] `\n` (backreferences)

---

## Phase 2: Extended Core Features

### Additional Command-Line Options
- [ ] `-i` (in-place editing without backup)
- [ ] `-i.SUFFIX` (in-place with backup)
- [ ] `-r` or `-E` (extended regex)
- [ ] `-s` (separate files)
- [ ] `-u` (unbuffered)

### Extended Addressing
- [ ] Regular expression addresses (`/pattern/`)
- [ ] Range with regex (`/start/,/end/`)
- [ ] Address step (`1~2`)
- [ ] Offset ranges (`/pattern/,+5`)

### Text Manipulation Commands
- [ ] `a\` - Append text
- [ ] `i\` - Insert text
- [ ] `c\` - Change lines
- [ ] `=` - Print line number
- [ ] `l` - List (visually unambiguous)

### Flow Control
- [ ] `:LABEL` - Define label
- [ ] `b LABEL` - Branch
- [ ] `t LABEL` - Branch on successful substitution
- [ ] `T LABEL` - Branch on failed substitution

### Extended Substitution Flags
- [ ] `i/I` - Case-insensitive matching
- [ ] `w FILE` - Write to file if substituted

### Extended Regex (with -r/-E)
- [ ] `+` (one or more)
- [ ] `?` (zero or one)
- [ ] `|` (alternation)
- [ ] `()` (groups without escaping)
- [ ] `{n,m}` (repetition without escaping)

---

## Phase 3: Advanced Features

### Hold Space Commands
- [ ] `h` - Copy to hold space
- [ ] `H` - Append to hold space
- [ ] `g` - Copy from hold space
- [ ] `G` - Append from hold space
- [ ] `x` - Exchange spaces

### Multi-line Commands
- [ ] `n` - Next (replace pattern space)
- [ ] `N` - Next (append to pattern space)
- [ ] `D` - Delete first line of pattern space
- [ ] `P` - Print first line of pattern space

### File I/O Commands
- [ ] `r FILE` - Read file
- [ ] `R FILE` - Read line from file
- [ ] `w FILE` - Write to file
- [ ] `W FILE` - Write first line to file

### Additional Commands
- [ ] `y///` - Transliterate
- [ ] `Q` - Quit silently
- [ ] `Q N` - Quit with exit code
- [ ] `q N` - Quit with exit code
- [ ] `{}` - Command grouping

### GNU Regex Extensions
- [ ] `\s` - Whitespace
- [ ] `\S` - Non-whitespace
- [ ] `\w` - Word character
- [ ] `\W` - Non-word character
- [ ] `\b` - Word boundary
- [ ] `\B` - Non-word boundary
- [ ] `\<` - Start of word
- [ ] `\>` - End of word

### Advanced Replacement Features
- [ ] `\L` - Lowercase until `\U` or `\E`
- [ ] `\U` - Uppercase until `\L` or `\E`
- [ ] `\E` - Stop case conversion
- [ ] `\l` - Lowercase next char
- [ ] `\u` - Uppercase next char

### Escape Sequences
- [ ] `\n` - Newline
- [ ] `\t` - Tab
- [ ] `\r` - Carriage return
- [ ] `\a` - Alert
- [ ] `\f` - Form feed
- [ ] `\v` - Vertical tab
- [ ] `\oNNN` - Octal
- [ ] `\xHH` - Hexadecimal

---

## Phase 4: Advanced Options & Modes

### Command-Line Options
- [ ] `-l N` - Line wrap length for `l` command
- [ ] `-z` - NUL-separated lines
- [ ] `--posix` - POSIX mode
- [ ] `--debug` - Debug mode with annotations

### Advanced Substitution Features
- [ ] `e` flag - Execute as shell command
- [ ] `m/M` flag - Multi-line mode
- [ ] Multiple numeric flags

### Regex Capture Groups
- [ ] `\1` through `\9` backreferences
- [ ] Nested capture groups
- [ ] Non-capturing groups (if implementing)

---

## Phase 5: Polish & Optimization

### Error Handling
- [ ] Graceful handling of missing files
- [ ] Invalid regex error messages
- [ ] Invalid address error messages
- [ ] Invalid command error messages
- [ ] Circular file operations detection
- [ ] Resource limits

### Performance Optimization
- [ ] Efficient regex compilation
- [ ] Memory-efficient pattern space
- [ ] Buffered I/O optimization
- [ ] Early exit optimization
- [ ] Address filtering optimization

### Compatibility
- [ ] POSIX sed compatibility mode
- [ ] GNU sed compatibility
- [ ] BSD sed compatibility notes
- [ ] Behavior flags for dialect switching

### User Experience
- [ ] Colored error messages (optional)
- [ ] Better error reporting
- [ ] Warning messages for common mistakes
- [ ] Dry-run mode for `-i`
- [ ] Verbose mode
- [ ] Progress indicators for large files

---

## Phase 6: Super Features (ssed Enhancements)

These are potential features that could make ssed "super" compared to regular sed:

### Modern Improvements
- [ ] JSON/YAML-aware operations
- [ ] UTF-8 native support
- [ ] Better Unicode handling
- [ ] Syntax highlighting for scripts
- [ ] Better error messages with context
- [ ] Warnings for common mistakes

### Enhanced Regex
- [ ] PCRE (Perl-Compatible Regular Expressions)
- [ ] Named capture groups
- [ ] Lookahead/lookbehind assertions
- [ ] Atomic groups
- [ ] Possessive quantifiers

### New Features
- [ ] Built-in variable system
- [ ] Arithmetic operations
- [ ] String functions (length, substr, etc.)
- [ ] Multiple pattern/hold spaces
- [ ] Subroutines/functions
- [ ] Include directive for scripts
- [ ] Conditional expressions
- [ ] Loop constructs

### Developer Experience
- [ ] Script formatter/linter
- [ ] Script debugger with breakpoints
- [ ] Step-through execution
- [ ] Watch pattern/hold space
- [ ] Script profiler
- [ ] Unit testing framework for scripts

### Safety Features
- [ ] Dry-run mode for in-place editing
- [ ] Confirmation prompts for destructive operations
- [ ] Automatic backups with timestamps
- [ ] Undo capability
- [ ] Safety limits (max recursion, max file size, etc.)

### Integration Features
- [ ] Git integration (edit in repo safely)
- [ ] Streaming JSON/CSV parsing
- [ ] Network stream support
- [ ] Parallel processing for multiple files
- [ ] Watch mode (re-run on file changes)

---

## Testing Priorities

### Unit Tests Needed
- [ ] Each command individually
- [ ] Each command-line option
- [ ] Address parsing and matching
- [ ] Regex compilation and matching
- [ ] Pattern/hold space operations
- [ ] File I/O operations
- [ ] Edge cases (empty files, single line, etc.)

### Integration Tests
- [ ] Common use cases (sed one-liners)
- [ ] Multi-command scripts
- [ ] Complex flow control
- [ ] Multi-line operations
- [ ] In-place editing
- [ ] Multiple file processing

### Compatibility Tests
- [ ] Compare output with GNU sed
- [ ] Compare output with BSD sed
- [ ] POSIX compliance tests
- [ ] Error handling compatibility

### Performance Tests
- [ ] Large files (GB+)
- [ ] Many files
- [ ] Complex regex patterns
- [ ] Deep recursion/branching
- [ ] Memory usage profiling

---

## Documentation Priorities

- [ ] README.md with quick start
- [ ] Installation instructions
- [ ] Basic usage examples
- [ ] Command reference
- [ ] Option reference
- [ ] Migration guide from sed
- [ ] Cookbook/recipes
- [ ] API documentation (if library)
- [ ] Contributing guidelines
- [ ] Changelog

---

## Notes

- Start with Phase 1 to get a working MVP
- Ensure each phase is fully tested before moving to the next
- Phase 6 features should be carefully considered for usefulness vs. complexity
- Maintain backward compatibility with sed where possible
- Document any intentional deviations from sed behavior
