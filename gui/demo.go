package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("demo.glade") //读取glade文件
	window := gtk.WindowFromObject(builder.GetObject("window"))
	b1 := gtk.ButtonFromObject(builder.GetObject("button1")) //获取按钮1
	b2 := gtk.ButtonFromObject(builder.GetObject("button2")) //获取按钮2

	//信号处理
	b1.Connect("clicked", func() {
		//获取按钮内容
		fmt.Println("button txt = ", b1.GetLabel())
	})

	b2.Connect("clicked", func() {
		//获取按钮内容
		fmt.Println("button txt = ", b2.GetLabel())
		gtk.MainQuit() //关闭窗口
	})

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)

	window.Show() //显示窗口

	gtk.Main() //主事件循环，等待用户操作
}
