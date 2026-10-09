package main

import "testing"

func TestAdd(t *testing.T) {
	//子测试
	//t.Run("Add1", func(t *testing.T) {
	//	res := Add(1, 2)
	//	if res != 3 {
	//		t.Errorf("测试失败")
	//		return
	//	}
	//	t.Logf("测试通过")
	//})
	//t.Run("Add2", func(t *testing.T) {
	//	res := Add(1, 1)
	//	if res != 2 {
	//		t.Errorf("测试失败")
	//		return
	//	}
	//	t.Logf("测试通过")
	//})
	//单元测试
	//res := Add(1, 2)
	//if res != 3 {
	//	t.Errorf("测试失败")
	//	return
	//}
	//t.Logf("测试通过")
	//测试用例很多
	cases := []struct {
		Name         string
		A, B, Expect int
	}{
		{"a1", 2, 3, 5},
		{"a2", 1, 2, 1},
		{"a3", 2, 0, 2},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			if Add(c.A, c.B) != c.Expect {
				t.Errorf("测试失败")
				return
			}
			t.Logf("测试通过")
		})
	}

}
