package urlshort

import (
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(response, request, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(response, request)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	handles, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	handleMap := handleToMap(handles)
	return MapHandler(handleMap, fallback), nil
}

func handleToMap(handles []handle) map[string]string {
	handleMap := make(map[string]string)
	for _, v := range handles {
		handleMap[v.Path] = v.URL
	}
	return handleMap
}

func parseYaml(yamlContent []byte) ([]handle, error) {
	var pa []handle
	err := yaml.Unmarshal(yamlContent, &pa)
	if err != nil {
		return nil, err
	}
	return pa, nil
}

type handle struct {
	//Path is the key to the URL
	Path string `yaml:"path"`
	//URL full address to redirect
	URL string `yaml:"url"`
}
