# DCI Manager

# ⚠ WIP ⚠

DCI Manager is a utility to create and manage your DCI support matrix


#### Environment Variables

| **variable**   | **description**                     | **default**   |
|----------------|-------------------------------------|---------------|
| DCIM_CONFIG    | Path to DCIM config file. yaml/json | `dcim.yaml`   |
| DCIM_STATE_DIR | Location of DCI state files         | `./.state/`   |
| DCIM_JOBS_DIR  | Location of DCI jobs                | `./jobs/`     |

### Available commands

#### `dcim`
```
Usage:
  dcim [command]

Available Commands:
  apply       Run diff on current state and see what will change. Approve to apply
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a config file
  help        Help about any command
  parse       Parse the given config file
  plan        Run diff on current state and see what will change
  state       Print the current matrix

Flags:
  -h, --help      help for dcim
  -v, --version   version for dcim

Use "dcim [command] --help" for more information about a command.
```

#### `parse`
```
Parse the given config file

Usage:
  dcim parse [flags]

Flags:
  -h, --help   help for parse

```

#### `generate`
```
Generate a config file

Usage:
  dcim generate <name> [flags]

Flags:
  -f, --format string   Output format. Can be json/yaml/yml. Defaults to yaml (default "yaml")
  -h, --help            help for generate

```

#### `state`
```
Print the current matrix

Usage:
  dcim state [flags]

Flags:
  -h, --help   help for state

```

#### `plan`
```
Run diff on current state and see what will change

Usage:
  dcim plan [flags]

Flags:
  -h, --help   help for plan
```

#### `apply`
```
Run diff on current state and see what will change. Approve to apply

Usage:
  dcim apply [flags]

Flags:
  -h, --help   help for apply
  -y, --yes    To auto approve
```

