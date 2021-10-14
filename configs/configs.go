package configs

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfigs struct {
	App App `json:"app"`
}

type MysqlConfigs struct {
	Mysql Mysql `json:"mysqldb"`
}

type Mysql struct {
	Driver         string `json:"driver"`
	Host           string `json:"host"`
	Database       string `json:"db_name"`
	Timeout        int    `json:"timeout"`
	Dial_Timeout   int    `json:"dial_timeout"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	MaxIdleConns   int    `json:"max_idle_conns"`
	MaxOpenConns   int    `json:"max_open_conns_int"`
	MaxLifeTimeout int    `json:"max_life_timeout"`
	SSL            string `json:"ssl_mode"`
	URI            string `json:"uri"`
}

type App struct {
	Name string `json:"app_name"`
	Port string `json:"port"`
}

func ReadConfigs() (*Mysql, *App, error) {
	// read file

	data, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return nil, nil, err
	}

	// unmarshaling
	var mysql MysqlConfigs
	var app AppConfigs

	err = json.Unmarshal(data, &mysql)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(data, &app)

	if err != nil {
		return nil, nil, err
	}

	return &mysql.Mysql, &app.App, nil

}
