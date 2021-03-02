package cache

// CheckConn check redis connection
func CheckConn() (res string) {
	res, err := client.Ping().Result()
	if err != nil {
		res = ""
	}
	return
}

// GetValue return a value
func GetValue(key string) (val string) {
	val, err := client.Get(key).Result()
	if err != nil {
		val = ""
	}
	return
}

// SetValue allow storage a value in redis
func SetValue(key string, val string) {
	_ = client.Set(key, val, 0).Err()
}

// SAdd adds an element in a set in redis
func SAdd(key, member string) {
	_ = client.SAdd(key, member)
}

// SRem remove an element set in redis
func SRem(key, member string) {
	_ = client.SRem(key, member)
}
