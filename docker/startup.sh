set -x
cd /filebeat
pwd
rm -f data/registry
/filebeat/filebeat version
./filebeat -e -c filebeat.yml # -d "*" # extra debug log option
