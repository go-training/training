# Concat String

Performance result with differnet concat string in Golang.

## Result

```
+ go test -v -bench=. ./example36-performance/...
goos: linux
goarch: amd64
pkg: training/example36-performance
BenchmarkString01
BenchmarkString01-48      17109832         69.8 ns/op       32 B/op        1 allocs/op
BenchmarkString02
BenchmarkString02-48       4786852          252 ns/op       64 B/op        3 allocs/op
BenchmarkString03
BenchmarkString03-48     122471509         9.71 ns/op        0 B/op        0 allocs/op
BenchmarkString04
BenchmarkString04-48      12525997         94.4 ns/op       48 B/op        2 allocs/op
BenchmarkString05
BenchmarkString05-48       7054936          175 ns/op       56 B/op        3 allocs/op
PASS
ok   training/example36-performance 7.607s
```
