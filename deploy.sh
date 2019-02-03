#! /bin/sh

kill $(pgrep webserver_learn)
cd ~/learnGoLang/
# git reset --hard
chmod -x webserver_learn/webserver_learn
git pull https://github.com/littleghost2016/learnGoLang.git
chmod +x webserver_learn/webserver_learn
webserver_learn/webserver_learn &