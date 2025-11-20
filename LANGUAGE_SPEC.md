# ssed Natural Language Specification

A human-friendly text transformation language that maps natural English to sed-like operations.

## Philosophy

**"No syntax to memorize, just describe what you want"**

ssed should be usable by anyone without reading documentation. Users express their intent in plain English, and ssed figures out what to do.

---

## Core Design Principles

1. **Natural Language First**: Queries should read like English sentences
2. **Forgiving Parser**: Accept multiple ways to express the same intent
3. **Interactive Clarification**: When ambiguous, ask the user
4. **Safe Defaults**: Non-destructive by default, require explicit confirmation for file changes
5. **Progressive Disclosure**: Start simple, reveal complexity only when needed

---

## Language Structure

Every ssed query has three main components:

```
[ACTION] [TARGET] [CONTEXT] [OPTIONS]
```

### Examples
```
replace "foo" with "bar" in file.txt
delete lines containing "error" from log.txt
show lines matching "TODO" in all .js files
insert "# Header" at the beginning of README.md
```

---

## 1. ACTIONS (Verbs)

### 1.1 Replace/Substitute
**Intent**: Change text from one form to another

**Variations**:
```
replace [pattern] with [text]
substitute [pattern] with [text]
change [pattern] to [text]
swap [pattern] for [text]
rename [pattern] to [text]
```

**Examples**:
```
replace "foo" with "bar"
replace all "color" with "colour"
replace first "hello" with "hi"
replace word "cat" with "dog"
replace line "old line" with "new line"
```

### 1.2 Delete/Remove
**Intent**: Remove text

**Variations**:
```
delete [target]
remove [target]
drop [target]
erase [target]
clear [target]
```

**Examples**:
```
delete line 5
delete lines containing "debug"
remove empty lines
remove trailing spaces
delete first 10 lines
delete from line 5 to line 10
```

### 1.3 Insert/Add
**Intent**: Add new text

**Variations**:
```
insert [text] [position]
add [text] [position]
append [text] [position]
prepend [text] [position]
```

**Examples**:
```
insert "new line" before line 5
add "footer" at the end
append "text" after line 10
prepend "header" at the beginning
insert "comment" before lines matching "function"
```

### 1.4 Extract/Show/Print
**Intent**: Display specific content (non-modifying)

**Variations**:
```
show [target]
print [target]
display [target]
extract [target]
get [target]
find [target]
list [target]
```

**Examples**:
```
show line 5
print lines containing "error"
display first 10 lines
extract lines matching "TODO"
find all email addresses
list line numbers of "function"
```

### 1.5 Transform
**Intent**: Convert text format/case

**Variations**:
```
convert [target] to [format]
transform [target] to [format]
make [target] [format]
```

**Examples**:
```
convert to uppercase
make line 5 lowercase
transform all words to title case
convert spaces to tabs
make URLs lowercase
```

### 1.6 Count
**Intent**: Count occurrences

**Variations**:
```
count [target]
number of [target]
how many [target]
```

**Examples**:
```
count lines
count words containing "test"
how many lines have "error"
number of blank lines
```

### 1.7 Duplicate
**Intent**: Repeat content

**Variations**:
```
duplicate [target]
repeat [target]
double [target]
```

**Examples**:
```
duplicate line 5
repeat lines matching "header"
double all blank lines
```

### 1.8 Move/Reorder
**Intent**: Change position of content

**Variations**:
```
move [target] to [position]
swap [target1] with [target2]
reverse [target]
sort [target]
```

**Examples**:
```
move line 5 to end
swap lines 1 and 10
reverse all lines
sort lines alphabetically
reverse word order in each line
```

---

## 2. TARGETS (What to Operate On)

### 2.1 Literal Text
Exact string matching

**Patterns**:
```
"exact text"
'exact text'
```

**Examples**:
```
replace "hello" with "hi"
delete "TODO:"
```

### 2.2 Patterns/Regex
Pattern matching (simplified regex)

**Natural Language Patterns**:
```
words starting with [text]
words ending with [text]
words containing [text]
text matching [pattern]
text between [start] and [end]
email addresses
URLs
phone numbers
numbers
dates
```

**Examples**:
```
replace words starting with "pre" with "post"
delete lines containing numbers
extract all email addresses
find URLs in the file
```

**Advanced Regex** (for power users):
```
pattern /regex/
regex /[a-z]+/
matching /\d{3}-\d{4}/
```

### 2.3 Line Specifications

