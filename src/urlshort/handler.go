package urlshort

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func parseYAML(yamlData []byte) ([]map[string]string, error) {
	var m []map[string]string
	if err := yaml.Unmarshal(yamlData, &m); err != nil {

		return nil, err
	}

	return m, nil
}

func buildMap(parsedYaml []map[string]string) (map[string]string, error) {
	if parsedYaml == nil {
		err := errors.New("Parsed YAML is empty")
		return nil, err
	}
	properMap := make(map[string]string)
	for _, paths := range parsedYaml {
		properMap[paths["path"]] = paths["url"]
	}

	return properMap, nil
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap, err := buildMap(parsedYaml)
	if err != nil {
		log.Fatal(err)
	}

	return MapHandler(pathMap, fallback), nil
}
