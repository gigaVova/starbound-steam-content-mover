# starbound-steam-content-mover

## How can I make use of it?

If you need to move `.pak` files from Steam's workshop folder to the native Starbound's mods folder, this app will do exactly this

## Command

To run this app follow this steps:

1. Open system terminal in the folder you have the executable file
2. Run one of the following commands:

This will create a `./content` folder

```bash
./starbound-steam-content-mover.exe -src %steamFolder%/steamapps/workshop/content/211820
```

Also you can specify the target directory

```bash
./starbound-steam-content-mover.exe -src %steamFolder%/steamapps/workshop/content/211820 -%userprofile%/Downloads
```
