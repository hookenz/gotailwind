# GoTailwind
A wrapper for the excellent [TailwindCSS](https://tailwindcss.com/) cli binary to facilitate it's use with golang as a go tool plugin.

# Latest release
[![latest version](https://img.shields.io/github/v/tag/hookenz/gotailwind?color=%2344cc11&label=Latest%20release&style=for-the-badge)](https://github.com/hookenz/gotailwind/releases/latest)

https://github.com/user-attachments/assets/f540a315-a70b-49f0-843a-f35e7520d5d5

# About

I created this tool to make it easier to use tailwindcss within golang projects such as with those 
using the go [templ](https://templ.guide/) template engine.

It basically downloads the tailwindcss cli standalone binary.  Nodejs is not required.

Note: requires go 1.24+

# Usage

To install the latest version of tailwindcss 4:
```
go get -tool github.com/hookenz/gotailwind/v4@latest
```

Or a specific version of tailwindcss:
```
go get github.com/hookenz/gotailwind/v4@v4.1.7
```
To run it:
```
go tool gotailwind
```

# How does it work? 
This tool is a thin go wrapper around the standalone tailwindcss binary.  It's just a go program that calls the appropriate tailwindcss 
binary that is downloaded and cached into a special folder.

The tagged version corresponds to the tailwindcss binary.

It has been tested under Linux.  It should work under mac and windows although I haven't tested it.

The version downloaded by this tool match the tailwindcss binary.
In linux they are placed into the versioned directories beneath
```
~/.cache/gotailwind/
```

i.e.  
```
~/.cache/gotailwind/v4.1.6/tailwindcss-linux-x64
```

# Contributing
Suggestions or improvements welcome.



