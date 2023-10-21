package dto

import "fr33d0mz/moneyflowx/models"

type UserRequestParams struct {
	UserID string `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Username string `form:"username"`
	Email    string `form:"email"`
}

type UserResponseBody struct {
	ID       string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserDetailResponse struct {
	ID       string         `json:"uuid"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Wallet   WalletResponse `json:"wallet"`
}

func FormatUser(user *models.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.ID
	formattedUser.Username = user.Username
	formattedUser.Email = user.Email
	return formattedUser
}

func FormatUsers(authors []*models.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

func FormatUserDetail(user *models.User, wallet *models.Wallet) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = user.ID
	formattedUser.Username = user.Username
	formattedUser.Email = user.Email
	formattedUser.Wallet = FormatWallet(wallet)
	return formattedUser
}
