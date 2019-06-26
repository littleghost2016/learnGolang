#! /bin/sh

kill $(pgrep webserver_learn)
cd ~/learnGoLang/
# git reset --hard
chmod -x webserver/webserver
git pull https://github.com/littleghost2016/learnGoLang.git
chmod +x webserver/webserver
webserver/webserver &