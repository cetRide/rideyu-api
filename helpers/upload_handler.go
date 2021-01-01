package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cetRide/rideyu-api/apihelpers"
	"github.com/h2non/filetype"
)

const MAX_MEMORY = 1 * 1024 * 1024

func UploadFiles(
	w http.ResponseWriter, r *http.Request,
	user uint64, category string) map[string]interface{} {
	// path := "storage" + string(os.PathSeparator) + dir + string(os.PathSeparator)

	// path := "storage" + string(os.PathSeparator) + user + string(os.PathSeparator) + "posts"

	path := "storage"

	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		return apihelpers.Message(false, "File size is too large")
	}

	formdata := r.MultipartForm
	description := r.FormValue("description")
	files := formdata.File["file"]
	var filenames []string
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			return apihelpers.Message(false, "Invalid request2")
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return apihelpers.Message(false, "Invalid request2")
		}

		kind, _ := filetype.Match(fileBytes)
		if kind == filetype.Unknown {
			return apihelpers.Message(false, "Unknown file type")
		}
		user_id := strconv.FormatUint(user, 10)

				tempFile, err := ioutil.TempFile(path, user_id+"_"+category+"_*."+kind.Extension)
				if err != nil {
					return apihelpers.Message(false, "Internal server error")
				}
				defer tempFile.Close()
				tempFile.Write(fileBytes)
				filenames = append(filenames, tempFile.Name())
				fmt.Println("the file name is", tempFile.Name())
		}

	response := make(map[string]interface{})
	response["description"] = description
	response["files"] = filenames
	return response
}