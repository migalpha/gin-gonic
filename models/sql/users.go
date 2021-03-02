package sql

// User struct to map account_users table
type User struct {
	ID        int    `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// JSONResult struct type response API
type JSONResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// JSONResultError struct type response API on error
type JSONResultError struct {
	Success bool        `json:"success" default:"false"`
	Msg     interface{} `json:"msg"`
}

// TableName set a table for User struct
func (User) TableName() string {
	return "users"
}

// GetUser get a specific user
func GetUser(email string) (user User) {
	db.Where("email = ?", email).First(&user)
	return
}

// GetUsers get all users for pagination
func GetUsers(pageNum int, pageSize int, maps interface{}) (users []User) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)

	return
}

// GetUsersTotal return total of users
func GetUsersTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)
	return
}

// EditUser allow edit an user
func EditUser(id int, data interface{}) {
	db.Model(&User{}).Where("id = ?", id).Updates(data)
}

// ExistUserByID checks if an user exists
func ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
