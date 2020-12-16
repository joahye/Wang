
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
