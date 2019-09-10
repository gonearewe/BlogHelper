package global

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	// mylog        *logs.BeeLogger //declare first here;先在外面声明，以保证外界可以访问
	confFileName = "default.conf"
)

type appConfig struct {
	ImgRepoPath    string //where you store your img
	BlogImgPath    string //where you put your img for blog background
	Author         string
	TemplateFile   string
	FilePath       string
	PictureDstPath string
}

var AppConfig *appConfig

func init() {
	mylog := NewLogger()
	conf, err := config.NewConfig("ini", confFileName)
	if err != nil {
		mylog.Error("cannot read config file", err)
	}

	AppConfig = new(appConfig)
	AppConfig.ImgRepoPath = conf.String("ImgRepoPath")
	AppConfig.BlogImgPath = conf.String("BlogImgPath")
	AppConfig.Author = conf.String("Author")
	AppConfig.TemplateFile = conf.String("TemplateFile")
	AppConfig.FilePath = conf.String("FilePath")
	AppConfig.PictureDstPath = conf.String("PictureDstPath")
}
func NewLogger() *logs.BeeLogger {

	//example showing how to use beego/logs
	//init process
	//later you may use mylog.Debug("information") and so on to record
	//mylog.Flush() may be calld to flush buf into output manually
	//mylog.Close() should be called at last to compete the whole process
	mylog := logs.NewLogger(10000)
	jsonConfig := `{
        "filename" : "test.log", 
        "maxlines" : 1000,       
        "maxsize"  : 10240       
	}`
	mylog.SetLogger("file", jsonConfig)
	mylog.SetLevel(logs.LevelDebug) // set the level above which will be recorded
	mylog.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	return mylog

}
