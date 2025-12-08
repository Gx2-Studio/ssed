Output compound_operations.gif

Set Shell "bash"
Set FontSize 16
Set Width 1000
Set Height 500
Set Theme "Dracula"

Type "# Compound Operations - ssed"
Enter
Sleep 500ms

Type "echo -e '# comment\\nvalue=old\\n# another\\ndata=old' | ssed 'delete lines starting with # then replace old with new'"
Enter
Sleep 1s

Type "echo -e '  hello world  \\n  goodbye world  ' | ssed 'trim then convert to uppercase'"
Enter
Sleep 1s

Type "echo -e '  # COMMENT  \\n  HELLO  \\n  # ANOTHER  \\n  WORLD  ' | ssed 'trim then delete lines starting with # then convert to lowercase'"
Enter
Sleep 1s

Sleep 2s
