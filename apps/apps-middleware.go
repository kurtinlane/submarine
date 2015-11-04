package apps

import (
	"net/http"
	"github.com/go-martini/martini"
)

func ResolveApp(context martini.Context, res http.ResponseWriter, req *http.Request){
	apiKey := req.Header.Get("API-KEY")
	
	if apiKey == "" {
		http.Error(res, "bad request", http.StatusBadRequest)
	} else {
		app, err := GetAppWithApiKey(apiKey)
		
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		} else {
			context.Map(app)
		}
	}
}