package commands

import (
	"fmt"
	"strings"

	"xcel/internal/utils"

	"github.com/spf13/cobra"
)

// NewConvertCommand 创建convert子命令
func NewConvertCommand() *cobra.Command {
	var outputFormat string
	var inputFiles []string

	cmd := &cobra.Command{
		Use:   "convert [input-files...]",
		Short: "Convert Excel and CSV files between formats",
		Long:  `Convert Excel files to CSV and vice versa. Support multiple input files.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// 处理输入文件，支持逗号分隔的多个文件
			for _, arg := range args {
				files := strings.Split(arg, ",")
				inputFiles = append(inputFiles, files...)
			}

			fmt.Printf("Converting %d files to %s format...\n", len(inputFiles), outputFormat)

			// 根据输出格式选择转换逻辑
			if outputFormat == "csv" {
				// Excel to CSV conversion
				for _, file := range inputFiles {
					err := utils.ConvertExcelToCSV(file)
					if err != nil {
						return fmt.Errorf("conversion failed for %s: %w", file, err)
					}
				}
			} else if outputFormat == "xlsx" {
				// CSV to Excel conversion
				err := utils.ConvertCSVToExcel(inputFiles)
				if err != nil {
					return fmt.Errorf("conversion failed: %w", err)
				}
			} else {
				return fmt.Errorf("unsupported output format: %s", outputFormat)
			}

			fmt.Println("Conversion completed successfully!")
			return nil
		},
	}

	// 添加命令选项
	cmd.Flags().StringVarP(&outputFormat, "format", "f", "csv", "Output format (default: csv)")
	// 保持向后兼容
	cmd.Flags().StringVarP(&outputFormat, "type", "t", "csv", "Output format (default: csv)")

	return cmd
}
