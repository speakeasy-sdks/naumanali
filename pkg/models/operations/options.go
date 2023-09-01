// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/naumanali/pkg/utils"
)

var ErrUnsupportedOption = errors.New("unsupported option")

const (
	SupportedOptionServerURL            = "serverURL"
	SupportedOptionRetries              = "retries"
	SupportedOptionAcceptHeaderOverride = "acceptHeaderOverride"
)

type AcceptHeaderEnum string

const (
	AcceptHeaderEnumApplicationJson            AcceptHeaderEnum = "application/json"
	AcceptHeaderEnumApplicationProblemPlusJson AcceptHeaderEnum = "application/problem+json"
	AcceptHeaderEnumApplicationYaml            AcceptHeaderEnum = "application/yaml"
)

func (e AcceptHeaderEnum) ToPointer() *AcceptHeaderEnum {
	return &e
}

type Options struct {
	ServerURL            *string
	Retries              *utils.RetryConfig
	AcceptHeaderOverride *AcceptHeaderEnum
}

type Option func(*Options, ...string) error

// WithServerURL allows providing an alternative server URL.
func WithServerURL(serverURL string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionServerURL) {
			return ErrUnsupportedOption
		}

		opts.ServerURL = &serverURL
		return nil
	}
}

// WithTemplatedServerURL allows providing an alternative server URL with templated parameters.
func WithTemplatedServerURL(serverURL string, params map[string]string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionServerURL) {
			return ErrUnsupportedOption
		}

		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		opts.ServerURL = &serverURL
		return nil
	}
}

// WithRetries allows customizing the default retry configuration.
func WithRetries(config utils.RetryConfig) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionRetries) {
			return ErrUnsupportedOption
		}

		opts.Retries = &config
		return nil
	}
}

func WithAcceptHeaderOverride(acceptHeaderOverride AcceptHeaderEnum) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionAcceptHeaderOverride) {
			return ErrUnsupportedOption
		}

		opts.AcceptHeaderOverride = &acceptHeaderOverride
		return nil
	}
}
