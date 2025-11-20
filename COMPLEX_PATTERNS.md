# Complex Patterns in Natural Language (Without AI)

How to express complex patterns using structured natural language that can be parsed deterministically.

## The Challenge

Complex regex like:
```regex
API_CALL\(\"([^\"]*)\", \"\/([^\"]*)\)\"
```

Needs to be expressible in natural language without requiring AI to interpret intent.

---

## Solution: Structured Pattern Templates

### Core Principle
Break complex patterns into **composable building blocks** that can be parsed deterministically.

---

## 1. Template Matching with Placeholders

### Syntax
```
in template "LITERAL {placeholder} LITERAL {placeholder}"
```

### Placeholders
- `{text}` - any non-whitespace text
- `{word}` - single word (alphanumeric + _)
- `{number}` - numeric value
- `{quoted}` - anything in quotes
- `{any}` - any characters
- `{line}` - rest of line

### Examples

**API_CALL Pattern:**
```bash
ssed "in template 'API_CALL({quoted}, {quoted})' replace second quoted text starting with / by removing the /"

# Or more explicitly:
ssed "match template 'API_CALL(\"{first}\", \"/{second}\")' replace with 'API_CALL(\"{first}\", \"{second}\")'"
```

**Key-Value Pairs:**
```bash
ssed "in template '{key}={value}' where key is 'port' replace value with '3000'"

# Matches: port=8080 → port=3000
```

**Function Calls:**
```bash
ssed "in template 'console.log({any})' delete line"

# Matches any console.log call
```

---

## 2. Positional Parameter System

### Syntax
```
in "FUNCTION(param1, param2, param3)"
```

### Operations on Parameters
- `first parameter`
- `second parameter`
- `parameter 2`
- `last parameter`
- `all parameters`

### Examples

**Modify Specific Parameter:**
```bash
ssed "in function calls to API_CALL, in second parameter, remove leading slash"

# Or:
ssed "in API_CALL calls, change second parameter from '/X' to 'X'"
```

**Swap Parameters:**
```bash
ssed "in function SWAP, swap first and second parameters"

# SWAP(a, b) → SWAP(b, a)
```

**Add Parameter:**
```bash
ssed "in setTimeout calls with 1 parameter, add second parameter '0'"

# setTimeout(fn) → setTimeout(fn, 0)
```

---

## 3. Structured Field Matching

For structured data (CSV, JSON-like, etc.)

### CSV/Delimited
```bash
ssed "in comma-separated lines, replace column 3 with 'N/A'"

ssed "in pipe-delimited lines where column 1 is 'ERROR', delete line"
```

### Key-Value
```bash
ssed "in lines with KEY=VALUE format where KEY is 'DEBUG', set VALUE to 'false'"

ssed "in KEY: VALUE format, if KEY starts with 'temp_', delete line"
```

---

## 4. Nested Pattern Composition

Combine multiple conditions with "and", "or", "where"

### Syntax
```
in [PATTERN1] where [CONDITION] and [CONDITION]
```

### Examples

**Multiple Conditions:**
```bash
ssed "in quoted strings where text starts with '/' and appears after comma, remove first character"

# In API_CALL("foo", "/bar") - removes / from "/bar"
```

**Complex Structure:**
```bash
ssed "in function calls where function name ends with 'Async' and has 2 parameters, add third parameter 'null'"
```

**Contextual:**
```bash
ssed "in lines containing 'function' where word after 'function' starts with capital letter, add comment above line"
```

---

## 5. Between/Within Delimiters

### Syntax
```
between "OPEN" and "CLOSE"
within "DELIM"
inside "WRAPPER"
```

### Examples

**Quotes:**
```bash
ssed "within double quotes, replace spaces with underscores"

# "hello world" → "hello_world"
```

**Parentheses:**
```bash
ssed "between '(' and ')' after 'API_CALL', in second comma-separated part, remove leading slash"

# API_CALL("x", "/y") → API_CALL("x", "y")
```

**Tags:**
```bash
ssed "between '<code>' and '</code>', escape special characters"
```

---

## 6. Character Class Specifications

