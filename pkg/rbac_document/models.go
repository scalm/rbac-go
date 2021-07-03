package rbac_document

type Document struct {
	Subjects    []*SubjectNode    `json:"subjects"`
	Roles       []*RoleNode       `json:"roles"`
	Permissions []*PermissionNode `json:"permissions"`
}

type SubjectNode struct {
	Id string `json:"id"`
	// Reference to RoleNode.Id
	Roles []string `json:"roles"`
}

type RoleNode struct {
	Id string `json:"id"`
	// Reference to PermissionNode.Id
	Permissions []string `json:"permissions"`
}

type PermissionNode struct {
	Id string `json:"id"`
}
