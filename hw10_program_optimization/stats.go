package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

const maxCapacity = 512 * 1024

var ErrInvalidEmailFormat = errors.New("invalid user email format")

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	scanner := bufio.NewScanner(r)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var i int
	var user User
	for scanner.Scan() {
		if err = json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return result, err
		}
		result[i] = user
		i++
	}
	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	domainPatern := "." + domain
	for _, user := range u {
		if strings.HasSuffix(user.Email, domainPatern) {
			if at := strings.IndexByte(user.Email, '@'); at > 0 {
				result[strings.ToLower(user.Email[at+1:])]++
			} else {
				return nil, fmt.Errorf("%w: %v does not contains @", ErrInvalidEmailFormat, user.Email)
			}
		}
	}
	return result, nil
}
