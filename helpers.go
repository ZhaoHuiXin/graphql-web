package main

import (
	"io/ioutil"
)


func getSchema(path string)(string, error){
	b, err := ioutil.ReadFile(path)
	if err!= nil {
		return "",err
	}
	return string(b), nil
}

func strP(s string) *string {
	return &s
}

func boolP(b bool) *bool {
	return &b
}

func int32P(i uint) *int32 {
	r := int32(i)
	return &r
}
