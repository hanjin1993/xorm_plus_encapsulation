package mysql

import (
	"go_batch_create_orders/utils"
)

type Options func(*Mysql)

type MysqlSystem struct {
	Types     string `json:"types"`
	Prefix    string `json:"prefix"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Port      string `json:"port"`
	Database  string `json:"database"`
	ParseTime string `json:"parseTime"`
	Charset   string `json:"charset"`
}

type Mysql struct {
	MysqlResoure string `json:"mysqlResoure"`
	MysqlSystem
	Runmode string `json:"runmode"`
}

func New() *Mysql {
	//初始化 mysql链接信息
	system := MysqlSystem{
		Types:     "",
		Prefix:    "",
		UserName:  "root",
		Password:  "root",
		Address:   "127.0.0.1",
		Port:      "3306",
		Database:  "test",
		ParseTime: "true",
		Charset:   "utf8mb4",
	}
	m := &Mysql{
		MysqlResoure: "",
		MysqlSystem:  system,
		Runmode:      "",
	}

	options := m.SetOptions("database")
	//fmt.Println(options)
	data := m.SetData(options)

	//查询当前的是否有状态配置信息
	config := utils.Config{Module: "database", Filed: "runmode"}
	fieldData := config.Config()
	if fieldData != "" {
		setOptions := data.SetOptions(fieldData)
		data = data.SetData(setOptions)
	}

	return data
}

func (m *Mysql) SetOptions(module string) []Options {
	config := utils.Config{Module: module}
	list := config.ConfigList()
	value := ""
	configs := utils.Config{}
	options := make([]Options, 0, len(list))
	for _, v := range list {
		switch v {
		case "types":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetTypes(value))
			break
		case "prefix":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetPrefix(value))
			break
		case "userName":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetUserName(value))
			break
		case "password":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetPassword(value))
			break
		case "address":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetAddress(value))
			break
		case "port":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetPort(value))
			break
		case "database":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetDatabase(value))
			break
		case "parseTime":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetParseTime(value))
			break
		case "mysqlResoure":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetMysqlResoure(value))
			break
		case "runmode":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetRunmode(value))
			break
		case "charset":
			configs = utils.Config{Module: module, Filed: v}
			value = configs.Config()
			options = append(options, m.SetCharset(value))
			break
		}

	}
	return options
}

func (m *Mysql) SetData(option []Options) *Mysql {
	for _, v := range option {
		v(m)
	}
	return m
}

func (mysql *Mysql) SetCharset(charset string) Options {
	return func(mysql *Mysql) {
		mysql.Charset = charset
	}
}

func (mysql *Mysql) SetTypes(types string) Options {
	return func(mysql *Mysql) {
		mysql.Types = types
	}
}

func (mysql *Mysql) SetPrefix(prefix string) Options {
	return func(mysql *Mysql) {
		mysql.Prefix = prefix
	}
}

func (mysql *Mysql) SetUserName(userName string) Options {
	return func(mysql *Mysql) {
		mysql.UserName = userName
	}
}

func (mysql *Mysql) SetPassword(password string) Options {
	return func(mysql *Mysql) {
		mysql.Password = password
	}
}

func (mysql *Mysql) SetAddress(address string) Options {
	return func(mysql *Mysql) {
		mysql.Address = address
	}
}

func (mysql *Mysql) SetPort(port string) Options {
	return func(mysql *Mysql) {
		mysql.Port = port
	}
}

func (mysql *Mysql) SetDatabase(database string) Options {
	return func(mysql *Mysql) {
		mysql.Database = database
	}
}

func (mysql *Mysql) SetParseTime(parseTime string) Options {
	return func(mysql *Mysql) {
		mysql.ParseTime = parseTime
	}
}

func (mysql *Mysql) SetMysqlResoure(mysqlResoure string) Options {
	return func(mysql *Mysql) {
		mysql.MysqlResoure = mysqlResoure
	}
}

func (mysql *Mysql) SetRunmode(runmode string) Options {
	return func(mysql *Mysql) {
		mysql.Runmode = runmode
	}
}
