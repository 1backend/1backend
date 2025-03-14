FROM node:22.4.1 as build

# # Install dependencies
# RUN apt-get update && apt-get install -y wget gnupg
# RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add -
# RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list'
# RUN apt-get update && apt-get install -y google-chrome-stable
#
# # Set environment variable for Chrome
# ENV CHROME_BIN=/usr/bin/google-chrome

WORKDIR /app
COPY . .

ENV CHROME_NO_SANDBOX=true

RUN npm i && \
    cd workspaces/angular-app && \
    npm run lint && \
    # npm install -g @angular/cli@latest && \
    # ng test && \
    cd ../.. && \
    npm run package:angular-app

FROM nginx:alpine

COPY --from=build /app/workspaces/angular-app/.dist/angular-app/browser /usr/share/nginx/html

COPY ./docker/entrypoint.sh /entrypoint.sh
COPY ./docker/nginx.conf /etc/nginx/nginx.conf

RUN if [ ! -f /entrypoint.sh ]; then \
    echo "ERROR: entrypoint.sh not found"; \
    exit 1; \
    fi

RUN chmod +x /entrypoint.sh

EXPOSE 80

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
