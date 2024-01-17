//go:build server
// +build server

package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/pkg/browser"
	"github.com/wailsapp/wails/v2/pkg/assetserver"
	"github.com/wailsapp/wails/v2/pkg/assetserver/webview"
	"github.com/wailsapp/wails/v2/pkg/menu"

	"github.com/wailsapp/wails/v2/internal/binding"
	"github.com/wailsapp/wails/v2/internal/frontend"
	wailsruntime "github.com/wailsapp/wails/v2/internal/frontend/runtime"
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

type Screen = frontend.Screen

var initOnce = sync.Once{}

const startURL = "wails://wails/"

// var secondInstanceBuffer = make(chan options.SecondInstanceData, 1)

type Frontend struct {

	// Context
	ctx context.Context

	frontendOptions *options.App
	logger          *logger.Logger
	debug           bool
	devtoolsEnabled bool

	// Assets
	assets   *assetserver.AssetServer
	startURL *url.URL
	exitChan chan bool
	// main window handle
	bindings   *binding.Bindings
	dispatcher frontend.Dispatcher
}

func (f *Frontend) RunMainLoop() {
	<-f.exitChan
	return
}
func (f *Frontend) BrowserOpenURL(url string) {
	// Specific method implementation
	_ = browser.OpenURL(url)
}

func (f *Frontend) OpenFileDialog(dialogOptions frontend.OpenDialogOptions) (result string, err error) {
	return "", nil
}

func (f *Frontend) MenuSetApplicationMenu(menu *menu.Menu) {

}

func (f *Frontend) MenuUpdateApplicationMenu() {

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func (f *Frontend) OpenMultipleFilesDialog(dialogOptions frontend.OpenDialogOptions) ([]string, error) {
	var userInput string

	done := make(chan bool)
	go func() {
		fmt.Print("请输入文件全路径，多个文件以分号分隔（60秒内）: ")
		fmt.Scanln(&userInput)
		done <- true
	}()
	var err error
	select {
	case <-done:
		fmt.Println("选择的文件是:", userInput)
	case <-time.After(60 * time.Second):
		os.Stdin.Close()
		err = errors.New("超时，未输入任何内容")
	}
	files := strings.Split(userInput, ";")
	for i := range files {
		if !fileExists(files[i]) {
			err = errors.New("文件不存在 " + files[i])
		}
	}
	return files, err
}

func (f *Frontend) OpenDirectoryDialog(dialogOptions frontend.OpenDialogOptions) (string, error) {
	return "", nil
}

func (f *Frontend) SaveFileDialog(dialogOptions frontend.SaveDialogOptions) (string, error) {
	return "", nil
}

func (f *Frontend) MessageDialog(dialogOptions frontend.MessageDialogOptions) (string, error) {
	return "", nil
}

func (f *Frontend) ClipboardGetText() (string, error) {
	var text string
	return text, nil
}

func (f *Frontend) ClipboardSetText(text string) error {
	return nil
}

func (f *Frontend) WindowClose() {

}

func NewFrontend(ctx context.Context, appoptions *options.App, myLogger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) *Frontend {
	result := &Frontend{
		frontendOptions: appoptions,
		logger:          myLogger,
		bindings:        appBindings,
		dispatcher:      dispatcher,
		exitChan:        make(chan bool),
		ctx:             ctx,
	}
	result.startURL, _ = url.Parse(startURL)

	if _starturl, _ := ctx.Value("starturl").(*url.URL); _starturl != nil {
		result.startURL = _starturl
	} else {
		if port, _ := ctx.Value("assetserverport").(string); port != "" {
			result.startURL.Host = net.JoinHostPort(result.startURL.Host+".localhost", port)
		}

		var bindings string
		var err error
		if _obfuscated, _ := ctx.Value("obfuscated").(bool); !_obfuscated {
			bindings, err = appBindings.ToJSON()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			appBindings.DB().UpdateObfuscatedCallMap()
		}
		assets, err := assetserver.NewAssetServerMainPage(bindings, appoptions, ctx.Value("assetdir") != nil, myLogger, wailsruntime.RuntimeAssetsBundle)
		if err != nil {
			log.Fatal(err)
		}
		result.assets = assets

		go result.startRequestProcessor()
	}

	go result.startMessageProcessor()

	var _debug = ctx.Value("debug")
	var _devtoolsEnabled = ctx.Value("devtoolsEnabled")

	if _debug != nil {
		result.debug = _debug.(bool)
	}
	if _devtoolsEnabled != nil {
		result.devtoolsEnabled = _devtoolsEnabled.(bool)
	}

	go result.startSecondInstanceProcessor()

	return result
}

func (f *Frontend) startMessageProcessor() {
	for message := range messageBuffer {
		f.processMessage(message)
	}
}

func (f *Frontend) WindowReload() {
	f.ExecJS("runtime.WindowReload();")
}

func (f *Frontend) WindowSetSystemDefaultTheme() {
	return
}

func (f *Frontend) WindowSetLightTheme() {
	return
}

func (f *Frontend) WindowSetDarkTheme() {
	return
}

func (f *Frontend) Run(ctx context.Context) error {
	f.ctx = ctx
	go func() {
		if f.frontendOptions.OnStartup != nil {
			f.frontendOptions.OnStartup(f.ctx)
		}
	}()

	return nil
}

func (f *Frontend) WindowCenter() {

}

func (f *Frontend) WindowSetAlwaysOnTop(b bool) {

}

func (f *Frontend) WindowSetPosition(x, y int) {

}
func (f *Frontend) WindowGetPosition() (int, int) {
	return 0, 0
}

func (f *Frontend) WindowSetSize(width, height int) {

}

func (f *Frontend) WindowGetSize() (int, int) {
	return 0, 0
}

func (f *Frontend) WindowSetTitle(title string) {

}

func (f *Frontend) WindowFullscreen() {
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = false;")
	}

}

