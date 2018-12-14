package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"sync"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/util/log"
)

//数据库实例
var dbClient *PDB
var dbArgs string

//避免并发情况
var once sync.Once

/***
  PostgresqlDB
*/
type PDB struct {
	*gorm.DB
}

/**
设置数据连接参数
*/
func ConfigDB(host string, port string, user string, password string, dbname string) {

	dbArgs = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password,
		host, port, dbname)
	log.Info("dbconnectconfig:" + dbArgs)
	fmt.Println("dbconnect : " + dbArgs)

	gormDB, err := gorm.Open("mysql", dbArgs)
	if err != nil {
		panic(err)
	}
	dbClient = &PDB{gormDB}
	if !dbClient.HasTable(&datamodel.UserInfo{}) {
		dbClient.CreateTable(&datamodel.UserInfo{})
		dbClient.Exec("alter table user_info AUTO_INCREMENT=10000;")
	} else {
		dbClient.AutoMigrate(&datamodel.UserInfo{})
	}
	if !dbClient.HasTable(&datamodel.Token{}) {
		dbClient.CreateTable(&datamodel.Token{})
	} else {
		dbClient.AutoMigrate(&datamodel.Token{})
	}
}

/***
使用默认参数
*/
func DB() *PDB {
	if dbArgs != "" {
		log.Info("dbconnectconfig:" + dbArgs)
	}
	return dbClient
}

/***
默认的UUID
*/
func (d *PDB) UUID() string {
	//27
	return ksuid.New().String()
}

/***
根据UUID查询指定数据
*/
func (d *PDB) Get(out interface{}, uuid string) *PDB {
	db := d.Where("id = ?", uuid).First(out)
	return &PDB{db}
}

/***
根据UUID删除指定表的数据
注意：尽量使用ID删除，避免传入空的对象，导致整个表被删除
*/
func (d *PDB) DeleteByUuid(value interface{}, uuid string) *PDB {
	db := d.Where("id = ?", uuid).Delete(value)
	return &PDB{db}
}

/***
添加
*/
func (d *PDB) Add(value interface{}) *PDB {
	db := d.Create(value)
	return &PDB{db}
}
