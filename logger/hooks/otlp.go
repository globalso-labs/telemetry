/*
 * telemetry
 * otlp.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sun, 15 Sep 2024 18:34:32 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package hooks

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
)

// OTLPHook struct defines a logger hook for the zerolog logging library.
type OTLPHook struct{}

// Run is the method that gets called on each log event.
// It converts the zerolog event to an OpenTelemetry log record and emits it using the global logger provider.
// In case of PanicLevel or FatalLevel log events, it also attempts to shut down the logger provider gracefully.
//
// Parameters:
// - event: The zerolog event that contains all the log information.
// - constants: The logging constants of the event (e.g., Info, Warn, Error).
// - message: The log message.
//
// The method extracts the context from the event, converts the event to an OpenTelemetry log record,
// and emits the record using the global logger provider. If the log constants is PanicLevel or FatalLevel,
// it shuts down the logger provider to ensure all logs are flushed before the application exits.
func (h OTLPHook) Run(event *zerolog.Event, level zerolog.Level, message string) {
	ctx := event.GetCtx()

	// Extract context from the event.
	record := h.convertEvent(event, level, message) // Convert zerolog event to OpenTelemetry log record.
	provider := global.GetLoggerProvider()          // Get the global logger provider.
	provider.Logger("global").Emit(ctx, record)     // Emit the log record.
}

// convertEvent transforms a zerolog event into an OpenTelemetry log record.
//
// This function takes a zerolog event, its constants, and a message string as inputs.
// It constructs an OpenTelemetry log record by setting the timestamp, message body,
// severity constants, observed timestamp, and any additional attributes extracted from the event.
//
// Parameters:
// - e *zerolog.Event: The zerolog event to be converted.
// - constants zerolog.Level: The logging constants of the event.
// - msg string: The log message.
//
// Returns:
// - log.Record: The constructed OpenTelemetry log record.
func (h OTLPHook) convertEvent(e *zerolog.Event, level zerolog.Level, msg string) log.Record {
	var record log.Record
	record.SetTimestamp(time.Now().UTC())         // Set the timestamp using zerolog's configured function.
	record.SetBody(log.StringValue(msg))          // Set the log message body.
	record.SetSeverity(convertSeverity(level))    // Convert and set the severity constants based on zerolog's constants.
	record.SetSeverityText(level.String())        // Set the severity text using zerolog's constants string.
	record.AddAttributes(convertOTLPFields(e)...) // Convert and add any additional fields  attributes.
	return record
}

// convertSeverity converts a zerolog logging constants to an OpenTelemetry log severity.
//
// This function maps zerolog's logging levels to OpenTelemetry's log severity levels.
// It ensures that logs are categorized correctly in the OpenTelemetry ecosystem
// according to their severity. The mapping is direct for most levels, but zerolog's
// PanicLevel, NoLevel, and Disabled are all mapped to OpenTelemetry's SeverityUndefined
// as they do not have direct equivalents.
//
// Parameters:
// - constants zerolog.Level: The zerolog logging constants to be converted.
//
// Returns:
// - log.Severity: The corresponding OpenTelemetry log severity constants.
func convertSeverity(level zerolog.Level) log.Severity {
	switch level {
	case zerolog.TraceLevel:
		return log.SeverityTrace1
	case zerolog.DebugLevel:
		return log.SeverityDebug1
	case zerolog.InfoLevel:
		return log.SeverityInfo1
	case zerolog.WarnLevel:
		return log.SeverityWarn1
	case zerolog.ErrorLevel:
		return log.SeverityError1
	case zerolog.FatalLevel:
		return log.SeverityFatal1
	case zerolog.PanicLevel:
		return log.SeverityFatal3
	case zerolog.NoLevel, zerolog.Disabled:
		return log.SeverityUndefined
	}

	return log.SeverityUndefined
}

// convertOTLPFields extracts and converts zerolog event fields to OpenTelemetry key-value pairs.
//
// This function iterates over all fields present in a zerolog event, converting each field
// to an OpenTelemetry log.KeyValue structure. The conversion process is handled by the
// convertValue function, which adapts the field's value to the appropriate OpenTelemetry
// log.Value type based on the value's underlying type.
//
// Parameters:
// - e *zerolog.Event: The zerolog event containing the fields to be converted.
//
// Returns:
// - []log.KeyValue: A slice of OpenTelemetry key-value pairs representing the converted fields.
func convertOTLPFields(e *zerolog.Event) []log.KeyValue {
	ev := fmt.Sprintf("%s}", reflect.ValueOf(e).Elem().FieldByName("buf"))
	data := make(map[string]interface{})
	if err := json.Unmarshal([]byte(ev), &data); err != nil {
		return nil
	}

	kvs := make([]log.KeyValue, 0)
	for k, v := range data {
		kvs = append(kvs, log.KeyValue{
			Key:   k,
			Value: convertValue(v),
		})
	}

	return kvs
}

// convertValue adapts a generic interface value to a specific OpenTelemetry log value type.
//
// This function takes a value of type interface{} and determines its actual type to convert it
// into the corresponding OpenTelemetry log value type. Supported types include basic Go types
// (bool, []byte, float64, int, int64, string) and complex types (struct, slice, array, map).
// For complex types, it recursively processes each element or field to ensure accurate
// representation in the log record. Unhandled types are converted to a string representation
// with a warning about the unhandled type.
//
// Parameters:
// - v interface{}: The value to be converted into an OpenTelemetry log value.
//
// Returns:
// - log.Value: The OpenTelemetry log value representation of the input.
func convertValue(v interface{}) log.Value {
	switch v := v.(type) {
	case bool:
		return log.BoolValue(v)
	case []byte:
		return log.BytesValue(v)
	case float64:
		return log.Float64Value(v)
	case int:
		return log.IntValue(v)
	case int64:
		return log.Int64Value(v)
	case string:
		return log.StringValue(v)
	}

	t := reflect.TypeOf(v)
	if t == nil {
		return log.Value{}
	}
	val := reflect.ValueOf(v)
	switch t.Kind() { //nolint:exhaustive // We only handle the types we care about.
	case reflect.Struct:
		return log.StringValue(fmt.Sprintf("%+v", v))
	case reflect.Slice, reflect.Array:
		items := make([]log.Value, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			items = append(items, convertValue(val.Index(i).Interface()))
		}
		return log.SliceValue(items...)
	case reflect.Map:
		kvs := make([]log.KeyValue, 0, val.Len())
		for _, k := range val.MapKeys() {
			var key string
			if k.Kind() == reflect.Struct {
				key = fmt.Sprintf("%+v", k.Interface())
			} else {
				key = fmt.Sprintf("%v", k.Interface())
			}
			kvs = append(kvs, log.KeyValue{
				Key:   key,
				Value: convertValue(val.MapIndex(k).Interface()),
			})
		}
		return log.MapValue(kvs...)
	case reflect.Ptr, reflect.Interface:
		return convertValue(val.Elem().Interface())
	default:
		return log.StringValue(fmt.Sprintf("unhandled attribute type: (%s) %+v", t, v))
	}
}
