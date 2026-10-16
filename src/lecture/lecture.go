package lecture

import (
	"io/fs"
	"os"
)

func Lecture(txtentre string) string {
	var fsys fs.FS = os.DirFS("src/fichiertxt")
	contenu, err := fs.ReadFile(fsys, txtentre)
	if err != nil {
		return "Erreur"
	} else {
		return string(contenu)
	}
}
