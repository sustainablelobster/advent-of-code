#!/usr/bin/env bash

if ! which standardrb 2> /dev/null; then
    gem install standardrb
fi

git diff --name-only --cached |
    xargs ls -1 2>/dev/null |
    grep '\.rb$' |
    xargs standardrb --fix
