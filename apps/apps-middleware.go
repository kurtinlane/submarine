package apps

import (
	"net/http"
	"github.com/go-martini/martini"
	"appengine"
)

func ResolveApp(context martini.Context, res http.ResponseWriter, req *http.Request){
	apiKey := req.Header.Get("API-KEY")
	
	if apiKey == "" {
		http.Error(res, "bad request", http.StatusBadRequest)
	} else {
		appEngineContext := appengine.NewContext(req) 
		app, err := GetAppWithApiKey(apiKey, appEngineContext)
		
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		} else {
			context.Map(app)
		}
	}
}