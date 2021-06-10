# go-wallpaper

go-wallpaper is a cli tool for setting random wallpapers for Windows. It is built with Golang and inspired by [styli.sh](https://github.com/thevinter/styli.sh)

## how it works

![go-wallpaper.gif](./go-wallpaper.gif)

## features

- currently it downloads images from [Unsplash](https://unsplash.com)
- it can take resolution as a flag e.g., `go-wallpaper.exe -r 1900x1080`
- it can take a search term/tag as a flag e.g., `go-wallpaper.exe -t sea,sunset`
- `go-wallpaper.exe --help` for help

## how to use

- If you already have Golang installed in your system, you can clone this repo and run the following commands
  - `go mod tidy` to install dependencies
  - `go install` to add the binary to your gopath
- Or you can download the pre-built binaries from releases sections and add it to your path (or somewhere you can use it from terminal, cmd, powershell etc...)

## features to come

- [ ] Set wallpapers from subreddits like `r/wallpapers` etc...
- [ ] Flags for saving the wallpaper permanently
- [ ] Writing tests and maybe publish it as package for others to use

## contribution

the project is a mess currently 😓, I will clean it, and make it easy for contribution, so until then I can not accept any contributions
