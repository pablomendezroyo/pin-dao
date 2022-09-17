FROM node:16.17.0-alpine3.16

# set working directory
WORKDIR /usr/src/app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
#COPY ../ui/package*.json ./
#RUN npm install react-scripts@5.0.1 -g --silent
#RUN npm install  npm@8.19.2 -g --silent

# add app
COPY ./ui ./

RUN npm install --silent

# start app
CMD ["npm", "start"]
