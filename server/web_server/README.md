# Web后端

## 一、接口概览
|  接口名  |  访问地址  | 简介 | 操作 |
|  ----  | ----  | ---- | ---- |
| 注册  | /register | 用户注册 | GET |
| 登录  | /login | 用户登录 | POST |
| 登出  | /logout | 用户登出 | POST |
| 信息  | /info | 获取用户信息与车辆状态 | GET |
| 充值  | /recharge | 用户充值 | PUT |
| 车牌操作 | /operator | 车牌增删改操作 | PUT |

## 二、接口详情
### 1. 注册
* 访问方式：`GET` -- `/register`
* 功能：用于用户注册(注册成功后需再次登录)
* 传参：
<br> `username` -- 用户名
<br> `password` -- 密码
* 返回：`{result}` -- 注册结果
* 注册结果：
<br>0 -- 注册成功
<br>1 -- 注册失败(用户已存在)
### 2. 登录
* 访问方式：`POST` -- `/login`
* 功能：用于用户登录
* 传参：
<br> `username` -- 用户名
<br> `password` -- 密码
* 返回：`{uid, result}` -- uid与登录结果
* 登录结果：
<br>0 -- 登录成功
<br>1 -- 登录失败(用户不存在)
<br>2 -- 登录失败(密码错误)
### 3. 登出
* 访问方式：`POST` -- `/logout`
* 功能：用于用户登出
* 传参：无
* 返回：无
### 4. 信息
* 访问方式：`POST` -- `/login`
* 功能：用于获取用户信息与车辆信息
* 传参：无
* 返回：`{{车牌信息}, ...(一或多组)}` -- 多组车辆信息
* 车牌信息：
```
License      string `json:"license"`        车牌
PTemperature int32  `json:"p_temperature"`  停车场温度
PHumidity    int32  `json:"p_humidity"`     停车场湿度
PWeather     int32  `json:"p_weather"`      停车场天气
PAddress     string `json:"p_address"`      停车场地址
SID          int32  `json:"s_id"`           车位ID
STemperature int32  `json:"s_temperature"`  车位温度
SHumidity    int32  `json:"s_humidity"`     车位湿度
SAlarm       int32  `json:"s_alarm"`        车位警告
```

### 5. 充值
* 访问方式：`PUT` -- `/recharge`
* 功能：用于用户充值余额
* 传参：
<br> `amount` -- 充值金额
* 返回：`{balance}` -- 用户最新余额
### 6. 车牌操作
* 访问方式：`PUT` -- `/operator`
* 功能：用于用户操作车辆
* 传参：
<br> `operation` -- 操作
<br> `license` -- 目标车牌
<br> `newLicense` -- 新车牌
* 返回：`{result}` -- 操作结果
* 操作：
<br> `add` -- 增加车牌
<br> `delete` -- 删除车牌
<br> `change` -- 更新车牌
* 操作结果：
<br>0 -- 增加成功
<br>1 -- 增加失败(车牌已存在)
<br>2 -- 增加失败(用户不存在)
<br>3 -- 删除成功
<br>4 -- 删除失败(车牌已入场)
<br>5 -- 删除失败(车牌不存在)
<br>6 -- 删除失败(用户不存在)
<br>7 -- 更新成功
<br>8 -- 更新失败(车牌已入场)
<br>9 -- 更新失败(车牌不存在)
<br>10 -- 更新失败(用户不存在)