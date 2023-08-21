#!/bin/bash

declare -a files=(
    "internal/generator/storage.go"
    "internal/generator/processes.go"
    "internal/generator/bootstraper.go"
    "internal/generator/projection.go"
)

path="$(pwd)"
for file in "${files[@]}"
do
   filepath="${file/internal/}"
   docker run --rm -v $path:$path github.com/mmadfox/gpsgend/mock -destination="$path/tests/mocks$filepath" -source="$path/$file"
done;