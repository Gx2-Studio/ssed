# sed Quick Reference

A concise reference for sed commands and syntax.

## Command Syntax
```
[address[,address]]command[arguments]
```

## Common Options
```
-n          Suppress automatic printing
-e SCRIPT   Add script
-f FILE     Read script from file
-i[.SUFFIX] Edit in-place (optional backup)
-r, -E      Extended regex
```

## Addresses
```
5           Line 5
5,10        Lines 5 to 10
$           Last line
/regex/     Lines matching regex
/start/,/end/  Range between patterns
1~2         Every other line (1,3,5...)
5!          All lines except 5
```

## Essential Commands

### Substitution
```
s/old/new/      Replace first match
s/old/new/g     Replace all matches
s/old/new/2     Replace 2nd match
s/old/new/i     Case-insensitive
s/old/new/p     Print if substituted
```

### Deletion
```
d               Delete line
D               Delete first line of pattern space
```

### Print
```
p               Print pattern space
P               Print first line of pattern space
l               List (show special chars)
=               Print line number
```

### Text Insertion
```
a\TEXT          Append text after line
i\TEXT          Insert text before line
c\TEXT          Change (replace) line
```

### Flow Control
```
:LABEL          Define label
b LABEL         Branch to label
t LABEL         Branch if substitution succeeded
T LABEL         Branch if substitution failed
q               Quit
Q               Quit without printing
```

### Next
```
n               Read next line (replace pattern space)
N               Append next line to pattern space
```

### Hold Space
```
h               Copy pattern to hold space
H               Append pattern to hold space
g               Copy hold space to pattern
G               Append hold space to pattern
x               Exchange pattern and hold spaces
```

### File I/O
```
r FILE          Read file and append
R FILE          Read one line from file
w FILE          Write pattern space to file
W FILE          Write first line to file
```

### Other
```
y/src/dst/      Transliterate characters
{}              Group commands
!               Negate address
#               Comment
```

## Regex Special Characters

### Basic (default)
```
.               Any character
*               Zero or more
^               Start of line
$               End of line
[abc]           Character class
[^abc]          Negated class
\(RE\)          Capture group (escaped)
\n              Backreference (n=1-9)
\{n,m\}         Repetition (escaped)
```

### Extended (-r or -E)
```
+               One or more
?               Zero or one
|               Alternation
(RE)            Capture group
{n,m}           Repetition
```

### GNU Extensions
```
\s \S           Whitespace / non-whitespace
\w \W           Word char / non-word char
\b \B           Word boundary / non-boundary
\< \>           Start/end of word
```

## Replacement Special Sequences
```
&               Matched string
\1..\9          Backreferences
\L..\E          Convert to lowercase
\U..\E          Convert to uppercase
\l              Lowercase next char
\u              Uppercase next char
```

## Common One-Liners

```bash
# Substitute
sed 's/foo/bar/g' file                    # Replace all foo with bar

# Delete
sed '/pattern/d' file                     # Delete matching lines
sed '5d' file                             # Delete line 5
sed '5,10d' file                          # Delete lines 5-10
sed '/^$/d' file                          # Delete empty lines

# Print
sed -n '5p' file                          # Print line 5
sed -n '5,10p' file                       # Print lines 5-10
sed -n '/pattern/p' file                  # Print matching lines

# Insert/Append
sed '5i\New line' file                    # Insert before line 5
sed '5a\New line' file                    # Append after line 5
sed '5c\Replacement' file                 # Replace line 5

# Multiple operations
sed -e 's/foo/bar/' -e 's/baz/qux/' file # Multiple substitutions
sed '/pattern/{s/foo/bar/; s/baz/qux/}' file # Group operations

# In-place editing
sed -i 's/foo/bar/g' file                 # Edit file in-place
sed -i.bak 's/foo/bar/g' file            # Edit with backup

# Line numbers
sed '=' file                              # Add line numbers
sed -n '$=' file                          # Count lines

# Extract sections
sed -n '/START/,/END/p' file              # Print between patterns

# Reverse file
sed '1!G;h;$!d' file                      # Reverse lines (tac)

# Remove duplicates
sed '$!N; /^\(.*\)\n\1$/!P; D' file      # Remove consecutive dupes

# Add spacing
sed G file                                # Double-space
sed 'G;G' file                            # Triple-space

# Remove spacing
sed '/^$/d' file                          # Remove blank lines
sed '/./,/^$/!d' file                    # Remove leading blanks

# Wrap lines
sed 's/.*/     &/' file                   # Indent by 5 spaces
```

## Examples with Explanation

### Basic Substitution
```bash
# Replace first occurrence on each line
sed 's/cat/dog/' pets.txt

# Replace all occurrences
sed 's/cat/dog/g' pets.txt

# Case-insensitive replacement
sed 's/cat/dog/gi' pets.txt
```

### Using Addresses
```bash
# Only line 5
sed '5s/cat/dog/' file.txt

# Lines 5-10
sed '5,10s/cat/dog/' file.txt

# Lines matching pattern
sed '/^#/d' file.txt  # Delete comment lines
```

### Using Backreferences
```bash
# Swap two words
sed 's/\([a-z]*\) \([a-z]*\)/\2 \1/' file.txt

# Add quotes around numbers
sed 's/\([0-9]*\)/"\1"/g' file.txt

# Extract email domain
sed 's/.*@\(.*\)/\1/' emails.txt
```

### Multi-line Operations
```bash
# Join lines ending with backslash
sed ':a; /\\$/N; s/\\\n//; ta' file.txt

# Append next line to lines matching pattern
sed '/pattern/{N;s/\n/ /}' file.txt

# Print paragraph containing pattern
sed '/./{H;$!d;};x;/pattern/!d' file.txt
```

### Using Hold Space
```bash
# Swap first and last lines
sed '1h;1d;$!H;$!d;G' file.txt

# Copy line 1 to end of file
sed '1h;$G' file.txt
```

## Tips

1. **Test First**: Use `-n` with `p` to preview changes
   ```bash
   sed -n 's/foo/bar/gp' file  # Show what would change
   ```

2. **Backup Files**: Always use `-i.bak` for in-place edits
   ```bash
   sed -i.bak 's/foo/bar/g' file  # Creates file.bak
   ```

3. **Escape Special Chars**: Use different delimiter if pattern contains `/`
   ```bash
   sed 's|/path/to/file|/new/path|' file  # Use | instead of /
   ```

4. **Debug Scripts**: Test each command separately
   ```bash
   sed 's/foo/bar/' file | sed 's/baz/qux/'  # Test in stages
   ```

5. **Quote Scripts**: Always quote sed scripts in shell
   ```bash
   sed 's/$HOME/~/' file  # Wrong - $HOME expanded by shell
   sed 's/$HOME/~/' file  # Correct
   ```
