package test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/araddon/dateparse"
	"golang.org/x/text/width"
)

type Data0 struct {
	Id         int
	SubMediaId int
	ProjectId  int
	StartTime  string
	EndTime    string
}

func main() {

	ex := "好贷天下信息技术（北京）有限公司"

	// 全角转半角
	fmt.Println(width.Narrow.String(ex))
	// 半角转全角
	fmt.Println(width.Widen.String(ex))

	t0, _ := dateparse.ParseLocal("2019-06-16")

	fmt.Println(t0.AddDate(0, 0, -6).Format("2006-01-02"))

	//第三种读文件方法
	if bytes, err := ioutil.ReadFile("src/test.txt"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(bytes))
	}

	//打开文件的两种方法
	//if file, err := os.Open("src/test.txt"); err == nil {
	if file, err := os.OpenFile("src/suning.txt", os.O_RDONLY, 0644); err == nil {
		defer file.Close()

		//第一种读 文件对象方法
		if bytes0, err := ioutil.ReadAll(file); err == nil {
			fmt.Println(string(bytes0))
		} else {
			fmt.Println(err.Error())
		}
		////第二种读 文件对象方法
		//buf := make([]byte, 10)
		//if n, err := file.Read(buf); err == nil {
		//fmt.Println(string(buf))
		//fmt.Println(n)
		//} else {
		//fmt.Println(err.Error())
		//}
		//
		////第三种读 文件对象方法 带有缓存
		//buf := make([]byte, 1024)
		//reader := bufio.NewReader(file)
		//if n, err := reader.Read(buf); err == nil {
		//fmt.Println(string(buf))
		//fmt.Println(n)
		//} else {
		//fmt.Println(err.Error())
		//}

	} else {
		fmt.Printf("文件打开失败 %s", err.Error())
	}

}
