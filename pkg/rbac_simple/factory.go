package rbac_simple

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

type CreateSubjectParams struct{ Id string }

type CreateRoleParams struct{ Id string }

type CreatePermissionParams struct{ Id string }

func (factory *Factory) CreateSubject(subject *CreateSubjectParams) *Subject {
	return NewSubject(subject.Id)
}

func (factory *Factory) CreateRole(role *CreateRoleParams) *Role {
	return NewRole(role.Id)
}

func (factory *Factory) CreatePermission(permission *CreatePermissionParams) *Permission {
	return NewPermission(permission.Id)
}

func (factory *Factory) CreateSubjectAssignment(subjectRef *SubjectRef, roleRef *RoleRef) *SubjectAssignment {
	return NewSubjectAssignment(subjectRef, roleRef)
}

func (factory *Factory) CreatePermissionAssignment(permissionRef *PermissionRef, roleRef *RoleRef) *PermissionAssignment {
	return NewPermissionAssignment(roleRef, permissionRef)
}
