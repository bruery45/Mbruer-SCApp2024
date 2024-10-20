package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	/*
		As the method signature does not contain orgIDs, we assume that only one
		instance of name and dst occur within the file structure.

		Otherwise, we would implement explicit checking when getting their folders
	*/

	// checking if moving folder into itself
	if name == dst {

		return nil, errors.New("error: cannot move a folder to itself")
	}

	// getting desired folders
	destination, destExists := f.getFolder(dst)
	source, sourceExists := f.getFolder(name)

	// checking if destination folder exists
	if !destExists {
		return nil, errors.New("error: destination folder does not exist")
	}

	// checking if source folder exists
	if !sourceExists {
		return nil, errors.New("error: source folder does not exist")
	}

	// checking if orgIDs match
	if source.OrgId != destination.OrgId {

		return nil, errors.New("error: cannot move a folder to a different organization")
	}

	// checking if moving folder to its own child
	if isChild(destination.Paths, source) {

		return nil, errors.New("error: cannot move a folder to a child of itself")
	}

	folders := f.folders

	res := make([]Folder, len(folders))

	copy(res, folders)

	for _, folder := range folders {

		// skip organisations folders from other organisations
		if source.OrgId != folder.OrgId {
			continue
		}

		// moving current folder or child folder into destination
		if folder.Name == source.Name || isChild(source.Paths, folder) {

			relativePath := subPath(folder, name)

			newPath := destination.Paths + "." + relativePath

			folder.Paths = newPath
		}
	}

	return res, nil

}

func (f *driver) getFolder(name string) (Folder, bool) {

	found := false
	folders := f.folders

	var folder Folder

	for _, item := range folders {

		// if same name, return
		if item.Name == name {
			found = true
			folder = item
			break
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
		}
	}

	return subPath
}
