package model

type Service struct {
	AppModel string `json:"AppModel"`
	HTTPPort int    `json:"HttpPort"`
}

type Redis struct {
	RedisAddr   string `json:"RedisAddr"`
	RedisPort   int    `json:"RedisPort"`
	RedisPW     string `json:"RedisPW"`
	RedisDBName int    `json:"RedisDBName"`
}

type Mysql struct {
	DBDrive string `json:"DBDrive"`
	DBAddr  string `json:"DBAddr"`
	DBPort  int    `json:"DBPort"`
	DBUser  string `json:"DBUser"`
	Dbpw    string `json:"DBPW"`
	DBName  string `json:"DBName"`
}

type Config struct {
	Service Service `json:"service"`
	Redis   Redis   `json:"redis"`
	Mysql   Mysql   `json:"Mysql"`
}
