# ---------- Base ----------
FROM node:22.17.0-slim AS base
RUN apt-get update && apt-get install -y bash
WORKDIR /app
RUN apt-get update && apt-get install -y \
    wget ca-certificates fonts-liberation libasound2 libatk1.0-0 libatk-bridge2.0-0 \
    libc6 libcairo2 libcups2 libdbus-1-3 libdrm2 libexpat1 libgbm1 libglib2.0-0 \
    libgtk-3-0 libnspr4 libnss3 libpango-1.0-0 libx11-6 libx11-xcb1 libxcb1 \
    libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 \
    libxrender1 libxss1 libxtst6 xdg-utils chromium \
 && rm -rf /var/lib/apt/lists/*

RUN npm install -g @angular/cli
ENV CHROME_BIN=/usr/bin/chromium
COPY package*.json ./
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ARG APP_NAME=app
ENV APP_NAME=${APP_NAME}

# ---------- Dev ----------
FROM base AS dev
COPY . .
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["npm", "start"]

# ---------- Build ----------
FROM base AS build
COPY . .
RUN npm ci
RUN test -f src/environments/environment.ts || cp src/environments/environment.dist src/environments/environment.ts
RUN npm run build

# ---------- Runtime SPA avec Nginx ----------
FROM nginx:alpine AS runtime-base
WORKDIR /usr/share/nginx/html
# Copier les fichiers buildés (SPA)
ARG APP_NAME=app
COPY --from=build /app/dist/${APP_NAME}/browser ./
# Configuration Nginx pour SPA
COPY nginx.conf /etc/nginx/nginx.conf
# Sécurité
RUN chown -R nginx:nginx /usr/share/nginx/html
USER nginx
CMD ["nginx", "-g", "daemon off;"]

# ---------- Environnements spécifiques ----------
FROM runtime-base AS prod
FROM runtime-base AS preprod
