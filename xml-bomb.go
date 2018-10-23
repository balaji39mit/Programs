package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	entityKey = "<!ENTITY"
)

// Complete the entityExpansion function below.
func entityExpansion(l int64, entities []string) {
	xmlEntity := make(map[string][]string, 0)
	parsingFailed := false
	for _, entity := range entities {
		items := strings.Fields(entity)
		if len(items) < 3 || items[0] != entityKey {
			parsingFailed = true
			return
		}
		entityValue := strings.TrimSuffix(items[2], ">")
		if number, err := strconv.Atoi(entityValue); err == nil { //given entity is a number
			xmlEntity[items[1]] = []string{string(number)}
		} else { //Given entity is another entity or combination of entities.
			keys := strings.Split(entityValue, ";")
			for _, key := range keys {
				value, ok := xmlEntity[key]
				if ok {
					x, ok := xmlEntity[items[1]]
					if ok {
						xmlEntity[items[1]] = append(x, value...)
					} else {
						xmlEntity[items[1]] = value
					}
				} else {
					parsingFailed = true
					x, ok := xmlEntity[items[1]]
					if ok {
						xmlEntity[items[1]] = append(x, key)
					} else {
						xmlEntity[items[1]] = []string{key}
					}
				}
			}
		}
	}

	//xmlParsing done successfully, check for the length.
	if parsingFailed {
		for key, value := range xmlEntity {
			for index, item := range value {
				val, ok := xmlEntity[item]
				if ok {
					xmlEntity[key] = append(xmlEntity[key][:index], xmlEntity[key][index+1:]...)
					xmlEntity[key] = append(value, val...)
				}
			}

		}
	}
	count := 0
	for _, value := range xmlEntity {
		count += len(value)

	}
	if int64(count) <= l {
		fmt.Println(1, count)
	} else {
		fmt.Println(0, count)
	}
}

func xmlParsingFailed() {
	fmt.Println("Parsingg failed")
	//find the number of lines.

	fmt.Println(0)
}

func main() {
	/*reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	l, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	entitiesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var entities []string

	for i := 0; i < int(entitiesCount); i++ {
		entitiesItem := readLine(reader)
		entities = append(entities, entitiesItem)
	}*/
	l := int64(5)
	entities := []string{"<!ENTITY a1 10>", "<!ENTITY a2  a1;a1>"}

	entityExpansion(l, entities)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
