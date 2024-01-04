package builder

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"xray-wrapper/vmess_maker/entity"
)

func RemoveRightPart(str, substring string) string {
	return str[:strings.Index(str, substring)]
}

func RemoveLeftPart(str, substring string) string {
	return str[strings.Index(str, substring)+len(substring):]
}

func GenerateRandomString(length int) string {

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	result := string(hex.EncodeToString(b))
	return result
}

func WriteFile(filename string, newVmess entity.VmessJson) error {

	file, err := json.MarshalIndent(newVmess, "", " ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Error during WriteFile(): ", err)
		return err
	}

	return nil
}
