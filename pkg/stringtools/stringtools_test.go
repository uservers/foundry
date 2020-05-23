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
package stringtools_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uServers/foundry/pkg/stringtool"
)

func TestIsValidDomain(t *testing.T) {

	validDomain := stringtool.IsValidDomain("uservers.net")
	domainWithSpace := stringtool.IsValidDomain("user vers.net")

	require.Nil(t, validDomain)
	require.Error(t, domainWithSpace)

	require.Equal(t, tc.expected, actual)
}
