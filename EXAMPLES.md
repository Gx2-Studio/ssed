# ssed Usage Examples

Real-world examples showing how to use ssed for common text processing tasks.

## Getting Started

### Command-Line Mode
```bash
# Direct command
ssed "replace foo with bar in file.txt"

# Preview changes before applying
ssed "replace foo with bar in file.txt show preview"

# Multiple files
ssed "delete empty lines in all .txt files"
```

### Interactive Mode
```bash
# Launch interactive mode
ssed

# Interactive session
ssed> What would you like to do?
```

---

## Basic Operations

### 1. Simple Text Replacement

**Replace first occurrence:**
```bash
ssed "replace hello with hi in greeting.txt"
```

**Replace all occurrences:**
```bash
ssed "replace all hello with hi in greeting.txt"
```

**Replace specific occurrence:**
```bash
ssed "replace 3rd hello with hi in greeting.txt"
```

**Case-insensitive replacement:**
```bash
ssed "replace all hello with hi ignore case in greeting.txt"
```

**Whole word only:**
```bash
ssed "replace whole word cat with dog in story.txt"
```

### 2. Deleting Content

**Delete specific line:**
```bash
ssed "delete line 5 from data.txt"
```

**Delete line range:**
```bash
ssed "delete lines 10 to 20 from data.txt"
```

**Delete first/last lines:**
```bash
ssed "delete first 5 lines from log.txt"
ssed "delete last line from file.txt"
```

**Delete lines matching pattern:**
```bash
ssed "delete lines containing debug from app.log"
ssed "delete lines starting with # from config.txt"
ssed "delete empty lines from document.txt"
```

**Delete lines NOT matching:**
```bash
ssed "delete lines not containing error from log.txt"
```

### 3. Inserting/Adding Text

**Insert before line:**
```bash
ssed "insert # Header before line 1 in README.md"
```

**Insert after line:**
```bash
ssed "insert footer text after last line in file.txt"
```

**Insert before/after pattern:**
```bash
ssed "insert TODO: before lines containing function in code.js"
ssed "insert ; after lines ending with statement in script.js"
```

### 4. Viewing/Extracting Content

**Show specific lines:**
```bash
ssed "show line 10 from data.csv"
ssed "show lines 5 to 15 from log.txt"
ssed "show first 20 lines from big-file.txt"
ssed "show last 10 lines from output.log"
```

**Show lines matching pattern:**
```bash
ssed "show lines containing error from app.log"
ssed "show lines starting with TODO from source.js"
ssed "extract email addresses from contacts.txt"
```

**Show with line numbers:**
```bash
ssed "show lines containing warning from log.txt with line numbers"
```

---

## Working with Patterns

### 5. Pattern-Based Operations

**Predefined patterns:**
```bash
ssed "extract all email addresses from document.txt"
ssed "find all URLs in webpage.html"
ssed "show lines with phone numbers from directory.txt"
ssed "remove all email addresses from privacy.txt"
```

**Natural language patterns:**
```bash
ssed "show words starting with pre from text.txt"
ssed "delete lines ending with semicolon from data.csv"
ssed "extract text between START and END from file.txt"
ssed "replace words containing test with REDACTED in log.txt"
```

**Multiple patterns:**
```bash
ssed "delete lines containing debug or test from code.js"
ssed "show lines matching error or warning from app.log"
```

---

## File Operations

### 6. In-Place Editing

**Modify file directly:**
```bash
ssed "replace old with new in config.txt modify file"
ssed "delete empty lines from data.csv in-place"
```

**With backup:**
```bash
ssed "replace all foo with bar in important.txt with backup"
ssed "delete lines containing test from code.js modify file create backup"
```

**Save to new file:**
```bash
ssed "remove comments from code.js save to clean-code.js"
ssed "extract errors from app.log save to errors.txt"
```

### 7. Multiple Files

**Process multiple files:**
```bash
ssed "replace all TODO with DONE in all .txt files"
ssed "delete empty lines in all .log files"
ssed "add header to all .md files in docs folder"
```

**With pattern matching:**
```bash
ssed "remove trailing spaces in all files matching test*.js"
ssed "uppercase first word in all .txt files in current directory"
```

