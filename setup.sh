#! /bin/bash

# Updating the available updates
sudo apt update
sudo apt upgrade -y
sudo apt autoremove -y

# Installing Go
mkdir Downloads
cd ./Downloads
curl -O https://dl.google.com/go/go1.18.3.linux-amd64.tar.gz
sudo tar -xvf ./go1.18.3.linux-amd64.tar.gz -C /usr/local

# Coming back to home
cd
echo "export GOPATH=$HOME/go" >> ~/.profile
echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.profile
source ~/.profile

# Checking Go version
go version


# Making the GOPATH
cd
mkdir go 
cd ./go/
mkdir src 
mkdir bin 
mkdir pkg

# Downloading the project
cd ./src/
git clone https://github.com/MohammadArik/gRPC-EC2_IP_Test.git
cd ./gRPC-EC2_IP_Test

echo "It's all yours now!"

# Making the network info file
cd
mkdir info
cd ./info/
printf "{ \n\t\"pub_ip\": \"\", \n\t\"pr_ip\": \"\", \n\t\"handlerAddress\": \"\" \n}" >> network.json
nano ./network.json