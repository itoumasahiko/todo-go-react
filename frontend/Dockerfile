FROM node:20

WORKDIR /app
COPY . .
RUN npm install
ENV CHOKIDAR_USEPOLLING=true
CMD ["npm", "start"]
