set -x
rm -f data/registry
./filebeat version
./filebeat -e -c filebeat.yml # -d "*" # extra debug log option
