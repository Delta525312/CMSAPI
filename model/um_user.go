package model

import (
	"time"
)

type Um_User_Login struct {
	ID                    string     `json:"id"`
	OrgID                 string     `json:"orgId"`
	OrgName               string     `json:"orgName"`
	DisplayName           string     `json:"displayName"`
	Title                 string     `json:"title"`
	FirstName             string     `json:"firstName"`
	MiddleName            *string    `json:"middleName"`
	LastName              string     `json:"lastName"`
	CitizenID             string     `json:"citizenId"`
	Bod                   time.Time  `json:"bod"`
	Blood                 string     `json:"blood"`
	Gender                string     `json:"gender"`
	MobileNo              *string    `json:"mobileNo"`
	Address               *string    `json:"address"`
	Photo                 *string    `json:"photo"`
	Username              string     `json:"username"`
	Password              string     `json:"password"`
	Email                 *string    `json:"email"`
	RoleID                string     `json:"roleId"`
	Permission            []string   `json:"permission"`
	RoleName              string     `json:"roleName"`
	UserType              string     `json:"userType"`
	EmpID                 string     `json:"empId"`
	DeptID                string     `json:"deptId"`
	CommID                string     `json:"commId"`
	StnID                 string     `json:"stnId"`
	Active                bool       `json:"active"`
	ActivationToken       *string    `json:"activationToken"`
	LastActivationRequest *int64     `json:"lastActivationRequest"`
	LostPasswordRequest   *int64     `json:"lostPasswordRequest"`
	SignupStamp           *int64     `json:"signupStamp"`
	IsLogin               bool       `json:"islogin"`
	LastLogin             *time.Time `json:"lastLogin"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
	CreatedBy             string     `json:"createdBy"`
	UpdatedBy             string     `json:"updatedBy"`
}

type Um_User struct {
	ID                    string     `json:"id"`
	OrgID                 string     `json:"orgId"`
	OrgName               string     `json:"orgName"`
	DisplayName           string     `json:"displayName"`
	Title                 string     `json:"title"`
	FirstName             string     `json:"firstName"`
	MiddleName            *string    `json:"middleName"`
	LastName              string     `json:"lastName"`
	CitizenID             string     `json:"citizenId"`
	Bod                   time.Time  `json:"bod"`
	Blood                 string     `json:"blood"`
	Gender                string     `json:"gender"`
	MobileNo              *string    `json:"mobileNo"`
	Address               *string    `json:"address"`
	Photo                 *string    `json:"photo"`
	Username              string     `json:"username"`
	Password              string     `json:"password"`
	Email                 *string    `json:"email"`
	RoleID                string     `json:"roleId"`
	RoleName              string     `json:"roleName"`
	UserType              string     `json:"userType"`
	EmpID                 string     `json:"empId"`
	DeptID                string     `json:"deptId"`
	CommID                string     `json:"commId"`
	StnID                 string     `json:"stnId"`
	Active                bool       `json:"active"`
	ActivationToken       *string    `json:"activationToken"`
	LastActivationRequest *int64     `json:"lastActivationRequest"`
	LostPasswordRequest   *int64     `json:"lostPasswordRequest"`
	SignupStamp           *int64     `json:"signupStamp"`
	IsLogin               bool       `json:"islogin"`
	LastLogin             *time.Time `json:"lastLogin"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
	CreatedBy             string     `json:"createdBy"`
	UpdatedBy             string     `json:"updatedBy"`
}

