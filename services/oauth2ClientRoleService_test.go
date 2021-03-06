/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package services

import (
	"fmt"
	"strconv"
	"testing"
)

var CID4 int64
var rID int64

func TestClientRoleService_AddClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var uri RedirectURI
	uri.URI = "http://googole.com"
	var uris []RedirectURI
	uris = append(uris, uri)
	var cc Client
	cc.Email = "ken@ken.com"
	cc.Enabled = true
	cc.Name = "A Big Company"
	cc.RedirectURIs = uris
	res := c.AddClient(&cc)
	fmt.Print("add client res: ")
	fmt.Println(res)
	CID4 = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientRoleService_AddClientRole(t *testing.T) {
	var c ClientRoleService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var cr ClientRole
	cr.Role = "user"
	cr.ClientID = CID4
	res := c.AddClientRole(&cr)

	fmt.Print("add client role res: ")
	fmt.Println(res)
	rID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientRoleService_GetClientRoleList(t *testing.T) {
	var c ClientRoleService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetClientRoleList(strconv.FormatInt(CID4, 10))
	fmt.Print("client role res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))

	if res == nil || len(*res) != 1 || (*res)[0].Role != "user" {
		t.Fail()
	}
}

func TestClientRoleService_DeleteClientRole(t *testing.T) {
	var c ClientRoleService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteClientRole(strconv.FormatInt(rID, 10))
	fmt.Print("res deleted client role: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestClientRoleService_DeleteClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteClient(strconv.FormatInt(CID4, 10))
	fmt.Print("res deleted client: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
