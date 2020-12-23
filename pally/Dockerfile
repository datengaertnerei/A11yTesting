FROM node:12-alpine

RUN apk add git python3

RUN git clone https://github.com/pa11y/pa11y-dashboard.git && cd pa11y-dashboard && npm i

COPY pa11y.config /pa11y-dashboard/config/production.json

ENTRYPOINT sh -c "cd /pa11y-dashboard && node index.js"