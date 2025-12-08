# ssed Examples

This directory contains animated examples demonstrating the various features of `ssed` - a natural language stream editor.

## Basic Replace

Simple text replacement operations.

![Basic Replace](basic_replace.gif)

```bash
echo 'hello world' | ssed 'replace world with universe'
echo 'foo bar baz' | ssed 'replace bar with qux'
echo 'old value here' | ssed 'replace old with new'
```

## Regex Replace

Pattern-based replacements using regular expressions.

![Regex Replace](regex_replace.gif)

```bash
echo 'order-123 and order-456' | ssed 'replace /[0-9]+/ with NUM'
echo 'error: something failed' | ssed 'replace /^error:/ with WARNING:'
```

## Delete Lines

Remove lines matching patterns or positions.

![Delete Lines](delete_lines.gif)

```bash
echo -e 'error: failed\ninfo: ok\nerror: again' | ssed 'delete error'
echo -e '# comment\ncode here\n# another' | ssed 'delete lines starting with #'
echo -e 'line 1\nline 2\nline 3\nline 4' | ssed 'delete first 2 lines'
echo -e 'line 1\nline 2\nline 3\nline 4' | ssed 'delete last 2 lines'
```

## Show/Filter Lines

Display only lines matching patterns or positions.

![Show Lines](show_lines.gif)

```bash
echo -e 'error: failed\ninfo: ok\nerror: again' | ssed 'show error'
echo -e '# comment\ncode here\n# another' | ssed 'show lines starting with #'
echo -e 'line 1\nline 2\nline 3\nline 4' | ssed 'show first 2 lines'
echo -e 'first\nsecond\nthird' | ssed 'show line numbers'
```

## Insert Text

Add new lines before or after matching patterns.

![Insert Text](insert_text.gif)

```bash
echo -e 'line 1\nline 2\nline 3' | ssed 'insert NEW before line 2'
echo -e 'line 1\nline 2\nline 3' | ssed 'insert NEW after line 2'
echo -e 'line 1\nline 2' | ssed 'insert HEADER first'
echo -e 'line 1\nline 2' | ssed 'insert FOOTER last'
```

## Transform Text

Case conversion and whitespace operations.

![Transform Text](transform_text.gif)

```bash
echo 'hello world' | ssed 'convert to uppercase'
echo 'HELLO WORLD' | ssed 'convert to lowercase'
echo 'hello world goodbye' | ssed 'convert to titlecase'
echo '   hello world   ' | ssed 'trim'
```

## Count Lines

Count lines matching patterns.

![Count Lines](count_lines.gif)

```bash
echo -e 'error here\nall good\nerror again\nfine' | ssed 'count lines containing error'
echo -e 'line 1\nline 2\nno number here' | ssed 'count /[0-9]+/'
```

## Compound Operations

Chain multiple operations with `then`.

![Compound Operations](compound_operations.gif)

```bash
echo -e '# comment\nvalue=old\n# another\ndata=old' | ssed 'delete lines starting with # then replace old with new'
echo -e '  hello world  \n  goodbye world  ' | ssed 'trim then convert to uppercase'
echo -e '  # COMMENT  \n  HELLO  \n  # ANOTHER  \n  WORLD  ' | ssed 'trim then delete lines starting with # then convert to lowercase'
```

## Regenerating GIFs

These GIFs were created using [VHS](https://github.com/charmbracelet/vhs). To regenerate them:

```bash
cd examples
for tape in *.tape; do vhs "$tape"; done
```
