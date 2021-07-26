# 技术文档

### 1.整体框架

解析ini配置文件，获取端口，再通过pflag从命令行传入json文件的路径，解析json文件获取数据，获取数据之后，根据客户端传来的请求参数，查询符合条件的数据，并返回给客户端。



### 2.目录结构

```
├── app
│   ├── httpserver
│   │   └── http.go
│   └── main.go
├── config
│   ├── NewConfig.json
│   ├── app.ini
│   └── config.army.model.json
├── firstmarkdown.md
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── soldierController.go
│   ├── model
│   │   └── soldierStruct.go
│   ├── router
│   │   └── routers.go
│   ├── service
│   │   ├── method.go
│   │   └── method_test.go
│   └── soldierEorrer
│       └── error.go
├── locust.py
├── report.html
└── utils
    ├── loadJson.go
    ├── pflagJson.go
    └── randIni.go

```



### 3.代码逻辑分层

| 层        | 文件夹                              | 主要职责                          | 调用关系                  | 其它说明     |
| --------- | ----------------------------------- | --------------------------------- | ------------------------- | ------------ |
| 应用层    | app/httpserver/http.go              | 启动服务器                        | 调用路由层                | 不可同层调用 |
| 路由层    | internal/router/routers.go          | 转发路由                          | 被应用层调用，调用控制层  | 不可同层调用 |
| 控制层    | internal/ctrl/soldierController.go  | 处理校验请求参数                  | 被路由层调用，调用handler | 不可同层调用 |
| Service层 | internal/service/method.go          | 处理具体业务                      | 被控制层调用              | 不可同层调用 |
| 工具层    | utils/loadJson.go、utils/randIni.go | 读取ini配置文件和处理json文件数据 | 被service调用             | 不可同层调用 |
| 压力测试  | locust.py                           | 进行压力测试                      | 无调用关系                | 不可同层调用 |
| 文件包    | config                              | 统一存放配置文件                  | 无调用关系                | 不可同层调用 |



### 4.存储设计

```go
//士兵结构体
type Soldier struct {
	Id           string  //id
	Rarity       string  //稀有度
	UnlockArena  string  //解锁阶段
	CombatPoints string  //战斗力
	Cvc          string  //客户端版本号
}
```

采用map存储数据，key为string类型的士兵id,value为Soldier结构体



### 5.接口设计

#### 1.请求符合输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵

##### 请求方式

http GET

##### 接口地址

localhost:8000/getLicitSoldier

##### 请求参数

| Key    | Value |
| ------ | ----- |
| rarity | 1     |
| unK    | 0     |
| cvc    | 1000  |

##### 请求响应

```json
{
    "Code": 200,
    "Data": {
        "10101": {
            "Id": "10101",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "167",
            "Cvc": "1000"
        },
        ....
    "Message": "ok"
}
```



#### 2.输入士兵id获取对应稀有度

##### 请求方式

http GET

##### 接口地址

localhost:8000/getRarity

##### 请求参数

| Key  | Value |
| ---- | ----- |
| Id   | 10101 |

##### 请求响应

```json
{
    "Code": 200,
    "Data": "1",
    "Message": "ok"
}
```



#### 3.输入士兵id获取对应战力

##### 请求方式

http GET

##### 接口地址

localhost:8000/getCombatPoints

##### 请求参数

| Key  | Value |
| ---- | ----- |
| Id   | 10101 |

##### 请求响应

```json
{
    "Code": 200,
    "Data": "167",
    "Message": "ok"
}
```



#### 4.输入cvc获取所有符合条件的士兵

##### 请求方式

http GET

##### 接口地址

localhost:8000/getCvcLicitSoldier

##### 请求参数

| Key  | Value |
| ---- | ----- |
| Cvc  | 2500  |

##### 请求响应

```json
{
    "Code": 200,
    "Data": {
        "19701": {
            "Id": "19701",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "256",
            "Cvc": "2500"
        },
        ...
    "Message": "ok"
}
```



#### 5.获取每个阶段解锁相应士兵的json数据

##### 请求方式

http GET

##### 接口地址

localhost:8000/getUnkSoldier

##### 请求参数

无

##### 请求响应

```json
{
    "Code": 200,
    "Data": {
       "1": [
            {
                "Id": "10401",
                "Rarity": "1",
                "UnlockArena": "1",
                "CombatPoints": "181",
                "Cvc": "1000"
            },
         ]
        ...
    "Message": "ok"
}
```

### 响应状态码

| 状态码 | 说明                     |
| ------ | ------------------------ |
| 200    | 成功                     |
| 1001   | 有参数为空               |
| 500    | 服务器内部错误           |
| 1002   | 数据为空，请检查输入参数 |



## 6、第三方库

### props/ini

```
github.com/tietang/props/ini
```

### pflag

```
https://github.com/spf13/pflag
```



### 7.如何编译执行

```
//编译可执行文件
go build ./app/main.go 
```

```
//编译执行
./main --filename="./config/config.army.model.json"
```



### 8.流程图

![未命名文件 (9)](https://user-images.githubusercontent.com/87186547/126963267-1d4edc21-eccb-4216-9375-546dae8f4776.jpg)

