package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Generate-Json-GOStructs/internal/codegen"
)

func main() {
	var u codegen.User
	in, _ := os.Open("apidata/instagram/user.json")
	b, _ := ioutil.ReadAll(in)
	json.Unmarshal(b, &u)
	fmt.Printf("%+v\n", u)
}
