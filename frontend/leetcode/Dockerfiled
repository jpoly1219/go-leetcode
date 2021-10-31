FROM node:16

WORKDIR /frontend

COPY package.json package-lock.json ./

RUN npm install

COPY . .

EXPOSE 3000
EXPOSE 15000

CMD ["npx", "svelte-kit", "dev", "--host", "0.0.0.0"]