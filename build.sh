set -x
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:/home/lieven/.local/bin:$GOPATH/bin
export BEATS=$HOME/workspace/beats

cd ~
cd workspace
git clone https://github.com/elastic/beats.git
git clone https://github.com/vortex314/superFilebeat
cd $BEATS/filebeat
mage build # delivers the Linux build
