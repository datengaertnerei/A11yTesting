FROM node:12-alpine

ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD=true \
    PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium-browser

# Installs latest Chromium package.
RUN apk add --no-cache \
	  git \
	  python3 \
      chromium \
      nss \
      freetype \
      freetype-dev \
      harfbuzz \
      ca-certificates \
      ttf-freefont \
      nodejs \
      yarn

RUN git clone https://github.com/pa11y/pa11y-dashboard.git && cd pa11y-dashboard && npm i

COPY pa11y.config /pa11y-dashboard/config/production.json

ENTRYPOINT sh -c "cd /pa11y-dashboard && node index.js"