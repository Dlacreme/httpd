package boot

import (
	"log"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/config/env"
	"github.com/Dlacreme/httpd/view/viewfunc/link"
	"github.com/Dlacreme/httpd/view/viewfunc/noescape"
	"github.com/Dlacreme/httpd/view/viewfunc/prettytime"
	"github.com/Dlacreme/httpd/view/viewmodify/authlevel"
	"github.com/Dlacreme/httpd/view/viewmodify/flash"
	"github.com/Dlacreme/httpd/view/viewmodify/uri"
	"github.com/Dlacreme/httpd/view/xsrf"
	"github.com/Dlacreme/httpd/webtools/form"
	"github.com/Dlacreme/httpd/webtools/pagination"
	"github.com/jmoiron/sqlx"
)

// LoadDefaultServices sets up the basic components. This function may be overload if you need more stuff.
func LoadDefaultServices(config *env.Info) *sqlx.DB {
	// Set up the session cookie store
	err := config.Session.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MySQL database
	mysqlDB, err := config.MySQL.Connect(true)
	if err != nil {
		panic(err)
	}

	// Set up the views
	config.View.SetTemplates(config.Template.Root, config.Template.Children)

	// Set up the functions for the views
	config.View.SetFuncMaps(
		config.Asset.Map(config.View.BaseURI),
		link.Map(config.View.BaseURI),
		noescape.Map(),
		prettytime.Map(),
		form.Map(),
		pagination.Map(),
	)

	// Set up the variables and modifiers for the views
	config.View.SetModifiers(
		authlevel.Modify,
		uri.Modify,
		xsrf.Token,
		flash.Modify,
	)

	// Store the variables in flight
	flight.StoreConfig(*config)

	// Store the database connection in flight
	flight.StoreDB(mysqlDB)

	// Store the csrf information
	flight.StoreXsrf(xsrf.Info{
		AuthKey: config.Session.CSRFKey,
		Secure:  config.Session.Options.Secure,
	})

	return mysqlDB
}
