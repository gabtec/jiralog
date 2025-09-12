#!/usr/bin/env bash
version="v0.1.0"

# if [ $1 != '' ];then
#     LATEST_TAG=$1
# else
#     LATEST_TAG=$(git tag --sort=-v:refname | head -n 1)
# fi

LATEST_TAG=$(cat constants/constants.go| grep Version | cut -d "=" -f2 | tr -d '"' | xargs)

echo "Releasing jiralog-${LATEST_TAG} locally..."

go build -o build/jiralog-${LATEST_TAG} main.go

sudo cp ./build/jiralog-${LATEST_TAG} /usr/local/bin/jiralog

echo "Done."

exit 0