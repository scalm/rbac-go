package rbac_simple

import "github.com/scalm/rbac/pkg/rbac"

type Role struct {
	Id string
}

func NewRole(id string) *Role {
	return &Role{id}
}

func (role *Role) IsEqual(other rbac.Role) bool {
	return *role == *(interface{}(other).(*Role))
}

type Permission struct {
	Id string
}

func NewPermission(id string) *Permission {
	return &Permission{id}
}

func (perm *Permission) IsEqual(other rbac.Permission) bool {
	return *perm == *(interface{}(other).(*Permission))
}

type Subject struct {
	Id string
}

func NewSubject(id string) *Subject {
	return &Subject{id}
}

func (subject *Subject) IsEqual(other rbac.Subject) bool {
	return *subject == *(interface{}(other).(*Subject))
}

type SubjectRef string

type RoleRef string

type PermissionRef string

type PermissionAssignment struct {
	roleRef       *RoleRef
	permissionRef *PermissionRef
}

func (pa *PermissionAssignment) GetPermissionRef() rbac.PermissionRef {
	return pa.permissionRef
}

func (pa *PermissionAssignment) GetRoleRef() rbac.RoleRef {
	return pa.roleRef
}

func NewPermissionAssignment(role *RoleRef, permission *PermissionRef) *PermissionAssignment {
	return &PermissionAssignment{role, permission}
}

type SubjectAssignment struct {
	subjectRef *SubjectRef
	roleRef    *RoleRef
}

func (sa *SubjectAssignment) GetSubjectRef() rbac.SubjectRef {
	return sa.subjectRef
}

func (sa *SubjectAssignment) GetRoleRef() rbac.RoleRef {
	return sa.roleRef
}

func NewSubjectAssignment(subjectRef *SubjectRef, roleRef *RoleRef) *SubjectAssignment {
	return &SubjectAssignment{subjectRef, roleRef}
}
