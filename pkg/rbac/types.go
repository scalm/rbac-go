package rbac

type Subject interface {
	IsEqual(other Subject) bool
}

type Role interface {
	IsEqual(other Role) bool
}

type Permission interface {
	IsEqual(other Permission) bool
}

type SubjectRef interface {

}

type RoleRef interface {

}

type PermissionRef interface {

}

type SubjectAssignment interface {
	GetSubjectRef() SubjectRef
	GetRoleRef() RoleRef
}

type PermissionAssignment interface {
	GetPermissionRef() PermissionRef
	GetRoleRef() RoleRef
}

type Storage interface {
	PermissionChecker
	EntityIterable
	EntityEditor
}

type PermissionChecker interface {
	HasSubjectRole(subjectRef SubjectRef, roleRef RoleRef) bool
	HasRolePermission(roleRef RoleRef, permissionRef PermissionRef) bool
	HasSubjectPermission(subject SubjectRef, permission PermissionRef) bool
}

type EntityIterable interface {
	IterateSubjects() SubjectIterator
	IterateSubjectRoles(subject Subject) RoleIterator
	IterateRolePermissions(role Role) PermissionIterator
}

type EntityEditor interface {
	AddSubject(subject Subject)
	AddRole(role Role)
	AddPermission(permission Permission)
	AddSubjectAssignment(sa SubjectAssignment)
	AddPermissionAssignment(pa PermissionAssignment)
}

type RoleIterator interface {
	FetchNext() bool
	Get() Role
}

type SubjectIterator interface {
	FetchNext() bool
	Get() Subject
}

type PermissionIterator interface {
	FetchNext() bool
	Get() Permission
}
