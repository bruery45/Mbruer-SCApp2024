package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string, orgID uuid.UUID) ([]Folder, error) {

	/*
		Added orgID to method signature as, otherwise, edge case of two different
		organisations both having folders (with same names as 'dst' and 'name')
		is not handled.
	*/

	folders := f.folders

	// checking if moving folder into itself
	if name == dst {

		return folders, errors.New("error: cannot move a folder to itself")
	}

	// getting desired folders
	destination, destExists := f.getFolder(dst, orgID)
	source, sourceExists := f.getFolder(name, orgID)

	// checking if destination folder exists
	if !destExists {
		return folders, errors.New("error: destination folder does not exist")
	}

	// checking if source folder exists
	if !sourceExists {
		return folders, errors.New("error: source folder does not exist")
	}

	// checking if dest orgIDs match
	if orgID != destination.OrgId {

		return folders, errors.New("error: cannot move a folder to a different organization")
	}

	// checking if source orgIDs match
	if orgID != source.OrgId {

		return folders, errors.New("error: cannot move a folder from a different organization")
	}

	// checking if moving folder to its own child
	if isChild(source.Paths, destination) {

		return folders, errors.New("error: cannot move a folder to a child of itself")
	}

	res := []Folder{}

	for _, folder := range folders {

		// skip organisations folders from other organisations
		if source.OrgId != folder.OrgId {
			res = append(res, folder)
			continue
		}

		// moving current folder or child folder into destination
		if folder.Name == source.Name || isChild(source.Paths, folder) {

			relativePath := subPath(folder, name)

			newPath := destination.Paths + "." + relativePath

			folder.Paths = newPath
			// folders[i].Paths = newPath

		}

		res = append(res, folder)
	}

	return res, nil

}

// Method no longer used due to more efficient approach being used
func (f *driver) getFolder(name string, orgID uuid.UUID) (Folder, bool) {

	found := false
	folders := f.folders

	var folder Folder

	for _, item := range folders {

		// if same name and organisation, return
		if item.Name == name {
			found = true
			folder = item

			// continue to collect orgID if invalid
			if orgID == item.OrgId {

				break
			}
		}
	}

	return folder, found
}
func subPath(folder Folder, parent string) string {

	pathFolders := strings.Split(folder.Paths, ".")

	var subPath string

	// finding relative path of folder in its path
	for index, folder := range pathFolders {

		if parent == folder {

			subPath = strings.Join(pathFolders[index:], ".")
			break
		}
	}

	return subPath
}
