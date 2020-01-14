package comm

import (
	"encoding/json"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"time"
)
var Log *logrus.Logger

type T_SVR struct {
	ip string				`json: "ip"`
	port string			`json: "port"`
}

type T_SVRS struct{
	svrs map[string]T_SVR
}

func ReadSvrAddr(svrName string) (string, string, error){
	file, _ := os.Open("../etc/svrs.json")
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil{
		return "","",err
	}
	conf := make(map[string]interface{})
	json.Unmarshal([]byte(data), &conf)

	ip := conf[svrName].(map[string]interface{})["ip"].(string)
	port := conf[svrName].(map[string]interface{})["port"].(string)
	return ip,port,nil
}

// config logrus log to local filesystem, with file rotation
func ConfigLocalFilesystemLogger(logFileName string){
	Log = logrus.New()
	logPath := "../log/"
	maxAge := time.Hour*24*365
	rotationTime := time.Hour*24
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d.log",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件

		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件

		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))

	}
	Log.SetLevel(logrus.DebugLevel)
	/*
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	Log.AddHook(lfHook)
	*/
	Log.SetOutput(writer)
	Log.Info("server logger started.")
}