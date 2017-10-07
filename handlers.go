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

package main

import (
	"net/http"
)

// user handlers-----------------------------------------------------
func handleIndex(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//page := vars["content"]
	//if page == "" {
	//	page = "home"
	//}
	//var c services.ContentService
	//c.ClientID = getAuthCodeClient()
	//c.APIKey = getGatewayAPIKey()
	//c.Host = getContentHost()
	//h, res := c.GetContentListCategory(getAuthCodeClient(), page)
	//var pg = new(pageContent)
	//pg.Cont = res
	//pg.MetaAuthor = h.MetaAuthor
	//pg.MetaKeyWords = h.MetaKeyWords
	//pg.MetaDesc = h.MetaDesc
	//pg.Title = h.Title
	templates.ExecuteTemplate(w, "index.html", nil)
}