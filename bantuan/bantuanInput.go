package bantuan

import(
	"minipro/user"
)

type GetBantuanDetailInput struct {
	ID int `uri:"id" validate:"required"`
}

type CreateBantuanInput struct {
	TittleBantuan             string `json:"tittle_bantuan" validate:"required"`
	DeskripsiSingkat 		  string `json:"short_description" binding:"required"`
	Deskripsi   	   		  string `json:"description" binding:"required"`
	JumlahTarget       	 	  int    `json:"goal_amount" binding:"required"`
	ListKondisi            	  string `json:"perks" binding:"required"`
	User             		  user.User
}