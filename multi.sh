set -x
rm -rf data/registry
./filebeat  -c multi.yml # contained -e
