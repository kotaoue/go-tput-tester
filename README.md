# go-tput-tester
This is the main.go when developing [kotaoue/go-tput](https://github.com/kotaoue/go-tput)

## Result
```ShellSession
$ go run main.go -func=Cols
Cols: 80
```
```ShellSession
$ go run main.go -func=HR
--------------------------------------------------------------------------------
```
```ShellSession
$ go run main.go -func=Setaf -arg=3
byte:  color: 3
```