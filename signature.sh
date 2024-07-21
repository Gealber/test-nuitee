#! /bin/bash

apikey="db11033c50b5ed53ab7b815cb1b2eaee"
secret="704773c0e3"
d=$(date +%s)

echo "computing for ts: ${d}"
# echo -n ${apiKey}${secret}${d}|sha256sum|awk '{ print $1}'
echo -n "${apiKey}${secret}${d}" |sha256sum | awk '{ print $1}'
