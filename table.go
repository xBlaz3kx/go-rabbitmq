package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Table stores user supplied fields of the following types:
//
//	bool
//	byte
//	float32
//	float64
//	int
//	int16
//	int32
//	int64
//	nil
//	string
//	time.Time
//	amqp.Decimal
//	amqp.Table
//	[]byte
//	[]interface{} - containing above types
//
// Functions taking a table will immediately fail when the table contains a
// value of an unsupported type.
//
// The caller must be specific in which precision of integer it wishes to
// encode.
//
// Use a type assertion when reading values from a table for type conversion.
//
// RabbitMQ expects int32 for integer values.
// The table also implements a propagation.TextMapCarrier interface for tracing, so trace and span id can be passed as headers.
type Table map[string]interface{}

func (t Table) Get(key string) string {
	v, ok := t[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func (t Table) Set(key string, value string) {
	t[key] = value
}

func (t Table) Keys() []string {
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

func tableToAMQPTable(table Table) amqp.Table {
	new := amqp.Table{}
	for k, v := range table {
		new[k] = v
	}
	return new
}

func mergeTables(tables ...Table) Table {
	newTable := Table{}
	for _, table := range tables {
		for k, v := range table {
			newTable[k] = v
		}
	}
	return newTable
}
