package conf

import (
	"fmt"
	"testing"
)

func testOnParseConfig(t *testing.T) {
	err := parse()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Conf)
}
