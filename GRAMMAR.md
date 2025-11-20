# ssed Complete Language Grammar Specification

**Version 1.0** - Formal syntax specification for the ssed natural language text transformation tool.

---

## Table of Contents

1. [Overview](#overview)
2. [Formal Grammar (BNF)](#formal-grammar-bnf)
3. [Lexical Elements](#lexical-elements)
4. [Complete Syntax Tree](#complete-syntax-tree)
5. [Pattern Matching System](#pattern-matching-system)
6. [Examples by Complexity](#examples-by-complexity)

---

## Overview

### Language Design Goals

1. **Deterministic Parsing** - No AI required, rule-based grammar
2. **Natural Language** - Readable English-like syntax
3. **Composable** - Build complexity from simple parts
4. **Complete** - Express all sed capabilities
5. **Progressive** - Simple things simple, complex things possible

### Query Structure

Every ssed query follows this pattern:

```
<action> <target> [context] [options]*
```

Where:
- `<action>` = what operation to perform (required)
- `<target>` = what to operate on (required)
- `[context]` = where to apply (optional, defaults to stdin/files)
- `[options]` = how to apply (optional, multiple allowed)

---

## Formal Grammar (BNF)

### Top-Level Grammar

```bnf
<query> ::= <action> <target> <context>? <option>*
         | <multi-step-query>
         | <sed-syntax>

<multi-step-query> ::= "step" <number> ":" <query> (<newline> "step" <number> ":" <query>)+

<sed-syntax> ::= <sed-command> <sed-options>* <file>*
```

### Actions

```bnf
<action> ::= <replace-action>
          | <delete-action>
          | <insert-action>
          | <show-action>
          | <transform-action>
          | <count-action>
          | <move-action>

<replace-action> ::= ("replace" | "substitute" | "change" | "swap") <scope>?

<delete-action> ::= ("delete" | "remove" | "drop" | "erase" | "clear")

<insert-action> ::= ("insert" | "add" | "append" | "prepend")

<show-action> ::= ("show" | "print" | "display" | "extract" | "get" | "find" | "list")

<transform-action> ::= ("convert" | "transform" | "make")

<count-action> ::= ("count" | "number of" | "how many")

<move-action> ::= ("move" | "swap" | "reverse" | "sort")

<scope> ::= "all" | "first" | "last" | <number> | <number> <ordinal>
```

### Targets

```bnf
<target> ::= <literal-target>
          | <pattern-target>
          | <line-target>
          | <template-target>
          | <positional-target>
          | <structured-target>

<literal-target> ::= <quoted-string>

<pattern-target> ::= <predefined-pattern>
                  | <natural-pattern>
                  | <regex-pattern>

<line-target> ::= <line-spec> | <line-range> | <line-condition>

<template-target> ::= "in template" <template-string>
                   | "match template" <template-string>

<positional-target> ::= "in" <structure> "," <position-spec>

<structured-target> ::= "in" <structure-type> "," <field-spec>
```

### Literal Targets

```bnf
<quoted-string> ::= '"' <text> '"' | "'" <text> "'"

<text> ::= <char>+
```

### Pattern Targets

```bnf
<predefined-pattern> ::= "email addresses"
                      | "URLs" | "links"
                      | "phone numbers"
                      | "numbers" | "integers" | "decimals"
                      | "dates" | "times"
                      | "IP addresses"
                      | "UUIDs"
                      | "hex colors"
                      | "file paths"
                      | "HTML tags"
                      | "comments"

<natural-pattern> ::= <text-pattern> | <line-pattern> | <word-pattern>

<text-pattern> ::= "text" <pattern-condition>
                | "characters" <pattern-condition>

<line-pattern> ::= "lines" <pattern-condition>

<word-pattern> ::= "words" | "word" <pattern-condition>

<pattern-condition> ::= "starting with" <quoted-string>
                     | "ending with" <quoted-string>
                     | "containing" <quoted-string>
                     | "matching" <quoted-string>
                     | "between" <quoted-string> "and" <quoted-string>

<regex-pattern> ::= "pattern" "/" <regex> "/"
                 | "regex" "/" <regex> "/"
                 | "matching" "/" <regex> "/"
```

### Line Targets

```bnf
<line-spec> ::= "line" <number>
             | "first line"
             | "last line"

<line-range> ::= "lines" <range-expr>
              | "first" <number> "lines"
              | "last" <number> "lines"

<range-expr> ::= <number> ("to" | "through" | "-") <number>
              | "from" <number> "onwards"
              | "until" <number>

<line-condition> ::= "lines" <condition>
                  | "empty lines"
                  | "blank lines"
                  | "non-empty lines"

<condition> ::= "containing" <pattern>
             | "matching" <pattern>
             | "starting with" <pattern>
             | "ending with" <pattern>
             | "with" <complex-condition>
```

### Template Targets

```bnf
<template-string> ::= '"' <template-element>+ '"'

<template-element> ::= <literal-text> | <placeholder>

<placeholder> ::= "{text}"
               | "{word}"
               | "{number}"
               | "{quoted}"
               | "{quoted" <number> "}"
               | "{any}"
               | "{line}"
               | <named-placeholder>

<named-placeholder> ::= "{" <identifier> "}"

<template-condition> ::= "where" <placeholder-name> <condition-op> <value>

<condition-op> ::= "starts with" | "ends with" | "contains" | "is" | "equals"
```

### Positional Targets

```bnf
<structure> ::= "function calls"
             | "function calls to" <identifier>
             | "function" <identifier>
             | "method calls"
             | "expressions"
             | "statements"

<position-spec> ::= <parameter-position>
                 | <field-position>
                 | <element-position>

<parameter-position> ::= <ordinal-or-number> "parameter"
                      | "parameters"
                      | "all parameters"

<field-position> ::= "column" <number>
                  | <ordinal-or-number> "column"
                  | <ordinal-or-number> "field"

<ordinal-or-number> ::= "first" | "second" | "third" | "last" | <number>
```

### Structured Targets

```bnf
<structure-type> ::= "comma-separated" | "CSV"
                  | "tab-separated" | "TSV"
                  | <delimiter> "-separated"
                  | "KEY=VALUE"
                  | "KEY: VALUE"

<field-spec> ::= "column" <number>
              | <ordinal-or-number> "field"
              | "where" <field-identifier> <condition>
```

### Context

```bnf
<context> ::= <file-context>
           | <range-context>
           | <scope-context>

<file-context> ::= ("in" | "from") <file-spec>

<file-spec> ::= <filename>
             | "all" <pattern> "files"
             | "all files in" <directory>
             | "all files matching" <glob-pattern>
             | "stdin"

<range-context> ::= "in range" <range-expr>
                 | "between" <pattern> "and" <pattern>
                 | "from" <pattern> "onwards"
                 | "until" <pattern>
                 | "before" <pattern>
                 | "after" <pattern>

<scope-context> ::= "only in" <target>
                 | "except in" <target>
                 | "but not in" <target>
```

### Options

```bnf
<option> ::= <matching-option>
          | <output-option>
          | <modification-option>
          | <iteration-option>

<matching-option> ::= "case-sensitive"
                   | "case-insensitive"
                   | "ignore case"
                   | "whole word"
                   | "whole words only"
                   | "exact match"

<output-option> ::= "show preview"
                 | "dry run"
                 | "show what would change"
                 | "show line numbers"
                 | "show matching context"
                 | "show only matches"
                 | "highlight matches"

<modification-option> ::= "in-place"
                       | "modify file"
                       | "edit file"
                       | "with backup"
                       | "create backup"
                       | "save to" <filename>
                       | "save backup as" <filename>

<iteration-option> ::= "repeat until no matches"
                    | "repeat" <number> "times"
                    | "recursively"
```

### Compound Operations

```bnf
<compound> ::= <query> <connector> <query> (<connector> <query>)*

<connector> ::= "and" | "then" | "," | ";"
```

### Conditional Operations

```bnf
<conditional> ::= "if" <condition> "then" <action>
               | "when" <condition> <action>
               | <action> "where" <condition>
               | <action> "only if" <condition>
```

### Captures and References

```bnf
<capture> ::= "capture" <target> "as" <identifier>

<reference> ::= "use" <identifier>
             | "{" <identifier> "}"
             | "$" <number>
             | "\" <number>
```

### Character Classes

```bnf
<char-class> ::= "letters" | "digits" | "alphanumeric"
              | "whitespace" | "spaces" | "tabs"
              | "punctuation"
              | "word characters"
              | "special characters"

<position-class> ::= "at start" | "at end" | "at beginning"
                  | "word boundary"
```

### Quantifiers

```bnf
<quantifier> ::= <number> "or more" <element>
              | "between" <number> "and" <number> <element>
              | "exactly" <number> <element>
              | "at least" <number> <element>
              | "at most" <number> <element>
              | "repeated" <element>
              | "consecutive" <element>
```

---

## Lexical Elements

### Keywords

**Actions:**
```
replace, substitute, change, swap, rename
delete, remove, drop, erase, clear
insert, add, append, prepend
show, print, display, extract, get, find, list
convert, transform, make
count, number, how many
move, reverse, sort, duplicate
```

**Quantifiers:**
```
all, every, each, any, some
first, second, third, last, next
one, two, three, ..., N
exactly, at least, at most, between
```

**Positions:**
```
before, after, at, in, from, to, until
beginning, start, end
line, word, character, parameter, column, field
```

**Conditions:**
```
containing, matching, starting, ending
where, if, when, only, except, but not
with, without
is, equals, starts, ends, contains
```

**Options:**
```
case-sensitive, case-insensitive, ignore case
whole word, exact match
show preview, dry run
in-place, modify file, with backup
repeat, recursively
```

**Patterns:**
```
email addresses, URLs, phone numbers, IP addresses
numbers, digits, letters, whitespace
dates, times, UUIDs, hex colors
comments, strings, functions, classes
```

### Operators

```
= (equals/assignment)
, (separator)
; (sequence separator)
: (label/step marker)
- (range)
/ (regex delimiter)
{ } (placeholder/capture)
" " (string delimiter)
' ' (string delimiter)
( ) (grouping)
```

### Identifiers

```bnf
<identifier> ::= <letter> (<letter> | <digit> | "_")*

<letter> ::= [a-zA-Z]
<digit> ::= [0-9]
<number> ::= <digit>+
```

### Comments

```bnf
<comment> ::= "#" <text> <newline>
```

---

## Complete Syntax Tree

### Level 1: Simple Literals (Phase 1)

```
Action: replace, delete, show, insert
Target: "literal text"
Context: in file.txt
Options: show preview

Example: replace "foo" with "bar" in file.txt
```

**Parse Tree:**
```
Query
├── Action: replace
├── Source: "foo"
├── Replacement: "bar"
└── Context
    └── File: "file.txt"
```

### Level 2: Predefined Patterns (Phase 2)

```
Action: extract, remove, find
Target: email addresses | URLs | phone numbers
Context: from file.txt
Options: save to output.txt

Example: extract all email addresses from contacts.txt
```

**Parse Tree:**
```
Query
├── Action: extract
├── Quantifier: all
├── Target: PredefinedPattern
│   └── Type: email addresses
└── Context
    └── File: "contacts.txt"
```

### Level 3: Natural Patterns (Phase 2)

```
Action: delete, show, replace
Target: lines|words|text [condition]
Condition: starting with|ending with|containing "pattern"
Context: in file.txt

Example: delete lines starting with "#" from config.txt
```

**Parse Tree:**
```
Query
├── Action: delete
├── Target: LinePattern
│   └── Condition
│       ├── Type: starting with
│       └── Pattern: "#"
└── Context
    └── File: "config.txt"
```

### Level 4: Template Matching (Phase 3)

```
Action: match template|in template
Target: "literal {placeholder} literal {placeholder}"
Condition: where placeholder [condition]
Replacement: "literal {placeholder} literal"

Example: in template "API_CALL({quoted1}, {quoted2})"
         replace with "API_CALL({quoted1}, {quoted2})"
```

**Parse Tree:**
```
Query
├── Action: in template
├── Template: "API_CALL({quoted1}, {quoted2})"
│   ├── Literal: "API_CALL("
│   ├── Placeholder: quoted1
│   ├── Literal: ", "
│   ├── Placeholder: quoted2
│   └── Literal: ")"
├── Action: replace with
└── Replacement: Template
    └── [same structure, potentially modified]
```

### Level 5: Positional Targeting (Phase 3)

```
Action: in [structure], [action]
Structure: function calls | function NAME
Position: first|second|Nth parameter
Condition: if|where [condition]
Modification: remove|replace|add

Example: in function calls to API_CALL,
         in second parameter,
         if starts with /,
         remove it
```

**Parse Tree:**
```
Query
├── Structure: function calls
│   └── Name: "API_CALL"
├── Position: parameter
│   └── Index: 2 (second)
├── Condition
│   ├── Type: starts with
│   └── Value: "/"
└── Action: remove
    └── Target: first character
```

### Level 6: Structured Data (Phase 3)

```
Action: in [data-format]
Format: CSV|TSV|KEY=VALUE
Field: column N | field N
Condition: where [field] [condition]

Example: in comma-separated lines,
         where column 1 is "ERROR",
         delete line
```

**Parse Tree:**
```
Query
├── DataFormat: CSV
├── Condition
│   ├── Field: column 1
│   ├── Operator: is
│   └── Value: "ERROR"
└── Action: delete line
```

### Level 7: Captures and References (Phase 4)

```
Action: capture [target] as NAME
Reference: use NAME | {NAME}
Operation: [action] with [reference]

Example: capture word after "class" as CLASSNAME,
         then replace "new CLASSNAME" with "new Modified_CLASSNAME"
```

**Parse Tree:**
```
CompoundQuery
├── Step1: Capture
│   ├── Target: word after "class"
│   └── Name: CLASSNAME
└── Step2: Replace
    ├── Pattern: "new {CLASSNAME}"
    └── Replacement: "new Modified_{CLASSNAME}"
```

### Level 8: Multi-Step Operations (Phase 4)

```
step 1: [action]
step 2: [action]
step 3: [action]

Example:
  step 1: find lines matching template "X({any})"
  step 2: within parentheses, find quoted text
  step 3: if starts with /, remove it
```

**Parse Tree:**
```
MultiStepQuery
├── Step1
│   ├── Number: 1
│   └── Query: [parse tree for step 1]
├── Step2
│   ├── Number: 2
│   └── Query: [parse tree for step 2]
└── Step3
    ├── Number: 3
    └── Query: [parse tree for step 3]
```

### Level 9: Compound Operations (Phase 3)

```
[query1] and [query2]
[query1] then [query2]
[query1], [query2], and [query3]

Example: delete empty lines and trim trailing spaces
```

**Parse Tree:**
```
CompoundQuery
├── Connector: and
├── Query1
│   ├── Action: delete
│   └── Target: empty lines
└── Query2
    ├── Action: trim
    └── Target: trailing spaces
```

### Level 10: Conditional Operations (Phase 4)

```
if [condition] then [action]
when [condition] [action]
[action] where [condition]
[action] only if [condition]

Example: if line contains "error" then make it uppercase
```

**Parse Tree:**
```
ConditionalQuery
├── Condition
│   ├── Target: line
│   ├── Operator: contains
│   └── Value: "error"
└── Action: make uppercase
```

---

## Pattern Matching System

### Simple Pattern Matching

**Direct Match:**
```
"literal text" → exact string match
```

**Wildcards (Basic):**
```
any character → . in regex
any text → .* in regex
digits → \d+ in regex
letters → [a-zA-Z]+ in regex
whitespace → \s+ in regex
```

### Template Matching

**Template Syntax:**
```
"LITERAL {placeholder} LITERAL"
```

**Placeholders:**

| Placeholder | Matches | Regex Equivalent |
|------------|---------|------------------|
| `{text}` | Non-whitespace text | `\S+` |
| `{word}` | Word characters | `\w+` |
| `{number}` | Numeric value | `\d+` |
| `{quoted}` | Content in quotes | `"([^"]*)"` |
| `{any}` | Any characters | `.*` |
| `{line}` | Rest of line | `.*$` |
| `{name}` | Named capture | `(?P<name>.*)` |

**Template Examples:**

```bash
# Function call with 2 parameters
Template: "function({param1}, {param2})"
Regex: function\(([^,]+), ([^)]+)\)

# Key-value pair
Template: "{key}={value}"
Regex: ([^=]+)=(.+)

# API call pattern
Template: "API_CALL({quoted1}, {quoted2})"
Regex: API_CALL\("([^"]+)", "([^"]+)"\)
```

### Positional Matching

**Function Parameters:**
```
in function NAME, parameter N
→ Parse function call, extract Nth parameter
```

**CSV/Delimited Fields:**
```
in comma-separated, column N
→ Split on comma, extract Nth field
```

**Sequential Elements:**
```
first|second|third|Nth|last element
→ Positional indexing (0-based or 1-based)
```

### Conditional Matching

**Simple Conditions:**
```
where X starts with Y
where X ends with Y
where X contains Y
where X is Y
where X equals Y
```

**Complex Conditions:**
```
where X starts with Y and X contains Z
where X is Y or X is Z
where X starts with Y and not contains Z
```

**Logical Operators:**
```
and → boolean AND
or → boolean OR
not → boolean NOT
```

---

## Examples by Complexity

### Level 1: Simple (Phase 1)

```bash
replace "hello" with "hi" in file.txt
delete line 5 from data.txt
show line 10 in log.txt
insert "header" before line 1 in README.md
```

### Level 2: Pattern-Based (Phase 2)

```bash
extract all email addresses from contacts.txt
delete lines containing "debug" from app.log
remove trailing spaces in all .js files
show lines starting with "ERROR" from log.txt
```

### Level 3: Template Simple (Phase 3)

```bash
in template "{key}={value}" where key is "port" replace value with "3000"
match template "console.log({any})" delete line
in template "import {module} from {path}" replace module with "newModule"
```

### Level 4: Template Complex (Phase 3)

```bash
in template "API_CALL({quoted1}, {quoted2})"
  where quoted2 starts with "/"
  replace with "API_CALL({quoted1}, {quoted2-without-first-char})"

match template "function {name}({params})"
  where name starts with capital letter
  add comment above line
```

### Level 5: Positional (Phase 3)

```bash
in function calls to setTimeout,
  if has 1 parameter,
  add second parameter "0"

in API_CALL function calls,
  in second parameter,
  remove leading slash

in comma-separated lines,
  replace column 3 with "N/A"
```

### Level 6: Multi-Step (Phase 4)

```bash
step 1: find lines matching template "API_CALL({any}, {any})"
step 2: in second parameter within quotes
step 3: if starts with /
step 4: remove the /
apply to all .hpp files
edit in place
```

### Level 7: Captures (Phase 4)

```bash
capture word after "class" as CLASSNAME,
  then replace "new CLASSNAME" with "new Modified_CLASSNAME"

capture first word as A and second word as B,
  swap them

in email addresses,
  capture domain as DOMAIN,
  show only DOMAIN
```

### Level 8: Compound (Phase 3)

```bash
delete empty lines and trim trailing spaces in all .txt files

replace "foo" with "bar" then sort lines then remove duplicates

remove comments, delete blank lines, and save to output.txt
```

### Level 9: Conditional (Phase 4)

```bash
if line contains "error" then make it uppercase

when line number is even then delete line

replace "debug" with "info" only if in first 10 lines

uppercase line where it starts with "header"
```

### Level 10: Complex Real-World (Phase 3-4)

```bash
# Fix API calls
in function calls to API_CALL with 2 parameters,
  in second parameter,
  if value in quotes starts with /,
  remove that /,
  in all .hpp files,
  modify files with backup

# Refactor imports
in lines starting with "import",
  capture module name as MOD,
  if MOD starts with "old_",
  replace with MOD but change "old_" to "new_",
  in all .js files

# Clean logs
delete lines containing "DEBUG" or "TRACE",
  then remove empty lines,
  then save to clean.log,
  from app.log

# Transform CSV
in comma-separated lines,
  where column 1 is "ACTIVE",
  set column 3 to "ENABLED",
  set column 4 to current date,
  save to output.csv
```

---

## Parser Implementation Notes

### Parsing Strategy

1. **Tokenization**: Break input into tokens (keywords, identifiers, strings, operators)
2. **Keyword Recognition**: Match against known action/pattern/option keywords
3. **Structure Detection**: Identify template patterns, function calls, data formats
4. **Tree Building**: Construct AST based on grammar rules
5. **Validation**: Check semantic validity
6. **Optimization**: Simplify/optimize the operation plan
7. **Execution**: Generate sed/internal operations

### Ambiguity Resolution

When multiple parses are possible:

1. **Prefer specific over general** - "email addresses" over "text matching .*@.*"
2. **Left-to-right parsing** - First valid parse wins
3. **Keyword priority** - Reserved keywords take precedence
4. **Context awareness** - Use file type to disambiguate
5. **Interactive clarification** - When truly ambiguous, ask user

### Error Handling

**Syntax Errors:**
```
Unknown action "foo" - did you mean "show", "find", "extract"?
```

**Semantic Errors:**
```
Line 1000 doesn't exist (file has only 50 lines)
```

**Suggestions:**
```
Did you mean: delete lines containing "error"
Instead of: delete line error
```

---

## Complete Grammar Summary

```bnf
<query> ::= <simple-query> | <template-query> | <positional-query>
         | <multi-step-query> | <compound-query> | <conditional-query>

<simple-query> ::= <action> <target> [<context>] [<option>]*

<template-query> ::= ("in" | "match") "template" <template-spec>
                     [<action>] [<replacement>]

<positional-query> ::= "in" <structure> "," <position> "," <action>

<multi-step-query> ::= ("step" <number> ":" <query>)+

<compound-query> ::= <query> <connector> <query>

<conditional-query> ::= ("if" | "when") <condition> ("then")? <action>
                     | <action> ("where" | "only if") <condition>

<action> ::= replace | delete | insert | show | transform | count | move

<target> ::= <literal> | <pattern> | <line-spec> | <template> | <position>

<pattern> ::= <predefined-pattern> | <natural-pattern> | <regex-pattern>

<context> ::= ("in" | "from") <file-spec> | <range-context> | <scope-context>

<option> ::= <matching-option> | <output-option> | <modification-option>
```

---

## Conclusion

This grammar provides:

✅ **Complete sed feature parity** through natural language
✅ **Deterministic parsing** without requiring AI
✅ **Progressive complexity** from simple to advanced
✅ **Composable operations** building complexity from basics
✅ **Clear syntax tree** for implementation
✅ **Formal specification** for parser development

Implementation can proceed phase-by-phase, with each level building on the previous, ensuring the tool remains usable and testable at every stage.
