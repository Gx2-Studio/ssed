Output regex_replace.gif

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 400
Set Theme "Dracula"

Type "# Regex Replace - ssed"
Enter
Sleep 500ms

Type "echo 'order-123 and order-456' | ssed 'replace /[0-9]+/ with NUM'"
Enter
Sleep 1s

Type "echo 'error: something failed' | ssed 'replace /^error:/ with WARNING:'"
Enter
Sleep 1s

Sleep 2s
