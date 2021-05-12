package helpers

import "path/filepath"

func Include(path string) []string {
	files,_ := filepath.Glob("site/views/templates/*.html")
	path_files,_ := filepath.Glob("site/views/" + path + "/*.html")
	for _,file := range path_files{
		files = append(files,file)
	}
	return files
}

