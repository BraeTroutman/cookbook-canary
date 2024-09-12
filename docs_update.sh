#!/bin/bash

printf "# Recipes\n\n" > docs/README.md

for FILE in recipes/***/*.cook
do
    TITLE=$(basename "${FILE}")
    TITLE=${TITLE%.cook}
    printf "## %s\n\n" "${TITLE}" >> docs/README.md
 done
