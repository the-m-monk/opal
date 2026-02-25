package api

import (
	"net/http"
	"opal/internal/librarymgmt"
	"opal/internal/usermgmt"
	"path"
)

// http://localhost:8096/Videos/84c329ff4d2221404996eb00855760db/stream.mp4
// Static=true
// mediaSourceId=84c329ff4d2221404996eb00855760db
// deviceId=TW96aWxsYS81LjAgKFgxMTsgTGludXggeDg2XzY0OyBydjoxNDcuMCkgR2Vja28vMjAxMDAxMDEgRmlyZWZveC8xNDcuMHwxNzcxMjg5NzM4NjY2
// ApiKey=9ab58966418e4178a11c97f6c0bff85e&Tag=ca89d7b53434618d799790ff89fbd0fe
func EndpointVideosUuidStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	apiKey := query.Get("ApiKey")

	isApiKeyValid := false
	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == apiKey {
			isApiKeyValid = true
		}
	}

	if !isApiKeyValid {
		http.Error(w, "Invalid api key", http.StatusUnauthorized)
		return
	}

	itemUuid := r.PathValue("uuid")

	item := librarymgmt.LibTreeMap[itemUuid]

	if item == nil || item.Type != "movie" {
		http.Error(w, "item uuid does not resolve to a valid media uuid", http.StatusNotFound)
		return
	}

	libraryRecord := librarymgmt.AllLibrariesMap[item.RootUuid]

	if libraryRecord == nil {
		http.Error(w, "Library not found", http.StatusNotFound)
		return
	}

	filePath := path.Join(libraryRecord.Path, item.Path)

	http.ServeFile(w, r, filePath)
}
