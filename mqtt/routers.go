/*
 * Mqtts API
 *
 * Mqtts service
 *
 * API version: 1.0.0
 */
package swagger

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	utils "github.com/menucha-de/utils"
)

//NewRouter create anew Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// AddRoutes add new routes
func AddRoutes(myroute []utils.Route) {
	routes = append(routes, myroute...)
}

var routes = []utils.Route{

	utils.Route{
		"AccesscontrolGet",
		strings.ToUpper("Get"),
		"/rest/mqtt/accesscontrol",
		accessControlGet,
	},
	utils.Route{
		"AccessControlSet",
		strings.ToUpper("Put"),
		"/rest/mqtt/accesscontrol",
		accessControlSet,
	},
	utils.Route{
		"Deleteacfile",
		strings.ToUpper("Delete"),
		"/rest/mqtt/accesscontrol/acfile",
		deleteACFile,
	},

	utils.Route{
		"Deletecafile",
		strings.ToUpper("Delete"),
		"/rest/mqtt/security/trust",
		deleteCAFile,
	},

	utils.Route{
		"Deletepassfile",
		strings.ToUpper("Delete"),
		"/rest/mqtt/accesscontrol/passwordfile",
		deletePassFile,
	},

	utils.Route{
		"Deleterevfile",
		strings.ToUpper("Delete"),
		"/rest/mqtt/security/revoclist",
		deleteRevFile,
	},

	utils.Route{
		"Deleteservfile",
		strings.ToUpper("Delete"),
		"/rest/mqtt/security/keystore",
		deleteServFile,
	},

	utils.Route{
		"Getacfile",
		strings.ToUpper("Get"),
		"/rest/mqtt/accesscontrol/acfile",
		getACFile,
	},

	utils.Route{
		"Getcertfile",
		strings.ToUpper("Head"),
		"/rest/mqtt/security/trust",
		getCertFile,
	},

	utils.Route{
		"Getpassfile",
		strings.ToUpper("Get"),
		"/rest/mqtt/accesscontrol/passwordfile",
		getPassFile,
	},

	utils.Route{
		"Getrevfile",
		strings.ToUpper("Head"),
		"/rest/mqtt/security/revoclist",
		getRevFile,
	},

	utils.Route{
		"Getservfile",
		strings.ToUpper("Head"),
		"/rest/mqtt/security/keystore",
		getServFile,
	},

	utils.Route{
		"Headacfile",
		strings.ToUpper("Head"),
		"/rest/mqtt/accesscontrol/acfile",
		headACFile,
	},

	utils.Route{
		"Headpassfile",
		strings.ToUpper("Head"),
		"/rest/mqtt/accesscontrol/passwordfile",
		headPassFile,
	},

	utils.Route{
		"Putacfile",
		strings.ToUpper("Put"),
		"/rest/mqtt/accesscontrol/acfile",
		putACFile,
	},

	utils.Route{
		"Putcertfile",
		strings.ToUpper("Put"),
		"/rest/mqtt/security/trust",
		putCertFile,
	},

	utils.Route{
		"Putpassfile",
		strings.ToUpper("Put"),
		"/rest/mqtt/accesscontrol/passwordfile",
		putPassFile,
	},

	utils.Route{
		"Putpassphrase",
		strings.ToUpper("Put"),
		"/rest/mqtt/security/passphrase",
		putPassPhrase,
	},

	utils.Route{
		"Putrevfile",
		strings.ToUpper("Put"),
		"/rest/mqtt/security/revoclist",
		putRevFile,
	},

	utils.Route{
		"Putservfile",
		strings.ToUpper("Put"),
		"/rest/mqtt/security/keystore",
		putServFile,
	},

	utils.Route{
		"SecurityGet",
		strings.ToUpper("Get"),
		"/rest/mqtt/security",
		securityGet,
	},
	utils.Route{
		"SecuritySet",
		strings.ToUpper("Put"),
		"/rest/mqtt/security",
		securitySet,
	},
}
