package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonearewe/BlogHelper/global"
	"github.com/gonearewe/BlogHelper/models"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing a article for your blog",
	Run:   initFunc,
}

func initFunc(c *cobra.Command, args []string) {
	mylog := global.NewLogger()

	ch := make(chan string, 1)
	go pickPicture(ch)

	var (
		title, subtitle string
		tags            []string
	)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入文章标题 :")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			title = input
			mylog.Debug("get article title :", input)
			break
		}
	}
	for {
		fmt.Print("请输入文章副标题 :")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			subtitle = input
			mylog.Debug("get article subtitle :", input)
			break
		}
	}
	for {
		fmt.Print("请输入文章标签 :")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			tags = strings.Fields(input)
			mylog.Debug("get article subtitle :", input)
			break
		}
	}

	article := models.NewArticle(
		title,
		subtitle,
		global.AppConfig.Author,
		<-ch,
		tags,
	)
	fmt.Println(article)

	var templateString string
	if tmp, err := ioutil.ReadFile(global.AppConfig.TemplateFile); err != nil {
		mylog.Error("fail to read template file")
		fmt.Errorf("\nfail to read template file\n")
		os.Exit(1)
	} else {
		templateString = string(tmp)
		templateString = article.FillTemplate(templateString)

		fileAddr := article.GetFileAddr(global.AppConfig.FilePath)
		f, err := os.Create(fileAddr)
		if err != nil {
			mylog.Error("fail to create article file")
			fmt.Errorf("\nfail to create article file\n")
			os.Exit(1)
		}
		defer f.Close()

		f.WriteString(templateString)
	}

	mylog.Flush()
	mylog.Close()
}

func pickPicture(ch chan string) {
	mylog := global.NewLogger()

	var dir string
	dir = global.AppConfig.ImgRepoPath
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil || len(fileinfo) == 0 {
		mylog.Error("fail to read picture directory")
		fmt.Errorf("\nfail to read picture directory\n")
		os.Exit(1)
	}

	srcAddr := filepath.Join(dir, fileinfo[0].Name())
	dstAddr := filepath.Join(global.AppConfig.PictureDstPath, fileinfo[0].Name())
	if err := os.Rename(srcAddr, dstAddr); err != nil {
		mylog.Error("fail to pick a picture")
		fmt.Errorf("\nfail to pick a picture\n")
		os.Exit(1)
	}

	// srcfile, err := os.Open(addr)
	// if err != nil {
	// 	mylog.Error("fail to open a picture")
	// 	fmt.Errorf("\nfail to open a picture\n")
	// 	os.Exit(1)
	// }
	// defer srcfile.Close()

	// dir = global.AppConfig.BlogImgPath
	// dstfile, err := os.Create(addr)
	// if err != nil {
	// 	mylog.Error("fail to copy a picture")
	// 	fmt.Errorf("\nfail to copy a picture\n")
	// 	os.Exit(1)
	// }
	// defer dstfile.Close()

	ch <- fileinfo[0].Name()

	// _, err = io.Copy(dstfile, srcfile)
	// if err != nil {
	// 	mylog.Error("fail to copy a picture")
	// 	fmt.Errorf("\nfail to copy a picture\n")
	// 	os.Exit(1)
	// }

}
