FROM node:18.18.2-alpine

RUN apk add --no-cache g++ make py3-pip libc6-compat
WORKDIR /app
COPY package.json ./package.json

ENV NODE_ENV=production
# RUN npm ci

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
USER nextjs

COPY --chown=nextjs:nodejs .next ./.next
COPY node_modules ./node_modules
COPY package.json ./package.json
COPY public ./public
EXPOSE 3000
CMD npm start
