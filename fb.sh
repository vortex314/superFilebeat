set -x
rm -rf data/registry
./filebeat -e -c fb.yml
