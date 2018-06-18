# Benchmark in Golang

```sh
$ go test -v -bench=. -run=none -benchmem .
```

result:

```
goos: darwin
goarch: amd64
BenchmarkPrintInt2String01-4    10000000               117 ns/op              16 B/op          2 allocs/op
BenchmarkPrintInt2String02-4    30000000                37.3 ns/op             3 B/op          1 allocs/op
BenchmarkPrintInt2String03-4    30000000                38.9 ns/op             3 B/op          1 allocs/op
PASS
```
