package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.{
//
//}
// var Reverse func(slice interface{})

func Reverse(obj interface{}) interface{} {

	typeOfObj := reflect.TypeOf(obj).Kind()

	if typeOfObj == reflect.Ptr {
		s := reflect.ValueOf(obj).Elem()

		middle := s.Len() / 2
		len := s.Len()

		fmt.Println("Size of slice = ", len)
		fmt.Println("Middle index = ", middle)
		fmt.Println(s.Index(0))

		newSlice := make([]interface{}, len)

		for i := 0; i < middle; i++ {
			j := len - i - 1
			x, y := s.Index(i).Interface(), s.Index(j).Interface()
			s.Index(i).Set(reflect.ValueOf(y))
			s.Index(j).Set(reflect.ValueOf(x))
		}

		for i := 0; i < len; i++ {
			fmt.Println("Index ", i, " ", s.Index(i))

		}

		fmt.Println("Reverse slice ", newSlice)
		return newSlice

	}

	return nil
}

func main() {
	var slice = []int{0, 1, 2}
	Reverse(&slice)
	// println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
