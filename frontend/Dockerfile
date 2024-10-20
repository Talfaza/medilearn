FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm install -g serve 

RUN npm run build

CMD ["serve", "-s", "dist"] 

EXPOSE 80

