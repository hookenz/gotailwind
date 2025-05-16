# GoTailwind
A wrapper for the excellent [TailwindCSS](https://tailwindcss.com/) cli to facilitate it's use with golang projects as a `go tool` plugin.

[![latest version](https://img.shields.io/github/v/tag/hookenz/gotailwind?color=%2344cc11&label=Latest%20release&style=for-the-badge)](https://github.com/hookenz/gotailwind/releases/latest)

https://github.com/user-attachments/assets/f540a315-a70b-49f0-843a-f35e7520d5d5

# About

I created this tool to make it easier to use tailwindcss within golang projects such as with those using the go [templ](https://templ.guide/) template engine or html/template etc.

It downloads the tailwindcss cli standalone binary automatically on first use.  Nodejs is not required.

Note: this requires go 1.24+.  Upgrade to go 1.24 or newer before using this tool.

# Usage

To install the latest version of tailwindcss 4:
```
go get -tool github.com/hookenz/gotailwind/v4@latest
```

Or a specific version of tailwindcss:
```
go get -tool github.com/hookenz/gotailwind/v4@v4.1.7
```
To run it:
```
go tool gotailwind
```

# How does it work? 
This tool is a thin go wrapper around the standalone tailwindcss cli.  It's just a go program that calls the appropriate tailwindcss 
binary that is downloaded and cached into a local cache folder.

The tagged version corresponds to the version tailwindcss cli.

It has been tested under Linux and Windows. It should also work under mac although is currently untested. Do let me know if it works or not.

The version downloaded by this tool should match the tailwindcss cli.
In linux they are placed into the versioned directories beneath
```
~/.cache/gotailwind/
```

i.e.  
```
~/.cache/gotailwind/v4.1.7/tailwindcss-linux-x64
```

Which means you can have different projects targetting different versions of tailwind.

Note: Versions of this tool prior to 4.1.6 do not work properly due to bug and broken go cache.  

# Contributing
Suggestions or improvements are more than welcome.
