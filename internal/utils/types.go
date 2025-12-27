package utils

// Bucket 定义桶统计结果的结构体
type Bucket struct {
	Range   string `json:"range"`
	Count   int    `json:"count"`
	Min     string `json:"min"`
	Max     string `json:"max"`
	Average string `json:"average,omitempty"`
}

// ColumnStats 定义列统计结果的结构体
type ColumnStats struct {
	ColumnName  string   `json:"column_name"`
	Unique      []string `json:"unique_values,omitempty"`
	Count       int      `json:"count"`
	UniqueCount int      `json:"unique_count,omitempty"`
	StatType    string   `json:"stat_type"`
	Buckets     []Bucket `json:"buckets,omitempty"`
	MinValue    string   `json:"min_value,omitempty"`
	MaxValue    string   `json:"max_value,omitempty"`
}
