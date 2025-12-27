package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// ConvertExcelToCSV 将Excel文件转换为CSV文件
func ConvertExcelToCSV(inputFile string) error {
	// 打开Excel文件
	f, err := excelize.OpenFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	// 获取所有Sheet名称
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return fmt.Errorf("no sheets found in Excel file")
	}

	// 设置buffer目录
	bufferDir := "buffer"

	// 获取文件名（不含扩展名）
	baseName := filepath.Base(inputFile)
	nameWithoutExt := baseName[:len(baseName)-len(filepath.Ext(baseName))]

	// 遍历所有Sheet
	for _, sheet := range sheets {
		fmt.Printf("Processing sheet: %s\n", sheet)

		// 获取Sheet中的所有行
		rows, err := f.GetRows(sheet)
		if err != nil {
			return fmt.Errorf("failed to get rows from sheet %s: %w", sheet, err)
		}

		if len(rows) == 0 {
			fmt.Printf("Sheet %s is empty, skipping...\n", sheet)
			continue
		}

		// 创建CSV文件，保存到buffer目录
		csvFileName := fmt.Sprintf("%s_%s.csv", nameWithoutExt, sheet)
		csvFilePath := filepath.Join(bufferDir, csvFileName)
		csvFile, err := os.Create(csvFilePath)
		if err != nil {
			return fmt.Errorf("failed to create CSV file %s: %w", csvFilePath, err)
		}
		defer csvFile.Close()

		// 创建CSV写入器
		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		// 写入所有行
		for _, row := range rows {
			if err := writer.Write(row); err != nil {
				return fmt.Errorf("failed to write row to CSV file %s: %w", csvFilePath, err)
			}
		}

		fmt.Printf("Created CSV file: %s\n", csvFilePath)
	}

	return nil
}

// ConvertCSVToExcel 将多个CSV文件转换为一个XLSX文件
func ConvertCSVToExcel(inputFiles []string) error {
	// 创建新的Excel文件
	f := excelize.NewFile()

	// 记录已创建的Sheet名称，用于处理重名
	sheetNames := make(map[string]int)

	// 标记是否有CSV文件名为Sheet1
	hasSheet1 := false

	// 遍历所有输入的CSV文件
	for _, csvFile := range inputFiles {
		// 打开CSV文件
		file, err := os.Open(csvFile)
		if err != nil {
			return fmt.Errorf("failed to open CSV file %s: %w", csvFile, err)
		}
		defer file.Close()

		// 创建CSV读取器
		reader := csv.NewReader(file)

		// 读取所有行
		rows, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read rows from CSV file %s: %w", csvFile, err)
		}

		if len(rows) == 0 {
			fmt.Printf("CSV file %s is empty, skipping...\n", csvFile)
			continue
		}

		// 获取CSV文件名（不含扩展名）作为Sheet名称
		baseName := filepath.Base(csvFile)
		sheetName := baseName[:len(baseName)-len(filepath.Ext(baseName))]

		// 检查是否有同名Sheet，如果有则添加数字后缀
		if count, exists := sheetNames[sheetName]; exists {
			// 重名，添加数字后缀
			sheetNames[sheetName]++
			sheetName = fmt.Sprintf("%s_%d", sheetName, count+1)
		} else {
			// 首次出现，记录计数
			sheetNames[sheetName] = 0
		}

		// 检查是否是Sheet1
		if sheetName == "Sheet1" {
			hasSheet1 = true
		}

		// 创建新Sheet
		sheetIndex, err := f.NewSheet(sheetName)
		if err != nil {
			return fmt.Errorf("failed to create sheet %s: %w", sheetName, err)
		}

		// 设置活动Sheet
		f.SetActiveSheet(sheetIndex)

		// 写入所有行到Sheet
		for i, row := range rows {
			for j, col := range row {
				cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
				f.SetCellValue(sheetName, cell, col)
			}
		}

		fmt.Printf("Added sheet %s from CSV file %s\n", sheetName, csvFile)
	}

	// 在末尾检查并删除默认Sheet1（如果没有同名CSV文件）
	if !hasSheet1 {
		f.DeleteSheet("Sheet1")
	}

	// 设置buffer目录
	bufferDir := "buffer"

	// 保存Excel文件到buffer目录
	excelFileName := "output.xlsx"
	excelFilePath := filepath.Join(bufferDir, excelFileName)
	if err := f.SaveAs(excelFilePath); err != nil {
		return fmt.Errorf("failed to save Excel file %s: %w", excelFilePath, err)
	}

	fmt.Printf("Created Excel file: %s\n", excelFilePath)
	return nil
}
