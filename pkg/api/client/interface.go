package client

import (
	"io"
	"net/http"

	"github.com/kubeshop/kubtest/pkg/api/kubtest"
)

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
}

type Client interface {
	GetScript(id string) (script kubtest.Script, err error)
	GetExecution(scriptID, executionID string) (execution kubtest.ScriptExecution, err error)
	ListExecutions(scriptID string) (executions kubtest.ScriptExecutions, err error)
	CreateScript(options CreateScriptOptions) (script kubtest.Script, err error)
	ExecuteScript(id, namespace, executionName string, executionParams map[string]string) (execution kubtest.ScriptExecution, err error)
	ListScripts(namespace string) (scripts kubtest.Scripts, err error)
}

// CreateScriptOptions - is mapping for now to OpenAPI schema for creating request
// if needed can beextended to custom struct
type CreateScriptOptions kubtest.ScriptCreateRequest
