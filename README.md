# Azure Functions Golang with scheduled trigger
## Prerequisites
- [Azure Functions Core Tools, v4](https://www.npmjs.com/package/azure-functions-core-tools)
- [Go SDK](https://golang.org/dl/)
- [Visual Studio Code](https://code.visualstudio.com) & [Azure Functions extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurefunctions)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) [optional]

> Mac or Linux machine, or Windows WSL2 are preferred.

## Local start/debug

First build the code by running: `go build handler.go weather.go fetch.go`. Then start the function with `func start`

## Debugging Go code

For breaking points debugging in VSCode a number of approaches can be used: attaching to process or connecting to server. In most cases, one would prefer restarting Go server without stoping and restarting `func start` session. Also, it's simpler to place a breakpoint, switch to `main.go` file, and launch debug. Such a scenario can be covered with:
