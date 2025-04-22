#!/bin/bash

DAC=${1:-./bin/dac}

update_example() {
  for file in $(find examples -name "*.json"); do
    echo "Updating $file"
    $DAC -i $file -o ${file%.json}.yaml -m ${file%.json}.md
  done
}

update_example
