package clickhouse_api

import (
	"crypto/tls"
	"github.com/rnbondarenko/clickhouse-go/v2"
	"github.com/rnbondarenko/clickhouse-go/v2/lib/driver"
	clickhouse_tests "github.com/rnbondarenko/clickhouse-go/v2/tests"
)

const TestSet string = "examples_clickhouse_api"

func GetNativeConnection(settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	return clickhouse_tests.GetConnection(TestSet, settings, tlsConfig, compression)
}

func GetNativeTestEnvironment() (clickhouse_tests.ClickHouseTestEnvironment, error) {
	return clickhouse_tests.GetTestEnvironment(TestSet)
}
