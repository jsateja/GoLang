package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string) int {

	return 0
}

func parseYAML(yamlData []byte) ([]map[string]string, error) {
	var m []map[string]string
	if err := yaml.Unmarshal(yamlData, &m); err != nil {

		return nil, err
	}

	return m, nil
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	fmt.Println(parsedYaml)
	properMap := make(map[string]string)
	for _, paths := range parsedYaml {
		fmt.Println(paths)
		for k, v := range paths {
			properMap[k] = v
		}
	}
	fmt.Println(properMap)

	return properMap
}

func YAMLHandler(yaml []byte) (int, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return 0, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap), nil
}
