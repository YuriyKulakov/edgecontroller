# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2020 Intel Corporation

FROM centos:7.6.1810 AS nfd-master

RUN yum install -y openssl sudo \
    && yum clean all

FROM nfd-master

ARG username=nfd-master
ARG user_dir=/home/$username

RUN useradd -d $user_dir -m -s /bin/bash $username
RUN groupadd sudo
RUN usermod -aG sudo $username
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER $username
WORKDIR $user_dir

COPY ./dist/nfd-master ./
COPY ./nfd-master/entrypoint.sh ./

CMD ["sudo", "-E", "./entrypoint.sh"]
