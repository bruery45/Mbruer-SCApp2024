package folder

import (
	"github.com/gofrs/uuid"

	"strings"

	"errors"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	inOrg := false
	folderExists := false

	folders := f.folders

	res := []Folder{}
	for _, folder := range folders {

		// checking folder existence for error handling
		if folder.Name == name {
			folderExists = true

			if folder.OrgId == orgID {
				inOrg = true
			}
		}

		// Folder must belong to same organisation
		if folder.OrgId == orgID {

			if isChild(name, folder) {
				res = append(res, folder)
			}
		}
	}

	if !folderExists {

		return nil, errors.New("Error: folder does not exist")
	}

	if !inOrg {
		return nil, errors.New("Error: no such folder in specified organisation")
	}

	return res, nil

	// return []Folder{}
}

func isChild(parentPath string, childCandidate Folder) bool {

	// checking if candidate child is the parent
	if childCandidate.Paths == parentPath {
		return false
	}

	// checking if file path includes parent candidate
	// adding '.' for similarly named file paths
	if strings.HasPrefix(childCandidate.Paths, parentPath+".") {
		return true
	}

	return false
}
