FROM browserless/chrome:latest
USER root

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - 
RUN echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list
RUN apt-get update -y && apt-get install -y dos2unix

RUN npm install -g lighthouse-batch 
RUN npm install -g json2csv

RUN mkdir lhbatch && cd lhbatch
COPY spider.sh /lhbatch/spider.sh
RUN dos2unix -u /lhbatch/spider.sh && chmod u+x /lhbatch/spider.sh

ENTRYPOINT "/lhbatch/spider.sh"