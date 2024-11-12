package model

import (
	"encoding/json"
	"io"
	"os"
)

/**
 * 获取费用规则
 * @return
 */
func GetFeeRules() ([]map[string]interface{}, error) {
	var feeReules []map[string]interface{}

	file, err := os.Open("feeReules.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &feeReules); err != nil {
		return nil, err
	}

	return feeReules, nil
}
