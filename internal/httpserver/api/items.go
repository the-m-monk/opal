package api

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"opal/internal/librarymgmt"
	"opal/internal/usermgmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

const imageNamespace = "9427cfd6-d8cf-47d0-82aa-86c6d17bb390"

func encodeAndServePNG(w http.ResponseWriter, img image.Image, cachePath string) {
	os.MkdirAll(filepath.Dir(cachePath), 0755)

	out, err := os.Create(cachePath)
	if err == nil {
		defer out.Close()
		png.Encode(out, img)
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

// http://localhost:7096/Items/eceec891-3674-5020-ad92-a9cc24d640b2/Images/Primary
// fillHeight=200
// fillWidth=355
// quality=96 - unused
// tag=backdrop.png
func EndpointItemsByUuidImagesPrimary(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	itemUuid := r.PathValue("uuid")
	tag := query.Get("tag")

	height, _ := strconv.Atoi(query.Get("fillHeight"))
	width, _ := strconv.Atoi(query.Get("fillWidth"))

	iNS := uuid.MustParse(imageNamespace)

	cacheUuidData := fmt.Sprintf("primary-%s-%d-%d-%s", itemUuid, width, height, tag)
	cacheUuid := uuid.NewSHA1(iNS, []byte(cacheUuidData)).String()

	imageCachePath := filepath.Join(librarymgmt.CacheDir, "images")
	imageFileCachePath := filepath.Join(imageCachePath, cacheUuid+".png")

	if _, err := os.Stat(imageFileCachePath); err == nil {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, imageFileCachePath)
		return
	}

	item := librarymgmt.LibTreeMap[itemUuid]
	if item == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	var sourcePath string
	switch tag {
	case "backdrop.png":
		sourcePath = filepath.Join(librarymgmt.MetadataDir, item.ImdbId, "backdrop.png")
	case "library.png":
		sourcePath = filepath.Join(librarymgmt.CacheDir, "images", item.RootUuid+".png")
	case "poster.png":
		sourcePath = filepath.Join(librarymgmt.MetadataDir, item.ImdbId, "poster.png")
	default:
		http.Error(w, "tag type not supported", http.StatusNotImplemented)
		return
	}

	img, err := imaging.Open(sourcePath)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	retImage := imaging.Resize(img, width, height, imaging.Lanczos)
	encodeAndServePNG(w, retImage, imageFileCachePath)
}

// http://localhost:7096/Items/f0776435-32bc-5a22-b72c-28d6815777c3/Images/Logo
// tag=logo.png
// quality=90 - unused
func EndpointItemsUuidImagesLogo(w http.ResponseWriter, r *http.Request) {
	//Note: this endpoint does not require authentication (i think)
	query := r.URL.Query()

	itemUuid := r.PathValue("uuid")
	tag := query.Get("tag")

	iNS := uuid.MustParse(imageNamespace)

	cacheUuidData := fmt.Sprintf("logo-%s-%s", itemUuid, tag)
	cacheUuid := uuid.NewSHA1(iNS, []byte(cacheUuidData)).String()

	imageFileCachePath := filepath.Join(librarymgmt.CacheDir, "images", cacheUuid+".png")

	if _, err := os.Stat(imageFileCachePath); err == nil {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, imageFileCachePath)
		return
	}

	item := librarymgmt.LibTreeMap[itemUuid]
	if item == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	imagePath := filepath.Join(librarymgmt.MetadataDir, item.ImdbId, "logo.png")
	img, err := imaging.Open(imagePath)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	encodeAndServePNG(w, img, imageFileCachePath)
}

func EndpointItemsUuid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	query := r.URL.Query()

	userId := query.Get("userId")
	aTokenString := getHeaderAuthToken(r)

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && userId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	itemUuid := r.PathValue("uuid")

	if librarymgmt.LibTreeMap[itemUuid] == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	switch librarymgmt.LibTreeMap[itemUuid].Type {
	case "folder":
		EndpointUsersItemsByUuidFolder(w, r, itemUuid)
	case "tvshow":
		EndpointUsersItemsByUuidTvshow(w, r, itemUuid)
	case "movie":
		EndpointUsersItemsByUuidMovie(w, r, itemUuid)
	}
}

func EndpointItemsUuidImagesBackdrop(w http.ResponseWriter, r *http.Request) {
	//Note: this endpoint does not require authentication (i think)
	query := r.URL.Query()

	itemUuid := r.PathValue("uuid")
	tag := query.Get("tag")

	iNS := uuid.MustParse(imageNamespace)

	cacheUuidData := fmt.Sprintf("logo-%s-%s", itemUuid, tag)
	cacheUuid := uuid.NewSHA1(iNS, []byte(cacheUuidData)).String()

	imageFileCachePath := filepath.Join(librarymgmt.CacheDir, "images", cacheUuid+".png")

	if _, err := os.Stat(imageFileCachePath); err == nil {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, imageFileCachePath)
		return
	}

	item := librarymgmt.LibTreeMap[itemUuid]
	if item == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	imagePath := filepath.Join(librarymgmt.MetadataDir, item.ImdbId, "backdrop.png")
	img, err := imaging.Open(imagePath)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	encodeAndServePNG(w, img, imageFileCachePath)
}
