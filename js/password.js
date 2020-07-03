function password_enc(salt,password){
// 省略具体加密过程
    return salt+password+salt
}
