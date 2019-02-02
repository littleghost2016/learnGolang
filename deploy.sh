#! /bin/sh

kill -9 $(pgrep webserver_learn)
cd ~/learnGoLang/
git reset --hard
git pull https://github.com/littleghost2016/learnGoLang.git
cd webserver_learn/
chmod +x webserver_learn
./webserver_learn &