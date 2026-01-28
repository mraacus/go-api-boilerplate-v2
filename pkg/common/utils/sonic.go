package utils

import (
	"go-api-boilerplate/pkg/common/logger"
	"go-api-boilerplate/pkg/common/verror"
	"go-api-boilerplate/pkg/constant"

	"github.com/bytedance/sonic"
)

// EncodeWithDefault return the string representation of an object in JSON format
// if there is an encoding error, default JSON string returns
func EncodeWithDefault(obj interface{}) string {
	b, err := marshal(obj)
	if err != nil {
		logger.Logger.Warn("[sonic] dump object by marshal failed", "object", obj)
		return constant.EmptyJson
	}
	return string(b)
}

// marshal - encode an object
func marshal(v interface{}) ([]byte, error) {
	var api = sonic.ConfigStd
	return api.Marshal(v)
}

// Decode returns error if decoding JSON string into v has error
func Decode(data string, ptr interface{}) error {
	err := unMarshal(data, ptr)
	if err != nil {
		logger.Logger.Error("[sonic] decode sonic json error", "error", err, "data", data)
		return verror.ServiceInternalError
	}
	return nil
}

// unMarshal - decode an object from string
func unMarshal(data string, v interface{}) error {
	var api = sonic.ConfigStd
	return api.UnmarshalFromString(data, v)
}

// Encode return the string representation of an object in JSON format
// nil => null
// plain value => plain value
// object => json format string
func Encode(obj interface{}) (string, error) {
	b, err := marshal(obj)
	if err != nil {
		logger.Logger.Error("[sonic] encode sonic json error", "error", err, "object", obj)
		return "", verror.ServiceInternalError
	}
	return string(b), nil
}