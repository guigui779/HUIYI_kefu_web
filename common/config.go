package common

import (
	"encoding/json"
	"goflylivechat/tools"
	"io/ioutil"
	"os"
)

type Mysql struct {
	Server   string
	Port     string
	Database string
	Username string
	Password string
}

func GetMysqlConf() *Mysql {
	var mysql = &Mysql{}
	// 优先从环境变量读取（Railway 等云平台）
	host := getEnv("MYSQL_HOST", "MYSQLHOST")
	if host != "" {
		mysql.Server = host
		mysql.Port = getEnv("MYSQL_PORT", "MYSQLPORT")
		mysql.Database = getEnv("MYSQL_DATABASE", "MYSQLDATABASE")
		mysql.Username = getEnv("MYSQL_USER", "MYSQLUSER")
		mysql.Password = getEnv("MYSQL_PASSWORD", "MYSQLPASSWORD")
		return mysql
	}
	// 本地开发从 mysql.json 读取
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return mysql
	}
	err = json.Unmarshal(info, mysql)
	return mysql
}

func getEnv(keys ...string) string {
	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}
