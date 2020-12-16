package test

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"common-libs/github.com/shopspring/decimal"
)

func TestExample(t *testing.T) {

	//s := []int{5, 2, 6, 3, 1, 4}
	//s1 := []int{5, 2, 6, 3, 1, 4}
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//sort.Ints(s1)
	//fmt.Println(s)
	//fmt.Println(s1)

	getType := reflect.TypeOf([]int{1})
	getValue := reflect.ValueOf([]int{1})

	fmt.Println("类型", getType.Name())
	fmt.Println("类型", getValue.Type().Name())
	fmt.Println("类型", getValue.Type())
}

func TestHello(t *testing.T) {

	user := User{1, "Allen.Wu", 25}

	//也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种
	//TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
	//ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
	getType := reflect.TypeOf(user)
	getValue := reflect.ValueOf(user)

	fmt.Println("类型", getType.Name())         //针对于结构体 切片或指针是没有类型名称的，只能返回空字符串
	fmt.Println("类型", getValue.Type().Name()) //针对于结构体 切片或指针是没有类型名称的，只能返回空字符串
	fmt.Println("类型", getValue.Type())        //引用类型
	fmt.Println("值", getValue)                //这里使用getValue.Elem() 报错

	//以下方法需要是非指针类型
	//field  n. 领域；牧场；旷野；战场；运动场；字段
	//这里取出来的 field 对象是 reflect.StructField 类型，但是它没有办法用来取得对应对象上的值
	filed, isexist := getType.FieldByName("Namee")
	if !isexist {
		fmt.Println("找不到该字段")
	} else {
		fmt.Println("字段名: ", filed.Name, "字段类型: ", filed.Type, "字段index: ", filed.Index, "字段大小: ", filed.Offset)
	}

	//这里取出来的 val 类型是 reflect.Value，它是一个具体的值，而不是一个可复用的反射对象了
	val := getValue.FieldByName("Name")
	fmt.Println("值: ", val, "值对应类型: ", val.Type())

	fmt.Println("--------我是分隔符----------")
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("%s: %v = %+v \n", field.Name, field.Type, value.Interface())
	}

	// 接口强制类型转换 realValue := value.Interface().(已知的类型)

	fmt.Println("--------我是分隔符1----------")
	for i := 0; i < getValue.NumField(); i++ {
		v := getValue.Field(i)
		f := getType.Field(i)
		fmt.Printf("%s: %v = %v \n", f.Name, f.Type, v.Interface())
	}

	fmt.Println("--------我是分隔符2----------")
	//注意，参数必须是指针才能修改其值
	//如果传入的参数不是指针，而是变量，那么 通过Elem获取原始值对应的对象则直接panic 通过CanSet方法查询是否可以根据元素重新赋值
	//指针、map、slice、channel、Array	Elem()类似于对指针做*操作
	getPointVale := reflect.ValueOf(&user)
	getElem := getPointVale.Elem()
	fmt.Println("使用Elem() :", getPointVale.Elem())
	fmt.Println("不使用Elem():", getPointVale)
	fmt.Println("类型", getPointVale.Type())
	fmt.Println("settability of pointer:", getElem.CanSet())
	fmt.Println("settability of pointer:", getValue.CanSet())
	fmt.Println("settability of pointer:", getPointVale.CanSet())
	// 重新赋值
	getElem.FieldByName("Id").SetInt(2)
	getElem.FieldByName("Name").SetString("Jack.Wu")
	getElem.FieldByName("Age").SetInt(22)
	fmt.Println("重新赋值后的value :", getPointVale.Elem())

	fmt.Println("--------我是分隔符3----------")
	// 如何通过反射来进行方法的调用？
	// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call
	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	valueObj := reflect.ValueOf(user)
	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := valueObj.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("Nancy"), reflect.ValueOf(30)}
	methodValue.Call(args) //不能使用methodValue.CallSlice(args)
	// 3. 再看看无参数的调用方法
	methodValue = getPointVale.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)

	//kind用法 针对基础类型
	fmt.Println("--------我是分隔符4----------")
	test := 0
	testValueOf := reflect.ValueOf(test)
	testTypeOf := reflect.TypeOf(test)
	fmt.Println(testValueOf.IsValid())
	fmt.Println(testValueOf.Kind().String())
	fmt.Println(testTypeOf.Kind().String())
	fmt.Println(getValue.Kind().String())

	//len()用法 结合数组使用
	fmt.Println("--------我是分隔符5----------")
	s := reflect.ValueOf("hello world!")
	fmt.Println(s.Len())
	for i := 0; i < s.Len(); i++ {
		fmt.Printf("%v \t", string(s.Index(i).Interface().(uint8)))
	}

}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u *User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs", u)
}

