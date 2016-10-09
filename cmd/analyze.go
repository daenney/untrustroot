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

package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/daenney/untrustroot/certificate"
	"github.com/daenney/untrustroot/security"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Output information about CA's found in the keychain",
	Long: `Provide keychain statistics.

This includes things like the issuing countries, signing algorithms
and other things of potential interest.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		keychain, _ := cmd.Flags().GetString("keychain")
		contents, err := security.ReadCertsFromKeychain(keychain)
		if err != nil {
			return err
		}
		certs := certificate.Decode(contents, nil)
		var ocsp = 0
		var crl = 0
		countries := make(map[string]int)
		sigalgs := make(map[string]int)
		for _, crt := range certs {
			if len(crt.Issuer.Country) > 0 {
				countries[strings.ToUpper(crt.Issuer.Country[0])]++
			}
			sigalgs[crt.SignatureAlgorithm.String()]++
			if len(crt.OCSPServer) > 0 {
				ocsp++
			}
			if len(crt.CRLDistributionPoints) > 0 {
				crl++
			}
		}
		fmt.Printf("There are %d CA's in your keychain:\n", len(certs))
		fmt.Printf("  • %d have CRL distribution points specified\n", crl)
		fmt.Printf("  • %d have OCSP responsders specified\n", ocsp)
		fmt.Println("Using the following signing algorithms:")
		printMapByPopularity(sigalgs)
		fmt.Println("Issued by entities in the following countries:")
		printMapByPopularity(countries)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
}

func printMapByPopularity(data map[string]int) {
	n := map[int][]string{}
	var a []int
	for k, v := range data {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("  • %s: %d\n", s, k)
		}
	}
}
