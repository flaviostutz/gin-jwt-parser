package jwtparser

"github.com/gin-gonic/gin"

//Config configuration properties for JWT Parser
type Config struct {
}

//Middleware Analyses http request, parse existing JWT tokens and set object "jwt" to gin context according to configuration.Middleware
//The jwt token claims can be later checked by request handlers with "c.GetString(...)"
func Middleware(config Config) gin.HandlerFunc {
}
