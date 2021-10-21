FROM node:16

WORKDIR /app

COPY package.json package-lock.json ./

RUN rm -rf build

RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000

CMD ["node", "build/index.js"]