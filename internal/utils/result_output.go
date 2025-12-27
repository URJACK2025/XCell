package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// outputResults 输出统计结果
func outputResults(inputFile string, columnStats []ColumnStats) error {
	// 转换为JSON格式
	jsonData, err := json.MarshalIndent(columnStats, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal column statistics to JSON: %w", err)
	}

	// 输出JSON结果
	fmt.Println("\nColumn Statistics:")
	fmt.Println(string(jsonData))

	// 保存JSON结果到buffer目录
	bufferDir := "buffer"
	baseName := filepath.Base(inputFile)
	nameWithoutExt := baseName[:len(baseName)-len(filepath.Ext(baseName))]
	jsonFileName := fmt.Sprintf("%s_col_stats.json", nameWithoutExt)
	jsonFilePath := filepath.Join(bufferDir, jsonFileName)

	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON result to file %s: %w", jsonFilePath, err)
	}

	fmt.Printf("\nSaved column statistics to %s\n", jsonFilePath)

	return nil
}
