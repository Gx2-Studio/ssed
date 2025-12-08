Output delete_lines.gif

Set Shell "bash"
Set FontSize 16
Set Width 900
Set Height 500
Set Theme "Dracula"

Type "# Delete Lines - ssed"
Enter
Sleep 500ms

Type "echo -e 'error: failed\\ninfo: ok\\nerror: again' | ssed 'delete error'"
Enter
Sleep 1s

Type "echo -e '# comment\\ncode here\\n# another' | ssed 'delete lines starting with #'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2\\nline 3\\nline 4' | ssed 'delete first 2 lines'"
Enter
Sleep 1s

Type "echo -e 'line 1\\nline 2\\nline 3\\nline 4' | ssed 'delete last 2 lines'"
Enter
Sleep 1s

Sleep 2s
