# starbound-steam-content-mover

[![Go Report Card](https://goreportcard.com/badge/github.com/gigavova/starbound-steam-content-mover)](https://goreportcard.com/report/github.com/gigavova/starbound-steam-content-mover)

## How can I make use of it?

If you need to move `.pak` files from Steam's workshop folder to the native Starbound's mods folder, this app will do exactly this

## Commands

To run this app follow this steps:

1. Open system terminal in the folder you have the executable file
2. Run one of the following commands:

This will create a `./content` folder

```shell
./starbound-steam-content-mover.exe -src %steamFolder%/steamapps/workshop/content/211820
```

Also, you can specify the target directory

```shell
./starbound-steam-content-mover.exe -src %steamFolder%/steamapps/workshop/content/211820 -target %userprofile%/Downloads
```

Using aforementioned ways this app will create a folder with `.pak` files with
a number which is actually a Steam workshop item ID, if you haven't changed anything.

So if you want your `.pak` files to be called similar to its names in the Steam
workshop, so add `--titles` flag. Your command will look something like:

```shell
./starbound-steam-content-mover.exe -src %steamFolder%/steamapps/workshop/content/211820 -target %userprofile%/Downloads -titles
```

BUT! This will create files without extensions (unexpectedly). They can be deleted by you.

## Todo

- [ ] Find the source of extensionless files
- [ ] Tell about this app on a forum