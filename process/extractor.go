package process

import (
	"bufio"
	"fmt"
	"github.com/Rom1-J/preprocessor/constants"
	"github.com/Rom1-J/preprocessor/logger"
	"github.com/Rom1-J/preprocessor/structs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Extractor(filePath string) ([]structs.MetadataStruct, error) {
	logger.Logger.Info().Msgf("Extractor starting on: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		var msg = fmt.Sprintf("Failed to open file: %v", err)
		logger.Logger.Error().Msgf(msg)
		fmt.Println(msg)

		return nil, nil
	}

	var metadataList []structs.MetadataStruct

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			return
		}
	}(file)

	reader := bufio.NewReader(file)
	offset := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			break
		}

		fragments := constants.FragmentPattern.FindAllString(line, -1)

		var splitPart = strings.Split(filepath.Ext(filePath), ".part")
		var part int
		if len(splitPart) != 2 {
			part = -1
		} else {
			part, err = strconv.Atoi(splitPart[1])
			if err != nil {
				part = -1
			}
		}

		metadata := structs.MetadataStruct{
			Part:      part,
			Offset:    offset,
			Fragments: fragments,
		}

		metadataList = append(metadataList, metadata)

		offset += len(line)
	}

	logger.Logger.Info().Msgf("Extractor finished on: %s", filePath)
	return metadataList, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
