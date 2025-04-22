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
	var markdown = flag.Bool("m", false, "output markdown")
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

	if *markdown {
		fmt.Printf("%s", dashboard.ToMarkdown())
	}

	// If output file is specified, write to file.
	if *output != "" {
		err = os.WriteFile(*output, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("%s", string(data))
	}
}
