// This file has been automatically generated. Don't edit it.

package general

/*
GetVersionParams represents the params body for the "GetVersion" request.
Gets data about the current plugin and RPC version.
*/
type GetVersionParams struct{}

/*
GetVersionResponse represents the response body for the "GetVersion" request.
Gets data about the current plugin and RPC version.
*/
type GetVersionResponse struct {
	// Array of available RPC requests for the currently negotiated RPC version
	AvailableRequests []string `json:"availableRequests,omitempty"`

	// Current OBS Studio version
	ObsVersion string `json:"obsVersion,omitempty"`

	// Current obs-websocket version
	ObsWebSocketVersion string `json:"obsWebSocketVersion,omitempty"`

	// Name of the platform. Usually `windows`, `macos`, or `ubuntu` (linux flavor). Not guaranteed to be any of those
	Platform string `json:"platform,omitempty"`

	// Description of the platform, like `Windows 10 (10.0)`
	PlatformDescription string `json:"platformDescription,omitempty"`

	// Current latest obs-websocket RPC version
	RpcVersion float64 `json:"rpcVersion,omitempty"`

	// Image formats available in `GetSourceScreenshot` and `SaveSourceScreenshot` requests.
	SupportedImageFormats []string `json:"supportedImageFormats,omitempty"`
}

// GetVersion sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetVersion(paramss ...*GetVersionParams) (*GetVersionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVersionParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetVersion", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetVersionResponse), nil
}