**Single Line**:
```
line [number]
line [number] (one-indexed)
first line
last line
```

**Line Ranges**:
```
lines [start] to [end]
lines [start] through [end]
lines [start]-[end]
first [n] lines
last [n] lines
```

**Line Conditions**:
```
lines containing [pattern]
lines matching [pattern]
lines starting with [pattern]
lines ending with [pattern]
empty lines
blank lines
non-empty lines
lines with [condition]
```

**Examples**:
```
delete line 5
show lines 10 to 20
remove first 5 lines
extract lines containing "error"
delete blank lines
```

### 2.4 Word/Character Specifications

**Words**:
```
word [number]
first word
last word
all words
words [condition]
```

**Characters**:
```
character [number]
first character
last character
characters [start] to [end]
```

**Examples**:
```
capitalize first word
delete last character
extract words starting with "test"
```

### 2.5 Structural Elements

```
beginning/start
end
header
footer
paragraph [number]
section [number]
```

---

## 3. CONTEXT (Where/When to Apply)

### 3.1 File/Stream Context

**Single File**:
```
in [filename]
from [filename]
in file [filename]
```

**Multiple Files**:
```
in all [pattern] files
in all files matching [pattern]
in *.txt
in all .js files
in files in [directory]
```

**Stdin/Stdout**:
```
from stdin
to stdout
from pipe
```

**Examples**:
```
replace "foo" with "bar" in file.txt
delete empty lines in all .log files
show line 5 from data.csv
```

### 3.2 Scope Modifiers

**Occurrence Scope**:
```
first [occurrence]
last [occurrence]
all [occurrences]
[nth] occurrence
every [nth]
```

**Conditional Scope**:
```
only in lines [condition]
except in lines [condition]
but not in [context]
```

**Examples**:
```
replace first "hello" with "hi"
delete all lines containing "debug"
replace "foo" with "bar" only in lines starting with "#"
remove trailing spaces except in code blocks
```

### 3.3 Range Context

```
in range [start] to [end]
between [pattern1] and [pattern2]
from [pattern] onwards
until [pattern]
before [pattern]
after [pattern]
```

**Examples**:
```
delete lines from line 5 to line 10
extract text between "START" and "END"
uppercase everything after line 20
remove comments between "/*" and "*/"
```

---

## 4. OPTIONS (How to Apply)

### 4.1 Matching Options

**Case Sensitivity**:
```
case-sensitive
case-insensitive
ignore case
```

**Word Boundaries**:
```
whole word
whole words only
exact match
partial match
```

**Examples**:
```
replace "test" with "TEST" case-sensitive
delete lines containing "error" ignore case
replace whole word "cat" with "dog"
```

### 4.2 Output Options

**Preview/Execute**:
```
show preview
dry run
show what would change
actually do it
for real
save changes
```

**Output Format**:
```
show line numbers
show matching context
show only matches
show count only
highlight matches
```

**Examples**:
```
replace "foo" with "bar" show preview
delete empty lines dry run
find "TODO" show line numbers
```

### 4.3 File Modification Options

**Backup**:
```
with backup
create backup
keep original
save backup as [filename]
```

**In-place**:
```
in-place
modify file
edit file
update file
save to [filename]
```

**Examples**:
```
replace "old" with "new" in file.txt in-place with backup
delete line 5 from data.csv modify file
```

### 4.4 Iteration Options

**Repetition**:
```
repeat until no matches
repeat [n] times
recursively
keep going until done
```

**Examples**:
```
replace "  " with " " repeat until no matches
delete empty lines recursively
```

---

## 5. COMPOUND OPERATIONS

Multiple operations in sequence

**Syntax**:
```
[action1] and [action2]
[action1] then [action2]
[action1], [action2], and [action3]
[action1]; [action2]
```

**Examples**:
```
delete empty lines and trim trailing spaces
replace "foo" with "bar" then sort lines
remove comments, delete blank lines, and save to output.txt
```

---

## 6. CONDITIONAL OPERATIONS

Operations with conditions

**Syntax**:
```
if [condition] then [action]
when [condition] [action]
[action] where [condition]
[action] only if [condition]
```

**Examples**:
```
if line contains "error" then delete it
uppercase line where it starts with "header"
replace "debug" with "info" only if in first 10 lines
when line number is even then delete line
```

---

## 7. COMMON PATTERNS (Predefined)

### 7.1 Text Patterns
```
email addresses
URLs / links
phone numbers
numbers
integers
decimals
dates
times
IP addresses
file paths
```

