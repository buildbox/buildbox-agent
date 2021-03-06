package api

import "fmt"

// Annotation represents a Buildkite Agent API Annotation
type Annotation struct {
	Body    string `json:"body,omitempty"`
	Context string `json:"context,omitempty"`
	Style   string `json:"style,omitempty"`
	Append  bool   `json:"append,omitempty"`
}

// Annotate a build in the Buildkite UI
func (c *Client) Annotate(jobId string, annotation *Annotation) (*Response, error) {
	u := fmt.Sprintf("jobs/%s/annotations", jobId)

	req, err := c.newRequest("POST", u, annotation)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req, nil)
}

// Remove an annotation from a build
func (c *Client) AnnotationRemove(jobId string, context string) (*Response, error) {
	u := fmt.Sprintf("jobs/%s/annotations/%s", jobId, context)

	req, err := c.newRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req, nil)
}
