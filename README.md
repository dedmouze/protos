# Proto for SSO

```go
/// Auth:
* Register(email, password string) (user_id int64)
* Login(app_id, email, password string) (token string)

/// UserInfo:
* User(email, token string) (user_id int64, email, password string)
* Admin(email, token string) (admin_id int64, email string, level int32)

/// Permission
* AddAdmin(email, token string)
* DeleteAdmin(email, token string)
```