type UserInput struct {
	DisplayName           string     `json:"displayName"`
	Title                 string     `json:"title"`
	FirstName             string     `json:"firstName"`
	MiddleName            string     `json:"middleName"`
	LastName              string     `json:"lastName"`
	CitizenID             string     `json:"citizenId"`
	Bod                   time.Time  `json:"bod"`
	Blood                 string     `json:"blood"`
	Gender                *int64     `json:"gender"`
	MobileNo              string     `json:"mobileNo"`
	Address               string     `json:"address"`
	Photo                 *string    `json:"photo"`
	Username              string     `json:"username"`
	Password              string     `json:"password"`
	Email                 string     `json:"email"`
	RoleID                string     `json:"roleId"`
	UserType              *int64     `json:"userType"`
	EmpID                 string     `json:"empId"`
	DeptID                string     `json:"deptId"`
	CommID                string     `json:"commId"`
	StnID                 string     `json:"stnId"`
	Active                bool       `json:"active"`
	LastActivationRequest *int64     `json:"lastActivationRequest"`
	LostPasswordRequest   *int64     `json:"lostPasswordRequest"`
	SignupStamp           *int64     `json:"signupStamp"`
	IsLogin               bool       `json:"islogin"`
	LastLogin             *time.Time `json:"lastLogin"`
}

type UserUpdate struct {
	DisplayName           string     `json:"displayName"`
	Title                 string     `json:"title"`
	FirstName             string     `json:"firstName"`
	MiddleName            string     `json:"middleName"`
	LastName              string     `json:"lastName"`
	CitizenID             string     `json:"citizenId"`
	Bod                   string     `json:"bod"`
	Blood                 string     `json:"blood"`
	Gender                *int64     `json:"gender"`
	MobileNo              string     `json:"mobileNo"`
	Address               string     `json:"address"`
	Photo                 *string    `json:"photo"`
	Username              string     `json:"username"`
	Password              string     `json:"password"`
	Email                 string     `json:"email"`
	RoleID                string     `json:"roleId"`
	UserType              *int64     `json:"userType"`
	EmpID                 string     `json:"empId"`
	DeptID                string     `json:"deptId"`
	CommID                string     `json:"commId"`
	StnID                 string     `json:"stnId"`
	Active                bool       `json:"active"`
	LastActivationRequest *int64     `json:"lastActivationRequest"`
	LostPasswordRequest   *int64     `json:"lostPasswordRequest"`
	SignupStamp           *int64     `json:"signupStamp"`
	IsLogin               bool       `json:"islogin"`
	LastLogin             *time.Time `json:"lastLogin"`
}

type UserContact struct {
	OrgID        string    `json:"orgId"`
	Username     string    `json:"username"`
	ContactName  string    `json:"contactName"`
	ContactPhone string    `json:"contactPhone"`
	ContactAddr  any       `json:"contactAddr"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
}

type UserSkill struct {
	OrgID     string    `json:"orgId"`
	UserName  string    `json:"userName"`
	SkillID   string    `json:"skillId"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
}

type UserSocial struct {
	OrgID      string    `json:"orgId"`
	Username   string    `json:"username"`
	SocialType string    `json:"socialType"`
	SocialID   string    `json:"socialId"`
	SocialName string    `json:"socialName"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CreatedBy  string    `json:"createdBy"`
	UpdatedBy  string    `json:"updatedBy"`
}

type UserContactInsert struct {
	OrgID        string `json:"orgId"`
	Username     string `json:"username"`
	ContactName  string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	ContactAddr  any    `json:"contactAddr"`
}

type UserContactInsertUpdate struct {
	ContactName  string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	ContactAddr  any    `json:"contactAddr"`
}

type UserSkillInsert struct {
	OrgID    string `json:"orgId"`
	UserName string `json:"userName"`
	SkillID  string `json:"skillId"`
	Active   bool   `json:"active"`
}

type UserSkillUpdate struct {
	SkillID string `json:"skillId"`
	Active  bool   `json:"active"`
}

type UserSocialInsert struct {
	OrgID      string `json:"orgId"`
	Username   string `json:"username"`
	SocialType string `json:"socialType"`
	SocialID   string `json:"socialId"`
	SocialName string `json:"socialName"`
}

type UserSocialUpdate struct {
	OrgID      string `json:"orgId"`
	Username   string `json:"username"`
	SocialType string `json:"socialType"`
	SocialID   string `json:"socialId"`
	SocialName string `json:"socialName"`
	UpdatedBy  string `json:"updatedBy"`
}
