package commands

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewHeadCommand 创建head命令
func NewHeadCommand() *cobra.Command {
	var lines int

	cmd := &cobra.Command{
		Use:   "head [file]",
		Short: "Display the first lines of a CSV file",
		Long:  `Display the first lines of a CSV file, including the header row. Default: 5 lines.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]
			return headCommand(filePath, lines)
		},
	}

	cmd.Flags().IntVarP(&lines, "lines", "n", 5, "Number of lines to display")

	return cmd
}

// NewTailCommand 创建tail命令
func NewTailCommand() *cobra.Command {
	var lines int

	cmd := &cobra.Command{
		Use:   "tail [file]",
		Short: "Display the last lines of a CSV file",
		Long:  `Display the last lines of a CSV file, including the header row. Default: 10 lines.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]
			return tailCommand(filePath, lines)
		},
	}

	cmd.Flags().IntVarP(&lines, "lines", "n", 10, "Number of lines to display")

	return cmd
}

// headCommand 实现head命令逻辑
func headCommand(filePath string, lines int) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 创建CSV读取器
	reader := csv.NewReader(file)

	// 读取并显示行
	count := 0
	for count < lines {
		record, err := reader.Read()
		if err != nil {
			break // 文件结束或读取错误
		}

		// 输出记录
		for i, field := range record {
			if i > 0 {
				fmt.Print(",")
			}
			fmt.Print(field)
		}
		fmt.Println()

		count++
	}

	return nil
}

// tailCommand 实现tail命令逻辑
func tailCommand(filePath string, lines int) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 创建CSV读取器
	reader := csv.NewReader(file)

	// 读取所有行
	allRecords, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(allRecords) == 0 {
		return nil // 空文件
	}

	// 标题行总是显示
	headerRow := allRecords[0]
	
	// 输出标题行
	for i, field := range headerRow {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(field)
	}
	fmt.Println()

	// 如果只有标题行，直接返回
	if len(allRecords) == 1 {
		return nil
	}

	// 计算数据行的起始位置
	dataLines := len(allRecords) - 1
	startIndex := 1 // 从数据行开始
	if dataLines > lines-1 {
		startIndex = len(allRecords) - (lines - 1)
	}

	// 输出数据行
	for i := startIndex; i < len(allRecords); i++ {
		record := allRecords[i]
		for j, field := range record {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Print(field)
		}
		fmt.Println()
	}

	return nil
}