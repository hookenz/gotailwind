# gotailwind
A wrapper for the excellent [TailwindCSS](https://tailwindcss.com/) cli binary to facilitate it's use with golang as a go tool plugin.

I created this tool to make it easier to use tailwind within golang projects such as with those 
using the go [templ](https://templ.guide/) template engine

It basically downloads the tailwindcss cli standalone binary.  Nodejs is not required.

Usage:

To install the latest version of tailwindcss 4:
```
go get -tool github.com/hookenz/gotailwind/v4@latest
```

Or a specific version of tailwindcss:
```
go get github.com/hookenz/gotailwind/v4@v4.1.5
```


To run it:
```
go tool gotailwind
```

# Contributing
Suggestions or improvements welcome.

# Notes: 
Unfortunately I made a mistake with version 4.1.4 and 4.1.5 and I can't fix it until a 
new version of tailwindcss is released.  But there is a workaround.

i.e. 
```
GOSUMDB=off GOPROXY=direct go get github.com/hookenz/gotailwind/v4@v4.1.5
```

Unfortunately the public go module cache will forever hold these broken versions.
Later versions won't have this problem.

The version downloaded by this tool match the tailwindcss binary.
In linux they are placed into the versioned directories beneath ~/.cache/gotailwind/`v4.1.5/tailwindcss-linux-x64
e.g.  ~/.cache/gotailwind/v4.1.5/tailwindcss-linux-x64



