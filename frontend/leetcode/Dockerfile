FROM node:16

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

EXPOSE 3000
EXPOSE 15000

CMD ["npm", "run", "dev", "--", "--open", "--host", "0.0.0.0"]