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
	services "ApiGatewayAdminPortal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleRedirectURLs(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		session.Values["userLoggenIn"] = true
		vars := mux.Vars(r)
		clientID := vars["clientId"]

		if clientID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient(clientID)
			//fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res

			var r services.RedirectURIService
			r.ClientID = getAuthCodeClient()
			r.Host = getOauthHost()
			r.Token = token.AccessToken
			res2 := r.GetRedirectURIList(clientID)
			page.RedirectURLs = res2
			if len(*res2) > 1 {
				page.CanDeleteRedirectURI = true
			}
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
		}
	}
}

func handleRedirectURLDelete(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		session.Values["userLoggenIn"] = true
		vars := mux.Vars(r)
		ID := vars["id"]
		clientID := vars["clientId"]

		if ID != "" && clientID != "" {
			token := getToken(w, r)
			var r services.RedirectURIService
			r.ClientID = getAuthCodeClient()
			r.Host = getOauthHost()
			r.Token = token.AccessToken
			dres := r.DeleteRedirectURI(ID)
			if dres.Success != true {
				fmt.Println(dres)
			}

			var c services.ClientService

			c.ClientID = getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient(clientID)
			//fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res

			res2 := r.GetRedirectURIList(clientID)
			page.RedirectURLs = res2
			if len(*res2) > 1 {
				page.CanDeleteRedirectURI = true
			}
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
		}
	}
}

func handleRedirectURLAdd(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {

		redirectURL := r.FormValue("redirectURL")
		fmt.Println(redirectURL)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		session.Values["userLoggenIn"] = true

		token := getToken(w, r)
		var r services.RedirectURIService
		r.ClientID = getAuthCodeClient()
		r.Host = getOauthHost()
		r.Token = token.AccessToken

		var rr services.RedirectURI
		rr.ClientID = clientID
		rr.URI = redirectURL
		ares := r.AddRedirectURI(&rr)
		if ares.Success != true {
			fmt.Println(ares)
		}

		var c services.ClientService

		c.ClientID = getAuthCodeClient()
		c.Host = getOauthHost()
		c.Token = token.AccessToken

		res := c.GetClient(clientIDStr)
		//fmt.Println(res)
		var page oauthPage
		page.OauthActive = "active"
		page.Client = res

		res2 := r.GetRedirectURIList(clientIDStr)
		page.RedirectURLs = res2
		if len(*res2) > 1 {
			page.CanDeleteRedirectURI = true
		}
		//fmt.Println(page)
		templates.ExecuteTemplate(w, "redirectUrls.html", &page)

	}
}