func (f *Frontend) WindowUnfullscreen() {
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = true;")
	}

}

func (f *Frontend) WindowReloadApp() {
	f.ExecJS(fmt.Sprintf("window.location.href = '%s';", f.startURL))
}

func (f *Frontend) WindowShow() {

}

func (f *Frontend) WindowHide() {

}

func (f *Frontend) Show() {

}

func (f *Frontend) Hide() {

}
func (f *Frontend) WindowMaximise() {

}
func (f *Frontend) WindowToggleMaximise() {

}
func (f *Frontend) WindowUnmaximise() {

}
func (f *Frontend) WindowMinimise() {

}
func (f *Frontend) WindowUnminimise() {

}

func (f *Frontend) WindowSetMinSize(width int, height int) {

}
func (f *Frontend) WindowSetMaxSize(width int, height int) {

}

func (f *Frontend) WindowSetBackgroundColour(col *options.RGBA) {
	if col == nil {
		return
	}
}

func (f *Frontend) ScreenGetAll() ([]Screen, error) {
	return nil, nil
}

func (f *Frontend) WindowIsMaximised() bool {
	return false
}

func (f *Frontend) WindowIsMinimised() bool {
	return false
}

func (f *Frontend) WindowIsNormal() bool {
	return false
}

func (f *Frontend) WindowIsFullscreen() bool {
	return false
}

func (f *Frontend) Quit() {
	f.exitChan <- true
}

func (f *Frontend) WindowPrint() {
	f.ExecJS("window.print();")
}

type EventNotify struct {
	Name string        `json:"name"`
	Data []interface{} `json:"data"`
}

func (f *Frontend) Notify(name string, data ...interface{}) {
	notification := EventNotify{
		Name: name,
		Data: data,
	}
	payload, err := json.Marshal(notification)
	if err != nil {
		f.logger.Error(err.Error())
		return
	}
	log.Println(`window.wails.EventsNotify('` + template.JSEscapeString(string(payload)) + `');`)
	// f.mainWindow.ExecJS(`window.wails.EventsNotify('` + template.JSEscapeString(string(payload)) + `');`)
}

func (f *Frontend) processMessage(message string) {
	// if message == "DomReady" {
	// 	if f.frontendOptions.OnDomReady != nil {
	// 		f.frontendOptions.OnDomReady(f.ctx)
	// 	}
	// 	return
	// }

	// if message == "drag" {
	// 	if !f.mainWindow.IsFullScreen() {
	// 		f.startDrag()
	// 	}
	// 	return
	// }

	// if message == "wails:showInspector" {
	// 	f.mainWindow.ShowInspector()
	// 	return
	// }

	// if strings.HasPrefix(message, "resize:") {
	// 	if !f.mainWindow.IsFullScreen() {
	// 		sl := strings.Split(message, ":")
	// 		if len(sl) != 2 {
	// 			f.logger.Info("Unknown message returned from dispatcher: %+v", message)
	// 			return
	// 		}
	// 		edge := edgeMap[sl[1]]
	// 		err := f.startResize(edge)
	// 		if err != nil {
	// 			f.logger.Error(err.Error())
	// 		}
	// 	}
	// 	return
	// }

	// if message == "runtime:ready" {
	// 	cmd := fmt.Sprintf(
	// 		"window.wails.setCSSDragProperties('%s', '%s');\n"+
	// 			"window.wails.flags.deferDragToMouseMove = true;", f.frontendOptions.CSSDragProperty, f.frontendOptions.CSSDragValue)
	// 	f.ExecJS(cmd)

	// 	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
	// 		f.ExecJS("window.wails.flags.enableResize = true;")
	// 	}
	// 	return
	// }

	// go func() {
	// 	result, err := f.dispatcher.ProcessMessage(message, f)
	// 	if err != nil {
	// 		f.logger.Error(err.Error())
	// 		f.Callback(result)
	// 		return
	// 	}
	// 	if result == "" {
	// 		return
	// 	}

	// 	switch result[0] {
	// 	case 'c':
	// 		// Callback from a method call
	// 		f.Callback(result[1:])
	// 	default:
	// 		f.logger.Info("Unknown message returned from dispatcher: %+v", result)
	// 	}
	// }()
}

func (f *Frontend) Callback(message string) {
	escaped, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	f.ExecJS(`window.wails.Callback(` + string(escaped) + `);`)
}

func (f *Frontend) startDrag() {
	// f.mainWindow.StartDrag()
}

func (f *Frontend) startResize(edge uintptr) error {
	// f.mainWindow.StartResize(edge)
	return nil
}

func (f *Frontend) ExecJS(js string) {
	// f.mainWindow.ExecJS(js)
}

var messageBuffer = make(chan string, 100)

// //export processMessage
// func processMessage(message *C.char) {
// 	goMessage := C.GoString(message)
// 	messageBuffer <- goMessage
// }

var requestBuffer = make(chan webview.Request, 100)

func (f *Frontend) startRequestProcessor() {
	for request := range requestBuffer {
		f.assets.ServeWebViewRequest(request)
	}
}

//export processURLRequest
func processURLRequest(request unsafe.Pointer) {
	// requestBuffer <- webview.NewRequest(request)
}

func (f *Frontend) startSecondInstanceProcessor() {
	// for secondInstanceData := range secondInstanceBuffer {
	// 	if f.frontendOptions.SingleInstanceLock != nil &&
	// 		f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch != nil {
	// 		f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch(secondInstanceData)
	// 	}
	// }
}
