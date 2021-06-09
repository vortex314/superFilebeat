# Filebeat with grok, javascript , avro schema
## TL;DR
Unzip the platform specific filebeat and get started on MacOs, Windows, Linux with a filebeat on steroids : *javascript*, *grok* and *avro codec* !
![The Flow](https://github.com/vortex314/superFilebeat/blob/master/docs/Filebeat%2B%2B.png)
## Motivation 
Elasticsearch needs correctly parsed data to be the most useful. If you have to integrate logging data in a big enterprise it needs to be aligned to a minimal datamodel, for the same reasons that Elastic developed ESC datamodel
(start rant) Elastic however believes in a centralized architecture where all parsing is done by logstash pipelines. In a big enterprise with agile teams with a lot of technical flexibility, you are not able to master that centrally. So we went for a decentralized solution where the teams have the responsibility to deliver log data respecting this model. Logstash is sooooo heavy, you just cannot run it on the same server as your application. Filebeat is.(end rant)
Luckily filebeat is open source and can be 'easily' extended. 
## The not so lightweight shipper
[Example](https://github.com/vortex314/beats/tree/master/filebeat/fb.yml)

The filebeat version that does everything to enable local processing before pushing events.

It contains :
-	grok pattern matching to extract fields from text
-	timestamp parser to extract time , date in native format
-	Javascript engine to do everything you cannot do with grok
-	AVRO codec to send this in a regular schema to kafka
## Build instructions 
#### System requirements
Started from a fresh Kubuntu 20.04 install on a desktop pc
https://kubuntu.org/getkubuntu/
#### Install GO
https://golang.org/doc/install?download=go1.14.4.linux-amd64.tar.gz
```
cd ~/Downloads
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
cd ..
```
add to .profile in $HOME and logout/login
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:/home/lieven/.local/bin:$GOPATH/bin
```
#### Install tools
- python3.8
- git
- mage
```
sudo apt-get install python3-pip
sudo apt-get install python git curl
sudo apt-get install python3-venv
go install github.com/magefile/mage@latest
go get github.com/elastic/beats
```
Check python versions
https://linuxconfig.org/how-to-change-from-default-to-alternative-python-version-on-debian-linux
#### install Docker => only needed for crossBuild to windows, macOs
```
sudo apt-get remove docker docker-engine docker.io containerd runc
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker lieven # should be your own user id, logout/login to take effect
```

#### Build standard filebeat -- test your install

```
cd $HOME
mkdir workspace
cd workspace
git clone https://github.com/elastic/beats.git
git clone https://github.com/vortex314/superFilebeat
export BEATS=$HOME/workspace/beats
cd $BEATS/filebeat
mage build # delivers the Linux build
```
#### Extract version 7.12.1
```
export BEATS=$HOME/workspace/beats
cd $BEATS
git checkout tags/v7.12.1
git switch -c my-branch
```
#### Build cross-platform standard builds for filebeat 
```
export BEATS=$HOME/workspace/beats
cd $BEATS/filebeat
mage crossBuild #delivers other platforms via Docker
```

#### Build customized filebeat
##### Download sub-packages first and install into elastic vendor packages
```
go get -v github.com/robertkrimen/otto
go get -v gopkg.in/sourcemap.v1
go get -v github.com/vjeantet/grok
go get -v github.com/linkedin/goavro

cd $HOME/go/src/github.com/
cp -r robertkrimen/ $HOME/workspace/beats/vendor/github.com
cp -r vjeantet $HOME/workspace/beats/vendor/github.com
cp -r linkedin $HOME/workspace/beats/vendor/github.com/linkedin
cd ../gopkg.in/
cp -r sourcemap.v1/ $HOME/workspace/beats/vendor/gopkg.in
```
##### Build custom build by changes sources
```
export BEATS=$HOME/workspace/beats
cd $HOME/workspace/superFilebeat
mkdir $BEATS/libbeat/outputs/codec/avro
cp avro.go $BEATS/libbeat/outputs/codec/avro
cp event.go $BEATS/libbeat/outputs/codec/avro
cp grok.go $BEATS/libbeat/processors/actions
cp javascript.go $BEATS/libbeat/processors/actions
cp includes.go $BEATS/libbeat/publisher/includes/includes.go
cd $BEATS/filebeat
mage build 
cp filebeat $HOME/workspace/superFilebeat
cd $HOME/workspace/superFilebeat
zip fb.zip filebeat
```
#### Test the new build
```
cd $HOME/workspace/superFilebeat
unzip fb.zip
./fb.sh # look at output
```
#### collect the crossbuilds
```
export CROSSBUILD=$HOME//workspace/beats/filebeat/build/golang-crossbuild
zip darwin-amd64.zip $CROSSBUILD/filebeat-darwin-amd64
zip linux-386.zip  $CROSSBUILD/filebeat-linux-386
zip linux-amd64.zip    $CROSSBUILD/filebeat-linux-amd64
zip linux-arm64.zip $CROSSBUILD/filebeat-linux-arm64
zip windows-386.zip   $CROSSBUILD/filebeat-windows-386.exe
zip windows-amd64     $CROSSBUILD/filebeat-windows-amd64.exe

```


