package dd

import (
	"encoding/json"
	"fmt"
	"os"
)

func Dump(data ...interface{}) {
	for _, d := range data {
		out, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Println(string(out))
		}
	}
}

func DumpAndDie(data ...interface{}) {
	Dump(data...)
	os.Exit(1)
}
