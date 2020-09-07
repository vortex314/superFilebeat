FROM irm2.docker.ing.net/ing-rhel8-openjdk11

ARG BUILD_NAME
ARG BUILD_VERSION

LABEL description=$BUILD_NAME-$BUILD_VERSION

RUN mkdir -p /filebeat

RUN chgrp -R 0 /filebeat && \
    chmod -R g=u /filebeat

WORKDIR /filebeat

COPY ./log.txt /filebeat/log.txt
COPY ./filebeat /filebeat/filebeat
COPY ./filebeat.yml /filebeat/filebeat.yml
COPY ./filebeat.js /filebeat/filebeat.js
COPY ./filebeat.avsc /filebeat/filebeat.avsc
COPY ./startup.sh /filebeat/startup.sh


ENTRYPOINT ["sh", "/filebeat/startup.sh"]

