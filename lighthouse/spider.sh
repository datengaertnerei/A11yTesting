#!/bin/sh

DEPTH=5
OUTPUT="./urls.txt"

apt-get --only-upgrade install google-chrome-stable > /dev/null

cd /lhbatch

wget -r --spider --delete-after --force-html -l $DEPTH $SITE 2>&1 \
    | grep '^--' | awk '{ print $3 }' | grep -v '\.\(css\|js\|svg\|ttf\|txt\|png\|gif\|xml\|json\|ico\|otf\|woff\|woff2\|jpg\)$' \
    | grep -v -F '?' | sort | uniq > $OUTPUT

lighthouse-batch -f $OUTPUT --no-report > lh.log 2>&1
json2csv -i report/lighthouse/summary.json -o report/lighthouse/summary.csv --flatten-objects
cp report/lighthouse/summary.csv /dev/stdout