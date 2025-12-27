package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readCSVFile 读取CSV文件并返回行数据和列名
func readCSVFile(inputFile string) ([][]string, []string, error) {
	// 打开CSV文件
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open CSV file %s: %w", inputFile, err)
	}
	defer file.Close()

	// 创建CSV读取器
	reader := csv.NewReader(file)

	// 读取所有行
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read rows from CSV file %s: %w", inputFile, err)
	}

	if len(rows) == 0 {
		return nil, nil, fmt.Errorf("CSV file %s is empty", inputFile)
	}

	// 获取列名
	columnNames := rows[0]

	return rows, columnNames, nil
}

// determineColumnsToAnalyze 确定要分析的列索引
func determineColumnsToAnalyze(columnSpec string, columnNames []string) ([]int, error) {
	columnsToAnalyze := make([]int, 0)
	if columnSpec == "" {
		// 分析所有列
		for i := 0; i < len(columnNames); i++ {
			columnsToAnalyze = append(columnsToAnalyze, i)
		}
	} else {
		// 解析列规范
		colIndex, err := parseColumnSpec(columnSpec, columnNames)
		if err != nil {
			return nil, err
		}
		columnsToAnalyze = append(columnsToAnalyze, colIndex)
	}

	return columnsToAnalyze, nil
}

// parseColumnSpec 解析列规范，支持列名或列号
func parseColumnSpec(columnSpec string, columnNames []string) (int, error) {
	// 尝试解析为数字
	colIndex, err := strconv.Atoi(columnSpec)
	if err == nil {
		// 列号从1开始，转换为从0开始的索引
		colIndex--
		if colIndex < 0 {
			return 0, fmt.Errorf("column number must be positive")
		}
		return colIndex, nil
	}

	// 尝试匹配列名
	for i, name := range columnNames {
		if strings.EqualFold(name, columnSpec) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("column not found: %s", columnSpec)
}
