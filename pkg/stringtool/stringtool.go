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
	"strconv"
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
	uuidPattern   = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	curpPattern   = `^([A-Z][AEIOUX][A-Z]{2}\d{2}(?:0[1-9]|1[0-2])(?:0[1-9]|[12]\d|3[01])[HM](?:AS|B[CS]|C[CLMSH]|D[FG]|G[TR]|HG|JC|M[CNS]|N[ETL]|OC|PL|Q[TR]|S[PLR]|T[CSL]|VZ|YN|ZS)[B-DF-HJ-NP-TV-Z]{3}[A-Z\d])(\d)$`
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

// IsValidUUID Indica si la string es un UUID valido
func IsValidUUID(uuid string) error {
	if len(uuid) != 36 {
		return errors.New("string is not a valid UUID")
	}
	loginRegex := regexp.MustCompile(uuidPattern)
	if !loginRegex.MatchString(uuid) {
		logrus.Debugf("%s is not a valid UUID", uuid)
		return errors.New("not a valid UUID")
	}

	return nil
}

// IsValidCURP Indica si la string es un curp valido
func IsValidCURP(curp string) error {
	// Checa el largo
	if len(curp) != 18 {
		return errors.New("string is not a valid CURP")
	}
	// Checa que empate en la regexp
	loginRegex := regexp.MustCompile(curpPattern)
	if !loginRegex.MatchString(curp) {
		logrus.Debugf("%s is not a valid CURP", curp)
		return errors.New("not a valid CURP")
	}

	// Ahora, el digito verificador:
	chars := `0123456789ABCDEFGHIJKLMNNOPQRSTUVWXYZ`
	// Otra version dice asi:
	// chars := "0123456789ABCDEFGHIJKLMÑNOPQRSTUVWXYZ"
	var suma, digito int
	for i := 0; i < 17; i++ {
		charIndex := strings.Index(chars, string(curp[i]))
		if string(curp[i]) == "Ñ" {
			charIndex = 23
		}
		if charIndex == -1 {
			return errors.New("string contains invalid chars")
		}
		// Hay que multiplcar el valor de la tabla por
		// la posicion en la cadena (de 18 para abajo)
		suma = suma + (charIndex * (18 - i))
		// fmt.Printf("Char: %s Posicion: %d Valor: %d\n", string(curp[i]), (18 - i), charIndex)
	}

	digito = 10 - (suma % 10)
	if digito == 10 {
		digito = 0
	}
	// fmt.Printf("Digito es %d vs %s", digito, string(curp[17]))
	lastInt, err := strconv.Atoi(string(curp[17]))
	if err != nil {
		return errors.New("Imposible encontrar digito")
	}
	if lastInt != digito {
		return errors.New("string is not a well formed curp")
	}
	return nil

	/*

			// GAVA760210HDFRYD06
		    //Validar que coincida el dígito verificador
		    function digitoVerificador(curp17) {
		        // Fuente https://consultas.curp.gob.mx/CurpSP/
		        var diccionario  = "0123456789ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
		            lngSuma      = 0.0,
		            lngDigito    = 0.0;
		        for(var i=0; i<17; i++)
		            lngSuma = lngSuma + diccionario.indexOf(curp17.charAt(i)) * (18 - i);
		        lngDigito = 10 - lngSuma % 10;
		        if (lngDigito == 10) return 0;
		        return lngDigito;
		    }

		    if (validado[2] != digitoVerificador(validado[1]))
		    	return false;
	*/
}
