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

// IsValidDomain Checks if a string is a vlaid domain
func IsValidDomain(domain string) error {
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

	logrus.Debugf("%s is a valid domain name", domain)
	return nil
}