Named character classes instead of regex:

### Built-in Classes
- `letters` = [a-zA-Z]
- `digits` = [0-9]
- `alphanumeric` = [a-zA-Z0-9]
- `whitespace` = [ \t\n]
- `punctuation` = [.,;:!?]
- `word characters` = [a-zA-Z0-9_]

### Position Classes
- `at start` = ^
- `at end` = $
- `word boundary` = \b

### Examples

```bash
ssed "replace 3 or more consecutive whitespace with single space"

ssed "at end of line, remove all punctuation"

ssed "replace sequences of digits with 'XXX'"
```

---

## 7. Repetition Patterns

### Syntax
- `N or more X`
- `between N and M X`
- `exactly N X`
- `repeated X`
- `consecutive X`

### Examples

```bash
ssed "replace 2 or more consecutive spaces with single space"

ssed "replace 3 to 5 consecutive digits with 'NUMBER'"

ssed "remove consecutive duplicate lines"
```

---

## 8. Capture and Reference

### Named Captures
```
capture [PATTERN] as NAME
use NAME
```

### Examples

**Save and Reuse:**
```bash
ssed "capture word after 'class' as CLASSNAME, then replace 'new CLASSNAME' with 'new Modified_CLASSNAME'"
```

**Swap Elements:**
```bash
ssed "capture word before comma as FIRST and word after comma as SECOND, swap them"

# foo, bar → bar, foo
```

**API_CALL Example:**
```bash
ssed "in API_CALL calls, capture first quoted text as ENDPOINT and second quoted text as PATH, if PATH starts with '/' remove it"
```

---

## 9. Context-Aware Patterns

### File-Type Specific
```bash
ssed --filetype=code "in string literals, escape backslashes"

ssed --filetype=csv "in field 3, remove quotes"
```

### Language Constructs (Predefined)
```
in function calls
in class definitions
in import statements
in string literals
in comments
```

### Examples

```bash
ssed "in C++ function calls to API_CALL, modify second parameter to remove leading slash"

ssed "in Python import statements, replace 'old_module' with 'new_module'"

ssed "in JavaScript string literals, replace single quotes with double quotes"
```

---

## 10. Multi-Step Pattern Description

Breaking complex operations into steps:

### Syntax
```
step 1: [action]
step 2: [action]
step 3: [action]
```

### Example - API_CALL Pattern

```bash
ssed "
  step 1: find lines matching template 'API_CALL(\"{any}\", \"/{any}\")'
  step 2: in second quoted section
  step 3: remove first character if it is slash
  in all .hpp files
  modify files
"
```

### Example - Complex Transformation

```bash
ssed "
  step 1: find function calls where function name is API_CALL
  step 2: within the parentheses, find comma-separated parts
  step 3: in second part, if text in quotes starts with /, remove that /
  save changes
"
```

---

## Real Examples: sed → ssed (Complex Patterns)

### 1. API_CALL Pattern

**sed:**
```bash
sed -i 's/API_CALL("\([^"]*\)", "\/\([^"]*\)"/API_CALL("\1", "\2"/g' *.hpp
```

**ssed (Option A - Template):**
```bash
ssed 'in template "API_CALL({quoted1}, {quoted2})" where quoted2 starts with "/" replace with "API_CALL({quoted1}, {quoted2-without-first-char})" in all .hpp files in-place'
```

**ssed (Option B - Structured):**
```bash
ssed 'in function calls to API_CALL, in second parameter, if value in quotes starts with slash, remove that slash, in all .hpp files, modify files'
```

**ssed (Option C - Step by Step):**
```bash
ssed '
  find API_CALL function calls with 2 parameters
  in second parameter within quotes
  if starts with /
  remove the /
  apply to all .hpp files
  edit in place
'
```

### 2. Swap Words

**sed:**
```bash
sed 's/\([a-z]*\) \([a-z]*\)/\2 \1/g'
```

**ssed:**
```bash
ssed 'capture first word as A and second word as B separated by space, replace with B space A'
```

### 3. Extract Domain from Email

**sed:**
```bash
sed 's/.*@\(.*\)/\1/'
```

