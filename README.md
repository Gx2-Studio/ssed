# ssed - Super sed

A modern, enhanced implementation of the classic sed (Stream EDitor) command-line tool, written in Go.

## Project Status

🚧 **In Development** - This project is in the planning and early development phase.

## What is ssed?

ssed aims to be a drop-in replacement for sed with enhanced features, better error messages, and modern improvements while maintaining backward compatibility with standard sed behavior.

### Goals

1. **Full sed Compatibility**: Support all standard sed commands and options
2. **Better UX**: Improved error messages, warnings for common mistakes
3. **Modern Features**: Native UTF-8 support, enhanced regex capabilities
4. **Performance**: Fast and efficient stream processing
5. **Safety**: Better safeguards for destructive operations
6. **Extensibility**: Additional "super" features beyond standard sed

## Why Another sed?

While sed is powerful, it has some limitations:
- Cryptic error messages
- Complex escape sequences
- Limited regex features (BRE by default)
- Platform inconsistencies (GNU vs BSD)
- No built-in safety features for destructive operations

ssed aims to address these while staying true to sed's philosophy of simple, powerful text transformation.

## Documentation

This repository contains comprehensive documentation about sed features:

- **[SED_FEATURES.md](SED_FEATURES.md)** - Complete feature mapping of sed
  - All commands, options, and capabilities
  - Detailed explanations and examples
  - Platform variations and compatibility notes

- **[IMPLEMENTATION_CHECKLIST.md](IMPLEMENTATION_CHECKLIST.md)** - Development roadmap
  - Phased implementation plan
  - Priority-ordered feature list
  - Testing and documentation requirements

- **[QUICK_REFERENCE.md](QUICK_REFERENCE.md)** - Concise sed reference
  - Quick syntax lookup
  - Common one-liners
  - Practical examples

## Planned Phases

### Phase 1: Core Functionality (MVP)
Basic sed operations: substitution, deletion, printing, basic addressing

### Phase 2: Extended Features
In-place editing, extended regex, flow control, text manipulation

### Phase 3: Advanced Features
Hold space, multi-line operations, file I/O, advanced addressing

### Phase 4: Advanced Options
Debug mode, NUL-separated lines, advanced substitution features

### Phase 5: Polish
Performance optimization, compatibility modes, excellent error messages

### Phase 6: Super Features
Modern enhancements beyond standard sed:
- JSON/YAML awareness
- PCRE regex support
- Enhanced debugging tools
- Safety features and dry-run modes
- Better Unicode handling

## Technology

- **Language**: Go
- **Target**: Cross-platform (Linux, macOS, Windows)
- **Philosophy**: Fast, simple, reliable

## Contributing

This project is in early stages. Contributions, ideas, and feedback are welcome!

## Compatibility Target

ssed aims to be compatible with:
- GNU sed (primary target)
- POSIX sed (standard compliance)
- BSD sed (where practical)

Intentional deviations from sed behavior will be clearly documented.

## License

TBD

## Resources

- [GNU sed Manual](https://www.gnu.org/software/sed/manual/sed.html)
- [POSIX sed Specification](https://pubs.opengroup.org/onlinepubs/9699919799/)
- [sed One-Liners](http://sed.sourceforge.net/sed1line.txt)

## Project Structure

```
ssed/
├── README.md                      # This file
├── SED_FEATURES.md               # Complete sed feature reference
├── IMPLEMENTATION_CHECKLIST.md   # Development roadmap
├── QUICK_REFERENCE.md            # Quick reference guide
├── docs/                         # Additional documentation
├── cmd/                          # Command-line interface
├── pkg/                          # Core packages
│   ├── parser/                   # Script parser
│   ├── lexer/                    # Tokenizer
│   ├── regex/                    # Regular expression engine
│   ├── executor/                 # Command executor
│   └── stream/                   # Stream processor
├── internal/                     # Internal packages
└── test/                         # Test suites
    ├── unit/                     # Unit tests
    ├── integration/              # Integration tests
    └── compatibility/            # Compatibility tests vs sed
```

## Quick Start (Future)

Once implemented:

```bash
# Install
go install github.com/yourusername/ssed@latest

# Use like sed
ssed 's/foo/bar/g' file.txt

# Enhanced features
ssed --dry-run -i 's/foo/bar/g' file.txt  # Preview changes
ssed --explain 's/foo/bar/g'              # Explain what script does
```

## Development Roadmap

- [x] Document sed features comprehensively
- [x] Create implementation checklist
- [ ] Set up Go project structure
- [ ] Implement lexer/tokenizer
- [ ] Implement parser
- [ ] Implement basic executor
- [ ] Phase 1: Core MVP
- [ ] Phase 2-6: Progressive enhancement

## Contact

TBD

---

**Note**: This project is in early planning stages. The feature set and implementation details are subject to change.
