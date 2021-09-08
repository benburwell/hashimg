#!/bin/sh

while read user
do
  ./hashimg -username "$user" | tee "$user.b64" | base64 -D > "$user.png"
done