**ssed:**
```bash
ssed 'in email addresses, show only domain part (after @)'

# Or:
ssed 'in template "{user}@{domain}", replace whole match with {domain}'
```

### 4. Remove HTML Tags

**sed:**
```bash
sed 's/<[^>]*>//g'
```

**ssed:**
```bash
ssed 'remove all text between < and > including the brackets'

# Or:
ssed 'remove HTML tags'  # predefined pattern
```

### 5. Comment Out Function Calls

**sed:**
```bash
sed 's/^\(\s*\)\(console\.log.*\)/\1\/\/ \2/'
```

**ssed:**
```bash
ssed 'in lines starting with optional whitespace followed by console.log, add // after the whitespace'

# Or simpler:
ssed 'in lines containing console.log at start (after whitespace), add // before console'
```

### 6. Add Quotes Around Numbers

**sed:**
```bash
sed 's/: \([0-9]*\)/: "\1"/g'
```

**ssed:**
```bash
ssed 'in template ": {number}", replace with ": \"{number}\""'

# Or:
ssed 'after colon and space, if followed by number, wrap number in quotes'
```

---

## Implementation Strategy

### Parser Components Needed:

1. **Template Parser**
   - Tokenize template strings
   - Extract placeholders
   - Build match patterns

2. **Placeholder Resolver**
   - `{quoted}` → match content in quotes
   - `{word}` → match \w+
   - `{number}` → match \d+
   - `{any}` → match .*

3. **Position Parser**
   - "first parameter" → capture group 1
   - "second quoted text" → second occurrence of quoted pattern
   - "after comma" → position relative to delimiter

4. **Condition Parser**
   - "where X starts with Y"
   - "if X contains Y"
   - "and", "or" logical operators

5. **Action Parser**
   - "remove" → delete matched text
   - "replace with" → substitute
   - "add" → insert

### Complexity Levels:

**Level 1 (Phase 1):** Simple literals
```bash
ssed "replace foo with bar"
```

**Level 2 (Phase 2):** Predefined patterns
```bash
ssed "remove email addresses"
```

**Level 3 (Phase 3):** Templates with simple placeholders
```bash
ssed 'in template "KEY={value}" replace value with "new"'
```

**Level 4 (Phase 3):** Positional with conditions
```bash
ssed "in API_CALL, second parameter, remove leading /"
```

**Level 5 (Phase 4):** Multi-step with captures
```bash
ssed "capture X as A, if A starts with /, remove it"
```

**Level 6 (Phase 5):** Fall back to regex
```bash
ssed "replace regex /complex[0-9]+/ with simple"
```

---

## Decision Tree for Complex Patterns

```
User Query: "Fix API_CALL slashes"

1. Is it a predefined pattern? NO
2. Can it be expressed as template? YES
   → Use template matching
3. Does it need positional awareness? YES
   → Parse function call structure
4. Are there conditions? YES ("starts with /")
   → Add conditional filter
5. Build operation: modify second parameter
6. Generate preview
7. Ask confirmation
```

---

## Benefits of This Approach

✅ **Deterministic** - No AI needed, rule-based parsing
✅ **Composable** - Build complexity from simple parts
✅ **Debuggable** - Each part can be tested independently
✅ **Extensible** - Add new templates and patterns
✅ **Understandable** - Reads like structured English
✅ **Parseable** - Clear grammar for each construct

---

## Limitations

⚠ **Verbose** - More typing than regex (but more readable)
⚠ **Learning Curve** - Still need to learn template syntax
⚠ **Edge Cases** - Very complex patterns might still need regex fallback

---

## Recommendation

For **ssed Phase 1-4**, implement:

1. **Simple patterns** - literal, predefined (Phase 1-2)
2. **Template matching** - with basic placeholders (Phase 3)
3. **Positional awareness** - parameters, fields (Phase 3)
4. **Multi-step descriptions** - break complex into simple (Phase 4)
5. **Regex fallback** - for truly complex cases (Phase 5)

This gives 80% of use cases without needing AI, while maintaining the natural language philosophy.
