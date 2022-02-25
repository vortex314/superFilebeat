set -v
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:/home/lieven/.local/bin:$GOPATH/bin
sudo apt-get install python3-pip
sudo apt-get install python git curl
sudo apt-get install python3-venv
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
go get github.com/elastic/beats
