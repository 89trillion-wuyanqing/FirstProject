## 1.整体框架

本服务通过gin框架提供http的web服务，通过启动命令行进行加载ini文件启动服务器。

## 2.目录结构

```
├── app
│   ├── http
│   │   └── httpServer.go #启动http
│   └── main.go  #程序入口
├── conf  #配置文件
│   ├── app.ini
│   ├── config.army.model.json
│   └── info.json
├── go.mod 
└── internal
    ├── ctrl #控制器层
    │   └── armyController.go
    ├── handler #handler层
    │   ├── armyHandler.go
    │   └── armyHandler_test.go
    ├── model #模型层
    │   └── Army.go
    ├── router #路由层
    │   └── routerArmy.go
    └── utils #工具层
        ├── ArmyJson.go
        └── IniUtils.go
```



## 3.代码逻辑分层

| 层        | 文件夹                  | 主要职责                             | 调用关系                  | 其他说明         |
| --------- | ----------------------- | ------------------------------------ | ------------------------- | ---------------- |
| 应用层    | /app/http/httpServer.go | 进程启动，server初始化               | 调用路由层                | 不可同层相互调用 |
| 路由层    | /internal/router        | 路由转发，http的path                 | 调用控制层，被应用层调用  | 不可同层相互调用 |
| 控制层    | /internal/ctrl          | 请求参数验证、处理请求后构造回复消息 | 调用handler，被路由层调用 | 不可同层相互调用 |
| handler层 | /internal/handler       | 处理具体业务逻辑                     | 调用模型层，被控制层调用  | 不可同层相互调用 |
| 模型层    | /internal/model         | 数据模型                             | 被业务逻辑层调用          | 不可同层相互调用 |
| 工具层    | /internal/utils         | 工具封装                             | 可以被各层调用            | 不可同层相互调用 |

## 4.存储设计

士兵信息设计

| 内容         | 数据库 | field       | 类型   |
| ------------ | ------ | ----------- | ------ |
| 士兵id       | 无     | Id          | string |
| 解锁阶段     | 无     | UnlockArena | string |
| 稀有度       | 无     | Rarity      | string |
| 客户端版本号 | 无     | Cvc         | string |
| 战力         | 无     | Atk         | string |

返回结果设计

| 内容     | 数据库 | field | 类型        |
| -------- | ------ | ----- | ----------- |
| 状态码   | 无     | Code  | string      |
| 返回消息 | 无     | Msg   | string      |
| 返回结果 | 无     | Data  | interface{} |



## 5.接口设计

### 5.1第一个接口 - go实现

#### 接口功能

根据输入的稀有度参数、当前解锁阶段参数和cvc参数，获取符合稀有度、cvc且已解锁的所有士兵数据并返回。

#### HTTP请求方式

POST

#### 请求地址

localhost:8000/getArmys

#### 请求参数

| 参数        | 必选 | 类型 | 说明                                   |
| ----------- | ---- | ---- | -------------------------------------- |
| rarity      | true | text | 稀有度（参数规范是输入数字）           |
| unlockArena | true | text | 当前解锁阶段参数（参数规范是输入数字） |
| cvc         | true | text | cvc（参数规范是输入数字）              |

```json
{
  "rarity":"4",
  "unlockArena":"0",
  "cvc":"1000"
}
```



#### 请求响应

```json
{
    "code": 200,
    "data": {
        "15001": {
            "id": "15001",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "234"
        },
        "15002": {
            "id": "15002",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "319"
        },
        "15003": {
            "id": "15003",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "481"
        },
        "15004": {
            "id": "15004",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "780"
        },
        "15005": {
            "id": "15005",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "1040"
        },
        "15006": {
            "id": "15006",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "1170"
        },
        "15007": {
            "id": "15007",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "1398"
        },
        "15008": {
            "id": "15008",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "1690"
        },
        "15009": {
            "id": "15009",
            "UnlockArena": "",
            "Rarity": "4",
            "Cvc": "1000",
            "Atk": "2044"
        }
    },
    "message": "OK"
}
```

### 5.2.第二个接口 - go实现

#### 接口功能

根据输入的士兵id参数获取该士兵的稀有度并返回

#### HTTP请求方式

POST

#### 请求地址

localhost:8000/getArmyRarity

#### 请求参数

| 参数 | 必选 | 类型 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | true | text | 士兵id |

```json
{
  "id":"10501"
}
```

#### 请求响应

```json
{
    "code": 200,
    "data": "2",
    "message": "OK"
}
```

### 5.3.第三个接口 - go实现

#### 接口功能

根据输入的id参数获取士兵的战斗力并返回

#### HTTP请求方式

POST

#### 请求地址

localhost:8000/getArmyAtk

#### 请求参数

| 参数 | 必选 | 类型 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | true | int  | 士兵id |

```json
{
  "id":"10502"
}
```



#### 请求响应

```json
{
    "code": 200,
    "data": "245",
    "message": "OK"
}
```

### 5.4.第四个接口 - go实现

#### 接口功能

根据输入的cvc参数，获取所有合法的士兵并返回

#### HTTP请求方式

POST

#### 请求地址

localhost:8000/getArmysByCvc

#### 请求参数

| 参数 | 必选 | 类型 | 说明 |
| ---- | ---- | ---- | ---- |
| cvc  | true | int  | cvc  |

```json
{
  "cvc":"1900"
}
```



#### 请求响应

