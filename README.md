# Pinger-logger

Простой скрипт, с помощью которого можно проверять работоспособность веб-приложения. 

```
> go run pinger-logger/main.go --log=./log.log --oStatus=./status --address=localhost:8888/status
> cat status
SUCCESS 2024/04/04 08:06:31
> cat log.log
2024/04/04 20:00:09 error during pinging localhost:8888/status: HTTP status 500
2024/04/04 20:01:30 error during pinging localhost:8888/status: HTTP status 500
2024/04/04 20:02:15 error during pinging localhost:8888/status: HTTP status 500
2024/04/04 20:02:54 error during pinging localhost:8888/status: HTTP status 500
2024/04/04 20:04:54 error during pinging localhost:8888/status: HTTP status 500
2024/04/04 20:08:13 Get "http://localhost:8888/status": dial tcp 127.0.0.1:8888: connect: connection refused
```

Принимает на вход три флага: путь до логгера, путь до файла с последним статусом и адрес для пингования.

В репе лежат два systemd файла и один таймер:

`webapp.service`, запускающий вебсервер после старта системы и сети, а так же `pinger-logger.service`, запускающий 
гошный скрипт `pinger-logger/main.go`. Таймер запускает его каждые 5 минут. Надо бы бинарник скрипта положить в `/opt/pinger-logger/`.

Чтобы их активировать надо поместить их в /etc/systemd/system/ и активировать и запустить

```
cd webapp && go build . && cp webapp /opt/webapp/webapp.bin && cd ..
cd pinger-logger && go build . && cp pinger-logger /opt/pinger-logger/pinger-logger && cd ..

cp pinger-logger.service /etc/systemd/system/
cp pinger-logger.timer /etc/systemd/system/
cp webapp.service /etc/systemd/system/

systemctl enable webapp.service
systemctl enable pinger-logger.timer
systemctl start webapp.service
systemctl start pinger-logger.timer
```