package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	action *bool = flag.Bool("d", false, "is decode?")
)

func handleError(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func main() {

	flag.Parse()

	secret := flag.Arg(0)

	reader := bufio.NewReader(os.Stdin)

	// decode
	if *action {
		s, err := ioutil.ReadAll(reader)

		if err != nil {
			handleError(err)
		}

		token, err := jwt.Parse(strings.TrimSpace(string(s)), func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("不予許的加密算法 %s \n", t.Header["alg"])
			}

			return []byte(secret), nil
		})

		if err != nil {
			handleError(err)
		}

		result, err := json.Marshal(token.Claims)

		if err != nil {
			handleError(err)
		}

		fmt.Println(string(result))

	} else {

		build := jwt.New(jwt.SigningMethodHS256)

		err := json.NewDecoder(reader).Decode(&build.Claims)
		if err != nil {
			handleError(err)
		}

		token, err := build.SignedString([]byte(secret))

		if err != nil {
			handleError(err)
		}

		fmt.Println(token)
	}

}
