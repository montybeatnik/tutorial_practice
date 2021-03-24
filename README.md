# Tutorial 1
# Overview
The thought process here is to get something down with which I can work to smooth out the patterns I'll use to create something great!

## packages
* devcon

## use TLS
[reference](https://stackoverflow.com/questions/63588254/how-to-set-up-an-https-server-with-a-self-signed-certificate-in-golang)
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

