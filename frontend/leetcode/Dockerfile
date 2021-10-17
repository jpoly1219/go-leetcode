FROM node:16

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install --produiction

COPY . .

EXPOSE 3000

EXPOSE 24678

CMD ["npm", "run", "dev", "--host", "0.0.0.0"]