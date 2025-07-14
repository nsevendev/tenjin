FROM node:22.17.0-alpine AS base
WORKDIR /app
COPY package*.json ./
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

FROM base AS dev
WORKDIR /app
COPY . .
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["npm", "run", "dev"]

FROM base AS build
WORKDIR /app
COPY . .
RUN npm ci
RUN npm run build

# Base runtime commune
FROM node:22.17.0-alpine AS runtime-base
WORKDIR /app
RUN apk add --no-cache dumb-init
COPY --from=build /app/dist ./dist
COPY --from=build /app/server ./server
COPY --from=build /app/package*.json ./
RUN npm ci --production
USER node
ENTRYPOINT ["dumb-init", "--"]
CMD ["npm", "run", "serve"]

# Environnements spécifiques
FROM runtime-base AS prod
FROM runtime-base AS preprod
