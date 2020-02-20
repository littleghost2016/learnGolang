#! /bin/sh

kill $(pgrep webserver)
cd ~/learnGolang/devops
# git reset --hard
chmod -x webserver/webserver
git pull https://github.com/littleghost2016/learnGoLang.git
chmod +x webserver/webserver
webserver/webserver &