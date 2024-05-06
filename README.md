## Create CA Certificate
```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./ca.key -out ./ca.crt -subj "/CN=backup.local/O=kubedb"
```

## Create certificate for Server:

### Create a CSR and server certificate with above `ca.key` and `ca.crt`

```bash
$ openssl req -newkey rsa:2048 -nodes -keyout server.key -out server.csr -subj "/CN=backup.local"
$ openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extfile ./altsubj.ext
```

## Create certificate for Client:

### Create a CSR and client certificate with above `ca.key` and `ca.crt`
```bash
$ openssl req -newkey rsa:2048 -nodes -keyout client.key -out client.csr -subj "/CN=clients"
$ openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extfile altsubj.ext
```
## CURL
```bash
$ curl  https://backup.local:8443 -X GET  --cacert ./certs/ca.crt
```
Note that, In the above command , We can skip `--cert ./certs/client-cert.pem --key ./certs/client-key.pem` this part, if the server was running with

- wither directly like this `srv.ListenAndServeTLS(s.cfg.CertFile, s.cfg.KeyFile))`
- or, `TLSConfig.ClientAuth = tls.NoClientCert(this is the default value)` is set.
- In other cases, like `VerifyClientCertIfGiven / RequestClientCert`, client flags are mandatory in curl command.

## Update Your newly created CA

```bash
$ cd /usr/local/share/ca-certificates
$ sudo vim test-ca.crt ## add here out newly created ca.crt

$ sudo update-ca-certificates
```

## After updating the default CA directory
```bash
$ curl  https://backup.local:8443
```

