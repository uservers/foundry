/*
Copyright 2020 U Servers Comunicaciones, S.C.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stringtool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	domainPattern = `^(?:(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)?[a-z0-9]\.)+[a-z]{2,15}$`
	useridPattern = `^[a-z][a-z0-9]{5}$`
	loginPattern  = `^[-_\.a-z0-9]{1,125}$`
)

func init() {
	message.SetString(language.Spanish, "Domain string is empty", "La cadena del dominio está vacía")
}

// IsValidDomain Checks if a string is a valid domain name
func IsValidDomain(domain string) error {
	if err := checkDomainChars(domain); err != nil {
		return err
	}

	// FIXME: Split el nombre

	logrus.Debugf("%s is a valid domain name", domain)
	return nil
}

// Check is string is a domain pattern
func checkDomainChars(domain string) error {
	if domain == "" {
		return errors.New("Domain string is empty")
	}

	domain = strings.ToLower(domain)
	domainRegex := regexp.MustCompile(domainPattern)
	// domainRegex.

	if !domainRegex.MatchString(domain) {
		logrus.Debugf("%s is not a valid domain name", domain)
		return fmt.Errorf("'%s' is not a valid domain name", domain)
	}

	return nil
}

// SplitDomain regresa el dominio y la TLD ( que determina el tipo)
func SplitDomain(domainName string) (sld, tld string, err error) {
	if err := checkDomainChars(domainName); err != nil {
		return sld, tld, err
	}

	cachos := strings.Split(domainName, ".")
	if len(cachos) < 2 {
		return sld, tld, errors.New("The string is not a valid domain name")
	}

	// Chequemos primero si es un mx
	if cachos[len(cachos)-1] == "mx" {
		second := cachos[len(cachos)-2]
		if second == "com" || second == "net" || second == "org" || second == "gob" || second == "edu" {
			if len(cachos) < 3 {
				return sld, tld, errors.New("Third level .mx domain malformed")
			}
			return cachos[len(cachos)-3], "com.mx", nil
		}

		return second, "mx", nil
	}

	// Aqui un return generico para dominios con forma tld.sld
	return cachos[len(cachos)-2], cachos[len(cachos)-1], nil
}

// ParseStorable divide un dominio en login + userid
func ParseStorable(storable string) (login, userid string, err error) {
	if len(storable) <= 6 {
		return "", "", errors.New("string is not a well formed storable")
	}
	if string(storable[len(storable)-7]) != "@" {
		return "", "", errors.New("string is not a well formed storable, cant find @")
	}

	// Checa que sea un userdi valido
	userid = storable[len(storable)-6:]
	if err := IsValidUserID(userid); err != nil {
		return "", "", errors.Wrap(err, "while checking the storable userid")
	}

	// Checa que sea un login valido
	login = storable[0 : len(storable)-7]
	if err := IsValidLogin(login); err != nil {
		return "", "", errors.Wrap(err, "while checking the storable login")
	}

	return login, userid, nil
}

// IsValidUserID Checks if a string is a valid userid
func IsValidUserID(userid string) error {

	useridRegex := regexp.MustCompile(useridPattern)
	// domainRegex.

	if !useridRegex.MatchString(userid) {
		logrus.Debugf("%s is not a valid domain name", userid)
		return errors.New("not a valid userid")
	}

	return nil
}

// IsValidLogin check if satring is a valid user login
func IsValidLogin(login string) error {
	if login == "" {
		return errors.New("empty string is not an valid login")
	}

	loginRegex := regexp.MustCompile(loginPattern)
	if !loginRegex.MatchString(login) {
		logrus.Debugf("%s is not a valid login name", login)
		return errors.New("not a valid login")
	}

	if strings.Contains(login, " ") {
		logrus.Debugf("%s contains spaces", login)
		return errors.New("string is not a valid string, contains spaces")
	}

	if strings.ContainsAny(string(login[0]), "-_.") {
		return errors.New("login cannot begin with a dash underscore or period")
	}

	if strings.ContainsAny(string(login[len(login)-1]), "-_.") {
		return errors.New("login cannot end with a dash underscore or period")
	}

	if len(login) == 1 && strings.ContainsAny(login, "0123456789") {
		return errors.New("login cannot be a single digit")
	}

	loginDigits := regexp.MustCompile(`^[0-9]+$`)
	if loginDigits.MatchString(login) {
		return errors.New("login cannot be a number")
	}

	logrus.Debugf("%s is a valid login", login)
	return nil
}
