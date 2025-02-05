package filesystem

import (
	"fmt"
)

// CheckFilesystem checks if a filesystem exists on a given drive using blkid.
// It returns nil if no filesystem is found, otherwise returns an error.
func CheckFilesystem(driveName string) error {
	cmd := "blkid"
	args := []string{
		"-o",
		"value",
		"-s",
		"TYPE",
		driveName,
	}

	fsOut, err := Command(cmd, args, "")
	if err != nil {
		if fsOut.Status == 2 {
			// No filesystem detected, return nil (OK)
			return nil
		}
		return err // Other errors should be returned
	}

	// If blkid returns any filesystem type, return an error
	if fsOut.Stdout != "" {
		return fmt.Errorf("Filesystem detected: %s", strings.TrimSpace(fsOut.Stdout))
	}

	return nil
}


//CreateFilesystem executes mkfs.<desired_filesystem> on the requested drive.
func CreateFilesystem(driveName string, desiredFs string, label string) error {
	cmd := "mkfs." + desiredFs
	args := []string{
		"-L",
		"GOAT-" + label,
		driveName,
	}

	if _, err := Command(cmd, args, ""); err != nil {
		return err
	}
	return nil
}
