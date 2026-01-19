package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Map与json互转
	data := map[string]interface{}{
		"name":    "Alice",
		"age":     30,
		"email":   "alice@qq.com",
		"hobbies": []string{"reading", "swimming", "coding"},
		"address": map[string]string{
			"city":    "Beijing",
			"country": "China",
		},
	}
	// 编码为json
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Json: %s\n", jsonData)

	// Json转Map
	var decodeData map[string]interface{}
	err = json.Unmarshal(jsonData, &decodeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("解码后的Map: %v\n", decodeData)

	// 访问解码后的数据
	fmt.Printf("姓名: %v\n", decodeData["name"])
	fmt.Printf("年龄: %v\n", decodeData["age"])

	// 处理嵌套数据
	if address, ok := decodeData["address"].(map[string]interface{}); ok {
		fmt.Printf("城市：%s\n", address["city"])
	}

	// 处理数组数据
	if hobbies, ok := decodeData["hobbies"].([]interface{}); ok {
		fmt.Println("爱好：")
		for _, hobby := range hobbies {
			fmt.Printf("  - %v\n", hobby)
		}
	}
}
