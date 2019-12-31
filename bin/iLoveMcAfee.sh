#!/bin/bash

if [ "$1" = "reset" ];then
    echo "reset McAfee"
    sudo mv /usr/local/McAfee_xx /usr/local/McAfee
    sudo mv /Library/McAfee_xx /Library/McAfee

    exit
fi

while true
do
    echo "try to kill McAfee..."

    for i in $(ps -ef | grep McA | grep -v grep | grep -v sh |awk '{print $2}')
    do 
        sudo kill -9 $i
    done 

    if [ -d "/usr/local/McAfee/" ];then
        sudo mv /usr/local/McAfee/ /usr/local/McAfee_xx 
    fi

    if [ "$?" != "0" ];then
        continue
    fi

    if [ -d "/Library/McAfee/" ];then
        sudo mv /Library/McAfee/ /Library/McAfee_xx
    fi

    if [ "$?" == "0" ];then
        echo "Success"
        break
    fi

done


