# cognito-go

JWT Authentication with AWS Cognito in Go + Gin Middleware

# Usage

## Single

```
import "github.com/leemeichin/cognito-go"

c, _ := cognito.NewCognito("ap-southeast-2", "cognito-app", "client-id")
token, err := c.VerifyToken("abc")
```

## Gin Middleware

```
import (
  "github.com/leemeichin/cognito-go"
  "github.com/gin-gonic/gin"
)

c, _ := cognito.NewCognito("ap-southeast-2", "cognito-app", "client-id")
r := gin.New()
r.GET("/protected", c.Authorize(), protectedEndpoint)
```
