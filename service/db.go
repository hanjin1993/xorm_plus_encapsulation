package service

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"go_batch_create_orders/utils/mysql"
	"regexp"
	"xorm.io/core"
)

type Db struct {
}

type Query struct {
	Resouces interface{}
	Where    string
	Join     [][4]string
	Page     int
	PageSize int
	OrderBy  string
	GroupBy  string
	Alias    string
	Table    string
	Unscoped bool //是否查询软性删除的数据
}

type SessionTruncation struct {
	xorm.Session
}

type Update struct {
	Resouces interface{}
	Where    string
	Table    string
}

type Delete struct {
	Resouces interface{}
	Where    string
	Table    string
	Unscoped bool //是否进行硬性删除 软删除需要配合 deleted标识
}

type Add struct {
	Resouces interface{}
	Table    string
}

var engine *xorm.Engine
var err error
var config *mysql.Mysql
var session *xorm.Session
var transactionSwitch bool

func init() {
	//首先加载数据库配置文件
	m := mysql.New()
	source := ""
	if m.Database != "" {
		//"root:2123wqwe@tcp(127.0.0.1:3306)/learn_go?charset=utf8"
		source = m.UserName + ":" + m.Password + "@tcp(" + m.Address + ":" + m.Port + ")/" + m.Database + "?charset=" + m.Charset
	} else {
		source = m.MysqlResoure
	}
	fmt.Println(m.Types)
	engine, err = xorm.NewMySQL(m.Types, source)
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.SetMaxIdleConns(5)  //设置最大空闲数
	engine.SetMaxOpenConns(10) //设置最大连接数
	mapper := core.NewPrefixMapper(core.SnakeMapper{}, m.Prefix)
	engine.SetTableMapper(mapper)
	engine.ShowSQL()
	config = m
	transactionSwitch = false //默认事务开关是关闭的
}

func (add *Add) Insert() (res int64, err error) {
	res, err = GetTableSession(add.Table).Insert(add.Resouces)
	return
}

func InitQuery() *Query {
	return &Query{
		Page:     1,
		PageSize: 10,
	}
}

func AddTablesPrefix(table string) string {
	return config.Prefix + table
}

//,join [][3]string,page int,pageSize int,orderBy string,groupBy string  where=>{'id':{'=',1},'age':{{'>',10},{'<',20},'and'}}
func (q *Query) Select() (err error) {
	//res := make([]resouces,0)
	table := GetTableSession(q.Table)
	table = table.Where(q.Where)
	//err = table.Where(where).Join("left", "go_order_goods","go_order_goods" ,"go_order_goods.order_id = go_order.id").
	//	Join("inner", "go_users","go_users", "go_users.id = go_order.user_id").Find(resouces)
	if len(q.Join) > 0 {
		for _, v := range q.Join {
			table.Join(v[0], AddTablesPrefix(v[1])+" as "+v[2], v[3])
		}
	}
	err = table.Alias(q.Alias).OrderBy(q.OrderBy).GroupBy(q.GroupBy).Limit(q.PageSize, (q.Page-1)*(q.PageSize)).Find(q.Resouces)
	if err != nil {
		err = errors.New(err.Error())
	}
	return
}

func filter(param string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		return false
	}
	matchString := re.MatchString(param)
	if matchString == true {
		return false
	}
	return true
}

func (update *Update) Update() (res int64, err error) {
	res, err = GetTableSession(update.Table).Where(update.Where).Update(update.Resouces)
	return
}

func GetTableSession(table string) *xorm.Session {
	if transactionSwitch == true {
		return session.Table(AddTablesPrefix(table))
	}
	return engine.Table(AddTablesPrefix(table))
}

func (delete *Delete) Delete() (res int64, err error) {
	session := GetTableSession(delete.Table)
	if delete.Unscoped == true { //是否进行硬性删除
		session = session.Unscoped()
	}
	res, err = session.Where(delete.Where).Delete(delete.Resouces)
	return
}

//开启事务
func StartTransaction() {
	session = engine.NewSession()
	err := session.Begin()
	if err != nil {
		panic(err.Error())
	}

	transactionSwitch = true
}

//事务提交
func CommitTransaction() error {
	//提交事务
	if transactionSwitch == true {
		err := session.Commit()
		if err != nil {
			return errors.New(err.Error())
		}
	}
	defer session.Close()
	return nil
}

//事务回退
func RollbackTransaction() error {
	if transactionSwitch == true {
		err := session.Rollback()
		if err != nil {
			return errors.New(err.Error())
		}
	}
	defer session.Close()
	return nil
}
