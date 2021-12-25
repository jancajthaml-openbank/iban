// Copyright (c) 2016-2021, Jan Cajthaml <jan.cajthaml@gmail.com>
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

package iban

// Calculate calculates IBAN given number and bank identity code
func Calculate(number string, bic string) string {
	switch bic {
	case "FIOBCZPP", "FIOBCZPPXXX":
		{
			return calculateCzech(number, "2010")
		}
	case "KOMBCZPP", "KOMBCZPPXXX":
		{
			return calculateCzech(number, "0100")
		}
	case "CEKOCZPP", "CEKOCZPPXXX":
		{
			return calculateCzech(number, "0300")
		}
	case "AGBACZPP", "AGBACZPPXXX":
		{
			return calculateCzech(number, "0600")
		}
	case "CNBACZPP", "CNBACZPPXXX":
		{
			return calculateCzech(number, "0710")
		}
	case "GIBACZPX", "GIBACZPXXXX":
		{
			return calculateCzech(number, "0800")
		}
	case "BOTKCZPP", "BOTKCZPPXXX":
		{
			return calculateCzech(number, "2020")
		}
	case "CITFCZPP", "CITFCZPPXXX":
		{
			return calculateCzech(number, "2060")
		}
	case "MPUBCZPP", "MPUBCZPPXXX":
		{
			return calculateCzech(number, "2070")
		}
	case "ARTTCZPP", "ARTTCZPPXXX":
		{
			return calculateCzech(number, "2220")
		}
	case "POBNCZPP", "POBNCZPPXXX":
		{
			return calculateCzech(number, "2240")
		}
	case "CTASCZ22", "CTASCZ22XXX":
		{
			return calculateCzech(number, "2250")
		}
	case "ZUNOCZPP", "ZUNOCZPPXXX":
		{
			return calculateCzech(number, "2310")
		}
	case "CITICZPX", "CITICZPXXXX":
		{
			return calculateCzech(number, "2600")
		}
	case "BACXCZPP", "BACXCZPPXXX":
		{
			return calculateCzech(number, "2700")
		}
	case "AIRACZPP", "AIRACZPPXXX":
		{
			return calculateCzech(number, "3030")
		}
	case "INGBCZPP", "INGBCZPPXXX":
		{
			return calculateCzech(number, "3500")
		}
	case "SOLACZPP", "SOLACZPPXXX":
		{
			return calculateCzech(number, "4000")
		}
	case "CMZRCZP1", "CMZRCZP1XXX":
		{
			return calculateCzech(number, "4300")
		}
	case "RZBCCZPP", "RZBCCZPPXXX":
		{
			return calculateCzech(number, "5500")
		}
	case "JTBPCZPP", "JTBPCZPPXXX":
		{
			return calculateCzech(number, "5800")
		}
	case "PMBPCZPP", "PMBPCZPPXXX":
		{
			return calculateCzech(number, "6000")
		}
	case "EQBKCZPP", "EQBKCZPPXXX":
		{
			return calculateCzech(number, "6100")
		}
	case "COBACZPX", "COBACZPXXXX":
		{
			return calculateCzech(number, "6200")
		}
	case "BREXCZPP", "BREXCZPPXXX":
		{
			return calculateCzech(number, "6210")
		}
	case "GEBACZPP", "GEBACZPPXXX":
		{
			return calculateCzech(number, "6300")
		}
	case "SUBACZPP", "SUBACZPPXXX":
		{
			return calculateCzech(number, "6700")
		}
	case "VBOECZ2X", "VBOECZ2XXXX":
		{
			return calculateCzech(number, "6800")
		}
	case "DEUTCZPX", "DEUTCZPXXXX":
		{
			return calculateCzech(number, "7910")
		}
	case "SPWTCZ21", "SPWTCZ21XXX":
		{
			return calculateCzech(number, "7940")
		}
	case "GENOCZ21", "GENOCZ21XXX":
		{
			return calculateCzech(number, "8030")
		}
	case "OBKLCZ2X", "OBKLCZ2XXXX":
		{
			return calculateCzech(number, "8040")
		}
	case "CZEECZPP", "CZEECZPPXXX":
		{
			return calculateCzech(number, "8090")
		}
	case "MIDLCZPP", "MIDLCZPPXXX":
		{
			return calculateCzech(number, "8150")
		}		
	default:
		{
			return ""
		}
	}
}
