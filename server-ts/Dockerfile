FROM node:18-slim as build
WORKDIR /app
COPY package*.json tsconfig.json ./
RUN npm install
COPY src ./src
RUN npm run build

FROM node:18-alpine 
WORKDIR /app
COPY package*.json ./
COPY --from=build package*.json ./
RUN npm install --omit=dev
COPY --from=build /app/dist ./dist

ENTRYPOINT ["npm", "run", "start"]