# Hello World 範例

## gofmt

請用 `gofmt` 指令修復 `main.go` 程式碼格式，顯示修正 diff 資料。`gofmt -d example01`

```diff
        fmt.Printf(HelloWorld("appleboy"))
        fmt.Println("一天就學會 Go 語言")

-       if (a >= 1) { fmt.Println("a >= 1") }
+       if a >= 1 {
+               fmt.Println("a >= 1")
+       }
```

透過 `-w` 自動修復並且存檔: `gofmt -w example01`

## golint

請用 `golint` 指令修復 `main.go` 程式碼品質

```
example01/main.go:13:1: exported function HelloWorld should have comment or be unexported
```
