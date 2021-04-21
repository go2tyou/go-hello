package cal

import (
	"fmt"
	"testing" // go 测试框架
)

func TestAddUpper(t *testing.T)  {
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("err respect=%v res=%v\n", 55, res)
		t.Fatalf("err respect=%v res=%v\n", 55, res)
	} else {
		// fmt.Printf("right respect=%v res=%v\n", 55, res)
		t.Logf("right respect=%v res=%v\n", 55, res)
	}
}

// =======================================================================
// === RUN   TestAddUpper
//     cal_test.go:15: right respect=55 res=55
// --- PASS: TestAddUpper (0.00s)
// PASS
// ok      gocode/helloworld/unit15/testing2/testcase      0.155s
// =======================================================================

// func TeOkAddUpper(t *testing.T)  {
// 	res := addUpper(10)
// 	if res != 55 {
// 		// fmt.Printf("err respect=%v res=%v\n", 55, res)
// 		t.Fatalf("err respect=%v res=%v\n", 55, res)
// 		} else {
// 			// fmt.Printf("right respect=%v res=%v\n", 55, res)
// 			t.Logf("right respect=%v res=%v\n", 55, res)
// 		}
// }
	
// =======================================================================
// testing: warning: no tests to run
// PASS
// ok      gocode/helloworld/unit15/testing2/testcase      0.154s
// =======================================================================

func TestHello(t *testing.T)  {
	fmt.Println("test hello")
}

// =======================================================================
// === RUN   TestHello
// test hello
// --- PASS: TestHello (0.00s)
// PASS
// ok      gocode/helloworld/unit15/testing2/testcase      0.153s
// =======================================================================

func TestGetSub(t *testing.T)  {
	res := getSub(3, 1)
	if res != 2 {
		t.Fatalf("err respect=%v res=%v\n", 2, res)
	} else {
		t.Logf("right respect=%v res=%v\n", 2, res)
	}
}
// =======================================================================
// === RUN   TestGetSub
//     cal_test.go:61: right respect=2 res=2
// --- PASS: TestGetSub (0.00s)
// PASS
// ok      gocode/helloworld/unit15/testing2/testcase      0.157s
// =======================================================================