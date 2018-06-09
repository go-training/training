# Deploy golang app using drone with Kubernetes

See the [chinese blog](https://blog.wu-boy.com/2018/06/drone-kubernetes-with-golang/)

## write simple golang app

see the [main.go](./main.go)

```go
package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var version = "master"

func showVersion(w http.ResponseWriter, r *http.Request) {
	log.Println(version)
	w.Write([]byte(version))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello, got the message: " + message
	log.Println(message)
	w.Write([]byte(message))
}

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	http.HandleFunc("/version", showVersion)
	http.HandleFunc("/", sayHello)
	log.Println("Listen server on " + port + " port")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
```

write app version using `-ldflags`. see the following example:

```makefile
build:
ifneq ($(DRONE_TAG),)
    go build -v -ldflags "-X main.version=$(DRONE_TAG)" -a -o release/linux/amd64/hello
else
    go build -v -ldflags "-X main.version=$(DRONE_COMMIT)" -a -o release/linux/amd64/hello
endif
```

how to build binary using drone?

```yaml
pipeline:
  build_linux_amd64:
    image: golang:1.10
    group: build
    environment:
      - GOOS=linux
      - GOARCH=amd64
      - CGO_ENABLED=0
    commands:
      - cd example19-deploy-with-kubernetes && make build
```

## upload to dockerhub

See the drone config:

```yaml
docker_golang:
  image: plugins/docker:17.12
  secrets: [ docker_username, docker_password ]
  repo: appleboy/golang-http
  dockerfile: example19-deploy-with-kubernetes/Dockerfile
  default_tags: true
  when:
    event: [ push, tag ]
```

## rolling update using kubernetes

Please preare the following drone secrets

1. KUBERNETES_SERVER
2. KUBERNETES_CERT
3. KUBERNETES_TOKEN

then see the drone config

```yaml
deploy:
  image: sh4d1/drone-kubernetes
  kubernetes_template: example19-deploy-with-kubernetes/deployment.yml
  kubernetes_namespace: default
  secrets: [ kubernetes_server, kubernetes_cert, kubernetes_token ]
```

see the `deployment.yaml`

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: frontend
  # these labels can be applied automatically
  # from the labels in the pod template if not set
  labels:
    app: gotraining
    tier: frontend
spec:
  # this replicas value is default
  # modify it according to your case
  replicas: 3
  # selector can be applied automatically
  # from the labels in the pod template if not set
  # selector:
  #   app: guestbook
  #   tier: frontend
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
  minReadySeconds: 5
  template:
    metadata:
      labels:
        app: gotraining
        tier: frontend
    spec:
      containers:
      - name: go-hello
        image: appleboy/golang-http:VERSION
        imagePullPolicy: Always
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        ports:
        - containerPort: 8080
        env:
        - name: FOR_GODS_SAKE_PLEASE_REDEPLOY
          value: 'THIS_STRING_IS_REPLACED_DURING_BUILD'
```
