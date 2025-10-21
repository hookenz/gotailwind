# GoTailwind   [![latest version](https://img.shields.io/github/v/tag/hookenz/gotailwind?color=%2344cc11&label=Latest%20release&style=for-the-badge)](https://github.com/hookenz/gotailwind/releases/latest)

A `go tool` plugin for the excellent [TailwindCSS](https://tailwindcss.com/) cli.  

# About

I created this tool to make it easier to use TailwindCSS within go projects such as with those using the go [templ](https://templ.guide/) template engine or html/template etc.

It downloads the TailwindCSS cli standalone binary automatically on first use.  Subsequent invocations will use the cached copy.  Nodejs is not required.

# Prerequisites
Go 1.24+.  You must upgrade to go 1.24 or newer before using this tool.

Please verify with `go version` that the version you have is at least 1.24 or it won't work.

# Usage

To install the latest version of TailwindCSS v4 into an existing go project:
```
go get -tool github.com/hookenz/gotailwind/v4@latest
```

Or a specific version of TailwindCSS:
```
go get -tool github.com/hookenz/gotailwind/v4@v4.1.15
```

To run it:
```
go tool gotailwind
```

# Demo

https://github.com/user-attachments/assets/f540a315-a70b-49f0-843a-f35e7520d5d5


# How does it work? 
GoTailwind is a thin go wrapper around the standalone TailwindCSS cli.  It's just a go program that calls the appropriate TailwindCSS 
binary that is downloaded and cached into a local cache folder.

The GoTailwind version matches the TailwindCSS cli version.

It has been tested under Linux and Windows. It should also work under mac although is currently untested. 
Do let me know if it works or not.

In linux the TailwindCSS cli is placed into a versioned directory beneath:
```
~/.cache/gotailwind/
```

i.e.
```
~/.cache/gotailwind/v4.1.7/tailwindcss-linux-x64
```

Which means you can have different projects targetting different versions of TailwindCSS.

Note: Versions of this tool prior to 4.1.6 do not work properly due to bug and broken go cache.  

# Contributing
Suggestions or improvements are more than welcome.
