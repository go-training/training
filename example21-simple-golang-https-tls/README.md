# Simple https-tls server example

## Generate SSL Certificate

Install [mkcert](https://github.com/FiloSottile/mkcert) tool.


```sh
$ mkcert -install
Using the local CA at "/Users/xxxxx/Library/Application Support/mkcert" ✨
```

create `localhost` certificate:

```sh
$ mkcert myapp.dev example.com
Using the local CA at "/Users/xxxxxx/Library/Application Support/mkcert" ✨

Created a new certificate valid for the following names 📜
 - "example.com"
 - "myapp.dev"

The certificate is at "./example.com+1.pem" and the key at "./example.com+1-key.pem" ✅
```
