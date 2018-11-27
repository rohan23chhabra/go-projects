package main

import (
	"github.com/ogier/pflag"
	"fmt"
	"os"
	"strings"
	"github.com/rohan23chhabra/gitgo/user"
	"encoding/json"
	"bytes"
)

var (
	userstring string
)

func main() {
	pflag.Parse()
	if pflag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		pflag.PrintDefaults()
		os.Exit(1)
	}

	users := strings.Split(userstring, ",")
	fmt.Printf("Searching users(s): %s\n", users)
	for _, username := range users {
		b, err := json.Marshal(user.GetUsers(username))
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
		fmt.Println(prettyPrintJSON(string(b)))
	}
}

func init() {
	pflag.StringVarP(&userstring, "user", "u", "", "Search Users")
}

func prettyPrintJSON(s string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(s), "", "\t")
	if err != nil {
		return s
	}
	return out.String()
}
