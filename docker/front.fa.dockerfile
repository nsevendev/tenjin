# ---------- Base ----------
FROM node:22.17.0-alpine AS base
WORKDIR /app
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
EXPOSE 3000
CMD ["npm", "run", "dev:ssr"]

# ---------- Build ----------
FROM base AS build
COPY . .
RUN npm ci
# Build SSR (adapter à tes scripts si différents)
# - Option 1 (CLI récent) :
# RUN npm run build -- --configuration=production --ssr
# - Option 2 (scripts générés par le CLI) :
RUN npm run build:ssr
# Si tu utilises le prerender :
# RUN npm run prerender

# ---------- Runtime commun ----------
FROM node:22.17.0-alpine AS runtime-base
WORKDIR /app
RUN apk add --no-cache dumb-init
# Copie des artefacts de build
# (browser + server pour SSR/hydratation)
COPY --from=build /app/dist ./dist
# On garde package.json pour installer les deps nécessaires au server Node
COPY --from=build /app/package*.json ./
# Déps prod uniquement (le serveur SSR a besoin de @angular/ssr + express & cie)
RUN npm ci --omit=dev
# Sécurité minimale
USER node
# Démarrage Node (SSR)
# Exemple de point d'entrée standard en Angular 17/18+ :
# node dist/<APP_NAME>/server/server.mjs
ENV NODE_ENV=production
ENTRYPOINT ["dumb-init", "--"]
CMD ["node", "dist/${APP_NAME}/server/server.mjs"]

# ---------- Environnements spécifiques ----------
FROM runtime-base AS prod
FROM runtime-base AS preprod
