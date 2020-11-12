package model

// Organization ...
type Organization struct {
	Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Code        Code   `json:"code" gorm:"column:code"`
	ParentID    ID     `json:"parent_id" gorm:"column:parent_id"`
}

// ResultOrganization 返回组织信息
type ResultOrganization struct {
	Organization       *Organization `json:"organization"`
	ParentOrganization *Organization `json:"parent_organization"`
}

// OrganizationRole 组织权限
type OrganizationRole struct {
	Model
	OrganizationID ID   `json:"organization_id" gorm:"column:organization_id"`
	RoleCode       Code `json:"role_code" gorm:"column:role_code"`
}
