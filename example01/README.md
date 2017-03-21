# Hello World

## gofmt

Try to fix `main.go` file using `gofmt` command and show diff data from `gofmt -d example01` command.

```diff
        fmt.Printf(HelloWorld("appleboy"))
        fmt.Println("一天就學會 Go 語言")

-       if (a >= 1) { fmt.Println("a >= 1") }
+       if a >= 1 {
+               fmt.Println("a >= 1")
+       }
```

Fix and save automatically using `-w` flag: `gofmt -w example01`

## golint

Try to improve code quailty using `golint` command.

```
example01/main.go:13:1: exported function HelloWorld should have comment or be unexported
example01/main.go:13:17: don't use underscores in Go names; func parameter user_name should be userName
```
