FROM node:16.17.0-alpine3.16

# set working directory
WORKDIR /usr/src/app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# add app
COPY ./ui ./

RUN npm install --silent

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

# start app
#ENTRYPOINT [ "entrypoint.sh" ]
CMD ["npm", "start"]