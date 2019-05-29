#!/usr/bin/env bash

DIST_HOST=127.0.0.1
DIST_USERNAME=root
DIST_PASSWORD=ebUM752FVCgmXD736hnmfEFC352Y
DIST_DATABASE=wdcj
SAVE_PATH=/tmp/mnt/

for table in $(cat 3.txt); do
    echo ${table}
    mysql -h${DIST_HOST} -u${DIST_USERNAME} -p${DIST_PASSWORD} \
        ${DIST_DATABASE} < ${SAVE_PATH}${table}.sql
    echo next
done

echo 'done'