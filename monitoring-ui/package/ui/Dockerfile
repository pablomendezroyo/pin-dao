#FROM node:18.9.0-alpine3.15
FROM node:13.12.0-alpine

# set working directory
WORKDIR /usr/src/app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package*.json ./
RUN npm install --silent
RUN npm install react-scripts@5.0.1 -g --silent
#RUN npm install  npm@8.19.2 -g --silent

# add app
COPY . ./

# start app
CMD ["npm", "start"]
