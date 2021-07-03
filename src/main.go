package main

func main() {
	//factory := rbac_simple.NewFactory()
	//storage := rbac_simple.Storage{}
	//controller := rbac_simple.NewController(factory, &storage)
	//
	//loginPerm := factory.CreatePermissionWithId("login")
	//fmt.Println(loginPerm)
	//controller.AddPermission(loginPerm)
	//
	//readPerm := factory.CreatePermissionWithId("read")
	//fmt.Println(readPerm)
	//controller.AddPermission(readPerm)
	//
	//writePerm := factory.CreatePermissionWithId("write")
	//fmt.Println(writePerm)
	//controller.AddPermission(writePerm)
	//
	//
	//adminRole := factory.CreateRoleWithId("admin")
	//fmt.Println(adminRole)
	//controller.AddRole(adminRole)
	//
	//readerRole := factory.CreateRoleWithId("reader")
	//fmt.Println(readerRole)
	//controller.AddRole(readerRole)
	//
	//editorRole := factory.CreateRoleWithId("editor")
	//fmt.Println(editorRole)
	//controller.AddRole(editorRole)
	//
	//
	//controller.AssignPermissions([]*rbac_simple.Permission{loginPerm, readPerm, writePerm}, adminRole)
	//controller.AssignPermissions([]*rbac_simple.Permission{loginPerm, readPerm}, readerRole)
	//
	//controller.AssignPermissions([]*rbac_simple.Permission{loginPerm, readPerm, writePerm}, editorRole)
	//
	//
	//johnDoe := factory.CreateSubjectWithId("john.doe@domain.com")
	//controller.AddSubject(johnDoe)
	//controller.AssignRoles([]*rbac_simple.Role{adminRole, editorRole}, johnDoe)
	//
	//marieStuart := factory.CreateSubjectWithId("marie.stuart@domain.com")
	//controller.AddSubject(marieStuart)
	//controller.AssignRole(marieStuart, readerRole)
	//
	//doctorWho := factory.CreateSubjectWithId("doctor.who@domain.com")
	//controller.AddSubject(doctorWho)
	//controller.AssignRole(doctorWho, editorRole)
	//
	//
	//roles := controller.GetSubjectRoles(johnDoe)
	//
	//fmt.Println("roles", roles)
	//
	//perms1 := controller.GetRolePermissions(adminRole)
	//fmt.Println("perms1", perms1)
	//
	//perms := controller.GetSubjectPermissions(johnDoe)
	//fmt.Println("perms", perms)
	//
	//
	//for _, subject := range controller.GetSubjects() {
	//	fmt.Println(subject)
	//	for _, role := range controller.GetSubjectRoles(subject.(*rbac_simple.Subject)) {
	//		fmt.Print("  ", role, ":")
	//		for _, perm := range controller.GetRolePermissions(role) {
	//			fmt.Print(" ", perm)
	//		}
	//		fmt.Println()
	//	}
	//	fmt.Print("  Total:")
	//	for _, perm := range controller.GetSubjectPermissions(subject.(*rbac_simple.Subject)) {
	//		fmt.Print(" ", perm)
	//	}
	//	fmt.Println()
	//}
	//
	//hasReadPerm := storage.HasRolePermission(adminRole, readPerm)
	//fmt.Println("", hasReadPerm)
	//
	//hasReadPerm2 := storage.HasRolePermission(adminRole, factory.CreatePermissionWithId("read"))
	//fmt.Println("", hasReadPerm2)
}
