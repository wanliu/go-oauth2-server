package web

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// Redirects to a new path while keeping current request's query string
func redirectWithQueryString(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s%s", to, getQueryString(query)), http.StatusFound)
}

// Redirects to a new path with the query string moved to the URL fragment
func redirectWithFragment(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s#%s", to, query.Encode()), http.StatusFound)
}

// Returns string encoded query string of the request
func getQueryString(query url.Values) string {
	encoded := query.Encode()
	if len(encoded) > 0 {
		encoded = fmt.Sprintf("?%s", encoded)
	}
	return encoded
}

// Helper function to handle redirecting failed or declined authorization
func errorRedirect(w http.ResponseWriter, r *http.Request, redirectURI *url.URL, err, state, responseType string) {
	query := redirectURI.Query()
	query.Set("error", err)
	if state != "" {
		query.Set("state", state)
	}
	if responseType == "code" {
		redirectWithQueryString(redirectURI.String(), query, w, r)
	}
	if responseType == "token" {
		redirectWithFragment(redirectURI.String(), query, w, r)
	}
}

func uploadFile(field string, r *http.Request) (string, error) {
	file, _, err := r.FormFile(field)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	// fmt.Fprintf(w, "%v", handler.Header)
	os.MkdirAll("./public/uploads/", os.ModePerm)
	filename := filepath.Join("./public/uploads/", buildFileName(".png"))

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	return filename, err
}

func buildFileName(ext string) string {
	return time.Now().Format("20060102150405") + ext
}
