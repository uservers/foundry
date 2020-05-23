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
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	// domainPattern = `^[a-z0-9]([-\.a-z0-9]*|[a-z]*)\.[a-z]{2,15}$`
	// domainPattern = `^(?:[a-z0-9]+(?:[a-z0-9-]{0,61}[a-z0-9])?)?[a-z0-9]\.[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
	domainPattern = `^(?:(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)?[a-z0-9]\.)+[a-z]{2,15}$`
)

func init() {
	message.SetString(language.Spanish, "Domain string is empty", "La cadena del dominio está vacía")
}

// IsValidDomain Checks if a string is a vlaid domain name
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
