# gotailwind
A wrapper for tailwindcss in order to use tailwindcss as a go tool

I created this tool to make it easier to use tailwind with golang projects.

It basically downloads the tailwindcss cli standalone binary.  Nodejs is not required.

Usage:
   github.com/hookenz/gotailwind/v4@v4.1.4

Note: I made a mistake.  You might need to put GOPROXY=direct
  i.e. GOPROXY=direct github.com/hookenz/gotailwind/v4@v4.1.4

Unfortunately the public go module cache will forever hold some broken versions.  
I've made this mistake with 4.1.4 and 4.1.5.  Going forward it should be fixed though.

The versions should ideally match the tailwindcss binary.
