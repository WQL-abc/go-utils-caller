package caller

import "testing"

func TestCaller(t *testing.T) {
	t.Run("GetCallFuncName", func(t *testing.T) {
		assertEquals(t, "get the function name of the caller", GetCallFuncName(), "func1")
	})
	t.Run("GetCallFuncRoute", func(t *testing.T) {
		skip := 1
		route := GetCallFuncRoute(skip)
		assertEquals(t, "get the function name of the caller", route[len(route)-2], "TestCaller")
	})
}

func assertEquals(t *testing.T, title string, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("%s: unexpected, actual: `%v`, expected: `%v`", title, actual, expected)
	}
}
