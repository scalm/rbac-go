package rbac

type AbstractFactory interface {
	CreateSubject(attributes interface{}) Subject
	CreateRole(attributes interface{}) Role
	CreatePermission(attributes interface{}) Permission
	CreateSubjectAssignment(subject SubjectRef, role RoleRef) SubjectAssignment
	CreatePermissionAssignment(permission PermissionRef, role RoleRef) PermissionAssignment
}
