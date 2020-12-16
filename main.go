package main

import (
	"encoding/json"
	"fmt"

	"adwetec.com/tools/constant"
	"adwetec.com/tools/oidb"
)

var (
	ProdSihai = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		"app_sihai_rw",
		"5T4G6)qXCxw$",
		"192.168.0.68",
		"3306",
		"adwetec_sihai_prod")

	ProdComm = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		"prodwetec",
		"Q!T09%RJhzDXkdC2",
		"192.168.0.18",
		"3306",
		"adwetec_prod")

	DevpComm = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		"app_apiserver_rw",
		"|~_>I8;pP:)/^E",
		"192.168.0.42",
		"3306",
		"adwetec_devp")

	DevpSihai = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		"app_sihai_rw",
		"jw0X0Oi71BwDL0Ux",
		"10.10.10.93",
		"3306",
		"adwetec_sihai_devp")
)

type T struct {
	AccId    string
	Id       int64
	EntityId int64
	AdvId    int64
	AccType  int64
	User     []int64
}

func main() {

	//dbcom, err := dao.NewDaoManager(ProdComm)
	//if err != nil {
	//	fmt.Println("数据库初始化错误: " + err.Error())
	//	return
	//}
	m := oidb.DashboardMap

	content := oidb.AdwetecDashboardJsonUtilfn.Query(constant.CONSTANT_MODE_PRODUCT, 712)
	if content != "" {
		err := json.Unmarshal([]byte(content), &m)
		if err != nil {
			return
		}
	}

	fmt.Println(m)

}
