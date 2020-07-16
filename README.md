# Filebeat with grok, javascript , avro schema
## The not so lightweight shipper
[Example](https://github.com/vortex314/beats/tree/master/filebeat/fb.yml)

The filebeat version that does everything to enable local processing before pushing events.

It contains :
-	grok pattern matching to extract fields from text
-	timestamp parser to extract time , date in native format
-	Javascript engine to do everything you cannot do with grok
-	AVRO codec to send this in a regular schema to kafka
![The Flow](https://github.com/vortex314/superFilebeat/blob/master/docs/Filebeat%2B%2B.png)
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
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
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
#### Extract version 7.8
```
export BEATS=$HOME/workspace/beats
cd $BEATS
git checkout tags/v7.8.0 my-branch
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
export CROSSBUILD=$HOME//workspace/beats/filebeat/build/golang-crossbuild
zip darwin-amd64.zip $CROSSBUILD/filebeat-darwin-amd64
zip linux-386.zip  $CROSSBUILD/filebeat-linux-386
zip linux-amd64.zip    $CROSSBUILD/filebeat-linux-amd64
zip linux-arm64.zip $CROSSBUILD/filebeat-linux-arm64
zip windows-386.zip   $CROSSBUILD/filebeat-windows-386.exe
zip windows-amd64     $CROSSBUILD/filebeat-windows-amd64.exe



