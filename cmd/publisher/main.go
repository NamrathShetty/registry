package main

import (
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/registry/cmd/publisher/commands"
)

// Version info for the MCP Publisher tool
// These variables are injected at build time via ldflags by goreleaser
var (
	// Version is the current version of the MCP Publisher tool
	Version = "dev"

	// BuildTime is the time at which the binary was built
	BuildTime = "unknown"

	// GitCommit is the git commit that was compiled
	GitCommit = "unknown"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	// Handle global help/version
	if cmd == "--version" || cmd == "-v" || cmd == "version" {
		log.Printf("mcp-publisher %s (commit: %s, built: %s)", Version, GitCommit, BuildTime)
		return
	}
	if cmd == "--help" || cmd == "-h" || cmd == "help" {
		printUsage()
		return
	}

	// Print command-specific help if --help or -h or help is the first argument to a subcommand
	if len(args) > 0 && (args[0] == "--help" || args[0] == "-h" || args[0] == "help") {
		printCommandHelp(cmd)
		return
	}

	var err error
	switch cmd {
	case "init":
		err = commands.InitCommand()
	case "login":
		err = commands.LoginCommand(args)
	case "logout":
		err = commands.LogoutCommand()
	case "publish":
		err = commands.PublishCommand(args)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printCommandHelp(command string) {
	switch command {
	case "init":
		_, _ = fmt.Fprintln(os.Stdout, "Usage: mcp-publisher init")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Description:")
		_, _ = fmt.Fprintln(os.Stdout, "  Create a server.json file template in the current directory.")
		_, _ = fmt.Fprintln(os.Stdout, "  This template includes all required and optional fields for")
		_, _ = fmt.Fprintln(os.Stdout, "  publishing your MCP server to the registry.")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Example:")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher init")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "The generated server.json file should be customized with your")
		_, _ = fmt.Fprintln(os.Stdout, "server's specific details before publishing.")
	case "login":
		_, _ = fmt.Fprintln(os.Stdout, "Usage: mcp-publisher login <method>")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Description:")
		_, _ = fmt.Fprintln(os.Stdout, "  Authenticate with the registry using one of the supported")
		_, _ = fmt.Fprintln(os.Stdout, "  authentication methods. Authentication is required before")
		_, _ = fmt.Fprintln(os.Stdout, "  you can publish servers to the registry.")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Available Methods:")
		_, _ = fmt.Fprintln(os.Stdout, "  github-at        - GitHub Personal Access Token")
		_, _ = fmt.Fprintln(os.Stdout, "  github-oidc      - GitHub OIDC (for CI/CD environments)")
		_, _ = fmt.Fprintln(os.Stdout, "  http             - HTTP-based authentication")
		_, _ = fmt.Fprintln(os.Stdout, "  dns              - DNS-based authentication")
		_, _ = fmt.Fprintln(os.Stdout, "  none             - No authentication (for testing)")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Example:")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher login github-at")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Credentials are securely stored locally for subsequent publishes.")
	case "logout":
		_, _ = fmt.Fprintln(os.Stdout, "Usage: mcp-publisher logout")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Description:")
		_, _ = fmt.Fprintln(os.Stdout, "  Clear saved authentication credentials from the local system.")
		_, _ = fmt.Fprintln(os.Stdout, "  This removes all stored authentication tokens and sessions.")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Example:")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher logout")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "After logging out, you will need to run 'mcp-publisher login'")
		_, _ = fmt.Fprintln(os.Stdout, "again before publishing.")
	case "publish":
		_, _ = fmt.Fprintln(os.Stdout, "Usage: mcp-publisher publish [options]")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Description:")
		_, _ = fmt.Fprintln(os.Stdout, "  Publish your MCP server to the registry using the server.json")
		_, _ = fmt.Fprintln(os.Stdout, "  file in the current directory. You must be authenticated before")
		_, _ = fmt.Fprintln(os.Stdout, "  publishing (see 'mcp-publisher login').")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Options:")
		_, _ = fmt.Fprintln(os.Stdout, "  --file <path>    - Specify a custom path to server.json")
		_, _ = fmt.Fprintln(os.Stdout, "  --dry-run        - Validate the server.json without publishing")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Example:")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher publish")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher publish --file ./my-server.json")
		_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher publish --dry-run")
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "The server.json file must conform to the registry schema.")
		_, _ = fmt.Fprintln(os.Stdout, "Run 'mcp-publisher init' to generate a template.")
	default:
		printUsage()
	}
}

func printUsage() {
	_, _ = fmt.Fprintln(os.Stdout, "MCP Registry Publisher Tool")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = fmt.Fprintln(os.Stdout, "Usage:")
	_, _ = fmt.Fprintln(os.Stdout, "  mcp-publisher <command> [arguments]")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = fmt.Fprintln(os.Stdout, "Commands:")
	_, _ = fmt.Fprintln(os.Stdout, "  init          Create a server.json file template")
	_, _ = fmt.Fprintln(os.Stdout, "  login         Authenticate with the registry")
	_, _ = fmt.Fprintln(os.Stdout, "  logout        Clear saved authentication")
	_, _ = fmt.Fprintln(os.Stdout, "  publish       Publish server.json to the registry")
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = fmt.Fprintln(os.Stdout, "Use 'mcp-publisher <command> --help' for more information about a command.")
}
