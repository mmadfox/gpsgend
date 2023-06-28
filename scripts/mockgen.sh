#!/bin/bash

declare -a files=(
    "internal/device/storage.go"
    "internal/device/query.go"
    "internal/device/generator.go"
    "internal/device/publisher.go"
    "internal/device/usecase.go"
)

path="$(pwd)"
for file in "${files[@]}"
do
   filepath="${file/internal/}"
   docker run --rm -v $path:$path github.com/mmadfox/gpsgend/mock -destination="$path/tests/mocks$filepath" -source="$path/$file"
done;