---

## Advanced Operations

### 8. Range Operations

**Replace in range:**
```bash
ssed "replace foo with bar in lines 10 to 50 of file.txt"
```

**Delete range:**
```bash
ssed "delete lines between START and END in document.txt"
```

**Extract range:**
```bash
ssed "show lines from first TODO to next DONE in tasks.txt"
```

**Conditional range:**
```bash
ssed "delete lines from line 5 onwards that contain debug"
ssed "replace old with new until line 100 in file.txt"
```

### 9. Text Transformation

**Case conversion:**
```bash
ssed "convert all text to uppercase in file.txt"
ssed "make first word of each line uppercase in document.txt"
ssed "convert to lowercase in lines containing HEADER"
ssed "convert to title case in file.txt"
```

**Format conversion:**
```bash
ssed "convert tabs to spaces in code.py"
ssed "convert spaces to tabs in Makefile"
ssed "convert DOS line endings to Unix in script.sh"
```

**Text manipulation:**
```bash
ssed "reverse word order in each line of file.txt"
ssed "sort lines alphabetically in names.txt"
ssed "remove duplicate consecutive lines from data.txt"
```

### 10. Conditional Operations

**If-then operations:**
```bash
ssed "if line contains error then make it uppercase in log.txt"
ssed "replace old with new only in lines starting with # in config.txt"
ssed "delete line only if it starts with // and contains TODO"
```

**Complex conditions:**
```bash
ssed "replace foo with bar in lines containing debug but not in lines with production"
ssed "uppercase lines with numbers except in first 5 lines"
```

### 11. Compound Operations

**Multiple operations:**
```bash
ssed "delete empty lines and trim trailing spaces in file.txt"
ssed "remove comments, delete blank lines, and trim whitespace in code.js"
ssed "replace foo with bar then sort lines then remove duplicates"
```

**Sequential operations:**
```bash
ssed "replace all TODO with DONE then show lines with DONE"
ssed "delete lines containing test, then number remaining lines"
```

---

## Real-World Scenarios

### 12. Code Cleanup

**Remove debug statements:**
```bash
ssed "delete lines containing console.log from all .js files"
ssed "remove lines starting with // DEBUG from code.cpp"
```

**Format code:**
```bash
ssed "trim trailing whitespace in all .py files"
ssed "replace all tabs with 4 spaces in source.java"
```

**Update imports:**
```bash
ssed "replace old-package with new-package in all .go files"
```

### 13. Log Processing

**Extract errors:**
```bash
ssed "show lines containing ERROR from app.log save to errors.txt"
ssed "extract lines with status 500 from access.log"
```

**Filter logs:**
```bash
ssed "show lines between 2024-01-01 and 2024-01-31 from server.log"
ssed "delete lines not containing ERROR or WARNING from debug.log"
```

**Analyze logs:**
```bash
ssed "count lines containing timeout in application.log"
ssed "show unique URLs from access.log"
```

### 14. Configuration Management

**Update config values:**
```bash
ssed "replace port=8080 with port=3000 in config.ini"
ssed "change database_host to localhost in settings.conf"
```

**Comment/uncomment:**
```bash
ssed "add # at beginning of lines containing debug_mode in app.conf"
ssed "remove # from lines starting with # production in config.yaml"
```

### 15. Data Processing

**CSV manipulation:**
```bash
ssed "delete first line from data.csv"  # Remove header
ssed "show column 3 from each line in data.csv"
ssed "replace , with | in data.csv"  # Change delimiter
```

**Extract data:**
```bash
ssed "extract all email addresses from customer-list.txt save to emails.txt"
ssed "show lines with phone numbers matching format XXX-XXX-XXXX"
```

**Clean data:**
```bash
ssed "remove duplicate lines from list.txt"
ssed "trim whitespace and remove empty lines from data.txt"
ssed "standardize dates to YYYY-MM-DD format in records.csv"
```

### 16. Documentation Tasks

**Update version numbers:**
```bash
ssed "replace version 1.0 with version 2.0 in all .md files"
ssed "change v1.2.3 to v1.2.4 in CHANGELOG.md"
```

