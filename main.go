package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zyy17/dac/pkg/grafana"
	"github.com/zyy17/dac/pkg/transform"
	"gopkg.in/yaml.v3"
)

func main() {
	var input = flag.String("i", "", "input file")
	flag.Parse()

	if input == nil {
		log.Fatal("input file is required")
	}

	content, err := os.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	var grafanaDashboard grafana.Dashboard
	err = json.Unmarshal(content, &grafanaDashboard)
	if err != nil {
		log.Fatal(err)
	}

	dashboard, err := transform.Transform(&grafanaDashboard)
	if err != nil {
		log.Fatal(err)
	}

	data, err := yaml.Marshal(dashboard)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(data))
}
