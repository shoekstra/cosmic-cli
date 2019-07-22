//
// Copyright © 2019 Stephen Hoekstra
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
//

package cosmic

import (
	"fmt"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
)

// WhoHasThisIP embeds *cosmic.WhoHasThisIP to allow additional fields.
type WhoHasThisIP struct {
	*cosmic.WhoHasThisIp
	Vpcname  string
	Zonename string
}

// WhoHasThisIPs exists to provide helper methods for []*WhoHasThisIP.
type WhoHasThisIPs []*WhoHasThisIP

// WhoHasThisMac embeds *cosmic.WhoHasThisMac to allow additional fields.
type WhoHasThisMac struct {
	*cosmic.WhoHasThisMac
	Vpcname  string
	Zonename string
}

// WhoHasThisMacs exists to provide helper methods for []*WhoHasThisMac.
type WhoHasThisMacs []*WhoHasThisMac

// ListIP returns a WhoHasThisIPs object using all configured *cosmic.CosmicClient objects.
func ListIP(clientMap map[string]*cosmic.CosmicClient, ipaddress string) (WhoHasThisIPs, error) {
	ips := []*WhoHasThisIP{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			// Zonename isn't returned in *cosmic.ListWhoHasThisIpResponse so we need to fetch it
			zoneparams := clientMap[client].Zone.NewListZonesParams()
			zoneresp, err := clientMap[client].Zone.ListZones(zoneparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			zonename := zoneresp.Zones[0].Name

			// VPCName isn't returned in *cosmic.ListWhoHasThisIpResponse so we need to fetch it
			netparams := clientMap[client].Network.NewListNetworksParams()
			netresp, err := clientMap[client].Network.ListNetworks(netparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			vpcparams := clientMap[client].VPC.NewListVPCsParams()
			vpcresp, err := clientMap[client].VPC.ListVPCs(vpcparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			params := clientMap[client].CloudOps.NewListWhoHasThisIpParams(ipaddress)
			resp, err := clientMap[client].CloudOps.ListWhoHasThisIp(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, ip := range resp.WhoHasThisIp {
				vpcid := ""
				vpcname := ""

				for _, n := range netresp.Networks {
					if n.Id == ip.Networkuuid {
						ip.Networkname = n.Name
						vpcid = n.Vpcid
						break
					}
				}
				for _, v := range vpcresp.VPCs {
					if v.Id == vpcid {
						vpcname = v.Name
						break
					}
				}

				// When returning a public IP address the network name is populated with the VPC name,
				// to avoid confusion we'll empty out the network name field so that only the VPC name
				// contains the VPC name.
				if ip.Networkname == ip.Vpcname {
					ip.Networkname = ""
					vpcname = ip.Vpcname
				}

				// When returning a public IP adress, the network mask is empty, so we populate it as
				// a /32.
				if ip.Netmask == "" {
					ip.Netmask = "255.255.255.255"
				}

				ips = append(ips, &WhoHasThisIP{
					WhoHasThisIp: ip,
					Vpcname:      vpcname,
					Zonename:     zonename,
				})
			}
		}(client)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return nil, err
		}
	}

	return ips, nil
}

// ListMAC returns a WhoHasThisMacs object using all configured *cosmic.CosmicClient objects.
func ListMAC(clientMap map[string]*cosmic.CosmicClient, macaddress string) (WhoHasThisMacs, error) {
	macs := []*WhoHasThisMac{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			// Zonename isn't returned in *cosmic.ListWhoHasThisMacResponse so we need to fetch it
			zoneparams := clientMap[client].Zone.NewListZonesParams()
			zoneresp, err := clientMap[client].Zone.ListZones(zoneparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			zonename := zoneresp.Zones[0].Name

			// VPCName isn't returned in *cosmic.ListWhoHasThisMacResponse so we need to fetch it
			netparams := clientMap[client].Network.NewListNetworksParams()
			netresp, err := clientMap[client].Network.ListNetworks(netparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			vpcparams := clientMap[client].VPC.NewListVPCsParams()
			vpcresp, err := clientMap[client].VPC.ListVPCs(vpcparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			params := clientMap[client].CloudOps.NewListWhoHasThisMacParams()
			params.SetMacaddress(macaddress)
			resp, err := clientMap[client].CloudOps.ListWhoHasThisMac(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, mac := range resp.WhoHasThisMac {
				vpcid := ""
				vpcname := ""

				for _, n := range netresp.Networks {
					if n.Id == mac.Networkuuid {
						mac.Networkname = n.Name
						vpcid = n.Vpcid
						break
					}
				}
				for _, v := range vpcresp.VPCs {
					if v.Id == vpcid {
						vpcname = v.Name
						break
					}
				}

				// When returning a MAC address the network name is populated with the VPC name,
				// to avoid confusion we'll empty out the network name field so that only the VPC name
				// contains the VPC name.
				if mac.Networkname == mac.Vpcname {
					mac.Networkname = ""
					vpcname = mac.Vpcname
				}

				// When returning a MAC adress, the network mask is empty, so we populate it as
				// a /32.
				if mac.Netmask == "" {
					mac.Netmask = "255.255.255.255"
				}

				macs = append(macs, &WhoHasThisMac{
					WhoHasThisMac: mac,
					Vpcname:       vpcname,
					Zonename:      zonename,
				})
			}
		}(client)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return nil, err
		}
	}

	return macs, nil
}
