package cgo
import "C"
/*
#include <stdio.h>
static void SayHello(const char* s) {
	puts(s);
}
*/
import "C"
func CgoTest01(){
	C.SayHello(C.CString("Hello, World\n"))
}

