//go:build dev || server
// +build dev server

package runtime

import _ "embed"

//go:embed ipc_websocket.js
var WebsocketIPC []byte
