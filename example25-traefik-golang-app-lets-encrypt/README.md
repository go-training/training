# Traefik Let's Encrypt & Docker

## Networking

```sh
$ docker network create web
```

## Start the traefik

Now, let's create a directory on the server where we will configure the rest of Traefik:

```sh
$ mkdir -p /opt/traefik
$ cp traefik/docker-compose.yml /opt/traefik/
$ cp traefik/traefik.toml /opt/traefik/
$ touch /opt/traefik/acme.json && chmod 600 /opt/traefik/acme.json
```

Run the traefik service

```sh
$ cd /opt/traefik && docker-compos up -d
```

## Start golang appa service

replace the following host name in docker-compose.yml

```
- "traefik.basic.frontend.rule=Host:demo1.ggz.tw"
- "traefik.basic.frontend.rule=Host:demo2.ggz.tw"
```

and start the app service

```sh
$ docker-compose up -d
```

## Test your app

```sh
$ curl -v https://demo1.ggz.tw
* Rebuilt URL to: https://demo1.ggz.tw/
*   Trying 54.255.136.214...
* Connected to demo1.ggz.tw (54.255.136.214) port 443 (#0)
* found 148 certificates in /etc/ssl/certs/ca-certificates.crt
* found 592 certificates in /etc/ssl/certs
* ALPN, offering http/1.1
* SSL connection using TLS1.2 / ECDHE_RSA_AES_128_GCM_SHA256
* 	 server certificate verification OK
* 	 server certificate status verification SKIPPED
* 	 common name: demo1.ggz.tw (matched)
* 	 server certificate expiration date OK
* 	 server certificate activation date OK
* 	 certificate public key: RSA
* 	 certificate version: #3
* 	 subject: CN=demo1.ggz.tw
* 	 start date: Sun, 20 Jan 2019 05:40:20 GMT
* 	 expire date: Sat, 20 Apr 2019 05:40:20 GMT
* 	 issuer: C=US,O=Let's Encrypt,CN=Let's Encrypt Authority X3
* 	 compression: NULL
* ALPN, server accepted to use http/1.1
> GET / HTTP/1.1
> Host: demo1.ggz.tw
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
< Date: Sun, 20 Jan 2019 08:25:01 GMT
<
* Connection #0 to host demo1.ggz.tw left intact
I love !, Hello World, traefik workshop!
```
