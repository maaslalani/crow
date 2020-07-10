package config

// IgnoredPaths is a list of paths to never watch for changes
var IgnoredPaths = []string{
	".git",
}

// Clear defines whether to clear the screen before running the next command
const Clear = true

// Restart defines whether to cancel the current running proccess before running again
const Restart = true
