## 练习项目--备忘录

### 使用 gin  gorm  jwt

`go get github.com/dgrijalva/jwt-go`

`go get github.com/sirupsen/logrus`

### 数据库 redis  mysql 


### 使用

API URL:`http://127.0.0.1:8081`

用户注册接口：POST `/api/v1/user/register`

body参数:`{
        "user_name": "Jayson001",  
        "password": "123456" 
         }`
    
用户登录接口：POST `/api/v1/user/login`   

body参数:`{
            "user_name": "Jayson003",
            "password": "123456"
        }` 
        
备忘录创建接口：POST `/api/v1/task`   

header参数（token）：`Authorization: ${Token}`

body参数:`{
            "title": "Test001",
            "content": "test001...."
        }` 
        
查看备忘录接口：Get `/api/v1/task/:id`   

header参数（token）：`Authorization: ${Token}`

查看所有备忘录接口：Get `/api/v1/tasks`   

header参数（token）：`Authorization: ${Token}`