# Overview

**Terminal UI for Odoo**

This application is calling the Odoo API to scrape the data, it makes an effort to show them on terminal with the ability of the navigation, it's destined for Odoo developers and customers, the first release is tested using the version Odoo 14, the next version will be more compatible with many other versions >= 12.0.


[![asciicast](https://asciinema.org/a/430567.svg)](https://asciinema.org/a/430567)

# Installation :

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


## Using docker

`docker pull chermed/kodoo`

```
docker run -it --rm -v $(pwd):/.kodoo --net host chermed/kodoo:latest init-config

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
   1. Provide a fast way to see IDs and metadata of records.
   2. Query very easily the data on invisible objects and invisible fields
2. For Odoo customers :
   1. Use the Zen mode to display data on a hanging screen (Kitchen for a restaurant, Work orders for the manufacturing, etc)

# Limitations :

1. Filtering data is basic:
   1. The values in the domain are sent to odoo as strings or list of strings
   2. If many filters are specified the logical operator is `AND`
2. Binary and reference fields value will not be shown or downloaded


## Roadmap for the the version v0.3.0

The features to be introduced are :

1.  Ensure compatibility with more Odoo versions
2.  Add the zen mode (focus on data)  
3.  Brew install
4.  Snap Install
