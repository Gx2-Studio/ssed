# sed Feature Mapping

This document maps all the features and capabilities of GNU sed to serve as a reference for the ssed project.

## Overview

sed (Stream EDitor) is a non-interactive text editor that processes text line-by-line using pattern matching and text transformation commands.

## Core Concepts

### 1. Pattern Space and Hold Space
- **Pattern Space**: The buffer containing the current line being processed
- **Hold Space**: A separate auxiliary buffer for storing text temporarily
- These two spaces allow complex multi-line text manipulations

### 2. Addressing
sed commands can be applied to specific lines using addresses:
- **Line numbers**: `5d` (delete line 5)
- **Line ranges**: `5,10d` (delete lines 5-10)
- **Regular expressions**: `/pattern/d` (delete lines matching pattern)
- **Range with regex**: `/start/,/end/d` (delete from start to end pattern)
- **Last line**: `$` (e.g., `$d` deletes last line)
- **Address step**: `1~2` (every other line starting from 1)
- **Negation**: `5!d` (delete all lines except line 5)

### 3. Execution Flow
- Reads input line by line
- Applies all commands in the script to each line
- Automatically prints pattern space (unless -n option used)
- Repeats for next line

---

## Command-Line Options

### Basic Options
- `-n, --quiet, --silent`: Suppress automatic printing of pattern space
- `-e SCRIPT`: Add script to commands to be executed
- `-f FILE`: Add contents of script-file to commands
- `-i[SUFFIX]`: Edit files in-place (optional backup with SUFFIX)
- `-r, -E`: Use extended regular expressions
- `--posix`: Disable GNU extensions

### Output Control
- `-l N`: Specify line wrap length for `l` command
- `-s, --separate`: Treat files as separate rather than continuous stream
- `-u, --unbuffered`: Load minimal amounts of data and flush output buffers more often
- `-z, --null-data`: Separate lines by NUL characters instead of newlines

### Debugging/Information
- `--debug`: Annotate program execution
- `--help`: Display help message
- `--version`: Output version information

---

## sed Commands

### Substitution Command (s)
The most commonly used sed command.

**Syntax**: `s/REGEXP/REPLACEMENT/FLAGS`

**Flags**:
- `g`: Global replacement (all occurrences on line)
- `NUMBER`: Replace only NUMBERth occurrence
- `p`: Print pattern space if substitution made
- `w FILE`: Write pattern space to FILE if substitution made
- `i` or `I`: Case-insensitive matching
- `m` or `M`: Multi-line mode
- `e`: Execute pattern space as shell command and use output
- `c`: Confirm each substitution (interactive)

**Examples**:
```bash
s/foo/bar/          # Replace first 'foo' with 'bar'
s/foo/bar/g         # Replace all 'foo' with 'bar'
s/foo/bar/2         # Replace second 'foo' with 'bar'
s/foo/bar/gi        # Replace all 'foo' with 'bar' (case-insensitive)
```

**Special Characters in Replacement**:
- `&`: The matched string
- `\n`: Backreference to nth captured group (n=1-9)
- `\L`: Convert to lowercase until `\U` or `\E`
- `\U`: Convert to uppercase until `\L` or `\E`
- `\E`: Stop case conversion
- `\l`: Convert next character to lowercase
- `\u`: Convert next character to uppercase

### Text Output Commands

#### p (Print)
Print pattern space
```bash
/pattern/p          # Print lines matching pattern
5p                  # Print line 5
```

#### l (List)
Print pattern space in visually unambiguous form (show non-printing chars)
```bash
l                   # List current line with special chars visible
l 40                # List with line wrap at 40 chars
```

#### = (Line Number)
Print current line number
```bash
=                   # Print line number
/pattern/=          # Print line numbers of matching lines
```

### Deletion Commands

#### d (Delete)
Delete pattern space and start next cycle
```bash
d                   # Delete line
/pattern/d          # Delete lines matching pattern
5,10d               # Delete lines 5-10
```

#### D (Delete First Line)
Delete first line of multi-line pattern space and restart cycle
```bash
D                   # Delete up to first newline
```

### Text Insertion Commands

#### a (Append)
Append text after current line
```bash
a\
TEXT                # Append TEXT after line
/pattern/a\
TEXT                # Append TEXT after lines matching pattern
```

#### i (Insert)
Insert text before current line
```bash
i\
TEXT                # Insert TEXT before line
```

