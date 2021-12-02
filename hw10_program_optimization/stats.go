package hw10programoptimization

import (
	"bufio"
	"fmt"
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
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) (users, error) {
	reader := bufio.NewReader(r)
	result := make(users, 0)
	var user *User
	var err error
	var line []byte
	for {
		user = &User{}

		line, _, err = reader.ReadLine()

		if err == io.EOF {
			break
		}

		err = easyjson.Unmarshal(line, user)

		if err != nil {
			return nil, err
		}
		result = append(result, *user)
	}

	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	for _, user := range u {
		if strings.Contains(user.Email, "."+domain) {
			result[strings.ToLower(strings.Split(user.Email, "@")[1])]++
		}
	}

	return result, nil
}
