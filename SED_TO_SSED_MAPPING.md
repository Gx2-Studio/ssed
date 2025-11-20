# Complete sed to ssed Mapping

This document maps every sed feature to its ssed natural language equivalent, ensuring feature parity.

## Basic Operations

### Substitution (s command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `s/old/new/` | `replace "old" with "new"` |
| `s/old/new/g` | `replace all "old" with "new"` |
| `s/old/new/2` | `replace 2nd "old" with "new"` |
| `s/old/new/p` | `replace "old" with "new" and show changed lines` |
| `s/old/new/i` | `replace "old" with "new" case-insensitive` |
| `s/old/new/gi` | `replace all "old" with "new" ignore case` |
| `s/old/new/w file` | `replace "old" with "new" and save matches to file` |

### Deletion (d command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `d` | `delete all lines` |
| `/pattern/d` | `delete lines containing "pattern"` |
| `5d` | `delete line 5` |
| `5,10d` | `delete lines 5 to 10` |
| `$d` | `delete last line` |
| `1,5d` | `delete first 5 lines` |
| `/^$/d` | `delete empty lines` |
| `/pattern/!d` | `delete lines not containing "pattern"` |

### Print (p command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `p` | `show all lines` (duplicate each) |
| `-n '5p'` | `show line 5` |
| `-n '5,10p'` | `show lines 5 to 10` |
| `-n '/pattern/p'` | `show lines containing "pattern"` |
| `-n '$p'` | `show last line` |
| `-n '/start/,/end/p'` | `show lines between "start" and "end"` |

### Append (a command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `5a\text` | `insert "text" after line 5` |
| `/pattern/a\text` | `insert "text" after lines containing "pattern"` |
| `$a\text` | `insert "text" at end` |

### Insert (i command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `5i\text` | `insert "text" before line 5` |
| `/pattern/i\text` | `insert "text" before lines containing "pattern"` |
| `1i\text` | `insert "text" at beginning` |

### Change (c command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `5c\text` | `replace line 5 with "text"` |
| `5,10c\text` | `replace lines 5 to 10 with "text"` |
| `/pattern/c\text` | `replace lines containing "pattern" with "text"` |

---

## Addressing

### Line Numbers

| sed Address | ssed Natural Language |
|------------|----------------------|
| `5` | `line 5` |
| `$` | `last line` |
| `1` | `first line` or `line 1` |
| `5,10` | `lines 5 to 10` |
| `5,$` | `lines 5 to end` |
| `1,10` | `first 10 lines` |

### Address Steps

| sed Address | ssed Natural Language |
|------------|----------------------|
| `1~2` | `every other line starting from 1` or `odd lines` |
| `0~2` | `every other line starting from 0` or `even lines` |
| `1~3` | `every 3rd line` |

### Regular Expression Addresses

| sed Address | ssed Natural Language |
|------------|----------------------|
| `/pattern/` | `lines containing "pattern"` |
| `/^pattern/` | `lines starting with "pattern"` |
| `/pattern$/` | `lines ending with "pattern"` |
| `/pattern1/,/pattern2/` | `lines between "pattern1" and "pattern2"` |
| `/pattern/,5` | `from first match of "pattern" to line 5` |
| `5,/pattern/` | `from line 5 to first match of "pattern"` |

### Address Negation

| sed Address | ssed Natural Language |
|------------|----------------------|
| `5!d` | `delete all lines except line 5` |
| `/pattern/!d` | `delete lines not containing "pattern"` |
| `1,10!p` | `show lines except 1 to 10` |

---

## Advanced Commands

### Next (n command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `n` | `skip to next line` |
| `/pattern/{n; s/old/new/}` | `for lines containing "pattern", skip next line and replace "old" with "new"` |

### Next Append (N command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `N` | `append next line` |
| `N; s/\n/ /` | `join current line with next line` |
| `/pattern/{N; s/\n//}` | `join lines starting with "pattern" with next line` |

### Line Number (= command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `=` | `show line numbers` |
| `/pattern/=` | `show line numbers of lines containing "pattern"` |
| `-n '$='` | `count total lines` |

