/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"github.com/infinitbyte/framework/core/api"
	"github.com/infinitbyte/framework/core/util"
	"github.com/medcl/elasticsearch-proxy/config"
	"net/http"
)

// API namespace
type API struct {
	api.Handler
	authConfig config.BasicAuthConfig
}

// InitAPI init apis
func InitAPI(authConfig config.BasicAuthConfig) {

	apis := API{authConfig: authConfig}

	//Index
	api.HandleAPIMethod(api.GET, "/", apis.IndexAction)
	//api.HandleAPIMethod(api.GET, "/favicon.ico", apis.FaviconAction)

	//Stats APIs
	api.HandleAPIMethod(api.GET, "/_proxy/stats", apis.StatsAction)
	api.HandleAPIMethod(api.GET, "/_proxy/queue/stats", apis.QueueStatsAction)

	// Handle proxy
	api.HandleAPIFunc("/", apis.BasicAuth(apis.ProxyAction))
}

func (api *API) BasicAuth(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if (!api.authConfig.Enabled) || (hasAuth && user == api.authConfig.User.Username && password == api.authConfig.User.Password) {
			// Delegate request to the given handle
			handler(w, r)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(util.ToJSONBytes(util.MapStr{
				"error": http.StatusText(http.StatusUnauthorized),
			}))
		}
	}
}