**Add headers:**
```bash
ssed "insert # Title at beginning of README.md"
ssed "add --- after first line in all .md files"
```

**Fix formatting:**
```bash
ssed "replace all TODO: with - [ ] in tasks.md"
ssed "convert all DONE to - [x] in checklist.md"
```

### 17. Batch Renaming Content

**Update references:**
```bash
ssed "replace all oldFileName with newFileName in all .html files"
ssed "change every occurrence of OldClass to NewClass in all .java files"
```

### 18. Privacy/Security

**Redact sensitive data:**
```bash
ssed "replace all email addresses with REDACTED in public-log.txt"
ssed "remove all phone numbers from shared-document.txt"
ssed "replace IP addresses with XXX.XXX.XXX.XXX in access-log.txt"
```

**Remove credentials:**
```bash
ssed "delete lines containing password= from config.txt"
ssed "remove lines with api_key from settings.ini"
```

---

## Interactive Mode Examples

### 19. Guided Workflow

```bash
$ ssed

ssed> What would you like to do?
(Type your request or press Tab for suggestions)

User: replace

ssed> What text do you want to replace?
User: old_function

ssed> What should it be replaced with?
User: new_function

ssed> Where should I look?
User: all .py files

ssed> Scanning 15 .py files...
Found 23 occurrences in 8 files:
  - utils.py: 5 matches
  - main.py: 3 matches
  - helpers.py: 7 matches
  ... (5 more files)

ssed> Preview changes? (yes/no/show-all)
User: yes

[Shows preview of first few changes]

ssed> Apply these changes? (yes/no/edit)
User: yes

ssed> ✓ Modified 8 files, 23 replacements made
     Created backups with .bak extension
```

### 20. Learning Mode

```bash
$ ssed

ssed> What would you like to do?
User: I want to clean up my log file

ssed> I can help with that! Common log cleanup tasks:
     1. Remove empty lines
     2. Delete debug messages
     3. Extract only errors
     4. Remove old entries
     5. Something else

User: 2

ssed> What text indicates a debug message?
User: lines with DEBUG

ssed> Should I:
     a) Delete entire lines containing "DEBUG"
     b) Just remove the word "DEBUG"
     c) Comment out lines with "DEBUG"

User: a

ssed> Which file?
User: app.log

ssed> [Preview]
Will delete 47 lines containing "DEBUG" from app.log

ssed> Proceed? (yes/no)
User: yes

ssed> ✓ Done! Deleted 47 lines from app.log
     Original saved as app.log.bak

ssed> The equivalent command is:
     ssed "delete lines containing DEBUG from app.log"

     Would you like to save this command? (yes/no)
```

---

## Advanced Pattern Examples

### 21. Email Extraction & Manipulation

```bash
# Extract emails
ssed "extract all email addresses from contacts.txt"

# Replace domains
ssed "replace @oldcompany.com with @newcompany.com in emails.txt"

# Redact emails
ssed "replace email addresses with [EMAIL] in public.txt"

# Show lines with specific email domain
ssed "show lines with email addresses ending with @gmail.com"
```

### 22. URL Manipulation

```bash
# Extract URLs
ssed "extract all URLs from document.html"

# Convert HTTP to HTTPS
ssed "replace http:// with https:// in all URLs"

# Remove URLs
ssed "delete all URLs from text.txt"

# Show lines with specific domain
ssed "show lines containing URLs with github.com"
```

### 23. Number Operations

```bash
# Extract numbers
ssed "extract all numbers from data.txt"

# Replace numbers
ssed "replace numbers with XXX in privacy-doc.txt"

# Show lines with numbers in range
ssed "show lines with numbers between 100 and 200"

# Format numbers
ssed "add commas to numbers over 1000 in report.txt"
```

---

## Error Handling Examples

### 24. Safe Operations

**Preview before applying:**
```bash
ssed "delete all lines from important.txt show preview"

# Output:
# ⚠ WARNING: This will delete ALL lines
# Preview:
# [Shows what would be deleted]
# Continue? (yes/no)
```

