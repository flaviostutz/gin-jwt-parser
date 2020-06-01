# jwt-parse-middleware
Golang Gin middleware lib for parsing and checking JWT tokens in http requests

It parses and validates any incoming JWT token from requests and if a valid token is found, sets token claims to gin.Context attribute, accessible by ctx.Get("<claimname>")

Check a full example at https://github.com/stutzlab/userme-demo-api/blob/master/api.go

## Usage

* Add module dependency

```sh
go get github.com/flaviostutz/gin-jwt-parser
```

* Configure Gin routes

```golang
import (
    jwtparser "github.com/flaviostutz/gin-jwt-parser"
)

func NewHTTPServer() *HTTPServer {
    router := gin.Default()

    router.Use(jwtparser.Middleware(jwtparser.Config{
        RequiredIssuer:   "Berimbal",
        RequiredType:     "access",
        FromBearer:       "Authorization",
        JWTSigningMethod: "ES256",
        JWTVerifyKeyFile: "/my-public-key",
    }))
}
```

* Check additional token data in specific API implementation

```golang
func listSomething() func(*gin.Context) {
    return func(c *gin.Context) {
      scope, _ := c.Get("scope")
      if scope != "admin" {
        return fmt.Errorf("User %s not authorized to access admin resource", sub)
      }

      sub, _ := c.Get("sub")
      logrus.Infof("User %s is listing items", sub)

      c.JSON(200, gin.H{})
      return
    }
}
```

* In this example, only JWT tokens coming from HTTP Header "Authorization", with claim "iss==Berimbal", "typ==access" and whose signature was checked against "/my-public-key" are accepted.

* After validating the JWT itself, if sets all claims as accessible properties from gin.Context

## API

* Check https://pkg.go.dev/github.com/flaviostutz/gin-jwt-parser

