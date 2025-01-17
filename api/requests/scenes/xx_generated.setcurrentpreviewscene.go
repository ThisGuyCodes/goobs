// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetCurrentPreviewScene request.
type SetCurrentPreviewSceneParams struct {
	// Scene to set as the current preview scene
	SceneName *string `json:"sceneName,omitempty"`
}

func NewSetCurrentPreviewSceneParams() *SetCurrentPreviewSceneParams {
	return &SetCurrentPreviewSceneParams{}
}
func (o *SetCurrentPreviewSceneParams) WithSceneName(x string) *SetCurrentPreviewSceneParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentPreviewSceneParams) GetRequestName() string {
	return "SetCurrentPreviewScene"
}

// Represents the response body for the SetCurrentPreviewScene request.
type SetCurrentPreviewSceneResponse struct {
	_response
}

/*
Sets the current preview scene.

Only available when studio mode is enabled.
*/
func (c *Client) SetCurrentPreviewScene(params *SetCurrentPreviewSceneParams) (*SetCurrentPreviewSceneResponse, error) {
	data := &SetCurrentPreviewSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
