package frontend_app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateUUID() string {
	return uuid.New().String()
}

func SetDefault(err error) error {
	return errors.New("default error")
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch v.(type) {
	case *bool, *int, *int32, *int64, *uint, *uint32, *uint64, *string:
		return false
	default:
		return true
	}
}

func Now() int64 {
	return time.Now().Unix()
}

func IntSliceContains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func IntSliceDiff(slice1, slice2 []int) []int {
	res := make([]int, 0)
	for _, item := range slice1 {
		if !IntSliceContains(slice2, item) {
			res = append(res, item)
		}
	}
	return res
}

func StrSliceContains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func StrSliceDiff(slice1, slice2 []string) []string {
	res := make([]string, 0)
	for _, item := range slice1 {
		if !StrSliceContains(slice2, item) {
			res = append(res, item)
		}
	}
	return res
}

func DBError(err error) error {
	if err == sql.ErrNoRows {
		return errors.New("no rows in result set")
	}
	return err
}

func GetEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

func GetEnvInt(key string, defaultVal int) int {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	if i, err := strconv.Atoi(val); err != nil {
		log.Printf("error parsing %s as int, using default: %d", key, defaultVal)
		return defaultVal
	}
	return i
}

func GetEnvInt64(key string, defaultVal int64) int64 {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	if i, err := strconv.ParseInt(val, 10, 64); err != nil {
		log.Printf("error parsing %s as int, using default: %d", key, defaultVal)
		return defaultVal
	}
	return i
}

func GetEnvBool(key string, defaultVal bool) bool {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	if b, err := strconv.ParseBool(val); err != nil {
		log.Printf("error parsing %s as bool, using default: %t", key, defaultVal)
		return defaultVal
	}
	return b
}

func GetEnvDuration(key string, defaultVal time.Duration) time.Duration {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	if d, err := time.ParseDuration(val); err != nil {
		log.Printf("error parsing %s as duration, using default: %s", key, defaultVal)
		return defaultVal
	}
	return d
}

func GetEnvStringSlice(key string, defaultVal []string) []string {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	return strings.Split(val, ",")
}

func GetEnvIntSlice(key string, defaultVal []int) []int {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make([]int, len(slice))
	for i, x := range slice {
		if i := GetEnvInt("INT"+strings.TrimSpace(x), 0); i != 0 {
			res[i] = i
		}
	}
	return res
}

func GetEnvInt64Slice(key string, defaultVal []int64) []int64 {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make([]int64, len(slice))
	for i, x := range slice {
		if i := GetEnvInt64("INT64"+strings.TrimSpace(x), 0); i != 0 {
			res[i] = i
		}
	}
	return res
}

func GetEnvBoolSlice(key string, defaultVal []bool) []bool {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make([]bool, len(slice))
	for i, x := range slice {
		if i := GetEnvBool("BOOL"+strings.TrimSpace(x), false); i != false {
			res[i] = i
		}
	}
	return res
}

func GetEnvDurationSlice(key string, defaultVal []time.Duration) []time.Duration {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make([]time.Duration, len(slice))
	for i, x := range slice {
		if i := GetEnvDuration("DURATION"+strings.TrimSpace(x), 0); i != 0 {
			res[i] = i
		}
	}
	return res
}

func GetEnvStringDict(key string, defaultVal map[string]string) map[string]string {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make(map[string]string)
	for _, x := range slice {
		pair := strings.SplitN(strings.TrimSpace(x), "=", 2)
		if len(pair) == 2 {
			res[pair[0]] = pair[1]
		}
	}
	return res
}

func GetEnvIntDict(key string, defaultVal map[string]int) map[string]int {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make(map[string]int)
	for _, x := range slice {
		pair := strings.SplitN(strings.TrimSpace(x), "=", 2)
		if len(pair) == 2 {
			res[pair[0]] = GetEnvInt("INT"+pair[1], 0)
		}
	}
	return res
}

func GetEnvInt64Dict(key string, defaultVal map[string]int64) map[string]int64 {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make(map[string]int64)
	for _, x := range slice {
		pair := strings.SplitN(strings.TrimSpace(x), "=", 2)
		if len(pair) == 2 {
			res[pair[0]] = GetEnvInt64("INT64"+pair[1], 0)
		}
	}
	return res
}

func GetEnvBoolDict(key string, defaultVal map[string]bool) map[string]bool {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make(map[string]bool)
	for _, x := range slice {
		pair := strings.SplitN(strings.TrimSpace(x), "=", 2)
		if len(pair) == 2 {
			res[pair[0]] = GetEnvBool("BOOL"+pair[1], false)
		}
	}
	return res
}

func GetEnvDurationDict(key string, defaultVal map[string]time.Duration) map[string]time.Duration {
	val := GetEnv(key, "")
	if val == "" {
		return defaultVal
	}
	slice := strings.Split(val, ",")
	res := make(map[string]time.Duration)
	for _, x := range slice {
		pair := strings.SplitN(strings.TrimSpace(x), "=", 2)
		if len(pair) == 2 {
			res[pair[0]] = GetEnvDuration("DURATION"+pair[1], 0)
		}
	}
	return res
}