### List (l command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `l` | `show lines with special characters visible` |
| `/pattern/l` | `show lines containing "pattern" with special characters` |

---

## Hold Space Commands

### Hold (h command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `h` | `save current line to memory` |
| `1h` | `save line 1 to memory` |
| `/pattern/h` | `save lines containing "pattern" to memory` |

### Hold Append (H command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `H` | `append current line to memory` |
| `/pattern/H` | `append lines containing "pattern" to memory` |

### Get (g command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `g` | `replace current line with saved line` |
| `$g` | `replace last line with saved line` |

### Get Append (G command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `G` | `append saved line to current line` |
| `$G` | `append saved line to end` |

### Exchange (x command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `x` | `swap current line with saved line` |

---

## Flow Control

### Branch (b command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `b` | `skip to end of script` |
| `b label` | `jump to label` |
| `/pattern/b` | `skip remaining commands for lines containing "pattern"` |

### Test (t command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `t` | `if last substitution succeeded, skip to end` |
| `t label` | `if last substitution succeeded, jump to label` |
| `s/old/new/; t done` | `replace "old" with "new", if successful skip to done` |

### Label (: command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `:label` | `define label "label"` |

---

## File Operations

### Read File (r command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `r file.txt` | `insert contents of file.txt after each line` |
| `5r file.txt` | `insert contents of file.txt after line 5` |
| `/pattern/r file.txt` | `insert contents of file.txt after lines containing "pattern"` |

### Write File (w command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `w output.txt` | `write all lines to output.txt` |
| `5w output.txt` | `write line 5 to output.txt` |
| `/pattern/w output.txt` | `write lines containing "pattern" to output.txt` |

---

## Other Commands

### Quit (q command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `q` | `quit after current line` |
| `5q` | `quit after line 5` or `show first 5 lines` |
| `/pattern/q` | `quit at first line containing "pattern"` |
| `q 42` | `quit with exit code 42` |

### Quit Silent (Q command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `Q` | `quit immediately without printing` |
| `5Q` | `quit silently after line 5` |

### Transliterate (y command)

| sed Command | ssed Natural Language |
|------------|----------------------|
| `y/abc/ABC/` | `replace 'a' with 'A', 'b' with 'B', 'c' with 'C'` |
| `y/aeiou/AEIOU/` | `uppercase all vowels` |

### Group Commands ({ })

| sed Command | ssed Natural Language |
|------------|----------------------|
| `/pattern/{s/old/new/; d}` | `for lines containing "pattern": replace "old" with "new" then delete line` |
| `5,10{s/^/> /}` | `for lines 5 to 10: add "> " at beginning` |

---

## Command-Line Options

### Basic Options

| sed Option | ssed Natural Language |
|-----------|----------------------|
| `-n` | (default behavior: show only explicit output) |
| `-e 'script'` | (just write the natural language query) |
| `-f file` | `run commands from file` |
| `-i` | `modify file in-place` or `edit file` |
| `-i.bak` | `modify file with backup` |

### Extended Regex

| sed Option | ssed Natural Language |
|-----------|----------------------|
| `-r` or `-E` | (automatically detected when needed) |

### Other Options

| sed Option | ssed Natural Language |
|-----------|----------------------|
| `-s` | `treat each file separately` |
| `-u` | `unbuffered mode` |
| `-z` | `use null character as line separator` |

---

## Regular Expression Features

### Basic Regex

| sed Pattern | ssed Natural Language |
|------------|----------------------|
| `.` | `any character` |
| `*` | `zero or more` |
| `^` | `starting with` or `at beginning` |
| `$` | `ending with` or `at end` |
| `[abc]` | `characters a, b, or c` |
| `[^abc]` | `any character except a, b, c` |
| `[a-z]` | `any lowercase letter` |
| `\(pattern\)` | `capture "pattern"` |
| `\1, \2, etc.` | `use captured group 1, 2, etc.` |

### Extended Regex (with -r)

