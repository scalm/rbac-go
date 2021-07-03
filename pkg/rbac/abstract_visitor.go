package rbac

type AbstractVisitor interface {
	VisitDocumentStart() error
	VisitSubject(subject Subject) error
	VisitRole(role Role) error
	VisitPermission(permission Permission) error
	VisitSubjectAssignment(sa SubjectAssignment) error
	VisitPermissionAssignment(pa PermissionAssignment) error
	VisitDocumentEnd() error
}
