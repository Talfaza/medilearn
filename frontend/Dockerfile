FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm install -g serve 

RUN npm run build

CMD ["npx", "serve", "-l", "80", "dist"]
EXPOSE 80

