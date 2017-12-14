#!/usr/bin/bash

echo 'Building your application...'
echo '---'

output_name=$1
if [ $# -eq 0 ]; then
    output_name='my-go-server'
fi

platforms=("windows/amd64" "windows/386" "darwin/amd64")

for platform in "${platforms[@]}"
do
    echo "Building for $platform..."
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$output_name'_'$GOOS'_'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    echo "Built done for $platform"
done

echo '---'
echo 'Buil process FINISHED'