#!/usr/bin/env bash

DIST_HOST=127.0.0.1
DIST_PASSWORD=root
DIST_USERNAME=deliang
DIST_DATABASE=crm
SAVE_PATH=/data/mnt/
mkdir -p ${SAVE_PATH}

mysql -h${DIST_HOST} -u${DIST_USERNAME} -p${DIST_PASSWORD} ${DIST_DATABASE} \
    -e "show tables" | sed -e '/^Tables_in_.*$/d' > 1.txt

for table in $(cat 1.txt); do
    echo ${table}
    mysqldump -h${DIST_HOST} -u${DIST_USERNAME} -p${DIST_PASSWORD} \
        ${DIST_DATABASE} ${table} > ${SAVE_PATH}${table}.sql
    echo next
done

echo 'done!!!'