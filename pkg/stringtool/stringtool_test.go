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
package stringtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uservers/foundry/pkg/stringtool"
)

func TestSplitDomain(t *testing.T) {
	sld, tld, err := stringtool.SplitDomain("uservers.com.mx")
	require.Equal(t, tld, "com.mx")
	require.Equal(t, sld, "uservers")
	require.Nil(t, err)

	sld, tld, err = stringtool.SplitDomain("uservers.mx")
	require.Equal(t, tld, "mx")
	require.Equal(t, sld, "uservers")
	require.Nil(t, err)

	sld, tld, err = stringtool.SplitDomain("uservers.com")
	require.Equal(t, tld, "com")
	require.Equal(t, sld, "uservers")
	require.Nil(t, err)

	sld, tld, err = stringtool.SplitDomain("with.subdomains.uservers.com")
	require.Equal(t, tld, "com")
	require.Equal(t, sld, "uservers")
	require.Nil(t, err)
}

func TestIsValidDomain(t *testing.T) {
	validDomain := stringtool.IsValidDomain("uservers.net")
	validComMx := stringtool.IsValidDomain("uservers.com.mx")
	endsWithDigit := stringtool.IsValidDomain("uservers1.com")
	extraLongDomain := stringtool.IsValidDomain("12345678901234567890123456789012345678901234567890123456789012345678901234567890.mx")
	domainWithSpace := stringtool.IsValidDomain("user vers.net")
	startsWithDash := stringtool.IsValidDomain("-achido.com")
	endsWithDash := stringtool.IsValidDomain("achido-.com")
	singleCharCom := stringtool.IsValidDomain("a.com")
	dashDomain := stringtool.IsValidDomain("-.com")
	emptyString := stringtool.IsValidDomain("")
	invalidChars := stringtool.IsValidDomain("9324759834_aasdasd.com")
	upperCaseDomain := stringtool.IsValidDomain("CHIDO1.COM")

	require.Nil(t, validDomain, "Dominio valido clasificado mal")
	require.Nil(t, singleCharCom, "Single char mal clasificado")
	require.Nil(t, validComMx, "Dominio com.mx es valido")
	require.Nil(t, endsWithDigit, "Dominio termina en digito y es valido")
	require.Error(t, extraLongDomain, "Dominio es demsiado largo validado")
	require.Error(t, dashDomain, "Dominio solo es guion")
	require.Error(t, startsWithDash, "Dominio empieza con guion")
	require.Error(t, endsWithDash, "Dominio termina con guion")
	require.Error(t, emptyString, "String en blanco pasada como valido")
	require.Error(t, domainWithSpace, "Dominio con espacio entendido como valido")
	require.Error(t, invalidChars, "Dominio tiene caracteres invalidos")
	require.Nil(t, upperCaseDomain, "Dominio es valido en mayusculs")

	// require.Equal(t, tc.expected, actual)
}

func TestIsValidUUID(t *testing.T) {
	require.Nil(t, stringtool.IsValidUUID("8f3d9af5-e51f-4d5c-9c1b-294466af8492"))
	require.NotNil(t, stringtool.IsValidUUID("8f3d9af5-e51f-4d5c-9c1b-294466af849")) // Longer
	require.NotNil(t, stringtool.IsValidUUID("8f3d9af5-e51f-4d5c-9c1b-2944"))        // Shorter
	require.Nil(t, stringtool.IsValidUUID("00000000-0000-0000-0000-000000000000"))   // Null UUID
	require.NotNil(t, stringtool.IsValidUUID("8f3d9af5-e51f-4d5c-9c1b-2A4466af849")) // Invalud char
	require.NotNil(t, stringtool.IsValidUUID(""))                                    // Empry String
}

func TestIsValidCURP(t *testing.T) {
	// Curps de prueba publicados aqui:
	// http://www.itchihuahua.edu.mx/wp-content/uploads/2016/02/aceptados_e@d_febrero_2016.pdf
	require.Nil(t, stringtool.IsValidCURP("AIHP911101MCHRRR03"))    // Valido Mujer
	require.Nil(t, stringtool.IsValidCURP("HOAE940218HCHLGR02"))    // Valido Hombre
	require.NotNil(t, stringtool.IsValidCURP("HOAE940218WCHLGR02")) // No es H/M
	require.NotNil(t, stringtool.IsValidCURP("HOAE940218WCHLGR03")) // Digito mal
	require.NotNil(t, stringtool.IsValidCURP("HOAE943618WCHLGR03")) // Fecha Mal
}

func TestParseStorable(t *testing.T) {
	// String valido
	login, userid, err := stringtool.ParseStorable("user@exampl")
	require.Nil(t, err)
	require.Equal(t, "user", login)
	require.Equal(t, "exampl", userid)

	// String muy corto
	login, userid, err = stringtool.ParseStorable("too@short")
	require.NotNil(t, err)

	// String muy corto
	login, userid, err = stringtool.ParseStorable("1@short")
	require.NotNil(t, err, "No pude cachar login corto")

	// User numerico
	login, userid, err = stringtool.ParseStorable("1@short")
	require.NotNil(t, err, "No pude cachar login corto")

	// Login invalido
	login, userid, err = stringtool.ParseStorable(".user@userid")
	require.NotNil(t, err, "Imposible cachar '.user' como login malo")

	// Userid invalido
	login, userid, err = stringtool.ParseStorable("username@not-bn")
	require.NotNil(t, err, "Falle encontrar 'not-bn' como mal userid")

}