```json
{
    "code": 200,
    "data": {
        "18901": {
            "id": "18901",
            "UnlockArena": "5",
            "Rarity": "2",
            "Cvc": "1900",
            "Atk": "125"
        },
        "18902": {
            "id": "18902",
            "UnlockArena": "5",
            "Rarity": "2",
            "Cvc": "1900",
            "Atk": "167"
        },
        "18903": {
            "id": "18903",
            "UnlockArena": "5",
            "Rarity": "2",
            "Cvc": "1900",
            "Atk": "250"
        },
        ......，
        "19309": {
            "id": "19309",
            "UnlockArena": "",
            "Rarity": "3",
            "Cvc": "1900",
            "Atk": "1500"
        }
    },
    "message": "OK"
}
```

### 5.5第五个接口 - go实现

#### 接口功能

获取每个阶段解锁的士兵的json数据（阶段 ： 士兵结构体）

#### HTTP请求方式

GET

#### 请求地址

localhost:8000/getArmysByStage

#### 请求参数

无

#### 请求响应

```json
{
    "code": 200,
    "data": {
        "0": [
            {
                "id": "19407",
                "UnlockArena": "",
                "Rarity": "2",
                "Cvc": "2000",
                "Atk": "525"
            },
            {
                "id": "19708",
                "UnlockArena": "",
                "Rarity": "2",
                "Cvc": "2500",
                "Atk": "720"
            },
           ...... ,
            {
                "id": "18606",
                "UnlockArena": "",
                "Rarity": "1",
                "Cvc": "1500",
                "Atk": "557"
            }
        ],
        "1": [
            {
                "id": "10602",
                "UnlockArena": "1",
                "Rarity": "2",
                "Cvc": "1000",
                "Atk": "200"
            },
            {
                "id": "10405",
                "UnlockArena": "1",
                "Rarity": "1",
                "Cvc": "1000",
                "Atk": "7650"
            },
          ...... ,
            {
                "id": "10402",
                "UnlockArena": "1",
                "Rarity": "1",
                "Cvc": "1000",
                "Atk": "1120"
            }
        ],
        "2": [
            {
                "id": "11009",
                "UnlockArena": "2",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "18888"
            },
            {
                "id": "11001",
                "UnlockArena": "2",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "2500"
            },
           ......,
            {
                "id": "13204",
                "UnlockArena": "2",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "2400"
            }
        ],
        "3": [
            {
                "id": "13302",
                "UnlockArena": "3",
                "Rarity": "2",
                "Cvc": "1000",
                "Atk": "143"
            },
            {
                "id": "14504",
                "UnlockArena": "3",
                "Rarity": "2",
                "Cvc": "1000",
                "Atk": "400"
            },
          ...... ,
            {
                "id": "13309",
                "UnlockArena": "3",
                "Rarity": "2",
                "Cvc": "1000",
                "Atk": "723"
            }
        ],
        "4": [
            {
                "id": "17203",
                "UnlockArena": "4",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "296"
            },
            {
                "id": "17208",
                "UnlockArena": "4",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "1040"
            },
            ...... ,
            {
                "id": "15909",
                "UnlockArena": "4",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "469"
            }
        ],
        "5": [
            {
                "id": "18703",
                "UnlockArena": "5",
                "Rarity": "3",
                "Cvc": "1500",
                "Atk": "361"
            },
            {
                "id": "16305",
                "UnlockArena": "5",
                "Rarity": "4",
                "Cvc": "1000",
                "Atk": "5100"
            },
            ...... ,
            {
                "id": "18307",
                "UnlockArena": "5",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": ""
            }
        ],
        "6": [
            {
                "id": "18402",
                "UnlockArena": "6",
                "Rarity": "4",
                "Cvc": "1000",
                "Atk": "530"
            },
            {
                "id": "17803",
                "UnlockArena": "6",
                "Rarity": "3",
                "Cvc": "1000",
                "Atk": "119"
            },
           ......,
            {
                "id": "18409",
                "UnlockArena": "6",
                "Rarity": "4",
                "Cvc": "1000",
                "Atk": "3395"
            }
        ],
        "7": [
            {
                "id": "19103",
                "UnlockArena": "7",
                "Rarity": "4",
                "Cvc": "1900",
                "Atk": "536"
            },
            {
                "id": "19109",
                "UnlockArena": "7",
                "Rarity": "4",
                "Cvc": "1900",
                "Atk": "2125"
            },
            ......,
            {
                "id": "19101",
                "UnlockArena": "7",
                "Rarity": "4",
                "Cvc": "1900",
                "Atk": "281"
            }
        ]
    },
    "message": "OK"
}
```

#### 响应状态码

| 状态码 | 说明                             |
| ------ | -------------------------------- |
| 200    | 成功                             |
| 201    | 获取不到json文件中士兵信息       |
| 202    | json文件中解锁阶段的数据格式错误 |
| 203    | 不存在该士兵                     |
| 204    | 对应客户端版本的士兵信息不存在   |
| 205    | 请正确填写rarity参数             |
| 206    | 请正确填写cvc参数                |
| 207    | 请正确填写unlockArena参数        |
| 208    | 请正确填写id参数                 |

#### 响应消息说明

| 返回字段 | 说明         |
| -------- | ------------ |
| code     | 状态码       |
| message  | 请求状态消息 |
| data     | 返回信息     |

## 6.第三方库

### gin

```
用于http服务创建
代码  https://github.com/gin-gonic/gin
```

### pflag

```
用于解析命令行参数
代码 https://github.com/spf13/pflag
```

### ini

```
用于解析ini文件
代码 https://github.com/go-ini/ini
```



## 7.如何编译执行

在app文件目录下编译main.go文件后生成可执行文件，将可执行文件文件放到根目录下进行执行。

在项目执行可执行文件的执行命令

```sh
./main -f="conf/config.army.model.json"
```

在app目录下执行main.go时需要设置working directory为项目根路径



在单元测试代码目录下执行go test时设置working directory为项目根路径



## 8.todo

1.后续的项目结构优化。

## 流程图

![第一题流程图](第一题流程图.png)





