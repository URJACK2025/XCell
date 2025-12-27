package utils

// AnalyzeCSVColumns 分析CSV文件的列统计信息并以JSON格式输出
func AnalyzeCSVColumns(inputFile string, statType string, bucketNumber int, columnSpec string) error {
	// 读取CSV文件
	rows, columnNames, err := readCSVFile(inputFile)
	if err != nil {
		return err
	}

	// 确定要分析的列索引
	columnsToAnalyze, err := determineColumnsToAnalyze(columnSpec, columnNames)
	if err != nil {
		return err
	}

	// 根据统计类型执行不同的统计逻辑
	columnStats, err := performStatistics(rows, columnNames, columnsToAnalyze, statType, bucketNumber)
	if err != nil {
		return err
	}

	// 输出统计结果
	return outputResults(inputFile, columnStats)
}
