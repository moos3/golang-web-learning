# Golang Web Starter Basic Edition

## Golang Starter project with Vue.js single page client

This is based off the Go Vue starter project by Mark Chenoweth, But I have changed the logic a great deal building off it.

### Features:
- Middleware: [Negroni](https://github.com/urfave/negroni)

- Router: [Gorilla](https://github.com/gorilla/mux)

- Orm: [Gorm](https://github.com/jinzhu/gorm) (sqlite or postgres)

- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)

- [Vue.js](https://vuejs.org/) spa client with webpack

- User management
  - Change user password

- Membership management

### TODO:
- email confirmation

- User Profiles (Needs profile editor and viewing finished)

- letsencrypt tls

### To get started:

``` bash
# clone repository
go get github.com/moos3/golang-web-learning
cd $GOPATH/src/github.com/moos3/golang-web-learning

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run server.go

# open a new terminal and change to the client dir
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

### License

MIT License  - see LICENSE for more details