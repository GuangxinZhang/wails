package server

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/samber/lo"
	"github.com/wailsapp/wails/v2/cmd/wails/flags"
	"github.com/wailsapp/wails/v2/cmd/wails/internal/gomod"
	"github.com/wailsapp/wails/v2/cmd/wails/internal/logutils"

	"github.com/wailsapp/wails/v2/pkg/commands/buildtags"

	"github.com/google/shlex"

	"github.com/pkg/browser"

	"github.com/wailsapp/wails/v2/internal/fs"
	"github.com/wailsapp/wails/v2/internal/process"
	"github.com/wailsapp/wails/v2/pkg/clilogger"
	"github.com/wailsapp/wails/v2/pkg/commands/build"
)

// Application runs the application in dev mode
func Application(f *flags.Dev, logger *clilogger.CLILogger) error {
	cwd := lo.Must(os.Getwd())

	// Update go.mod to use current wails version
	err := gomod.SyncGoMod(logger, !f.NoSyncGoMod)
	if err != nil {
		return err
	}

	// Run go mod tidy to ensure we're up-to-date
	err = runCommand(cwd, false, "go", "mod", "tidy")
	if err != nil {
		return err
	}

	buildOptions := f.GenerateBuildOptions()
	buildOptions.Logger = logger

	userTags, err := buildtags.Parse(f.Tags)
	if err != nil {
		return err
	}

	buildOptions.UserTags = userTags

	// Setup signal handler
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	exitCodeChannel := make(chan int, 1)

	// Build the frontend if requested, but ignore building the application itself.
	ignoreFrontend := buildOptions.IgnoreFrontend
	if !ignoreFrontend {
		buildOptions.IgnoreApplication = true
		if _, err := build.Build(buildOptions); err != nil {
			return err
		}
		buildOptions.IgnoreApplication = false
	}

	legacyUseDevServerInsteadofCustomScheme := false

	// Do initial build but only for the application.
	logger.Println("Building application for development...")
	buildOptions.IgnoreFrontend = true
	debugBinaryProcess, appBinary, err := restartApp(buildOptions, nil, f, exitCodeChannel, legacyUseDevServerInsteadofCustomScheme)
	buildOptions.IgnoreFrontend = ignoreFrontend || f.FrontendDevServerURL != ""
	if err != nil {
		return err
	}
	defer func() {
		if err := killProcessAndCleanupBinary(debugBinaryProcess, appBinary); err != nil {
			logutils.LogDarkYellow("Unable to kill process and cleanup binary: %s", err)
		}
	}()

	// open browser
	if f.Browser {
		err = browser.OpenURL(f.DevServerURL().String())
		if err != nil {
			return err
		}
	}

	logutils.LogGreen("Using DevServer URL: %s", f.DevServerURL())
	if f.FrontendDevServerURL != "" {
		logutils.LogGreen("Using Frontend DevServer URL: %s", f.FrontendDevServerURL)
	}
	logutils.LogGreen("Using reload debounce setting of %d milliseconds", f.Debounce)

	// Show dev server URL in terminal after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		logutils.LogGreen("\n\nTo develop in the browser and call your bound Go methods from Javascript, navigate to: %s", f.DevServerURL())
	}()

	// Kill the current program if running and remove dev binary
	if err := killProcessAndCleanupBinary(debugBinaryProcess, appBinary); err != nil {
		return err
	}

	// Reset the process and the binary so defer knows about it and is a nop.
	debugBinaryProcess = nil
	appBinary = ""

	logutils.LogGreen("Development mode exited")

	return nil
}

func killProcessAndCleanupBinary(process *process.Process, binary string) error {
	if process != nil && process.Running {
		if err := process.Kill(); err != nil {
			return err
		}
	}

	if binary != "" {
		err := os.Remove(binary)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	return nil
}

func runCommand(dir string, exitOnError bool, command string, args ...string) error {
	logutils.LogGreen("Executing: " + command + " " + strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(string(output))
		println(err.Error())
		if exitOnError {
			os.Exit(1)
		}
		return err
	}
	return nil
}

// restartApp does the actual rebuilding of the application when files change
func restartApp(buildOptions *build.Options, debugBinaryProcess *process.Process, f *flags.Dev, exitCodeChannel chan int, legacyUseDevServerInsteadofCustomScheme bool) (*process.Process, string, error) {
	appBinary, err := build.Build(buildOptions)
	println()
	if err != nil {
		logutils.LogRed("Build error - " + err.Error())

		msg := "Continuing to run current version"
		if debugBinaryProcess == nil {
			msg = "No version running, build will be retriggered as soon as changes have been detected"
		}
		logutils.LogDarkYellow(msg)
		return nil, "", nil
	}

	// Kill existing binary if need be
	if debugBinaryProcess != nil {
		killError := debugBinaryProcess.Kill()

		if killError != nil {
			buildOptions.Logger.Fatal("Unable to kill debug binary (PID: %d)!", debugBinaryProcess.PID())
		}

		debugBinaryProcess = nil
	}

	// parse appargs if any
	args, err := shlex.Split(f.AppArgs)
	if err != nil {
		buildOptions.Logger.Fatal("Unable to parse appargs: %s", err.Error())
	}

	// Set environment variables accordingly
	os.Setenv("loglevel", f.LogLevel)
	os.Setenv("assetdir", f.AssetDir)
	os.Setenv("devserver", f.DevServer)
	os.Setenv("frontenddevserverurl", f.FrontendDevServerURL)

	// Start up new binary with correct args
	newProcess := process.NewProcess(appBinary, args...)
	err = newProcess.Start(exitCodeChannel)
	if err != nil {
		// Remove binary
		if fs.FileExists(appBinary) {
			deleteError := fs.DeleteFile(appBinary)
			if deleteError != nil {
				buildOptions.Logger.Fatal("Unable to delete app binary: " + appBinary)
			}
		}
		buildOptions.Logger.Fatal("Unable to start application: %s", err.Error())
	}

	return newProcess, appBinary, nil
}
