FROM golang:latest
MAINTAINER treeui.old@gmail.com
COPY ../gear /data/gear
CMD [ "go","env" ]
ENTRYPOINT ["/data/gear/run.sh"]