#### c (Change)
Replace entire line(s) with text
```bash
c\
TEXT                # Replace line with TEXT
5,10c\
TEXT                # Replace lines 5-10 with TEXT
```

### Next Commands

#### n (Next)
Read next line into pattern space (replacing current)
```bash
n                   # Process next line
```

#### N (Next Append)
Append next line to pattern space (with newline separator)
```bash
N                   # Append next line to current pattern space
```

### File I/O Commands

#### r (Read)
Read file and append to pattern space
```bash
r FILE              # Read FILE and append after line
/pattern/r FILE     # Read FILE after lines matching pattern
```

#### R (Read Line)
Read single line from file
```bash
R FILE              # Read one line from FILE
```

#### w (Write)
Write pattern space to file
```bash
w FILE              # Write current line to FILE
/pattern/w FILE     # Write matching lines to FILE
```

#### W (Write First Line)
Write first line of pattern space to file
```bash
W FILE              # Write first line to FILE
```

### Hold Space Commands

#### h (Hold)
Copy pattern space to hold space (replacing)
```bash
h                   # Copy pattern space to hold space
```

#### H (Hold Append)
Append pattern space to hold space (with newline)
```bash
H                   # Append pattern space to hold space
```

#### g (Get)
Copy hold space to pattern space (replacing)
```bash
g                   # Copy hold space to pattern space
```

#### G (Get Append)
Append hold space to pattern space (with newline)
```bash
G                   # Append hold space to pattern space
```

#### x (Exchange)
Exchange pattern space and hold space
```bash
x                   # Swap pattern and hold spaces
```

### Flow Control Commands

#### b (Branch)
Branch to label or end of script
```bash
b                   # Branch to end of script
b LABEL             # Branch to :LABEL
/pattern/b          # Skip remaining commands if pattern matches
```

#### t (Test)
Branch if substitution was made since last line read or test
```bash
t LABEL             # Branch to LABEL if substitution succeeded
s/foo/bar/; t done  # Branch to :done if substitution worked
```

#### T (Test Negative)
Branch if NO substitution was made
```bash
T LABEL             # Branch to LABEL if no substitution
```

#### : (Label)
Define a label for branching
```bash
:LABEL              # Define label named LABEL
```

### Transformation Command

#### y (Transliterate)
Character-by-character translation (like tr)
```bash
y/abc/ABC/          # Translate a→A, b→B, c→C
y/aeiou/AEIOU/      # Uppercase vowels
```

### Other Commands

#### q (Quit)
Exit sed immediately
```bash
q                   # Quit after current line
5q                  # Quit after line 5
q 42                # Quit with exit code 42
```

#### Q (Quit Silently)
Quit immediately without printing pattern space
```bash
Q                   # Quit without printing
Q 42                # Quit with exit code 42
```

#### # (Comment)
Comment line (ignored)
```bash
# This is a comment
```

#### { } (Group)
Group multiple commands
```bash
/pattern/{          # Apply multiple commands to pattern
  s/foo/bar/
  p
}
```

#### ! (Negate)
Negate address selection
```bash
/pattern/!d         # Delete lines NOT matching pattern
5,10!p              # Print lines NOT in range 5-10
```

---

## Regular Expression Features

### Basic Regular Expression (BRE) - Default
- `.`: Any single character
- `*`: Zero or more of preceding
- `^`: Start of line
- `$`: End of line
- `[...]`: Character class
- `[^...]`: Negated character class
- `\(RE\)`: Capturing group (escaped parens)
- `\n`: Backreference (n=1-9)
- `\{n,m\}`: Repetition (escaped braces)

### Extended Regular Expression (ERE) - With -r/-E
- `+`: One or more of preceding
- `?`: Zero or one of preceding
- `|`: Alternation
- `(RE)`: Capturing group (no escape needed)
- `{n,m}`: Repetition (no escape needed)

### GNU Extensions
- `\s`: Whitespace
- `\S`: Non-whitespace
- `\w`: Word character [a-zA-Z0-9_]
- `\W`: Non-word character
- `\b`: Word boundary
- `\B`: Non-word boundary
- `\<`: Start of word
- `\>`: End of word
- `\``: Start of pattern space
- `\'`: End of pattern space

---

## Advanced Features

### Multi-line Processing
- Use `N` to append next line to pattern space
- Use `D` to delete first line and restart
- Use `P` to print first line only
- Enables processing multiple lines as a unit

