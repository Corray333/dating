package types

import "time"

const (
	SexMale = iota
	SexFemale
	SexNonBinary
)

const (
	OrientationGetero = iota
	OrientationHomo
	OrientationBisexual
	OrientationPansexual
	OrientationAsexual
	OrientationHelicopter
	OrientationAnimal
)

const (
	SearchAnybody = 1 << iota
	SearchPair
	SearchFriend
	SearchSexPartner
	SearchInCity
	SearchMale
	SearchFemale
	SearchNonBinary
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
	ByReferal   string `json:"by_referal" db:"by_referal"`
	Search      int    `json:"search" db:"search"`
	Searching   bool   `json:"searching" db:"searching"`
	Orientation int    `json:"orientation" db:"orientation"`
	Interests   []int  `json:"interests" db:"-"`

	EmailVerified bool `json:"email_verified" db:"email_verified"`
	PhoneVerified bool `json:"phone_verified" db:"phone_verified"`

	BirthTime time.Time `json:"-" db:"birth"`
}
