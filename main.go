package main

import (
	"fmt"
	cli "github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"strings"
)

func main() {

	low()
}
func listIsIn(strs []string, text string) bool {

	for _, value := range strs {
		if value == text {
			return true
		}
	}
	return false
}
func saveToFile(str string, path string) bool {

	if strings.HasSuffix(path, ".txt") != true {
		path = path + ".txt"
	}
	err := ioutil.WriteFile(path, []byte(str), 0666)
	if err != nil {
		fmt.Println("保存失败!请注意文件名 ", path)
		return false
	}
	return true
}
func low() {
	//监听所有事件
	EvChan := hook.Start()
	defer hook.End()
	robotgo.SetKeyDelay(200)
	robotgo.SetMouseDelay(200)
	var oldText []string //存历史数据防止重复
	ok := false
	fmt.Println("本工具为一键复制浏览器内某个元素的链接地址,就是自动右键菜单然后按e键,然后保存到文档中")
	fmt.Println("---    操作说明     ---")
	fmt.Println("--- 按  f9  键复制  ---")
	fmt.Println("--- 按  f8  键保存  ---")
	fmt.Println("--- 按  f7  键退出  ---")

	for ev := range EvChan {
		fmt.Println(ev)
		//打印事件信息
		if ev.Kind == 4 {
			if ev.Keycode == 67 {
				if ok == false {
					ok = true
					robotgo.MouseClick("right")
					robotgo.KeyTap("e")
					//获取剪辑版文件
					text, err := cli.ReadAll()
					if err != nil {
						continue
					}
					if listIsIn(oldText, text) == false && text != "" && strings.HasPrefix(text, "http") == true {
						oldText = append(oldText, text)

						txt := ""
						for _, value := range oldText {
							txt += value + "\r\n"
						}
						path := "old_uls.txt"
						if saveToFile(txt, path) == true {
							fmt.Println("数量:", len(oldText), "保存成功! ", path)
						}

					}
					ok = false
				}
			} else if ev.Keycode == 66 {
				if len(oldText) > 0 {
					ok = true
					txt := ""
					for _, value := range oldText {
						txt += value + "\r\n"
					}
					path := ""
					fmt.Println("将所有复制的链接保存到哪里? 输入文件名(回车确认): ")
					_, cErr := fmt.Scanf("%s", &path)
					if cErr != nil {
						fmt.Println("输入发生错误!")
					}
					if saveToFile(txt, path) == true {
						oldText = []string{}
						fmt.Println("保存成功!并且已经清空记录 ", path)
					}
					ok = false
				}
			} else if ev.Keycode == 65 {
				return
			}

		}

	}
}
