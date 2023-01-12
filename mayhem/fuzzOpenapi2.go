package fuzzOpenapi2

import "strconv"
import "github.com/grokify/spectrum/openapi2"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {
    
    case 0:
        content := string(bytes)
        openapi2.MergeDirectory(content)
        return 0

    case 1:
        content := string(bytes)
        openapi2.ReadOpenAPI2SpecFile(content)
        return 0

    default:
        content := string(bytes)
        openapi2.FilenameIsYAML(content)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}