| sed Pattern | ssed Natural Language |
|------------|----------------------|
| `+` | `one or more` |
| `?` | `zero or one` |
| `{n}` | `exactly n times` |
| `{n,m}` | `between n and m times` |
| `(pattern)` | `capture "pattern"` |
| `\|` or `\|` | `or` |

### GNU Extensions

| sed Pattern | ssed Natural Language |
|------------|----------------------|
| `\s` | `whitespace` |
| `\S` | `non-whitespace` |
| `\w` | `word character` |
| `\W` | `non-word character` |
| `\b` | `word boundary` |
| `\<` | `start of word` |
| `\>` | `end of word` |

### ssed Predefined Patterns

Instead of regex, ssed offers predefined patterns:

| Pattern Type | ssed Natural Language |
|-------------|----------------------|
| Email | `email addresses` |
| URL | `URLs` or `links` |
| Number | `numbers` |
| Date | `dates` |
| Phone | `phone numbers` |
| IP Address | `IP addresses` |
| UUID | `UUIDs` |
| Hex Color | `hex colors` |

---

## Replacement Features

### Special Replacement

| sed Replacement | ssed Natural Language |
|----------------|----------------------|
| `&` | `the matched text` |
| `\1, \2, etc.` | (automatically captured in natural language) |
| `\L...\E` | `convert to lowercase` |
| `\U...\E` | `convert to uppercase` |
| `\l` | `lowercase next character` |
| `\u` | `uppercase next character` |

### ssed Transformations

| Operation | ssed Natural Language |
|----------|----------------------|
| Uppercase | `convert to uppercase` or `make uppercase` |
| Lowercase | `convert to lowercase` or `make lowercase` |
| Title Case | `convert to title case` or `capitalize words` |
| Capitalize | `capitalize first letter` |
| Reverse | `reverse text` or `reverse order` |

---

## Complex sed Scripts → ssed

### Remove Duplicate Lines

```bash
# sed
sed '$!N; /^\(.*\)\n\1$/!P; D'

# ssed
remove duplicate consecutive lines
```

### Reverse File

```bash
# sed
sed '1!G;h;$!d'

# ssed
reverse all lines
```

### Print Last 10 Lines

```bash
# sed
sed -n ':a;N;$!ba;$s/.*//;x;p'

# ssed
show last 10 lines
```

### Double-Space File

```bash
# sed
sed G

# ssed
add blank line after each line
```

### Number Lines

```bash
# sed
sed = file | sed 'N; s/\n/\t/'

# ssed
show line numbers
```

### Remove HTML Tags

```bash
# sed
sed 's/<[^>]*>//g'

# ssed
remove HTML tags
```

### Convert DOS to Unix

```bash
# sed
sed 's/\r$//'

# ssed
remove carriage returns
# or
convert DOS line endings to Unix
```

### Trim Whitespace

```bash
# sed (leading)
sed 's/^[ \t]*//'

# sed (trailing)
sed 's/[ \t]*$//'

# sed (both)
sed 's/^[ \t]*//; s/[ \t]*$//'

# ssed
remove leading whitespace
remove trailing whitespace
trim whitespace
```

### Extract Email Addresses

```bash
# sed (complex regex)
sed -n 's/.*\([a-zA-Z0-9._-]*@[a-zA-Z0-9._-]*\).*/\1/p'

# ssed
extract email addresses
```

### Comment Out Lines

```bash
# sed
sed 's/^/# /'

# ssed
add "# " at beginning of each line
# or
comment out all lines with "#"
```

### Uncomment Lines

```bash
# sed
sed 's/^# //'

# ssed
remove "# " from beginning of lines
# or
uncomment lines
```

### Join Lines

```bash
# sed
sed ':a;N;$!ba;s/\n/ /g'

# ssed
join all lines with space
```

### Swap Two Words

```bash
# sed
sed 's/\(word1\) \(word2\)/\2 \1/'

# ssed
swap "word1" and "word2"
```

### Delete Lines in Range

