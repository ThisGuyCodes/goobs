// This file has been automatically generated. Don't edit it.

package ui

/*
GetMonitorListParams represents the params body for the "GetMonitorList" request.
Gets a list of connected monitors and information about them.
*/
type GetMonitorListParams struct{}

/*
GetMonitorListResponse represents the response body for the "GetMonitorList" request.
Gets a list of connected monitors and information about them.
*/
type GetMonitorListResponse struct {
	// a list of detected monitors with some information
	Monitors []interface{} `json:"monitors,omitempty"`
}

// GetMonitorList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetMonitorList(paramss ...*GetMonitorListParams) (*GetMonitorListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetMonitorListParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetMonitorList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetMonitorListResponse), nil
}
