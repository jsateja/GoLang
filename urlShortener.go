package main

import (
	"fmt"
	"log"
	"urlshort"
)

func main() {
	yamlData := `
    - path: '/google'
      url: 'https://google.com/'
    - path: '/stackoverflow'
      url: 'https://stackoverflow.com/'
      `

	x, err := urlshort.YAMLHandler([]byte(yamlData))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
}