### 7.2 Code Patterns
```
comments
strings
variables
functions
classes
imports
```

### 7.3 Formatting Patterns
```
whitespace
spaces
tabs
newlines
blank lines
empty lines
trailing spaces
leading spaces
duplicate spaces
```

**Examples**:
```
extract all email addresses
delete all comments
remove trailing spaces
find all URLs
```

---

## 8. QUANTIFIERS

### 8.1 Numeric Quantifiers
```
one
two, three, four, ... (spelled out)
1, 2, 3, ... (numeric)
first [n]
last [n]
next [n]
```

### 8.2 Universal Quantifiers
```
all
every
each
any
```

### 8.3 Existential Quantifiers
```
some
at least [n]
at most [n]
exactly [n]
more than [n]
fewer than [n]
```

**Examples**:
```
delete first 5 lines
replace all "foo" with "bar"
show lines with at least 3 words
```

---

## 9. VARIABLES & REFERENCES

For advanced queries (Phase 2+)

**Capture & Reuse**:
```
save [target] as [name]
use [name]
with captured [name]
```

**Examples**:
```
replace "Hello (.*)" with "Hi $1"
swap first word with last word
reverse order of words in line
```

---

## 10. INTERACTIVE MODE

When launched without arguments: `ssed`

### 10.1 Interactive Prompts

```
ssed> What would you like to do?
User: replace foo with bar

ssed> In which file?
User: data.txt

ssed> [Preview of changes]
     Line 5:  "foo bar baz" → "bar bar baz"
     Line 12: "foo test" → "bar test"

ssed> Apply these changes? (yes/no/edit)
User: yes

ssed> ✓ Modified 2 lines in data.txt
```

### 10.2 Guided Mode

```
ssed> What would you like to do?
User: [presses Tab]

ssed> Common actions:
     - replace text
     - delete lines
     - insert text
     - extract/show lines
     - transform text
     - count items

User: replace

ssed> What text do you want to replace?
User: old_variable

ssed> What should it be replaced with?
User: new_variable

ssed> Where should I look?
User: all .go files

ssed> [Shows preview and asks for confirmation]
```

### 10.3 Command History & Learning

```
ssed> Recent commands:
     1. replace "TODO" with "DONE" in tasks.txt
     2. delete empty lines in all .log files
     3. extract email addresses from contacts.txt

ssed> Type a number to repeat a command, or describe a new action
```

---

## 11. EXAMPLES: sed → ssed Translation

### Simple Substitution
```bash
# sed
sed 's/foo/bar/g' file.txt

# ssed
replace all "foo" with "bar" in file.txt
```

### Delete Lines
```bash
# sed
sed '/pattern/d' file.txt

# ssed
delete lines containing "pattern" in file.txt
```

### Print Specific Lines
```bash
# sed
sed -n '5,10p' file.txt

# ssed
show lines 5 to 10 in file.txt
```

### In-place Edit with Backup
```bash
# sed
sed -i.bak 's/old/new/g' file.txt

# ssed
replace all "old" with "new" in file.txt in-place with backup
```

### Delete Empty Lines
```bash
# sed
sed '/^$/d' file.txt

# ssed
delete empty lines in file.txt
```

### Insert Text
```bash
# sed
sed '5i\New Line' file.txt

# ssed
insert "New Line" before line 5 in file.txt
```

### Complex: Delete Range
```bash
# sed
sed '/START/,/END/d' file.txt

# ssed
delete lines between "START" and "END" in file.txt
```

### Case Conversion (GNU sed)
```bash
# sed (GNU)
sed 's/.*/\U&/' file.txt

# ssed
convert all lines to uppercase in file.txt
```

### Extract Pattern
```bash
# sed
sed -n '/pattern/p' file.txt

# ssed
show lines matching "pattern" in file.txt
```

### Multiple Operations
```bash
# sed
sed -e 's/foo/bar/' -e '/baz/d' file.txt

# ssed
replace "foo" with "bar" and delete lines containing "baz" in file.txt
```

---

## 12. IMPLEMENTATION GRAMMAR (For Parser)

### 12.1 Core Grammar (BNF-like)

