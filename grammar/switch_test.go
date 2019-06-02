package grammar

import "testing"

func TestSwitch(t *testing.T) {
	var name *int = nil
	println(name)
	var str = "2"
	switch str {
	case "1":
		println("yes")
	default:
		println("default")
	}
}
