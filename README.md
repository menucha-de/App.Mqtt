# App.Mqtt
## Manual upgrade
    NAME=mqtt
    ctr -n system i pull -k ghcr.io/peramic/$NAME:latest
    systemctl stop $NAME
    ctr -n system c rm $NAME
    ctr -n system c create --label NAME=$NAME --label IS_ACTIVE=true --label org.opencontainers.image.exposedPorts='{"1883/tcp":{},"8883/tcp":{},"9001/tcp":{}}' --env LOGHOST=$NAME --with-ns network:/var/run/netns/$NAME ghcr.io/peramic/$NAME:latest $NAME
    systemctl start $NAME
