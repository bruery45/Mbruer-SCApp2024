package main

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	childFolder, err := folderDriver.GetAllChildFolders(orgID, "evident-silver-centurion")

	if err == nil {
		folder.PrettyPrint(childFolder)
	}
	// fmt.Print(folderDriver)

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	// folder.PrettyPrint(orgFolder)
}
