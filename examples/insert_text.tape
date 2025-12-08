Output insert_text.gif

Set Shell "bash"
Set FontSize 16
Set Width 900
Set Height 500
Set Theme "Dracula"

Type "# Insert Text - ssed"
Enter
Sleep 500ms

Type "echo -e 'line 1\\nline 2\\nline 3' | ssed 'insert NEW before line 2'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2\\nline 3' | ssed 'insert NEW after line 2'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2' | ssed 'insert HEADER first'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2' | ssed 'insert FOOTER last'"
Enter
Sleep 1s

Sleep 2s
