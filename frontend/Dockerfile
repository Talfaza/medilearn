# Use the official Node.js image
FROM node:16-alpine

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the application
RUN npm run build

# Serve the application
CMD ["npx", "serve", "dist"]

# Expose the port the app runs on
EXPOSE 80

