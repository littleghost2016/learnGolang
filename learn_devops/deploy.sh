#! /bin/sh

kill $(pgrep webserver_learn)
cd ~/learnGolang/learn_devops
# git reset --hard
chmod -x webserver/webserver
git pull https://github.com/littleghost2016/learnGoLang.git
chmod +x webserver/webserver
webserver/webserver &