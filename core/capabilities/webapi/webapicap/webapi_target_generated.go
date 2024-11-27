// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package webapicap

import (
	"encoding/json"
	"fmt"
)

// A target that sends HTTP requests to a URL
type Target struct {
	// Config corresponds to the JSON schema field "config".
	Config TargetConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs TargetPayload `json:"inputs" yaml:"inputs" mapstructure:"inputs"`
}

type TargetConfig struct {
	// The delivery mode for the request. Defaults to SingleNode
	DeliveryMode *string `json:"deliveryMode,omitempty" yaml:"deliveryMode,omitempty" mapstructure:"deliveryMode,omitempty"`

	// The number of times to retry the request. Defaults to 0 retries
	RetryCount *uint8 `json:"retryCount,omitempty" yaml:"retryCount,omitempty" mapstructure:"retryCount,omitempty"`

	// The timeout in milliseconds for the request. If set to 0, the default value is
	// 30 seconds
	TimeoutMs *uint32 `json:"timeoutMs,omitempty" yaml:"timeoutMs,omitempty" mapstructure:"timeoutMs,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetConfig) UnmarshalJSON(b []byte) error {
	type Plain TargetConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.RetryCount != nil && 10 < *plain.RetryCount {
		return fmt.Errorf("field %s: must be <= %v", "retryCount", 10)
	}
	if plain.TimeoutMs != nil && 600000 < *plain.TimeoutMs {
		return fmt.Errorf("field %s: must be <= %v", "timeoutMs", 600000)
	}
	*j = TargetConfig(plain)
	return nil
}

type TargetPayload struct {
	// The body of the request
	Body *string `json:"body,omitempty" yaml:"body,omitempty" mapstructure:"body,omitempty"`

	// The headers to include in the request
	Headers TargetPayloadHeaders `json:"headers,omitempty" yaml:"headers,omitempty" mapstructure:"headers,omitempty"`

	// The HTTP method to use for the request
	Method *string `json:"method,omitempty" yaml:"method,omitempty" mapstructure:"method,omitempty"`

	// The URL to send the request to
	Url string `json:"url" yaml:"url" mapstructure:"url"`
}

// The headers to include in the request
type TargetPayloadHeaders map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetPayload) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["url"]; raw != nil && !ok {
		return fmt.Errorf("field url in TargetPayload: required")
	}
	type Plain TargetPayload
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TargetPayload(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Target) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Target: required")
	}
	if _, ok := raw["inputs"]; raw != nil && !ok {
		return fmt.Errorf("field inputs in Target: required")
	}
	type Plain Target
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Target(plain)
	return nil
}
