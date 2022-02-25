# 71.6.3 upgrade

```
sed -i 's;http://archive.debian.org/debian;http://deb.debian.org/debian;' /etc/apt/sources.list
apt update
```

## Upgrade go
As super user : 
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
```
$ go mod tidy 


debconf: delaying package configuration, since apt-utils is not installed

```
cd $BEATS
mkdir save
mv go.mod  save
git checkout v7.16.3
mv go.sum save
git checkout v7.16.3
export BEATS=$HOME/workspace/beats
cd $BEATS/filebeat
mage build # delivers the Linux build
mage crossBuild
export CROSSBUILD=$HOME/workspace/beats/filebeat/build/golang-crossbuild
cd $CROSSBUILD
```
