package parser

import (
	"appmagic-entrance/models"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

func ExportJSON(statistics *models.Statistics, filename string) error {
	// encoded, err := json.Marshal(*statistics)
	encoded, err := json.MarshalIndent(*statistics, "", "	")
	if err != nil {
		return errors.New("unexpected error: " + err.Error())
	}

	err = os.WriteFile(filename, encoded, fs.ModePerm)
	if err != nil {
		return errors.New("error writing file: " + err.Error())
	}
	return nil
}
