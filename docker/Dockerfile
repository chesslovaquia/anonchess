FROM debian:trixie-20240812-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="240827"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

ENV DEBIAN_FRONTEND noninteractive

ENV APT_INSTALL bash openssl ca-certificates build-essential rustup nodejs npm

RUN apt-get clean \
	&& apt-get update -yy \
	&& apt-get dist-upgrade -yy --purge \
	&& apt-get install -yy --no-install-recommends ${APT_INSTALL} \
	&& apt-get clean \
	&& apt-get autoremove -yy --purge \
	&& rm -rf /var/lib/apt/lists/* \
		/var/cache/apt/archives/*.deb \
		/var/cache/apt/*cache.bin

#
# devel user
#
ARG DEVEL_UID=1000
ARG DEVEL_GID=1000

ENV DEVEL_UID ${DEVEL_UID}
ENV DEVEL_GID ${DEVEL_GID}

RUN groupadd -o -g ${DEVEL_GID} devel \
	&& useradd -o -d /home/devel -m -c 'devel' -g ${DEVEL_GID} -u ${DEVEL_UID} devel \
	&& chmod -v 0750 /home/devel

RUN printf 'umask %s\n' '027' >>/home/devel/.profile
RUN printf "export PS1='%s '\n" '\u@\h:\W\$' >>/home/devel/.profile

#
# rust setup
#
ADD ./build-deps.sh /var/tmp/build-deps.sh
RUN chmod -v 555 /var/tmp/build-deps.sh

#
# devel user
#
USER devel:devel
WORKDIR /home/devel

ENV USER devel
ENV HOME /home/devel

#
# rust setup
#
ENV RUST_VERSION 1.80.1

RUN rustup --version && \
	rustup default ${RUST_VERSION}

RUN rustc --version && \
	cargo --version

RUN /var/tmp/build-deps.sh

RUN printf 'export PATH=%s\n' '/usr/local/bin:/usr/bin:/home/devel/.cargo/bin' >>/home/devel/.profile

ENV RUST_BACKTRACE=1

CMD /bin/bash -i -l
