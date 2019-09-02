package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gonearewe/BlogHelper/models"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing a article for your blog",
	Run:   initFunc,
}

func initFunc(c *cobra.Command, args []string) {
	article := new(models.Article)

	reader := bufio.NewReader(os.Stdin)
	go pickPicture()
	for {
		fmt.Print("请输入文章标题 :")
		input, err := reader.ReadString('\n')
		if err == io.EOF && input != "" {
			article.Title = input
			Mylog.Debug("get article title :", input)
			break
		}
	}
	for {
		fmt.Print("请输入文章副标题 :")
		input, err := reader.ReadString('\n')
		if err == io.EOF && input != "" {
			article.Subtitle = input
			Mylog.Debug("get article subtitle :", input)
			break
		}
	}
	Mylog.Flush()
	Mylog.Close()
}

func pickPicture() {
	ioutil.ReadDir()
}
