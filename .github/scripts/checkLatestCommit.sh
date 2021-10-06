#!/bin/bash

# shellcheck disable=SC2034
commit=$1

# shellcheck disable=SC2034
contains="feat"

if [[ $(commit) == *$(contains)* ]]
then
  exit 1
fi
