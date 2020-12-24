FROM node:12-alpine

# Installs latest Chromium package.
RUN apk add --no-cache \
	  wget \
      chromium \
      nss \
      freetype \
      freetype-dev \
      harfbuzz \
      ca-certificates \
      ttf-freefont \
      nodejs \
      yarn

RUN mkdir lhbatch && cd lhbatch && npm install lighthouse-batch -g && npm install -g json2csv

COPY spider.sh /lhbatch/spider.sh

RUN dos2unix -u /lhbatch/spider.sh && chmod u+x /lhbatch/spider.sh

ENTRYPOINT "/lhbatch/spider.sh"