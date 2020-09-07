set -x
cd /filebeat
pwd
ls -l
ls -l ..
rm -f data/registry
/filebeat/filebeat version
./filebeat -e -c filebeat.yml # -d "*" # extra debug log option
