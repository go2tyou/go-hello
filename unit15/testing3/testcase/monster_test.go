package monster

import (
	"testing"
)

// func TestStore(t *testing.T)  {
// 	monster := &Monster {
// 		Name : "ssr",
// 		Age : 1000,
// 		Skill : "huhuhu~",
// 	}
// 	res := monster.Store()
// 	if !res {
// 		t.Fatalf("monster.Store() test err espect=%v res=%v\n", true, res)
// 	} else {
// 		t.Logf("monster.Store() test success")
// 	}
// }

// =======================================================================
// === RUN   TestStore
//     monster_test.go:17: monster.Store() test success
// --- PASS: TestStore (0.00s)
// PASS
// ok      gocode/helloworld/unit15/testing3/testcase      0.161s
// =======================================================================

func TestReStore(t *testing.T)  {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() test err espect=%v res=%v\n", true, res)
	}
	if monster.Name != "ssr" {
		t.Fatalf("monster.ReStore() test err espect=%v res=%v\n", "ssr", monster.Name)
	}
	t.Logf("monster.ReStore() test success")
}

// =======================================================================
// === RUN   TestReStore
//     monster_test.go:36: monster.ReStore() test err espect=ssr res=xxx
// --- FAIL: TestReStore (0.00s)
// FAIL
// exit status 1
// FAIL    gocode/helloworld/unit15/testing3/testcase      0.157s
// =======================================================================