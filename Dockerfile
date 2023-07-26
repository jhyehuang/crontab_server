FROM debian:bullseye

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

COPY ./letsfil-job /usr/bin/letsfil-job
ENTRYPOINT [ "/usr/bin/letsfil-job" ]
