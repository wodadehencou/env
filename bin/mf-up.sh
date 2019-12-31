#!/bin/bash

echo "reset McAfee"
sudo mv /usr/local/McAfee_xx /usr/local/McAfee
sudo mv /Library/McAfee_xx /Library/McAfee


echo "=============Moving================="
sudo mv /Library/LaunchDaemons.bak/* /Library/LaunchDaemons/
sudo mv /Library/LaunchAgents.bak/* /Library/LaunchAgents/

echo "=============After Moving, ORIG====="
ls -al /Library/LaunchDaemons/*fee*
ls -al /Library/LaunchAgents/*fee*

echo "=============After Moving, BAK======"
ls -al /Library/LaunchDaemons.bak/*fee*
ls -al /Library/LaunchAgents.bak/*fee*

