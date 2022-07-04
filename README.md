# Alias
A tool for Windows cmd to make permanent aliases

## Usage
Download the latest version from the releases page. 
Put it in a directory and add that directory to the Windows PATH.

Example:
```
alias purge "rmdir /s /q" -a
```
This makes an alias "purge" that can be used to delete a directory and its contents.

## Arguments
### -a, --args
Enables arguments for the alias.

### -b, --bat
Tells the program to use the .bat extension instead of .cmd

### -e, --echo
Turns **echo** on for the alias.
