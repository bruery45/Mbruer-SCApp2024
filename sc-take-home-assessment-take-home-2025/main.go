package main

import (
	"log"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID_A := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID_B := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")

	// res := folder.GetAllFolders()

	res := []folder.Folder{
		// org_A
		{Name: "a", Paths: "a", OrgId: orgID_A},

		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "be", Paths: "a.be", OrgId: orgID_A},
		{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

		// org_B
		{Name: "f", Paths: "f", OrgId: orgID_B},

		{Name: "g", Paths: "f.g", OrgId: orgID_B},
		{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
		{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

		{Name: "j", Paths: "f.j", OrgId: orgID_B},

		{Name: "k", Paths: "f.k", OrgId: orgID_B},
		{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
	}

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	// childFolder, err := folderDriver.GetAllChildFolders(orgID_B, "b")

	newFolders, err := folderDriver.MoveFolder("g", "j")

	if err == nil {
		folder.PrettyPrint(newFolders)
	} else {
		log.Fatalf("%v", err)
	}
	// fmt.Print(folderDriver)

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	// folder.PrettyPrint(orgFolder)
}
