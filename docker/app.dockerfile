# ---------- Base ----------
FROM node:22.17.0-slim AS base
WORKDIR /app
RUN apt-get update && apt-get install -y \
    python3 \
    make \
    g++ \
    && rm -rf /var/lib/apt/lists/*
RUN npm install -g @angular/cli

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
