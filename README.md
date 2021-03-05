# corpid poc api

[TOC]

本文档主要介绍基于 chaincode api  封装的业务 api。

本文档包含两大类 api：

- 账户管理
- 业务处理

---

## 系统信息

### API HOST

> 测试环境	http://23.102.230.245:10000



---

### HTTP Request Headers 设置

api 的数据传递使用 `application/json` 的  content-type 格式。

除了 `login` 及 `token check` 接口，其他接口均需要配置一个 `X-TOKEN` 的 header 信息。

比如 setHeader('X-TOKEN', "some token value")

其中 token 值从 login 接口获取。



---

## 账户管理

### /api/jwt/user/login 用户登录接口(获取 token)

```shell
POST /api/jwt/user/login

# 输入 body
{
    "username": "orgadmin",
    "password": "passwd"
}

# 主要注意的是，初始化的管理账户密码是 orgadmin/passwd
# 其他账户的分配可以在登录了管理员账户后进行操作。

# 一个返回例子
# 其中 data.token 的值即为其他接口的 headers 中的 X-TOKEN 的值
{
    "code": 200,
    "msg": "OK",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJvcmdhZG1pbiIsInVzZXJuYW1lIjoib3JnYWRtaW4iLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2MTU4MjAwNzYsImlhdCI6MTYxNDk1NjA3Nn0.hsMeNDR0tKoSar9Ea0yT-7kqHFurfNIjjJUF6EiMqog",
        "user": {
            "id": "orgadmin",
            "_rev": "1-b2e377ee93a5a177c9e0466abc83f22e",
            "username": "orgadmin",
            "role": "admin",
            "valid": true,
            "created": {
                "second": 1614956076,
                "humanTime": "2021-03-05 14:54:36"
            },
            "updated": {
                "second": 1614956076,
                "humanTime": "2021-03-05 14:54:36"
            }
        }
    }
}
```



---

###  /api/jwt/user/create 新建用户

```shell
# 需要管理员权限

POST /api/jwt/user/create

# 输入 body
{
    "username": "devtest",
    "password": "passwd",
    "role": "client"
}

# role 有几个可选值，目前直接填写为 client 即可。
```



---

### /api/jwt/token/check token 验证接口

```shell
POST /api/jwt/token/check

# 输入 body
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJvcmdhZG1pbiIsInVzZXJuYW1lIjoib3JnYWRtaW4iLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2MTU4MjAwNzYsImlhdCI6MTYxNDk1NjA3Nn0.hsMeNDR0tKoSar9Ea0yT-7kqHFurfNIjjJUF6EiMqog"
}

# 返回例子
{
    "code": 200,
    "msg": "OK",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJvcmdhZG1pbiIsInVzZXJuYW1lIjoib3JnYWRtaW4iLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2MTU4MjAwNzYsImlhdCI6MTYxNDk1NjA3Nn0.hsMeNDR0tKoSar9Ea0yT-7kqHFurfNIjjJUF6EiMqog",
        "valid": true,
        "claim": {
            "userId": "orgadmin",
            "username": "orgadmin",
            "role": "admin",
            "exp": 1615820076,
            "iat": 1614956076
        }
    }
}
```



---

## 业务处理

三类 schema 数据，分别包含三个接口

- 新建数据接口
- 按数据 id 读取单条数据接口
- 读取本类所以数据接口

下面分别介绍各自的路径



### Credential Data Schema 相关接口

#### /api/poc/credentialdata/new

```shell
# 新建数据接口
POST /api/poc/credentialdata/new

# 输入 body
{
  "id": "697c09ac-3df2-491f-ad96-c6838768e4c1",
  "object": "ProfileCredential",
  "type": "credentialSchema",
  "createdAt": "2018-09-14T21:19:10Z",
  "schema": [
    {
      "key": "englishName",
      "datatype": "string",
      "isCredAttr": true
    },
    {
      "key": "hkid",
      "datatype": "string",
      "isCredAttr": true
    },
    {
      "key": "dateOfBirth",
      "datatype": "string",
      "isCredAttr": true
    },
    {
      "key": "email",
      "datatype": "string",
      "isCredAttr": false
    },
    {
      "key": "phone",
      "datatype": "string",
      "isCredAttr": false
    },
    {
      "key": "address",
      "datatype": "string",
      "isCredAttr": false
    }
  ]
}

# 成功时返回例子
{
    "code": 200,
    "msg": "OK",
    "data": "697c09ac-3df2-491f-ad96-c6838768e4c1"
}
```



