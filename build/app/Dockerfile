# -------------->
FROM node:16.13-alpine as installer

USER node

WORKDIR /home/node

COPY --chown=node:node package*.json ./

RUN npm ci

COPY --chown=node:node src/shared/infra/prisma/schema.prisma ./prisma/

RUN npx prisma generate

# -------------->
FROM node:16.13-alpine as builder

USER node

WORKDIR /home/node

COPY --chown=node:node --from=installer /home/node/node_modules ./node_modules
COPY --chown=node:node . .

RUN nest build

# -------------->
FROM node:16.13-alpine as runner

RUN apk add dumb-init

ENV NODE_ENV 'production'

USER node

WORKDIR /home/node

COPY --chown=node:node --from=installer /home/node/node_modules ./node_modules
COPY --chown=node:node --from=builder /home/node/dist ./dist
COPY --chown=node:node src/shared/infra/prisma/migrations ./migrations
COPY --chown=node:node src/shared/infra/prisma/schema.prisma ./
COPY --chown=node:node package.json ./
COPY --chown=node:node wait-for.sh ./

CMD ["dumb-init", "node", "dist/main/server.js"]