**Dry run:**
```bash
ssed "replace all old with new in config.txt dry run"

# Output:
# DRY RUN - No changes will be made
# Would modify 15 lines:
#   Line 3: "old value" → "new value"
#   Line 7: "old setting" → "new setting"
#   ...
```

**Validation:**
```bash
ssed "delete line 1000 from small-file.txt"

# Output:
# ⚠ Error: small-file.txt only has 50 lines
# Did you mean:
#   - delete line 50 (last line)
#   - delete last 10 lines
#   - delete all lines
```

---

## Performance & Optimization

### 25. Large Files

```bash
# Show progress for large files
ssed "replace all old with new in huge-log.txt show progress"

# Process first N lines only
ssed "replace foo with bar in first 1000 lines of big-file.txt"

# Stop after first match
ssed "show first line containing ERROR from gigantic.log"
```

---

## Integration Examples

### 26. Piping & Streaming

```bash
# From stdin
echo "hello world" | ssed "replace hello with hi"

# Pipeline
cat file.txt | ssed "delete empty lines" | ssed "trim whitespace"

# With other commands
grep ERROR app.log | ssed "extract timestamps"
```

### 27. Script Integration

```bash
#!/bin/bash
# cleanup-logs.sh

# Use ssed in a script
ssed "delete lines older than 7 days from /var/log/app.log in-place"
ssed "remove sensitive data from /tmp/export.txt"
ssed "archive errors to /backup/errors-$(date +%Y%m%d).txt"
```

---

## Tips & Best Practices

### Do's ✓
- Always preview changes to important files
- Use backups when modifying files in-place
- Start with "show" commands to understand your data
- Use descriptive patterns for clarity
- Test on sample data first

### Don'ts ✗
- Don't modify files without backup on first try
- Don't use vague patterns on multiple files
- Don't skip previewing destructive operations
- Don't forget to quote special characters
- Don't process files you don't understand

---

## Comparison with Traditional Tools

### vs. sed
```bash
# sed
sed -i.bak 's/foo/bar/g' file.txt

# ssed
ssed "replace all foo with bar in file.txt with backup"
```

### vs. awk
```bash
# awk
awk '/pattern/ {print $3}' file.txt

# ssed
ssed "show word 3 from lines containing pattern in file.txt"
```

### vs. grep
```bash
# grep
grep "pattern" file.txt

# ssed
ssed "show lines containing pattern from file.txt"
```

### vs. multiple tools
```bash
# Traditional
grep ERROR log.txt | sed 's/^/[ERROR] /' | sort | uniq

# ssed
ssed "show lines with ERROR from log.txt, add [ERROR] at start, sort and remove duplicates"
```

---

## Future Features (Coming Soon)

### Variables
```bash
ssed "save lines with ERROR as errors, then show count of errors"
```

### Macros
```bash
ssed "define cleanup as: delete empty lines, trim spaces, remove comments"
ssed "run cleanup on all .txt files"
```

### AI-Assisted
```bash
ssed "make this code more readable" --ai
ssed "fix formatting issues automatically" --ai
```

---

## Getting Help

### Built-in Help
```bash
ssed help                    # Show general help
ssed examples               # Show examples
ssed "help with patterns"   # Help on patterns
```

### Interactive Help
```bash
ssed> help
ssed> examples replace
ssed> what can I do with patterns?
```

---

## Quick Reference Card

```
COMMON OPERATIONS:
  replace "X" with "Y"              - Substitute text
  delete lines containing "X"       - Remove matching lines
  show lines X to Y                 - Display range
  insert "text" before line X       - Add content

PATTERNS:
  lines containing "X"              - Match pattern
  lines starting with "X"           - Match start
  email addresses                   - Predefined pattern
  words ending with "X"             - Word pattern

SCOPE:
  first/last/all                    - Occurrence
  in file.txt                       - Single file
  in all .txt files                 - Multiple files
  with line numbers                 - Show numbers

OPTIONS:
  show preview                      - Preview changes
  dry run                           - Test without changes
  with backup                       - Create backup
  ignore case                       - Case-insensitive
  modify file                       - Save changes
```
