#!/usr/bin/env sh

# cli
docker build -t exercism_build dockerfiles/exercism_build

# language images
declare -a supported_languages=("go" "python")

for i in "${supported_languages[@]}";
do
    echo "$i"
    docker build -t "$i"_test dockerfiles/"$i"_test_environment
done
