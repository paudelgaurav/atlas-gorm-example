# gin-api-permissions
This is reference repo for my medium article : "Simple API level permission in gin"

1. Clone the repo ```git clone https://github.com/paudelgaurav/gin-api-permissions```
2. Install dependencies ```go mod download```
3. Populate the .env file using the .env.sample file
4. Run server: ```go run main.go```

---
- [CreateUser Controller](https://github.com/paudelgaurav/gin-api-permissions/blob/main/controller/user.go#L23)
- [BasicAuthPermission Middleware](https://github.com/paudelgaurav/gin-api-permissions/blob/main/middleware/auth.go#L12)
