package cli

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"github.com/Carina-labs/nova/x/airdrop/types"
	"os"
)

func ReadUserStateFromFile(filename string) ([]*types.UserState, error) {
	var userState []*types.UserState
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &userState)
	if err != nil {
		return nil, err
	}

	return userState, nil
}

func ReadUserAddressFromCSVFile(filename string) ([]string, error) {
	var result []string
	file, _ := os.Open(filename)
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		result = append(result, row[0])
	}

	return result, nil
}
