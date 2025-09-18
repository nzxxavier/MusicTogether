package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflakeID() (uint, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return uint(node.Generate().Int64()), nil
}
