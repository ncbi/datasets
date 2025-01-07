#!/bin/bash

VERSION=7.5.0
CLI=openapi-generator-cli-${VERSION}.jar

if [[ ! -e ${CLI} ]]; then
    curl "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${VERSION}/${CLI}" > ${CLI}
fi

for template in go; do
    PATCH_FILES=$(find ${template} -type f)
    rm -rf a b
    mkdir -p a b
    ( cd a; echo ${PATCH_FILES} | xargs -r unzip ../${CLI} )
    ln -s ../${template} b/${template}
    diff -Naur a b > ${template}.patch
    rm -rf a b
done
