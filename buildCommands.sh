#!/usr/bin/env bash

i=0

docker_apps="fdk-harvest-admin"

for docker_app in $docker_apps; do
    i=$((i+1))
    BUILD_APPS[$i]=$docker_app
    BUILD_CMD[$i]="( ./buildDocker.sh )"
done

export BUILD_APPS
export BUILD_CMD