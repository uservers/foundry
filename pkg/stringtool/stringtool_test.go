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
	"github.com/uServers/foundry/pkg/stringtool"
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
