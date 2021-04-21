package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string
}

func (m Monster) Println() {
	fmt.Println(m)
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, sex string) {
	m.Name = name
	m.Age = age
	m.Sex = sex
}

func TestStruct(a interface{})  {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) // struct has 3 fields
	for i := 0; i < num; i++ {
		// fmt.Printf("Field %d: v=%v\n", i, val.Field(i)) // Field 0: v=ssr  Field 1: v=17   Field 2: v=male

		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			// fmt.Printf("Filed %d: tag=%v\n", i, tagVal) // Filed 0: tag=name       Filed 1: tag=age
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod) // struct has 3 methods
	val.Method(1).Call(nil) // 获取第二个方法 // {ssr 17 male} ? -> 函数的方法排序 -> 按照函数名排序(ASCII码)
	
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int()) // res= 50
}

func main()  {
	var a Monster = Monster {
		Name : "ssr",
		Age : 17,
		Sex : "male",
	}
	TestStruct(a)
}