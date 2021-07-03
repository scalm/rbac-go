package rbac_simple

import (
	"github.com/scalm/rbac/pkg/rbac"
)

type Controller struct {
	factory *Factory
	storage rbac.Storage
}

func NewController(factory *Factory, storage rbac.Storage) *Controller {
	return &Controller{factory, storage}
}

func (controller *Controller) AddRole(role *rbac.Role) {
	controller.storage.AddRole(*role)
}

func (controller *Controller) AddRoles(roles []*Role) {
	for _, role := range roles {
		var r rbac.Role = role
		controller.AddRole(&r)
	}
}

func (controller *Controller) AddSubjects(subjects []*Subject) {
	for _, subject := range subjects {
		controller.AddSubject(subject)
	}
}

func (controller *Controller) AddPermissions(permissions []*Permission) {
	for _, permission := range permissions {
		controller.AddPermission(permission)
	}
}

func (controller *Controller) AddSubject(subject *Subject) {
	controller.storage.AddSubject(subject)
}

func (controller *Controller) AddPermission(permission *Permission) {
	controller.storage.AddPermission(permission)
}

func (controller *Controller) AssignRole(subjectRef *SubjectRef, roleRef *RoleRef) *SubjectAssignment {
	sa := controller.factory.CreateSubjectAssignment(subjectRef, roleRef)
	controller.storage.AddSubjectAssignment(sa)
	return sa
}

func (controller *Controller) AssignPermission(permissionRef *PermissionRef, roleRef *RoleRef) *PermissionAssignment {
	pa := controller.factory.CreatePermissionAssignment(permissionRef, roleRef)
	controller.storage.AddPermissionAssignment(pa)
	return pa
}

func (controller *Controller) AssignPermissions(permissionRefs []*PermissionRef, roleRef *RoleRef) []*PermissionAssignment {
	assigments := make([]*PermissionAssignment, 0, len(permissionRefs))
	for _, permRef := range permissionRefs {
		assigment := controller.AssignPermission(permRef, roleRef)
		assigments = append(assigments, assigment)
	}
	return assigments
}

func (controller *Controller) AssignRoles(roleRefs []*RoleRef, subjectRef *SubjectRef) []*SubjectAssignment {
	assignments := make([]*SubjectAssignment, 0, len(roleRefs))
	for _, roleRef := range roleRefs {
		assigment := controller.AssignRole(subjectRef, roleRef)
		assignments = append(assignments, assigment)
	}
	return assignments
}

func (controller *Controller) GetSubjects() []rbac.Subject {
	list := make([]rbac.Subject, 0)
	for it := controller.storage.IterateSubjects(); it.FetchNext(); {
		list = append(list, it.Get())
	}
	return list
}

func (controller *Controller) GetSubjectRoles(subject *Subject) []*Role {
	roleSet := make(map[rbac.Role]bool)
	for it := controller.storage.IterateSubjectRoles(subject); it.FetchNext(); {
		role := it.Get()
		roleSet[role] = true
	}

	return castTypeRoleSlice(convertRoleSetToRoleList(roleSet))
}

func castTypeRoleSlice(source []rbac.Role) []*Role {
	list := make([]*Role, 0, len(source))
	for _, e := range source {
		list = append(list, e.(*Role))
	}
	return list
}

func (controller *Controller) GetRolePermissions(role *Role) []*Permission {
	permSet := make(map[rbac.Permission]bool)
	for it := controller.storage.IterateRolePermissions(role); it.FetchNext(); {
		permission := it.Get()
		permSet[permission] = true
	}

	return castTypePermissionSlice(convertPermSetToPermList(permSet))
}

func castTypePermissionSlice(source []rbac.Permission) []*Permission {
	list := make([]*Permission, 0, len(source))
	for _, e := range source {
		list = append(list, e.(*Permission))
	}
	return list
}

func (controller *Controller) GetSubjectPermissions(subject *Subject) []rbac.Permission {
	permSet := make(map[rbac.Permission]bool)
	// It is ugly and it has too many memory allocations. What about iterators? Does Golang know about it? Terribly.
	for _, role := range controller.GetSubjectRoles(subject) {
		for _, perm := range controller.GetRolePermissions(role) {
			permSet[perm] = true
		}
	}

	return convertPermSetToPermList(permSet)
}

func convertPermSetToPermList(set map[rbac.Permission]bool) []rbac.Permission {
	// Why does not Golang have a function to do this???
	list := make([]rbac.Permission, 0, len(set))
	for perm, _ := range set {
		list = append(list, perm)
	}

	return list
}

func convertRoleSetToRoleList(set map[rbac.Role]bool) []rbac.Role {
	// Why does not Golang have a function to do this???
	list := make([]rbac.Role, 0, len(set))
	for perm, _ := range set {
		list = append(list, perm)
	}

	return list
}
