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

const (
	userSession       = "ugw-user-session"
	sessingTimeToLive = (120 * 60) //120 minutes -- 2 hours

	//http
	schemeDefault = "http://"

	//OAuth Auth Code
	authCodeClient      = "403"
	authCodeSecret      = "554444vfg55ggfff22454sw2fff2dsfd"
	authCodeState       = "ghh66555h"
	authCodeRedirectURI = "/tokenHandler"
)
