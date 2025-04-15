package main

//#include <stdio.h>
//static void run_on_load(void) __attribute__((constructor));
//static void run_on_load(void) {
//    // Call the Go function
//    extern void WfrpCoreInit();
//    WfrpCoreInit();
//}

import "C"
import "github.com/Pretty-IT/wfrp-core/api"

//export WfrpCoreInit
func WfrpCoreInit() string {
	return api.HelloWorld()
}

func main() {}
