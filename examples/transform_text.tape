Output transform_text.gif

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 500
Set Theme "Dracula"

Type "# Transform Text - ssed"
Enter
Sleep 500ms

Type "echo 'hello world' | ssed 'convert to uppercase'"
Enter
Sleep 1s

Type "echo 'HELLO WORLD' | ssed 'convert to lowercase'"
Enter
Sleep 1s

Type "echo 'hello world goodbye' | ssed 'convert to titlecase'"
Enter
Sleep 1s

Type "echo '   hello world   ' | ssed 'trim'"
Enter
Sleep 1s

Sleep 2s
