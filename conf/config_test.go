package conf

import (
	"fmt"
	"testing"
)

func TestOnParseConfig(t *testing.T) {
	err := parse()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
		return
	}
	fmt.Println(Conf)
}
