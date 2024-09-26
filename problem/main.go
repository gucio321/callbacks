package main

// extern int goCallback(int a, int b);
// typedef int (*callback)(int, int);
//
// static void test(callback cb) {
//     cb(1, 2);
// }
import "C"
import "fmt"

//export goCallback
func goCallback(a, b C.int) C.int {
	fmt.Println("Go callback called with", a, b)
	return a + b
}

func main() {
	C.test(C.callback(C.goCallback))

	// Problem:
	// How to do this
	myPrivCallback := func(a, b C.int) C.int {
		fmt.Println("This is a private function called with", a, b)
		return a + b
	}

	_ = myPrivCallback // do not cause crash

	// C.test(C.callback(myPrivCallback)) // This will not work
}