```bnf
<query> ::= <action> <target> <context>? <options>*

<action> ::= <replace> | <delete> | <insert> | <show> | <transform> | <count> | <move>

<replace> ::= ("replace" | "change" | "substitute") <pattern> ("with" | "to") <text>

<delete> ::= ("delete" | "remove" | "drop") <target>

<insert> ::= ("insert" | "add" | "append" | "prepend") <text> <position>

<show> ::= ("show" | "print" | "display" | "extract") <target>

<target> ::= <literal> | <line_spec> | <pattern_spec> | <range_spec>

<literal> ::= '"' <text> '"' | "'" <text> "'"

<line_spec> ::= "line" <number> | "lines" <range> | "lines" <condition>

<pattern_spec> ::= "words" <pattern_condition> | "text" <pattern_condition> | <predefined_pattern>

<context> ::= ("in" | "from") <file_spec>

<file_spec> ::= <filename> | "all" <pattern> "files"

<options> ::= <scope_option> | <matching_option> | <output_option> | <modification_option>

<scope_option> ::= "first" | "last" | "all" | <number> "occurrence"

<matching_option> ::= "case-sensitive" | "case-insensitive" | "whole word"

<output_option> ::= "show preview" | "dry run" | "show line numbers"

<modification_option> ::= "in-place" | "with backup" | "save to" <filename>
```

### 12.2 Pattern Recognition Priority

1. **Exact quoted strings**: `"foo"`, `'bar'`
2. **Predefined patterns**: `email addresses`, `URLs`, `numbers`
3. **Natural language patterns**: `words starting with`, `lines containing`
4. **Regex patterns**: `/pattern/`, `regex /pattern/`
5. **Line specifications**: `line 5`, `lines 1-10`

---

## 13. ERROR HANDLING & SUGGESTIONS

### 13.1 Ambiguity Resolution

```
User: replace foo in file.txt

ssed: ⚠ What should "foo" be replaced with?
     Suggestion: replace "foo" with "___" in file.txt
```

### 13.2 Did You Mean?

```
User: delte line 5

ssed: Did you mean "delete line 5"? (yes/no)
```

### 13.3 Invalid Patterns

```
User: delete line 1000 from file.txt

ssed: ⚠ file.txt only has 50 lines. Did you mean:
     - delete line 50 (last line)
     - delete all lines
     - delete lines 10 to 50
```

### 13.4 Dangerous Operations

```
User: delete all lines in project.txt

ssed: ⚠ This will delete ALL lines in project.txt
     Are you sure? (yes/no/preview)
```

---

## 14. FUTURE EXTENSIONS

### 14.1 Clipboard Integration
```
copy lines containing "TODO" to clipboard
paste from clipboard after line 10
```

### 14.2 Undo/Redo
```
undo last change
redo
show history
```

### 14.3 Macros
```
define macro "cleanup" as "delete empty lines and trim spaces"
run macro "cleanup" on all .txt files
```

### 14.4 AI-Enhanced
```
User: make this file more readable

ssed: [AI analyzes and suggests]:
     - Add spacing between functions
     - Fix inconsistent indentation
     - Remove trailing whitespace

     Apply these changes? (yes/no/show details)
```

---

## 15. IMPLEMENTATION PHASES

### Phase 1: MVP
- Basic CRUD operations (replace, delete, insert, show)
- Literal text matching
- Line specifications
- Single file operations
- Interactive preview/confirmation

### Phase 2: Pattern Matching
- Predefined patterns (email, URLs, etc.)
- Natural language patterns (starting with, containing, etc.)
- Case-insensitive matching
- Word boundaries

### Phase 3: Advanced Operations
- Transform operations (case conversion, etc.)
- Multiple files
- Range operations
- Compound operations

### Phase 4: Interactive Mode
- Guided prompts
- Tab completion
- Command history
- Suggestions

### Phase 5: Power User Features
- Regex support
- Variables & captures
- Macros
- Complex conditionals

### Phase 6: AI Enhancement
- Intent recognition
- Smart suggestions
- Error correction
- Learning from usage

---

## Appendix: Design Decisions

### Why Natural Language?

1. **Lower Barrier to Entry**: Anyone can use it without training
2. **Self-Documenting**: Queries read like documentation
3. **Discoverable**: Tab completion reveals possibilities
4. **Forgiving**: Multiple ways to express same intent
5. **Safe**: Preview before execution reduces mistakes

### Trade-offs

**Pros**:
- Extremely user-friendly
- No syntax to memorize
- Self-explanatory command history
- Accessible to non-programmers

**Cons**:
- More verbose than sed
- Parsing complexity
- Potential ambiguity
- May be slower for expert users

**Mitigation**:
- Allow shorthand aliases for common operations
- Support partial sed syntax for power users
- Fast parsing with good error messages
- Learn user preferences over time
