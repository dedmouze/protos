# Proto for SSO

```Go
/// Auth:
* Register(email, password string) (user_id int64)
* RegisterApp(name string) (api_key string)
* Login(app_id, email, password string) (token string)

/// UserInfo:
* User(email, api_key string) (user_id int64, email string, created_at, visited_at *timestamppb.Timestamp)
* Admin(email, api_key string) (admin_id int64, email string, level int32)

/// Permission
* AddAdmin(email, api_key string)
* DeleteAdmin(email, api_key string)

/// timestamppb.Timestamp
struct Timestamp {
    Seconds int64
    Nanos int32
}
```