FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build
RUN http-server .

CMD ["npx", "serve", "dist"]

EXPOSE 80

