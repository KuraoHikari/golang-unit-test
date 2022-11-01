package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T){
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		// panic("Result is not Hello Eko")
		// t.Fail()
		t.Error("harusnya Hello Eko")
	}
	fmt.Println("Test HelloWorldEko Done")
}

func TestHelloWorldKurao(t *testing.T){
	result := HelloWorld("Kurao")
	if result != "Hello Kurao" {
		// unit test failed
		// t.FailNow()
		t.Fatal("harusnya Hello Kurao YGY")
	}
	fmt.Println("Test HelloWorldKurao Done")
}