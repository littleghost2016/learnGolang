#! /bin/sh

kill -9 $(pgrep webserver_learn)
cd ~/learnGoLang/
git pull https://github.com/littleghost2016/learnGoLang.git
cd webserver_learn/
chmod +x webserver_learn
./webserver_learn &