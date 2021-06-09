set -x
rm -rf ../data/registry
../filebeat -e -c filebeat.yml
