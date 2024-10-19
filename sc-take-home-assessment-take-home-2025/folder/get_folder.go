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

	folders := f.folders

	parentPath, inOrg, folderExists := f.getPath(orgID, name)

	// checking folder existence
	if !folderExists {

		return nil, errors.New("error: folder does not exist")
	}

	// checking existence in organisation
	// also returns if organisation does not exist
	if !inOrg {
		return nil, errors.New("error: no such folder in specified organisation")
	}

	res := []Folder{}
	for _, folder := range folders {

		// Folder must belong to same organisation
		if folder.OrgId == orgID {

			if isChild(parentPath, folder) {
				res = append(res, folder)
			}
		}
	}

	return res, nil

	// return []Folder{}
}

func isChild(parentPath string, childCandidate Folder) bool {

	// checking if candidate child is the parent
	if childCandidate.Paths == parentPath {
		return false
	}

	// checking if file path includes parent candidate filepath
	if strings.HasPrefix(childCandidate.Paths, parentPath+".") {

		return true
	}

	return false
}

func (f *driver) getPath(orgID uuid.UUID, name string) (string, bool, bool) {

	inOrg := false
	folderExists := false

	var filepath string
	folders := f.folders

	for _, folder := range folders {

		// checking correct folder
		if folder.Name == name {

			folderExists = true

			// checking correct organisation
			if folder.OrgId == orgID {

				inOrg = true

				// collect filepath
				filepath = folder.Paths
				break
			}
		}
	}

	return filepath, inOrg, folderExists
}
