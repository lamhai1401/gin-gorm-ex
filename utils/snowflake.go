package utils

import "github.com/bwmarrin/snowflake"

var snowflakeGenerator *snowflake.Node

func InitSnowflakeGenerator(id int64) error {
	var err error
	snowflakeGenerator, err = snowflake.NewNode(id)
	if err != nil {
		return err
	}
	return nil
}

func GetSnowflakeID() string {
	return snowflakeGenerator.Generate().String()
}

func ParseSnowflakeID(data string) (snowflake.ID, error) {
	id, err := snowflake.ParseString(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}
