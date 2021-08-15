# Overview

**Terminal UI for Odoo**

This application is calling the Odoo API to scrape the data, it makes an effort to show them on terminal with the ability of the navigation, it's destined for Odoo developers and customers, the first release is tested using the version Odoo 14, the next version will be more compatible with many other versions >= 12.0.


[![asciicast](https://asciinema.org/a/430567.svg)](https://asciinema.org/a/430567)

# Installation :

Check the [release page](https://github.com/chermed/kodoo/releases)  

# Get started

1. Install the binary using the instructions on the release page
2. Initialize the configuration using the command `kodoo init-config` 
3. Edit the configuration (only the list of servers is mandatory)
4. Run the command `kodoo`
5. Start to query the Odoo API


# Features

1. Switch between many Odoo servers
2. Manage and run macros
3. Query objects and automatic refresh
4. Pagination support
5. Quick access to related records
6. Auto detection of fields to show as columns in the table
7. Run remote function on a selected record
8. Sort and filter records

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


## Roadmap for the the version v0.2.0

The features to be introduced are :

1. Show shortcuts for different relations
2. Call a remote function for many records (multiple selection)
3. See details of a record (show all fields values)
4. See metadata of a record
5. List and change the databases
6.  Ensure compatibility with more Odoo versions
7.  Add the readonly mode
8.  Add the zen mode (focus on data)    