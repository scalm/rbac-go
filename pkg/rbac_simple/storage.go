package rbac_simple

import (
	"github.com/scalm/rbac/pkg/rbac"
)

type Storage struct {
	subjects              []rbac.Subject
	roles                 []rbac.Role
	permissions           []rbac.Permission
	subjectAssignments    []rbac.SubjectAssignment
	permissionAssignments []rbac.PermissionAssignment
}

func (storage *Storage) AddSubject(subject rbac.Subject) {
	storage.subjects = append(storage.subjects, subject)
}

func (storage *Storage) AddRole(role rbac.Role) {
	storage.roles = append(storage.roles, role)
}

func (storage *Storage) AddPermission(permission rbac.Permission) {
	storage.permissions = append(storage.permissions, permission)
}

func (storage *Storage) AddSubjectAssignment(sa rbac.SubjectAssignment) {
	storage.subjectAssignments = append(storage.subjectAssignments, sa)
}

func (storage *Storage) AddPermissionAssignment(pa rbac.PermissionAssignment) {
	storage.permissionAssignments = append(storage.permissionAssignments, pa)
}

func (storage *Storage) HasSubjectRole(subjectRef rbac.SubjectRef, roleRef rbac.RoleRef) bool {
	for _, assignment := range storage.subjectAssignments {
		if subjectRef.IsEqual(assignment.GetSubjectRef()) && roleRef.IsEqual(assignment.GetRoleRef()) {
			return true
		}
	}
	return false
}

func (storage *Storage) HasRolePermission(roleRef rbac.RoleRef, permissionRef rbac.PermissionRef) bool {
	for _, assignment := range storage.permissionAssignments {
		if roleRef.IsEqual(assignment.GetRoleRef()) && permissionRef.IsEqual(assignment.GetPermissionRef()) {
			return true
		}
	}
	return false
}

func (storage *Storage) HasSubjectPermission(subjectRef rbac.SubjectRef, permissionRef rbac.PermissionRef) bool {
	for _, assignment := range storage.subjectAssignments {
		if subjectRef.IsEqual(assignment.GetSubjectRef()) && storage.HasRolePermission(assignment.GetRoleRef(), permissionRef) {
			return true
		}
	}
	return false
}

func (storage *Storage) IterateSubjects() rbac.SubjectIterator {
	return &SimpleStorageSubjectIterator{storage.subjects, -1}
}

type SimpleStorageSubjectIterator struct {
	list  []rbac.Subject
	index int
}

func (it *SimpleStorageSubjectIterator) FetchNext() bool {
	it.index++
	return it.index < len(it.list)
}

func (it *SimpleStorageSubjectIterator) Get() rbac.Subject {
	return it.list[it.index]
}

func (storage *Storage) IterateSubjectRoles(subject rbac.Subject) rbac.RoleIterator {
	return &SimpleStorageRoleIterator{storage.subjectAssignments, subject, -1}
}

type SimpleStorageRoleIterator struct {
	list    []rbac.SubjectAssignment
	subject rbac.Subject
	index   int
}

func (it *SimpleStorageRoleIterator) FetchNext() bool {
	it.index++
	for l := len(it.list); it.index < l; it.index++ {
		e := it.list[it.index]
		if it.subject.IsEqual(e.GetSubjectRef()) {
			return true
		}
	}
	return false
}

func (it *SimpleStorageRoleIterator) Get() rbac.Role {
	return it.list[it.index].GetRoleRef()
}

func (storage *Storage) IterateRolePermissions(role rbac.Role) rbac.PermissionIterator {
	return &StorageRolePermissionIterator{storage.permissionAssignments, role, -1}
}

type StorageRolePermissionIterator struct {
	list  []rbac.PermissionAssignment
	role  rbac.Role
	index int
}

func (it *StorageRolePermissionIterator) FetchNext() bool {
	it.index++
	for l := len(it.list); it.index < l; it.index++ {
		e := it.list[it.index]
		if e.GetRoleRef().IsEqual(it.role) {
			return true
		}
	}
	return false
}

func (it *StorageRolePermissionIterator) Get() rbac.Permission {
	return it.list[it.index].GetPermissionRef()
}
