#!/bin/bash


echo "try to kill McAfee..."

# for i in $(ps -ef | grep McA | grep -v grep | grep -v sh |awk '{print $2}')
# do 
#     sudo kill -9 $i
# done 


while true
do

	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')
	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')
	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')


	if [ -d "/usr/local/McAfee/" ];then
	    sudo mv /usr/local/McAfee /usr/local/McAfee_xx 
	fi

	if [ "$?" == "0" ];then
	    break
	fi

done

while true
do

	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')
	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')
	sudo kill -9 $(ps -ef | grep fee | awk '{print $2}')


	if [ -d "/Library/McAfee/" ];then
	    sudo mv /Library/McAfee/ /Library/McAfee_xx
	fi

	if [ "$?" == "0" ];then
	    echo "Success"
	    break
	fi

done

echo "=============Running================"
ps -ef | grep fee
echo "=============Moving================="
sudo mkdir -p /Library/LaunchDaemons.bak/
sudo mkdir -p /Library/LaunchAgents.bak/
sudo mv /Library/LaunchDaemons/*fee* /Library/LaunchDaemons.bak/
sudo mv /Library/LaunchAgents/*fee* /Library/LaunchAgents.bak/

echo "=============After Moving, ORIG====="
ls -al /Library/LaunchDaemons/*fee*
ls -al /Library/LaunchAgents/*fee*

echo "=============After Moving, BAK======"
ls -al /Library/LaunchDaemons.bak/*fee*
ls -al /Library/LaunchAgents.bak/*fee*

