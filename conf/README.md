# 自签CA证书、server、client双向认证

> fix go1.15bug
>
> openssl 生成证书上 grpc 报 legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
> 最近用传统的方式 生成的证书上用golang 1.15. 版本 grpc 通信报上面错误

参考：

> https://www.cnblogs.com/jackluo/p/13841286.html
> https://blog.csdn.net/ma_jiang/article/details/111992609

> https://segmentfault.com/a/1190000016601810
> https://www.cnblogs.com/devhg/p/13751770.html



## 创建双方SAN证书

> go version go1.15.3 darwin/amd64
> 上面调用的时候报错了

> rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
>
> 如果出现上述报错，是因为 go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 如果想兼容之前的方式，需要设置环境变量 GODEBUG 为 x509ignoreCN=0。

 

### 什么是 SAN
SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

下面简单示例如何用 openssl 生成 ca 和双方 SAN 证书。



### 创建CA证书

根证书 根证书（root certificate）是属于根证书颁发机构（CA）的公钥证书。我们可以通过验证 CA 的签名从而信任 CA ，任何人都可以得到 CA 的证书（含公钥），用以验证它所签发的证书（客户端、服务端）

它包含的文件如下：

```shell
➜  conf git:(master) ✗ openssl                                                                                                  
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

 

### 修改配置

准备默认 OpenSSL 配置文件于当前目录

* linux系统 : /etc/pki/tls/openssl.cnf
* Mac系统: /System/Library/OpenSSL/openssl.cnf
* Windows：安装目录下 openssl.cfg 比如 D:\Program Files\OpenSSL-Win64\bin\openssl.cfg

1. 拷贝配置文件到项目 然后修改
   `cp /System/Library/OpenSSL/openssl.cnf ./`
2. 找到 [ CA_default ]，打开 copy_extensions = copy
3. 找到[ req ]，打开 req_extensions = v3_req 
4. 找到[ v3_req ]，添加 subjectAltName = @alt_names

5. 添加新的标签 [ alt_names ] , 和标签字段

```
[ alt_names ]
DNS.1 = localhost
DNS.2 = *.custer.fun
```


接着使用这个临时配置生成证书：



### 生成Server

```bash
$ openssl genpkey -algorithm RSA -out server.key
 
$ openssl req -new -nodes -key server.key -out server.csr -days 3650 -subj "/C=cn/OU=devhg/O=devhg/CN=localhost" -config ./openssl.cnf -extensions v3_req
 
$ openssl x509 -req -days 3650 -in server.csr -out server.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
```

输出
```
OpenSSL> genpkey -algorithm RSA -out server.key
..........................................................................+++
..............................................................................+++
OpenSSL> req -new -nodes -key server.key -out server.csr -days 3650 -subj "/C=cn/OU=devhg/O=devhg/CN=localhost" -config ./openssl.cnf -extensions v3_req
OpenSSL> x509 -req -days 3650 -in server.csr -out server.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
Signature ok
subject=/C=cn/OU=devhg/O=devhg/CN=localhost
Getting CA Private Key
```



### 生成client

```bash
$ openssl genpkey -algorithm RSA -out client.key
 
$ openssl req -new -nodes -key client.key -out client.csr -days 3650 -subj "/C=cn/OU=devhg/O=devhg/CN=localhost" -config ./openssl.cnf -extensions v3_req
 
$ openssl x509 -req -days 3650 -in client.csr -out client.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
```

输出
```
OpenSSL> genpkey -algorithm RSA -out client.key
..................+++
...........................................................................................+++
OpenSSL> req -new -nodes -key client.key -out client.csr -days 3650 -subj "/C=cn/OU=devhg/O=devhg/CN=localhost" -config ./openssl.cnf -extensions v3_req
OpenSSL> x509 -req -days 3650 -in client.csr -out client.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
Signature ok
subject=/C=cn/OU=devhg/O=devhg/CN=localhost
Getting CA Private Key
OpenSSL> 
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

# 