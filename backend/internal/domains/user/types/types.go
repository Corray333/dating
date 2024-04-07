package types

type User struct {
	ID            int    `json:"id,omitempty" db:"user_id"`
	Username      string `json:"username" db:"username"`
	Email         string `json:"email" db:"email"`
	Password      string `json:"password,omitempty" db:"password"`
	Avatar        string `json:"avatar" db:"avatar"`
	Name          string `json:"name" db:"name"`
	Surname       string `json:"surname" db:"surname"`
	Patronymic    string `json:"patronymic" db:"patronymic"`
	City          string `json:"city" db:"city"`
	Bio           string `json:"bio" db:"bio"`
	Sex           string `json:"sex" db:"sex"`
	Referal       string `json:"referal" db:"referal"`
	OrientationID int32  `json:"orientation_id" db:"orientation_id"`
	IsSubmitted   bool   `json:"is_submitted" db:"is_submitted"`
}

type Orientation struct {
	ID          int32  `json:"id" db:"orientation_id"`
	Orientation string `json:"orientation" db:"orientation"`
}

type Interest struct {
	ID       int32  `json:"id" db:"interest_id"`
	Interest string `json:"interest" db:"interest"`
	Icon     string `json:"icon" db:"icon"`
}
