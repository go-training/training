# Golang Training

[![Run Tests](https://github.com/go-training/training/actions/workflows/go.yml/badge.svg)](https://github.com/go-training/training/actions/workflows/go.yml)

Learning basic [Golang](https://go.dev/) in one day

## Example code

- [example01](./example01-hello-world): Show hello world. Try to use [golint][1] and [gofmt][2] tool.
- [example02](./example02-golang-package): Let's [write a library][3] and use it from the hello program.
- [example03](./example03-if-switch-const): How to use `if`, `switch` and `const` in Go.
- [example04](./example04-constructor-and-struct): How to initializing `constructor` in Go.
- [example05](./example05-interface): Create interface for multiple struct.
- [example06](./example06-go-concurrency): How to use Goroutines in Go, share by communicating.
- [example07](./example07-errors-handler): Errors handler in Go.
- [example08](./example08-type-assertions): A type assertion provides access to an interface value.
- [example09](./example09-command-line-tool): Write command line tool in Go.
- [example10](./example10-simple-http-server): Simple http server and show example using [gin][4] framework.
- [example11](./example11-cross-build): Learn go build constraints.
- [example12](./example12-build-with-docker): Building minimal docker containers for go application.
- [example13](./example13-share-golang-package-to-c): Sharing Golang packages to C and Go.
- [example14](./example14-go-func): How to use func in go?
- [example15](./example15-pass-slice-as-function-args): How to pass slice as function arguments?
- [example16](./example16-init-func): What is init function in Go?
- [example17](./example17-json): How to handle json in Go?
- [example18](./example18-write-testing-and-doc): How to write testingin Go?
- [example19](./example19-deploy-with-kubernetes): Deploy golang app using [drone](https://drone.io/) + [kubernetes](https://kubernetes.io/).
- [example20](./example20-write-benchmark): write simple benchmark.
- [exmaple21](./example21-simple-golang-https-tls): Simple https-tls server example using [mkcert](https://github.com/FiloSottile/mkcert).
- [exmaple22](./example22-go-module-in-go.11): How to use [go module](https://github.com/golang/go/wiki/Modules) in go version 1.11.
- [example23](./example23-deploy-go-application-with-up): Deploy go application useing [apex/up](https://github.com/apex/up) tool.
- [example24](./example24-debug-go-code-using-vs-code): Debug golang code in vscode using [dlv](https://github.com/go-delve/delve).
- [example25](./example25-traefik-golang-app-lets-encrypt): Running golang app using [traefik & Let's Encrypt & Docker](https://docs.traefik.io/user-guide/docker-and-lets-encrypt/).
- [example27](./example27-how-to-load-env): How to loads environment variables from `.env`.
- [example28](./example28-webserver-with-gracefull-shutdown): Go webserver with gracefull shutdown.
- [example29](./example29-handle-multiple-channel): How to handle multiple Go channel?
- [example30](./example30-context-timeout): Simplest Way to Handle Timeouts?
- [example31](./example31-job-queue): How to implements a simple job queue? [Youtube](https://www.youtube.com/watch?v=wvdbobFiXNg)
- [example32](./example32-what-is-select): four tips with `select` in golang.
- [example33](./example33-share-memory-by-communicating): share memory by communicating.
- [example34](./example34-graceful-shutdown-with-worker): graceful shutdown with multiple workers.
- [example35](./example35-goroutine-with-context): introduction to context package.
- [example36](./example36-performance): performance result with differnet concat string method.
- [example38](./example38-concurrency-is-still-not-easy): concurrency is still not easy. [English Blog](https://utcc.utoronto.ca/~cks/space/blog/programming/GoConcurrencyStillNotEasy), [Chinese Blog](https://blog.wu-boy.com/2020/09/limit-concurrency-in-golang/)
- [example39](./example39-select-random-channel): select random channel value?
- [example40](./example40-embedding-files): [go 1.16](https://tip.golang.org/doc/go1.16) support [Embedding Files](https://tip.golang.org/pkg/embed/), [Chinese Blog](https://blog.wu-boy.com/2020/12/embedding-files-in-go-1-16/)
- [example41](./example41-using-buffer-channel-signal) why use buffered channel signal? [Chinese Blog](https://blog.wu-boy.com/2021/03/why-use-buffered-channel-in-signal-notify/)
- [example42](./example42-three-goroutine-interview/README.md) Please use three goroutines to run 10 times each, and output `cat`, `dog`, `bird` in order
- [example43](./example43-progress-bar/README.md) graceful shutdown with progress bar. [Chinese Blog](https://blog.wu-boy.com/2021/05/graceful-shutdown-with-progress-bar-in-golang/)
- [example44](./example44-signal-with-docker-compose) graceful stop service using docker-compose. [Chinese Blog](https://blog.wu-boy.com/2021/06/graceful-stop-service-using-docker-compose-in-golang/)
- [example45](./example45-go-1.18-generics) introduction to go1.18 generics. [Chinese Blog](https://blog.wu-boy.com/2022/02/introduction-to-golang-1-18-generics/)
- [example46](./example46-go-1.18-generics-part2) introduction to go1.18 generics constraints, slices and maps packages. [Chinese Blog](https://blog.wu-boy.com/2022/02/golang-1-18-generics-constraints-slices-maps/)
- [example47](./example47-pyroscope-adhoc-profiling) pyroscope adhoc profiling in go. [Chinese Blog](https://blog.wu-boy.com/2022/02/pyroscope-adhoc-profiling-in-golang/)
- [example48](./example48-pub-sub-pattern) simple pub/sub pattern in go. [Chinese Blog](https://blog.wu-boy.com/2022/04/simple-publish-subscribe-pattern-in-golang/)
- [example49](./example49-dependency-injection/) Dependency Injection using wire tool. [Chinese Blog](https://blog.wu-boy.com/2022/09/dependency-injection-in-go/)
- [exmaple50](./example50-consumer-producer-pattern/) simple scheduler with multiple worker. [Chinese Blog](https://blog.wu-boy.com/2022/11/simple-scheduler-with-multiple-worker-using-golang/)
- [example51](./example51-canceler/) How to cancel a running task between backend and worker service. [Chinese Blog](https://blog.wu-boy.com/2022/12/system-design-how-to-cancel-a-running-task-in-golang/)
- [example52](./example52-ring-buffer-queue/) ring buffer queue with fixed size. [Chinese Blog](https://blog.wu-boy.com/2023/01/ring-buffer-queue-with-fixed-size-in-golang/)
- [example53](./example53-container-base-image-updates/) A container-based solution for automating Docker container base image updates. [Chinese Blog](https://blog.wu-boy.com/2023/02/automating-docker-container-base-image-updates-cht/), [English Blog](https://blog.wu-boy.com/2023/02/automating-docker-container-base-image-updates-eng/)
- [exmaple54](./example54-hot-reload-for-development) introduction live reload tool: air [Chinese Blog](https://blog.wu-boy.com/2023/10/live-reload-in-go/)
- [example55](./example55-cache-hotspot-invalid/) how to invalidate cache hotspot in golang. [Chinese Blog](https://blog.wu-boy.com/2024/02/how-to-reslove-the-hotspot-invalid-using-singleflight/), [English Blog](https://blog.wu-boy.com/2024/02/how-to-reslove-the-hotspot-invalid-using-singleflight-en/)
- [example56](./example56-context-timeout/) how to use context with timeout in golang. [Chinese Blog](https://blog.wu-boy.com/2024/11/learning-go-language-with-gitlab-api-zh-tw/), [English Blog](https://blog.wu-boy.com/2024/11/learning-go-language-with-gitlab-api-en/) [Chinese Video](https://www.youtube.com/watch?v=nPn-pnOz7qc)

[1]: https://github.com/golang/lint
[2]: https://golang.org/cmd/gofmt/
[3]: https://golang.org/doc/code.html#Library
[4]: https://github.com/gin-gonic/gin
