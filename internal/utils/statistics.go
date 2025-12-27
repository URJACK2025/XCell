package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// performStatistics 根据统计类型执行不同的统计逻辑
func performStatistics(rows [][]string, columnNames []string, columnsToAnalyze []int, statType string, bucketNumber int) ([]ColumnStats, error) {
	// 初始化列统计结果
	columnStats := make([]ColumnStats, 0, len(columnsToAnalyze))
	for _, colIndex := range columnsToAnalyze {
		columnStats = append(columnStats, ColumnStats{
			ColumnName: columnNames[colIndex],
			Count:      0,
			StatType:   statType,
		})
	}

	// 根据统计类型执行不同的统计逻辑
	if statType == "set" {
		return performSetStatistics(rows, columnsToAnalyze, columnStats)
	} else if statType == "bucket" {
		return performBucketStatistics(rows, columnsToAnalyze, columnStats, bucketNumber)
	} else {
		return nil, fmt.Errorf("unsupported statistic type: %s", statType)
	}
}

// performSetStatistics 执行set统计逻辑
func performSetStatistics(rows [][]string, columnsToAnalyze []int, columnStats []ColumnStats) ([]ColumnStats, error) {
	rowCount := len(rows)

	// 初始化唯一值map
	uniqueMaps := make([]map[string]bool, len(columnsToAnalyze))
	for i := range uniqueMaps {
		uniqueMaps[i] = make(map[string]bool)
	}

	// 遍历所有数据行
	for i := 1; i < rowCount; i++ {
		row := rows[i]
		for j, colIndex := range columnsToAnalyze {
			var value string
			if colIndex < len(row) {
				value = row[colIndex]
			} else {
				// 处理行长度不足的情况
				value = ""
			}
			// 更新唯一值map
			uniqueMaps[j][value] = true
			columnStats[j].Count++
		}
	}

	// 计算每列的唯一值
	for i := range columnStats {
		// 将map转换为切片
		uniqueValues := make([]string, 0, len(uniqueMaps[i]))
		for value := range uniqueMaps[i] {
			uniqueValues = append(uniqueValues, value)
		}
		// 排序唯一值
		sort.Strings(uniqueValues)

		columnStats[i].Unique = uniqueValues
		columnStats[i].UniqueCount = len(uniqueValues)
	}

	return columnStats, nil
}

// performBucketStatistics 执行bucket统计逻辑
func performBucketStatistics(rows [][]string, columnsToAnalyze []int, columnStats []ColumnStats, bucketNumber int) ([]ColumnStats, error) {
	rowCount := len(rows)

	// 初始化值列表，用于计算桶统计
	valuesList := make([][]float64, len(columnsToAnalyze))
	for i := range valuesList {
		valuesList[i] = make([]float64, 0, rowCount-1)
	}

	// 遍历所有数据行，收集数值
	for i := 1; i < rowCount; i++ {
		row := rows[i]
		for j, colIndex := range columnsToAnalyze {
			var valueStr string
			if colIndex < len(row) {
				valueStr = strings.TrimSpace(row[colIndex])
			} else {
				continue
			}

			// 尝试转换为数值
			value, err := strconv.ParseFloat(valueStr, 64)
			if err == nil {
				valuesList[j] = append(valuesList[j], value)
				columnStats[j].Count++
			}
		}
	}

	// 为每列计算桶统计
	for i, values := range valuesList {
		if len(values) == 0 {
			continue
		}

		// 排序数值
		sort.Float64s(values)

		// 计算最小值和最大值
		minVal := values[0]
		maxVal := values[len(values)-1]
		columnStats[i].MinValue = fmt.Sprintf("%f", minVal)
		columnStats[i].MaxValue = fmt.Sprintf("%f", maxVal)

		// 计算桶宽
		rangeVal := maxVal - minVal
		bucketWidth := rangeVal / float64(bucketNumber)

		// 初始化桶
		buckets := make([]Bucket, bucketNumber)
		for j := 0; j < bucketNumber; j++ {
			bucketMin := minVal + float64(j)*bucketWidth
			bucketMax := minVal + float64(j+1)*bucketWidth
			buckets[j] = Bucket{
				Range: fmt.Sprintf("%.2f-%.2f", bucketMin, bucketMax),
				Min:   fmt.Sprintf("%.2f", bucketMin),
				Max:   fmt.Sprintf("%.2f", bucketMax),
				Count: 0,
			}
		}

		// 分配值到桶
		for _, value := range values {
			// 特殊处理最大值，确保它被分配到最后一个桶
			if value == maxVal {
				buckets[bucketNumber-1].Count++
				continue
			}

			// 计算桶索引
			bucketIndex := int((value - minVal) / bucketWidth)
			// 确保索引在有效范围内
			if bucketIndex >= bucketNumber {
				bucketIndex = bucketNumber - 1
			}
			buckets[bucketIndex].Count++
		}

		columnStats[i].Buckets = buckets
	}

	return columnStats, nil
}
