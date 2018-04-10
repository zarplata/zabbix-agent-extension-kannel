package main

import (
	"encoding/json"
	"fmt"
)

func discovery(
	stats *Stats,
) error {
	discoveryData := make(map[string][]map[string]string)
	var discoveredItems []map[string]string

	for _, provider := range stats.SMSC {
		discoveredItem := make(map[string]string)
		discoveredItem["{#PROVIDER}"] = provider.ID
		discoveredItems = append(discoveredItems, discoveredItem)
	}

	discoveryData["data"] = discoveredItems

	out, err := json.Marshal(discoveryData)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)

	return nil
}
