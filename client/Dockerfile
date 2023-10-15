FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH - allow running 'vite'
ENV PATH /app/node_modules/.bin:$PATH

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install -d tailwindcss@latest postcss@latest autoprefixer@latest

RUN npm install axios vue-router@4

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

CMD [ "npm", "run", "dev" ]