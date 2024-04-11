package types

import "time"

const (
	SexMale = iota + 1
	SexFemale
	SexNonBinary
)

const (
	OrientationGetero = iota + 1
	OrientationHomo
	OrientationBisexual
	OrientationAsexual
	OrientationHelicopter
	OrientationAnimal
)

const (
	SearchAnybody = iota + 1
	SearchPair
	SearchFriend
	SearchInCity
	SearchSexPartner
)

type User struct {
	ID          int    `json:"id,omitempty" db:"user_id"`
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password,omitempty" db:"password"`
	Avatar      string `json:"avatar" db:"avatar"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Birth       int    `json:"birth" db:"-"`
	Phone       string `json:"phone" db:"phone"`
	City        string `json:"city" db:"city"`
	Bio         string `json:"bio" db:"bio"`
	Sex         int    `json:"sex" db:"sex"`
	Referal     string `json:"referal" db:"referal"`
	Orientation int    `json:"orientation" db:"orientation"`
	IsSubmitted bool   `json:"is_submitted" db:"is_submitted"`
	Interests   []int  `json:"interests" db:"-"`

	BirthTime time.Time `json:"-" db:"birth"`
}