```bash
# sed
sed '/START/,/END/d'

# ssed
delete lines between "START" and "END"
```

### Insert Header/Footer

```bash
# sed
sed '1i\HEADER TEXT'
sed '$a\FOOTER TEXT'

# ssed
insert "HEADER TEXT" at beginning
insert "FOOTER TEXT" at end
```

### Replace in Range

```bash
# sed
sed '10,20s/old/new/g'

# ssed
replace all "old" with "new" in lines 10 to 20
```

### Conditional Replace

```bash
# sed
sed '/pattern/s/old/new/'

# ssed
replace "old" with "new" only in lines containing "pattern"
```

### Delete Trailing Blank Lines

```bash
# sed
sed -e :a -e '/^\n*$/{$d;N;ba' -e '}'

# ssed
remove trailing blank lines
```

### Delete Leading Blank Lines

```bash
# sed
sed '/./,$!d'

# ssed
remove leading blank lines
```

### Extract Lines with Pattern

```bash
# sed
sed -n '/pattern/p'

# ssed
show lines containing "pattern"
```

### Count Matches

```bash
# sed
sed -n '/pattern/p' | wc -l

# ssed
count lines containing "pattern"
```

---

## Complete Feature Coverage Matrix

| sed Feature | ssed Equivalent | Implementation Priority |
|------------|-----------------|------------------------|
| Basic substitution | ✅ Natural language replace | Phase 1 |
| Global substitution | ✅ "replace all" | Phase 1 |
| Line deletion | ✅ "delete line/lines" | Phase 1 |
| Line printing | ✅ "show line/lines" | Phase 1 |
| Insert/append | ✅ "insert before/after" | Phase 1 |
| Line ranges | ✅ "lines X to Y" | Phase 1 |
| Pattern matching | ✅ "lines containing" | Phase 2 |
| Case-insensitive | ✅ "ignore case" | Phase 2 |
| In-place editing | ✅ "modify file" | Phase 2 |
| Backup creation | ✅ "with backup" | Phase 2 |
| Multiple files | ✅ "in all files" | Phase 2 |
| Next line (n) | ✅ "skip to next" | Phase 3 |
| Append next (N) | ✅ "join with next line" | Phase 3 |
| Hold space (h/H/g/G/x) | ✅ "save to memory" | Phase 3 |
| Branching (b) | ✅ Conditional operations | Phase 3 |
| Test (t) | ✅ "if successful then" | Phase 3 |
| File I/O (r/w) | ✅ "insert from file" | Phase 3 |
| Quit (q/Q) | ✅ "stop at line" | Phase 2 |
| Transliterate (y) | ✅ Transform operations | Phase 3 |
| Line numbers (=) | ✅ "show line numbers" | Phase 2 |
| List (l) | ✅ "show special chars" | Phase 3 |
| Extended regex | ✅ Predefined patterns | Phase 2 |
| Capture groups | ✅ Smart capture | Phase 4 |
| Command grouping | ✅ "and then" | Phase 3 |
| Address negation | ✅ "except" / "not" | Phase 2 |
| Step addresses | ✅ "every nth line" | Phase 3 |
| Range patterns | ✅ "between X and Y" | Phase 2 |

---

## Design Notes

### Philosophy Differences

**sed**: Terse, cryptic, powerful
- Optimized for expert users
- Minimal typing
- High learning curve

**ssed**: Verbose, clear, accessible
- Optimized for first-time users
- Self-documenting
- Zero learning curve

### When to Use Which

**Use sed when**:
- You're an expert and need speed
- Writing complex automation scripts
- Working in resource-constrained environments
- Need exact POSIX compliance

**Use ssed when**:
- Learning text processing
- Occasional text manipulation
- Prefer clarity over brevity
- Want to understand old commands
- Working with non-technical users

### Compatibility Mode

ssed should also accept sed syntax for compatibility:
```bash
ssed 's/old/new/g' file.txt  # Works like sed
ssed "replace all old with new in file.txt"  # Natural language
```

This allows gradual migration and use of ssed as a drop-in sed replacement.
