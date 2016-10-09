// Copyright © 2016 Daniele Sluijters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// Decode takes bytes and will search for any and all certificates in them.
func Decode(certificates []byte, res []*x509.Certificate) []*x509.Certificate {
	block, rest := pem.Decode(certificates)
	if block == nil {
		fmt.Println("No (further) certificates in keychain")
		return res
	}
	crt, _ := x509.ParseCertificates(block.Bytes)
	res = append(res, crt[0])
	if len(rest) != 0 {
		return Decode(rest, res)
	}
	return res
}
