/*
‘…’ 其实是go的一种语法糖。
它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
第二个用法是slice可以被打散进行传递。

形参的参数前的三个点，表示可以传0到多个参数
变量后三个点表示将一个切片或数组变成一个一个的元素，即打散．
*/

//生成md5
md5 := md5.New()
io.WriteString(md5, "学生注册实验自动生成账户")
MD5Str := hex.EncodeToString(md5.Sum(nil))

fmt.Println(MD5Str)


arg := "https://share.todoen.com/mtopic/guide3.html?utm_source=baidu&utm_medium=sem&utm_device=mobile&utm_campaign=318&utm_group=11&utm_keyword=58938&renqun_youhua=1851573&bd_vid=11348576986599726827"

params, _ := url.Parse(arg)

m, _ := url.ParseQuery(params.RawQuery)

fmt.Println(m["utm_keyword"][0])
fmt.Println(m["utm_campaign"][0])
fmt.Println(m["utm_group"][0])
fmt.Println(m.Get("utm_group"))
fmt.Println(m.Get("utm_keyword"))

hans := "中国人"
a := pinyin.NewArgs()
fmt.Println(pinyin.Pinyin(hans, a))
fmt.Println(pinyin.Convert("世界", nil))
//Convert 跟 Pinyin 的唯一区别就是 a 参数可以是 nil

//range方法的index和value是值复制
for index, dashArr := range dashUserAccMap {
   //遍历删除slice用for循环
   for i := 0; i < len(dashArr); {
      for acckey := range dashArr[i].AccountIds {
         if accIdsMap[acckey] {
            delete(dashArr[i].AccountIds, acckey)
         }
      }

      if len(dashArr[i].AccountIds) == 0 {
         dashArr = append(dashArr[:i], dashArr[i+1:]...)
      } else {
         i++
      }
   }
   
   dashUserAccMap[index] = dashArr
   if len(dashArr) == 0 {
      delete(dashUserAccMap, index)
   }
}


// make出来的是切片    指针类型
// []int这种new出来的是数组 非指针类型


// copy(a,b)
// a此时必须有值！！！！！否则复制出来的a依旧为空
// 如果b的长度大于a  a的长度不会发生改变 且 a改变的是与b位置相对应的位置的值
// 总结：copy不对a扩容 只会用a原本的容量

// append出来的是一个新的切片
eg: a := make([]int, 1)

f := func(t []int) {
   t = append(t, 1)
   //t[0] = 1
}

f(a)
fmt.Println(a)

slice := append([]int{1,2,3},4,5,6)
fmt.Println(slice) //[1 2 3 4 5 6]
//第二个参数也可以直接写另一个切片，将它里面所有元素拷贝追加到第一个切片后面。要注意的是，这种用法函数的参数只能接收两个slice，并且末尾要加三个点

//这种切片的复制 用的是同一个数组底
slice1 := make([]int, 5, 5)
slice2 := slice1
slice2[2] = 1
fmt.Println(slice1) //[0 1 0 0 0]
fmt.Println(slice2) //[0 1 0 0 0]


//正则表达式
f, _ := regexp.MatchString("p([a-z]+)ch", "peach")
fmt.Println(f)

r, _ := regexp.Compile("p([a-z]+)ch")
fmt.Println(r.MatchString("peach"))

test := "Hello 世界！123 Go."

reg := regexp.MustCompile(`[a-z]+`)
fmt.Println(reg.FindAllString(test, -1))

//eval包用法 计算string的计算表达式
str := "22.056*2+3.28*(6.7-1)/2"

r, err := evaler.Eval(str)
if err != nil {
   fmt.Println(err.Error())
} else {
   value := evaler.BigratToFloat(r)
   fmt.Println(value)
}

fmt.Println(22.056*2 + 3.28*(6.7-1)/2)

//flag的用法
//使用方式 --filepath ""
var file = flag.String("filepath", "", "目标文件名: ../test.csv")
//不给值 默认是8
var a = flag.Int("b", 8, "")
//给任意值 都是true 不给值就是false
var b = flag.Bool("c", false, "")

flag.Parse()

//如果空 或者 0就报错
if *file == "" || *a == 0 {
   flag.PrintDefaults()
   os.Exit(1)
}
