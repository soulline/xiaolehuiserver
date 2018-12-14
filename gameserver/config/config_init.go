package config

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	logSys "log"
	"os"
	"path"
	"time"
	"xiaolehuigo/accountserver/util"
	"xiaolehuigo/gameserver/database"
	"xiaolehuigo/gameserver/util/log"
)

const (
	maxAge       = 7 * 24 * time.Hour //保存一周
	rotationTime = 1 * 24 * time.Hour //一天切割一次文件
)

const (
	CONFIG_SERVER_KEY = "CONFIG_SERVER"
	//CONGIG_SERVER_ENV     = "DOCKER_ENV"
)

var (
	Log              *logrus.Logger
	AppName          string
	logFilePath      string
	infoLogFileName  string
	errorLogFileName string
)

func Init() {

	viper.AddConfigPath("gameserver/jsonconfig")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	AppName = viper.GetString("root")
	logFilePath = viper.GetString("local.log.filePath")
	infoLogFileName = viper.GetString("local.log.info.fileNameBase")
	errorLogFileName = viper.GetString("local.log.error.fileNameBase")
	fmt.Println("AppName : " + AppName)
	fmt.Println("logFilePath : " + logFilePath)
	fmt.Println("infoLogFileName : " + infoLogFileName)
	fmt.Println("errorLogFileName : " + errorLogFileName)
	if AppName == "" || infoLogFileName == "" || errorLogFileName == "" {
		log.Panic("config.json is not correct")
	}
	logSys.Printf("logFilePath:%s", logFilePath)

	initLog()
	host := viper.GetString("local.mysql.host")
	port := viper.GetString("local.mysql.port")
	user := viper.GetString("local.mysql.user")
	password := viper.GetString("local.mysql.password")
	dbname := viper.GetString("local.mysql.dbname")
	initDatabaseConfig(host, port, user, password, dbname)
}

func initDatabaseConfig(host string, port string, user string, password string, dbname string) {
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Panic("mysql config is empty")
	}
	dizhudatabase.ConfigDB(host, port, user, password, dbname)
}

func initLog() {
	log.SetFormatter(&logrus.JSONFormatter{})
	//如果日志路径不存在则创建
	if ok, _ := util.Exists(logFilePath); !ok {
		err := os.MkdirAll(logFilePath, 777)
		util.CheckErr(err)
	}
	InfoPath := path.Join(logFilePath, infoLogFileName)
	ErrorPath := path.Join(logFilePath, errorLogFileName)
	InfoWriter, infoerr := rotatelogs.New(
		InfoPath,
		rotatelogs.WithLinkName(InfoPath),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	util.CheckErr(infoerr)
	ErrorWriter, errorerr := rotatelogs.New(
		ErrorPath,
		rotatelogs.WithLinkName(ErrorPath),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	util.CheckErr(errorerr)
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: os.Stdout, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  InfoWriter,
		logrus.WarnLevel:  InfoWriter,
		logrus.ErrorLevel: ErrorWriter,
		logrus.FatalLevel: ErrorWriter,
		logrus.PanicLevel: ErrorWriter,
	}, &logrus.JSONFormatter{})
	log.AddHook(lfHook)
}