### In-place Editing
```bash
sed -i.bak 's/foo/bar/g' file.txt    # Edit file, save backup as file.txt.bak
sed -i 's/foo/bar/g' file.txt        # Edit file without backup
```

### Multiple Scripts
```bash
sed -e 's/foo/bar/' -e 's/baz/qux/' file.txt
sed -f script1.sed -f script2.sed file.txt
```

### Range Contexts
```bash
/START/,/END/ { commands }           # Apply commands between patterns
/START/,+5 { commands }              # Apply to matching line and next 5
1~3 { commands }                     # Apply to every 3rd line starting at 1
```

### Address Ranges with Step
```bash
1~2d                # Delete odd lines (1, 3, 5, ...)
0~2d                # Delete even lines (2, 4, 6, ...)
```

### Zero-Address Commands
Some commands don't operate on addresses:
- `:` (label)
- `#` (comment)
- `}` (closing brace)

### Escape Sequences
In text commands (a, i, c), pattern space, and replacement strings:
- `\a`: Alert (bell)
- `\f`: Form feed
- `\n`: Newline
- `\r`: Carriage return
- `\t`: Tab
- `\v`: Vertical tab
- `\oNNN`: Octal character
- `\xHH`: Hexadecimal character

---

## Common Use Cases

### 1. Simple Substitution
```bash
sed 's/old/new/' file.txt              # Replace first occurrence per line
sed 's/old/new/g' file.txt             # Replace all occurrences
```

### 2. Delete Lines
```bash
sed '5d' file.txt                      # Delete line 5
sed '/pattern/d' file.txt              # Delete lines matching pattern
sed '5,10d' file.txt                   # Delete lines 5-10
```

### 3. Insert/Append Text
```bash
sed '5i\New Line' file.txt             # Insert before line 5
sed '5a\New Line' file.txt             # Append after line 5
```

### 4. Print Specific Lines
```bash
sed -n '5p' file.txt                   # Print only line 5
sed -n '/pattern/p' file.txt           # Print matching lines
sed -n '5,10p' file.txt                # Print lines 5-10
```

### 5. Multiple Edits
```bash
sed -e 's/foo/bar/' -e 's/baz/qux/' file.txt
```

### 6. In-place Editing
```bash
sed -i 's/old/new/g' *.txt             # Edit all .txt files
```

### 7. Complex Pattern Matching
```bash
sed '/start/,/end/d' file.txt          # Delete from start to end pattern
sed '1,/pattern/d' file.txt            # Delete from line 1 to pattern
```

### 8. Swap Lines
```bash
sed -n '1!G;h;$p' file.txt             # Reverse file (tac)
```

### 9. Remove Blank Lines
```bash
sed '/^$/d' file.txt                   # Delete empty lines
sed '/^\s*$/d' file.txt                # Delete blank lines (with spaces)
```

### 10. Number Lines
```bash
sed '=' file.txt | sed 'N;s/\n/\t/'    # Add line numbers
```

---

## Performance Considerations

1. **Address Ranges**: More specific addresses reduce processing
2. **Early Quit**: Use `q` to exit early when possible
3. **Minimal Pattern Space**: Avoid unnecessary multi-line operations
4. **Command Order**: Most frequent operations first
5. **Compiled Scripts**: Use `-f` for complex/reusable scripts

---

## Limitations & Edge Cases

1. **Line-oriented**: Best for line-based operations
2. **Limited Arithmetic**: No built-in arithmetic operations
3. **No Variables**: Only pattern/hold space for storage
4. **Regex Dialect**: BRE by default (verbose escaping)
5. **Backslash Hell**: Escaping can get complex
6. **Memory**: Entire pattern space held in memory
7. **Binary Files**: Not designed for binary data
8. **Portability**: GNU extensions not available in all sed versions

---

## Platform Variations

### GNU sed (Linux)
- Most feature-rich
- Extended regex with -r or -E
- In-place editing with -i
- Many GNU extensions

### BSD sed (macOS)
- Requires argument for -i: `-i ''` or `-i .bak`
- Different regex behavior
- Fewer extensions
- Some syntax differences

### POSIX sed
- Most portable
- Basic features only
- Limited regex support
- No GNU extensions

---

## References

- GNU sed Manual: https://www.gnu.org/software/sed/manual/sed.html
- POSIX sed Specification: https://pubs.opengroup.org/onlinepubs/9699919799/
- sed One-Liners: http://sed.sourceforge.net/sed1line.txt
