package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

//easyjson:json
type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	reader := bufio.NewReader(r)
	result := make(DomainStat)
	var user *User
	var err error
	var line []byte
	for {
		user = &User{}

		line, _, err = reader.ReadLine()

		if errors.Is(err, io.EOF) {
			break
		}

		err = easyjson.Unmarshal(line, user)

		if err != nil {
			return nil, err
		}
		if strings.Contains(user.Email, "."+domain) {
			result[strings.ToLower(strings.Split(user.Email, "@")[1])]++
		}
	}

	return result, nil
}
