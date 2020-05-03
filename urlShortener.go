package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Urls struct {
	URLs struct {
		Dogs []string
		Cats []string
	}
}

//func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
//
//	return nil
//}
//

func parseYaml() {

	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading yaml file: %s\n", err)
		return
	}

	var u Urls
	err = yaml.Unmarshal(yamlFile, &u)
	if err != nil {
		fmt.Printf("Error parsing yaml: %s\n", err)
	}

	d, err := yaml.Marshal(&u)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- wuut:\n%s\n\n", string(d))

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
	parseYaml()
}
