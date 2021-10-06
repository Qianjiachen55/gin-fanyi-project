#!/bin/bash

# check latest commit message

# shellcheck disable=SC2034
commit=`cat $1`
echo $commit
# shellcheck disable=SC1036
# shellcheck disable=SC1036
# shellcheck disable=SC1073
# shellcheck disable=SC1072
#msg_re="^(feat|fix)\([A-Za-z0-9_]*\)\:\s[A-Za-z0-9_]*"
msg_re="^(feat|fix|docs|style|refactor|perf|test|workflow|build|ci|chore|release|workflow)(\(.+\))?: .{1,100}"
if [[ $commit =~ $msg_re ]]
then
  echo "commit success"
else
  exit 1
fi




