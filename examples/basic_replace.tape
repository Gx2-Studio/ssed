Output basic_replace.gif

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 400
Set Theme "Dracula"

Type "# Basic Replace - ssed"
Enter
Sleep 500ms

Type "echo 'hello world' | ssed 'replace world with universe'"
Enter
Sleep 1s

Type "echo 'foo bar baz' | ssed 'replace bar with qux'"
Enter
Sleep 1s

Type "echo 'old value here' | ssed 'replace old with new'"
Enter
Sleep 1s

Sleep 2s
