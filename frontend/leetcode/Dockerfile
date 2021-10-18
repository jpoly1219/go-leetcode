FROM node:latest

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install

COPY . .

EXPOSE 3000

EXPOSE 24678

CMD ["npx", "svelte-kit", "dev", "--host", "0.0.0.0"]