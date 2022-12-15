package dto

type UserRegistry struct {
	Account    string `json:"account"`
	Key        string `json:"key"`
	Class      string `json:"class"`
	Schoolname string `json:"schoolname"`
	Telephone  string `json:"telephone"`
	Outtime    string `json:"outtime"`
}

type UserLogin struct {
	Account string `json:"account"`
	Key     string `json:"key"`
}

// type UserDto struct {
//     Account string `json:"account"`
// }

// func ToUserDto(user model.User) UserDto {
//     return UserDto{
//         Account: user.Account,
//     }
// }
