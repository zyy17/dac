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
	var output = flag.String("o", "", "output file")
	var markdown = flag.String("m", "", "output path for markdown")
	flag.Parse()

	if input == nil {
		log.Fatal("input file is required")
	}

	content, err := os.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	var grafanaDashboard grafana.Dashboard
	if err = json.Unmarshal(content, &grafanaDashboard); err != nil {
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

	if *markdown != "" {
		if err := os.WriteFile(*markdown, []byte(dashboard.ToMarkdown()), 0644); err != nil {
			log.Fatal(err)
		}
	}

	// If output file is specified, write to file.
	if *output != "" {
		if err := os.WriteFile(*output, data, 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("%s", string(data))
	}
}
