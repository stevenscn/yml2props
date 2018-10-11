package properties

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func Write(propFile string, dict map[string]string) {
	fmt.Printf("Write to %s\n", propFile)
	f, err := os.Create(propFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var keys []string
	for key, _ := range dict {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		_, err := f.WriteString(fmt.Sprintf("%s=%s\n", key, dict[key]))
		if err != nil {
			log.Fatal(err)
		}
		f.Sync()
	}
}
