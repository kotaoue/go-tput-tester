# go-tput-tester
This is the main.go when developing [kotaoue/go-tput](https://github.com/kotaoue/go-tput)

## Result
```ShellSession
$ go run main.go -fg=terminfo
cols: 80
--------------------------------------------------------------------------------
```
```ShellSession
$ go run main.go -fg=effects
color is red
set underline
```
```ShellSession
$ go run main.go -fg=clear
```