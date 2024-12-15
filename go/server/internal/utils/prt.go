package utils

import (
	"encoding/json"
	"fmt"
)

func Prt(in any) {
	j, _ := json.MarshalIndent(in, "\t", "")
	fmt.Println(string(j))
}
