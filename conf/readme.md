# 自签CA证书、server、client双向认证

> https://segmentfault.com/a/1190000016601810
> https://www.cnblogs.com/devhg/p/13751770.html

## CA

根证书 根证书（root certificate）是属于根证书颁发机构（CA）的公钥证书。我们可以通过验证 CA 的签名从而信任 CA ，任何人都可以得到 CA 的证书（含公钥），用以验证它所签发的证书（客户端、服务端）

它包含的文件如下：

* 公钥
* 密钥

```shell
➜  grpc-demo git:(master) ✗ openssl                                                                                                     
OpenSSL> genrsa -out ca.key 2048
Generating RSA private key, 2048 bit long modulus
...........................................................................................................+++
........+++
e is 65537 (0x10001)
OpenSSL> req -new -x509 -days 7200 -key ca.key -out ca.pem
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:  
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:ihuidev 
Organizational Unit Name (eg, section) []:ihuidev
Common Name (eg, fully qualified host name) []:localhost
Email Address []:
OpenSSL> 
```

## Server

创建server.key，生成 CSR CSR 是 Cerificate Signing Request 的英文缩写，为证书请求文件。主要作用是 CA 会利用 CSR 文件进行签名使得攻击者无法伪装或篡改原有证书

```shell
OpenSSL> genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus
....................+++
.........................................................................+++
e is 65537 (0x10001)
OpenSSL> req -new -key server.key -out server.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:ihuidev
Organizational Unit Name (eg, section) []:ihuidev
Common Name (eg, fully qualified host name) []:localhost
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
OpenSSL> 
```

基于 CA 签发

```shell
OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
Signature ok
subject=/O=ihuidev/OU=ihuidev/CN=localhost
Getting CA Private Key
```

## Client

```shell
OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
Signature ok
subject=/O=ihuidev/OU=ihuidev/CN=localhost
Getting CA Private Key
OpenSSL> 
OpenSSL> 
OpenSSL> 
OpenSSL> ecparam -genkey -name secp384r1 -out client.key
OpenSSL> req -new -key client.key -out client.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:ihuidev
Organizational Unit Name (eg, section) []:ihuidev
Common Name (eg, fully qualified host name) []:localhost
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
OpenSSL> 
```

基于 CA 签发

```shell
OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
Signature ok
subject=/O=ihuidev/OU=ihuidev/CN=localhost
Getting CA Private Key
```

## 整理目录

```
➜  conf git:(master) ✗ tree                  
.
├── ca.key
├── ca.pem
├── ca.srl
├── client
│   ├── client.csr
│   ├── client.key
│   └── client.pem
├── readme.md
└── server
    ├── server.csr
    ├── server.key
    └── server.pem

```



# 自签证书

> https://www.cnblogs.com/devhg/p/13751770.html

```shell
openssl
genrsa -des3 -out server.key 2048
req -new -key server.key -out server.csr
rsa -in server.key -out server_no_password.key
x509 -req -days 3650 -in server.csr -signkey server_no_password.key -out server.crt
```



# Bug

openssl 生成证书上 grpc 报 legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
最近用传统的方式 生成的证书上用golang 1.15. 版本 报 grpc 上面

[fix go1.15 bug](./go1.15+bug.md)