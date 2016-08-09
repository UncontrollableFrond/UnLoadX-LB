#builds the apiserver image referenced in docker-compose.yml

FROM node:6.3.1
Maintainer UnloadX <aebrownz11@gmail.com>

#replace sh with bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# Install project dependencies
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app

# Install Node and project deps.
RUN npm install -g nodemon \
  && npm install -g gulp \
  && npm install --dev \
  && npm install \
  && npm install -g mocha

# Expose port 80 (http)
EXPOSE 80
