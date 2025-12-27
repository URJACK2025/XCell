package commands

import (
	"fmt"

	"xcel/internal/utils"

	"github.com/spf13/cobra"
)

// NewColStatCommand 创建col_stat子命令
func NewColStatCommand() *cobra.Command {
	var statType string
	var bucketNumber int
	var columnSpec string

	cmd := &cobra.Command{
		Use:   "col_stat [input-file]",
		Short: "Analyze CSV files by columns with various statistic types",
		Long:  `Analyze CSV files by columns, support different statistic types (set, bucket), and output the results in JSON format.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFile := args[0]

			// 检查使用set统计类型时是否指定了列
			if statType == "set" && columnSpec == "" {
				return fmt.Errorf("set statistic type requires specifying a column with -c/--column flag")
			}

			fmt.Printf("Analyzing %s for column statistics with type %s...\n", inputFile, statType)

			// 调用CSV统计工具
			err := utils.AnalyzeCSVColumns(inputFile, statType, bucketNumber, columnSpec)
			if err != nil {
				return fmt.Errorf("analysis failed: %w", err)
			}

			return nil
		},
	}

	// 添加命令选项
	cmd.Flags().StringVarP(&statType, "type", "t", "set", "Statistic type: set (default), bucket")
	cmd.Flags().IntVarP(&bucketNumber, "number", "n", 10, "Number of buckets for bucket statistic (default: 10)")
	cmd.Flags().StringVarP(&columnSpec, "column", "c", "", "Specify column by name or number (e.g., '性别' or 3)")

	return cmd
}
