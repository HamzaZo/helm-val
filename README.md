[![Licence](https://img.shields.io/badge/licence-Apache%202.0-green)]()
[![Helm](https://img.shields.io/badge/release-0.1.0-brightgreen)]()

# helm-val

`helm-val` is a helm plugin to fetch values from a previous release.

## Getting started

### Installation

To install the plugin:
```shell
$ helm plugin install https://github.com/HamzaZo/helm-val
```
Update to latest
```shell
$ helm plugin update val
```
Install a specific version
```shell
$ helm plugin install https://github.com/HamzaZo/helm-val --version 0.3.0
```
You can also verify it's been installed using
```shell
$ helm plugin list
```

### Usage

```
$ helm val 
The val plugin helps you to get values from an old release

Usage:
  val [command]

Available Commands:
  fetch       Fetch values from previous release
  help        Help about any command

Flags:
  -h, --help   help for val

Use "val [command] --help" for more information about a command.


```

#### fetch values

```
$ helm val fetch -h
Fetch values from previous release with(or without) specific revision using '--revision'.

Examples:
        
    $ helm val fetch RELEASE-NAME -r/--revision 1

    $ helm val fetch RELEASE-NAME -n/--namespace <ns>
 
    $ helm val fetch RELEASE-NAME -c/--kube-context <ctx>
 
    $ helm val fetch RELEASE-NAME -k/--kubeconfig <kcfg>

Usage:
  val fetch RELEASE-NAME [flags]

Flags:
  -h, --help                  help for fetch
  -c, --kube-context string   name of the kubeconfig context to use
  -k, --kubeconfig string     path to the kubeconfig file
  -n, --namespace string      namespace scope for this request
  -o, --output format         prints the outputs in the specified fomat. Allowed format: table,json,yaml (default table)
  -r, --revision int          number of revision release

```
**Note:** In case you don't specify a revision it'll automatically fetch values of the preceding release.