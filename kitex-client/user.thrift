namespace go usermodel

struct User {
    1: required i64 ID,
    2: required string loginName,
    3: required string password,
    4: required string nickName,
}

service UserService {
    // 登录接口
    bool signIn(1: User user),
    // 注册接口
    bool signUp(1: User user),
    // 获取所有用户
    list<User> getUserList(),
    // 根据UserId获取用户
    User getUserById(1: i64 id),
}