---

#### /api/poc/credentialdata/info/{id}  按 id 读取数据

```shell
GET /api/poc/credentialdata/info/{id}

# 比如 
GET /api/poc/credentialdata/info/697c09ac-3df2-491f-ad96-c6828768e3c1

# 成功时返回例子
{
    "code": 200,
    "msg": "OK",
    "data": {
        "id": "697c09ac-3df2-491f-ad96-c6828768e3c1",
        "type": "credentialSchema",
        "object": "ProfileCredential",
        "createdAt": "2018-09-14T21:19:10Z",
        "schema": [
            {
                "key": "englishName",
                "datatype": "string",
                "isCredAttr": true
            },
            {
                "key": "hkid",
                "datatype": "string",
                "isCredAttr": true
            },
            {
                "key": "dateOfBirth",
                "datatype": "string",
                "isCredAttr": true
            },
            {
                "key": "email",
                "datatype": "string",
                "isCredAttr": false
            },
            {
                "key": "phone",
                "datatype": "string",
                "isCredAttr": false
            },
            {
                "key": "address",
                "datatype": "string",
                "isCredAttr": false
            }
        ]
    }
}
```



---

#### /api/poc/credentialdatas 读取全部数据

```shell
GET api/poc/credentialdatas

# 返回例子
# 注意，目前是全部返回数据，无分页，无总计

```



---

### Authentication DID 相关接口

#### /api/poc/authenticationdid/new

```shell
POST /api/poc/authenticationdid/new

# 输入 body
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:corpidpoc:55de6b88-352d-4d7b-829d-74aa2cb4d963",
  "type": "Issuer",
  "created": "2018-09-14T21:19:10Z",
  "name": "XYZ Issuer",
  "publicKey": "3052301006072a8648ce3d020106052b81040003033e00040308df471044339418da340406ce4d389bd0e7bb28993d38494bf63f4e5d3746ce7fd0e830522ae2c1f3793466c85db55fd9adb7dce95f3d1b13a000"
}
```



---

#### /api/poc/authenticationdid/info/{id} 按 id 读取数据

```shell
GET /api/poc/authenticationdid/info/{id}

# 比如
GET /api/poc/authenticationdid/info/did:corpidpoc:55de6b88-352d-4d7b-829d-74aa2cb4d963

# 返回例子
{
    "code": 200,
    "msg": "OK",
    "data": {
        "id": "did:corpidpoc:55de6b88-352d-4d7b-829d-74aa2cb4d963",
        "@context": "https://www.w3.org/ns/did/v1",
        "type": "Issuer",
        "created": "2018-09-14T21:19:10Z",
        "name": "XYZ Issuer",
        "publicKey": "3052301006072a8648ce3d020106052b81040003033e00040308df471044339418da340406ce4d389bd0e7bb28993d38494bf63f4e5d3746ce7fd0e830522ae2c1f3793466c85db55fd9adb7dce95f3d1b13a000"
    }
}
```



---

#### /api/poc/authenticationdids 读取全部数据

```shell
GET /api/poc/authenticationdids
```



---

### Credential DID

#### /api/poc/credentialdid/new

```shell
POST /api/poc/credentialdid/new

# 输入 body
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:corpidpoc:55de6b88-352d-4d7b-829d-74aa2bb4d965",
  "type": [
    "VerifiableCredential",
    "CorpIDCredential"
  ],
  "created": "2021-02-25T16:00:00Z",
  "credentialStatus": "Issued",
  "hash": "688c5af1a36630048e56b95d974dda655818b924d9edb1c96323b6e0e3265b98",
  "holderSignature": "3040021e15700d55911e105692f23108b7d44d7ec4f47f4a87d148ab48f7d6472ebd021e12c2fb7d2265b0b91efa914b2460cab45dbe9ecf4cc40f73ef9217cf20fa",
  "issuerSignature": "3052301006072a8648ce3d020106052b81040003033e000467413ed92a18c0bccada7a7bc210f8e002aee7ae0f4d97c51e17f53674ab07176fc20b62d60ed9468735943622328e59804623ff2c969605108c87f2"
}
```



---

#### /api/poc/credentialdid/info/{id}

```shell
GET /api/poc/credentialdid/info/{id}

# 例子
GET /api/poc/credentialdid/info/did:corpidpoc:55de6b88-352d-4d7b-829d-74aa2bb4d965
```



---

#### /api/poc/credentialdids

```shell
GET /api/poc/credentialdids
```

