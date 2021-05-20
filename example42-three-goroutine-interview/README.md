# Question

Please use three goroutines to run 10 times each, and output cat, dog, bird in order

1. goroutine 01 -> println("cat")
2. goroutine 02 -> println("dog")
3. goroutine 03 -> println("bird")

output:

```sh
cat
dog
bird
cat
...
```

## solution 01

use unbuffered channel. see [the example](./solution01/main.go).
