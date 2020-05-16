package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

//func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
//
//	return nil
//}
//

type Urls struct {
	Path []map[string]string
}

func parseYaml(yamlData []byte) {
	var u Urls
	if err := yaml.Unmarshal(yamlData, u); err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}

//func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
//	parsedYaml, err := parseYAML(yaml)
//	if err != nil {
//		return nil, err
//	}
//	pathMap := buildMap(parsedYaml)
//	return MapHandler(pathMap, fallback), nil
//}

func main() {
	yamlData := `
    - path: '/google'
      url: 'https://google.com/'
    - path: '/stackoverflow'
      url: 'https://stackoverflow.com/'
      `

	parseYaml([]byte(yamlData))
}
