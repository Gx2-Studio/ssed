SSED - SIMPLE SED
=======================

Text transformation using plain English instead of cryptic sed syntax.

INSTALLATION

    go build -o ssed cmd/ssed/main.go

USAGE

    ssed "<command>" [file]
    cat file | ssed "<command>"

COMMANDS

    replace X with Y          Replace text
    delete X                  Delete lines containing X
    show X                    Show lines containing X
    insert X before Y         Insert text before pattern
    insert X after Y          Insert text after pattern
    convert to uppercase      Change case
    convert to lowercase
    trim                      Remove whitespace
    count X                   Count matching lines

PATTERNS

    delete lines starting with "#"
    delete lines ending with ";"
    delete lines containing "debug"
    delete lines not containing "important"
    show lines containing whole word "cat"
    show first 10 lines
    delete last 5 lines
    show line 5
    delete lines 10 to 20

CHAINING

    Use "then" to chain commands:

    ssed "trim then delete lines starting with '#' then convert to lowercase"

OPTIONS

    -i, --in-place    Edit file directly
    -b, --backup      Backup suffix (e.g., .bak)
    -p, --preview     Preview changes
    -q, --quiet       Suppress output

EXAMPLES

    ssed "replace foo with bar" file.txt
    ssed "delete lines starting with '#'" config.txt
    ssed "show error" app.log
    ssed "trim then convert to uppercase" data.txt
    ssed -i --backup .bak "replace old with new" file.txt

LICENSE

    MIT
