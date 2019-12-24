package multichain

//Creates a new stream on the blockchain called name. For now, always pass the value "stream" in the type parameter – this is designed for future functionality. If open is true then anyone with global send permissions can publish to the stream, otherwise publishers must be explicitly granted per-stream write permissions.
func (client *Client) Create(typeName, name string, open bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      CONST_ID,
		"method":  "create",
		"params":  []interface{}{typeName, name, open},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
