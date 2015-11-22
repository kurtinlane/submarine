package server

import (
	"github.com/go-martini/martini"
	"github.com/kurtinlane/submarine/apps"
	"github.com/kurtinlane/submarine/keys"
	"net/http"
)


func init() {
	//Separate apis because use different middleware to authenticate requests... might be a better way to handle this
	publicApi := martini.Classic()
	publicApi.Use(apps.ResolveApp)
	
	privateApi := martini.Classic()
	
	apps.RegisterWebService(privateApi)
	keys.RegisterWebService(publicApi)
	
	http.Handle("/api/v1/", publicApi)
	http.Handle("/", privateApi)
}
