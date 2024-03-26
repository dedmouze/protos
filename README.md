# Proto for SSO

```Go
/// Auth:
* Register(email, password string) (user_id int64)
* RegisterApp(name string) (api_key, user_key string)
* Login(app_id, email, password string) (token string)

/// UserInfo:
* User() (user User)
* UserByEmail() (email string) (user User)
* UserByID(user_id int64) (user User)
* UserRoles(user_id int64) (roles []Role)
* UserRolesByApp(user_id, app_id int64) (roles []Role)

/// AppInfo:
* App() (app App)
* AppByID(appID int64) (app App)
* AppRoles(appID int64) (role Role)
* AppRolesByApp(targetAppID, appID int64) (roles []Role)

/// Role
* CreateRole(name, object, actions string)
* DeleteRole(name string)
* Role(name string) (role Role)
* RoleByID(appID int64, name string) (role Role)
* Roles() (roles []Role)
* RolesByID(appID int64) (roles []Role)

/// Permission
* BanUser(userID, appID int64)
* UnbanUser(userID, appID int64)
* BanApp(banAppID, appID int64)
* UnbanApp(unbanAppID, appID int64)

/// timestamppb.Timestamp
struct Timestamp {
    Seconds int64
    Nanos int32
}

struct User {
    userID     int64
    email      string
    created_at google.protobuf.Timestamp
    visited_at google.protobuf.Timestamp
}

struct App {
    appID      int64 
    name       string
    created_at google.protobuf.Timestamp
    used_at    google.protobuf.Timestamp
}

struct Role {
    ID  int64
    appID   int64
    name    string
    object  string
    actions string
}
```