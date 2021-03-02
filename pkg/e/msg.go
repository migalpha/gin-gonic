package e

var MsgFlags = map[int]string {
    SUCCESS : "ok",
    ERROR : "fail",
    ERROR_MISSING_TOKEN: "Missing token",
    ERROR_AUTH_TOKEN: "Invalid token",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Expired token",
    ERROR_AUTH_CHECK_TOKEN_FAIL: "Attemp to check token fails, please refresh the token",
    
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}