package cmd

import (
	"github.com/astaxie/beego/logs"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "blog",
	Short: "A little helper for managing Blog",
	Long: `BLOG v0.5
	This is a CLI program for better managing Blog
	Developed with cobra`,
}

var Mylog *logs.BeeLogger //declare first here;先在外面声明，以保证外界可以访问

func init() {
	RootCmd.AddCommand(initCmd)

	//example showing how to use beego/logs
	//init process
	//later you may use Mylog.Debug("information") and so on to record
	//Mylog.Flush() may be calld to flush buf into output manually
	//Mylog.Close() should be called at last to compete the whole process
	Mylog := logs.NewLogger(10000)
	jsonConfig := `{
        "filename" : "test.log", 
        "maxlines" : 1000,       
        "maxsize"  : 10240       
	}`
	Mylog.SetLogger("file", jsonConfig)
	Mylog.SetLevel(logs.LevelDebug) // set the level above which will be recorded
	Mylog.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

}
