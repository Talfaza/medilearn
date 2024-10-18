FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run serve 

CMD ["npx", "serve", "dist"]

EXPOSE 80

