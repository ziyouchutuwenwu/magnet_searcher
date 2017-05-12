package parser

import (
	"encoding/hex"
)

func Parse(metadata interface{}, infoHash []byte) *BitTorrent {

	info := metadata.(map[string]interface{})

	btObject := new(BitTorrent)

	btObject.InfoHash = hex.EncodeToString(infoHash)

	if name, isExisted := info["name"]; isExisted {
		btObject.Name = name.(string)
	}

	//要么存在files节点，要么是外面的length节点
	if value, ok := info["files"]; ok {

		countedFileSize := 0

		files := value.([]interface{})
		btObject.Files = make([]File, len(files))

		for index, item := range files {
			file := item.(map[string]interface{})

			filePath := file["path"].([]interface{})
			fileLength := file["length"].(int)

			countedFileSize += fileLength

			fileObject := File{
				Path:   filePath,
				Length: fileLength,
			}
			btObject.Files[index] = fileObject
		}
		btObject.Length = countedFileSize

	}else if length, isExisted := info["length"]; isExisted {
		btObject.Length = length.(int)
	}

	return btObject
}
