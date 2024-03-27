#!/bin/bash

# Print and get input for repository name
echo -n "URL of your repository (without https://): "
read repositoryUrl

# Print and get input for provider name
echo -n "Name of the provider (e.g. Docker Provider): "
read providerName

# Replace ocurrences of github.com/daytonaio/daytona-provider-sample with the repository name
find . -type d \( -name "hack" -o -name ".git" \) -prune -o -type f -exec sed -i "s|github.com/daytonaio/daytona-provider-sample|$repositoryUrl|g" {} +
echo "Replaced github.com/daytonaio/daytona-provider-sample with $repositoryUrl"

# Replace occurrences of "provider-sample" with formatted provider name
find . -type d \( -name "hack" -o -name ".git" \) -prune -o -type f -exec sed -i "s/provider-sample/$(echo "$providerName" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')/g" {} +
echo "Replaced provider-sample with $(echo "$providerName" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')"

# Replace ocurrences of "SampleProvider" with the provider name
find . -type d \( -name "hack" -o -name ".git" \) -prune -o -type f -exec sed -i "s/SampleProvider/$(echo "$providerName" | tr -d ' ')/g" {} +
echo "Replaced SampleProvider with $(echo "$providerName" | tr -d ' ')"

# Replace occurrences of "Provider Sample" with the provider name
find . -type d \( -name "hack" -o -name ".git" \) -prune -o -type f -exec sed -i "s/Provider Sample/$providerName/g" {} +
echo "Replaced Provider Sample with $providerName"