func TestTicker(t *testing.T) {

	//t1 := time.NewTimer(time.Second * 2)

	ticker := time.NewTicker(2 * time.Second)

	ch := make(chan bool)
	go func(t *time.Ticker) {
		defer t.Stop()
		for {
			select {
			//case <-t1.C:
			case <-ticker.C:

				fmt.Println("timer running....")
				// 需要重置Reset 使 t 重新开始计时
				//t.Reset(time.Second * 2)
			case stop := <-ch:
				if stop {
					fmt.Println("timer Stop")
					return
				}
			}
		}
	}(ticker)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)

}

func TestAfter(t *testing.T) {
	t1 := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t1)
	//阻塞3秒
	fmt.Println("t=", <-t1)
}

func TestTickerFunc(t *testing.T) {

	SetTicker(time.Second*3, 6)
}

func SetTicker(dur time.Duration, sum int) {

	sum -= 1

	ch := make(chan bool, 10)

	ff := func() bool {

		fmt.Println(time.Now())

		//func() 有return

		if time.Now().Minute() == 42 {
			return true
		}

		return false
	}

	ch <- ff()

	for i := 0; i < sum; i++ {
		if <-ch {
			break
		}
		<-time.After(dur)
		ch <- ff()
	}

}

type TimePeriodSlice [][2]string

func (p TimePeriodSlice) Len() int           { return len(p) }
func (p TimePeriodSlice) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p TimePeriodSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestIsTimePeriodCross(t *testing.T) {

	a := TimePeriodSlice{
		{"2019-08-07", "2020-02-09"},
		{"2020-11-07", "2021-10-21"},
		{"2020-01-02", "2020-01-16"},
		{"2020-05-06", "2020-09-20"},
		{"2020-03-11", "2020-04-18"},
		{"2020-10-04", "2021-04-18"},
		{"2018-02-04", "2018-12-28"},
		{"2020-02-04", "2020-03-11"},
		{"2020-08-30", "2020-10-01"},
		{"2020-04-21", "2020-12-30"},
	}

	sort.Sort(a)

	fmt.Println(a)

	m := make(map[int][2]string, 0) //待检索 最终剩余的非重叠
	//res := make(map[int][2]string, 0) //重叠
	res := make([][2]string, 0)

	for index, i := range a {
		m[index] = i
	}

	for k := range a {

		if _, ok := m[k]; !ok {
			continue
		}

		if func() bool {
			f := false
			for _, i := range res {
				//st, _ := time.ParseInLocation("2006-01-02", i[0], time.Local)
				//et, _ := time.ParseInLocation("2006-01-02", i[1], time.Local)
				//t0, _ := time.ParseInLocation("2006-01-02", m[k][0], time.Local)
				//t1, _ := time.ParseInLocation("2006-01-02", m[k][1], time.Local)
				if m[k][0] <= i[1] && m[k][1] >= i[0] {
					f = true
					break
				}
			}
			return f
		}() {

			//res[k] = [2]string{m[k][0], m[k][1]}
			res = append(res, [2]string{m[k][0], m[k][1]})
			delete(m, k)

		} else if _, ok := m[k+1]; ok {

			//st, _ := time.ParseInLocation("2006-01-02", m[k+1][0], time.Local)
			//et, _ := time.ParseInLocation("2006-01-02", m[k+1][1], time.Local)
			//t0, _ := time.ParseInLocation("2006-01-02", m[k][0], time.Local)
			//t1, _ := time.ParseInLocation("2006-01-02", m[k][1], time.Local)
			if m[k][0] <= m[k+1][1] && m[k][1] >= m[k+1][0] {
				//res[k] = [2]string{m[k][0], m[k][1]}
				//res[k+1] = [2]string{m[k+1][0], m[k+1][1]}
				res = append(res, [2]string{m[k][0], m[k][1]})
				res = append(res, [2]string{m[k+1][0], m[k+1][1]})
				delete(m, k)
				delete(m, k+1)
			}

		}
	}

	fmt.Println(m)
	fmt.Println(res)

}

func TestEx(t *testing.T) {

	df1 := decimal.NewFromFloat(float64(0.0232222) * float64(1.234))

	df := decimal.NewFromFloat(float64(0.0232222))

	fmt.Println(df.Mul(decimal.NewFromFloat(1.234)).Div(decimal.NewFromInt(100)))

	fmt.Println(df1.Div(decimal.NewFromInt(100)))

	fmt.Println((decimal.NewFromFloat(float64(2324) / 100)))

	fmt.Println(decimal.NewFromFloat(2324).Div(decimal.NewFromInt(100)))

	/**********/
	fmt.Println(len(strings.Split(" ", ",")))

	fmt.Println(len(strings.Split("", ",")))

}
