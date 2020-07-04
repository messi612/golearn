package main

import (
	"fmt"
	"reflect"
)

type tt interface {
	hello() string
}

type T struct {
	Name string
}

func (t T) hello() string {
	return "hello"
}

func (r T) Read(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	fmt.Println("test float reflect....")
	var x float64 = 3.41
	xt := reflect.TypeOf(x)
	fmt.Println("xt.Kind = ", xt.Kind())
	fmt.Println("xt.Name = ", xt.Name())
	//fmt.Println(p.Elem())  //Elem of invalid type

	xv := reflect.ValueOf(x)
	fmt.Println("xv.CanSet = ", xv.CanSet())
	//call of reflect.Value.Elem on float64 Value
	//fmt.Println("xv.Elem = ", xv.Elem())
	fmt.Println("xv.Kind = ", xv.Kind())
	fmt.Println("xv.Type = ", xv.Type())

	fmt.Println("test float pointer reflect....")
	pxt := reflect.TypeOf(&x)
	fmt.Println("pxt.Kind = ", pxt.Kind())
	fmt.Println("pxt.Name = ", pxt.Name())
	fmt.Println("pxt.Elem = ", pxt.Elem())
	pxv := reflect.ValueOf(&x)
	fmt.Println("pxv.CanSet = ", pxv.CanSet())
	fmt.Println("pxv.Elem = ", pxv.Elem())
	fmt.Println("pxve.CanSet = ", pxv.Elem().CanSet())
	fmt.Println("before setFloat x = ", x)
	pxv.Elem().SetFloat(7.2)
	fmt.Println("after setFloat x = ", x)
	fmt.Println("pxv.Kind = ", pxv.Kind())
	fmt.Println("pxv.Type = ", pxv.Type())

	fmt.Println("test struct reflect....")
	s := T{Name: "xiaoluo"}
	st := reflect.TypeOf(s)
	fmt.Println("st.Kind = ", st.Kind())
	fmt.Println("st.Name = ", st.Name())

	sv := reflect.ValueOf(s)
	fmt.Println("sv.CanSet = ", sv.CanSet())
	//call of reflect.Value.Elem on struct Value
	//fmt.Println("sv.Elem = ", sv.Elem())
	fmt.Println("sv.Kind = ", sv.Kind())
	fmt.Println("sv.Type = ", sv.Type())

	psv := reflect.ValueOf(&s)
	fmt.Println("psv.CanSet = ", psv.CanSet())
	fmt.Println("psv.Elem = ", psv.Elem())
	fmt.Println("psv.Kind = ", psv.Kind())
	fmt.Println("psv.Type = ", psv.Type())

	psve := psv.Elem()
	fmt.Println("psve.CanSet = ", psve.CanSet(), ",psve.Type = ", psve.Type())
	typeOfse := psve.Type()
	for i := 0; i < psve.NumField(); i++ {
		f := psve.Field(i)
		fmt.Printf("%d: %s %s = %v \n", i, typeOfse.Field(i).Name, f.Type(), f.Interface())
		fmt.Println("before set Name s.Name = ", s.Name)
		f.SetString("messi")
		fmt.Println("after set Name s.Name = ", s.Name)
	}
}
