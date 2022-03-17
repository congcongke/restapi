// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ListTasksURL generates an URL for the list tasks operation
type ListTasksURL struct {
	ExecutionID       int64
	PreheatPolicyName string
	ProjectName       string

	Page     *int64
	PageSize *int64
	Q        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListTasksURL) WithBasePath(bp string) *ListTasksURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListTasksURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListTasksURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks"

	executionID := swag.FormatInt64(o.ExecutionID)
	if executionID != "" {
		_path = strings.Replace(_path, "{execution_id}", executionID, -1)
	} else {
		return nil, errors.New("executionId is required on ListTasksURL")
	}

	preheatPolicyName := o.PreheatPolicyName
	if preheatPolicyName != "" {
		_path = strings.Replace(_path, "{preheat_policy_name}", preheatPolicyName, -1)
	} else {
		return nil, errors.New("preheatPolicyName is required on ListTasksURL")
	}

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{project_name}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on ListTasksURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v2.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt64(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("page_size", pageSizeQ)
	}

	var qQ string
	if o.Q != nil {
		qQ = *o.Q
	}
	if qQ != "" {
		qs.Set("q", qQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListTasksURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListTasksURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListTasksURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListTasksURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListTasksURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListTasksURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}