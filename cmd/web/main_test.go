package main
import "testing"

func Test_run(t *testing.T){
	err := run()
	if(err !=nil){
		t.Error("failed run()")
	}
}