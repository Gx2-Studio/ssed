Output count_lines.gif

Set Shell "bash"
Set FontSize 16
Set Width 900
Set Height 400
Set Theme "Dracula"

Type "# Count Lines - ssed"
Enter
Sleep 500ms

Type "echo -e 'error here\\nall good\\nerror again\\nfine' | ssed 'count lines containing error'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2\\nno number here' | ssed 'count /[0-9]+/'"
Enter
Sleep 1s

Sleep 2s
