FROM node:12 AS build

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm i

COPY tsconfig.json .
COPY spec ./spec
COPY config ./config
COPY test ./test
COPY src ./src

RUN npm run build
RUN npm t
#we will copy node modules later, so they need to be pruned first
RUN npm prune --production



#FINAL CONTAINER
FROM node:12-alpine

ENV TZ=Europe/Oslo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/package.json ./
COPY --from=build /usr/src/app/config ./config
COPY --from=build /usr/src/app/spec ./spec
COPY --from=build /usr/src/app/dist/src ./dist/src
COPY --from=build /usr/src/app/node_modules ./node_modules

CMD ["npm", "run", "prod"]
