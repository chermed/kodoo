# Overview

**Terminal UI for Odoo**

The `kodoo` tool is a Terminal UI for Odoo that allow to show and query the data, it's destined for Odoo developers and for end users.


[![asciicast](https://asciinema.org/a/430567.svg)](https://asciinema.org/a/430567)

# Installation :

## Visit the last release page

The [release page](https://github.com/chermed/kodoo/releases/latest) contains all compiled binaries of this tool

## Via script for macOS and Linux

```
curl -sSL https://raw.githubusercontent.com/chermed/kodoo/main/install.sh | sh
```

## Via Homebrew for macOS or LinuxBrew for Linux

```
brew tap chermed/kodoo
brew install chermed/kodoo/kodoo
```

### Upgrade :

```
brew upgrade chermed/kodoo/kodoo
```

## Via a GO install

```
go get -u github.com/chermed/kodoo
```

## Using snap

```
snap install kodoo --channel=beta
```

## For Windows

Get the executable from the [release page](https://github.com/chermed/kodoo/releases/latest)

## Using docker

The image is `docker.io/chermed/kodoo`

```
docker run -it --rm -v $(pwd):/.kodoo --net host chermed/kodoo:latest init-config
```

Edit the generated file, then run:

```
docker run -it --rm -v $(pwd):/.kodoo --net host chermed/kodoo:latest
```

# Get started

1. Install the command following the instructions
2. Initialize the configuration using the command `kodoo init-config` 
3. Edit the configuration (only the list of servers is mandatory)
4. Run the command `kodoo`
5. Type `?` to see the help page, and `ESC` to go back to the main page
6. Start to query the data from a database


# Features

1. Switch between many Odoo servers
2. Manage and run macros
3. Query objects and automatic refresh
4. Pagination support
5. Quick access to related records
6. Auto detection of fields to show as columns in the table
7. Run remote function on a selected records
8. Sort and filter records
9. Show metadata and details of a record
10. Change dynamically the database or the user to use

# Use cases 

1. For developers :
   1. Provide a fast way to see technical data (IDs and metadata) of records.
   2. Query and show invisible objects and invisible fields
2. For Odoo customers :
   1. Use the Zen mode to display data on a hanging screen (Kitchen for a restaurant, Work orders for the manufacturing, etc)
   2. Use macros to access easily to custom views and data

# Limitations :

1. Filtering data is basic:
   1. The values in the domain are sent to odoo as strings or list of strings
   2. If many filters are specified the logical operator is `AND`
2. Binary fields value will not be shown or downloaded


## Thanks

Thanks to [derailed](https://github.com/derailed) for his awesome [k9s tool](https://github.com/derailed/k9s), it gave me the idea to build this tool.