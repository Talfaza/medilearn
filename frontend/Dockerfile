FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .
RUN npm install -g http-server
RUN npm run build
RUN cd dist ;http-server .

CMD ["npx", "serve", "dist"]

EXPOSE 80

