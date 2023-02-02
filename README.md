# emojiterm
A simple cli tool that encodes words into emojis and echos them to the terminal. use this in your scripts.

![using emojiterm](vhs/emojiterm-running.gif)


## Usage

```shell
./emojiterm beer sushi bento
 🍺  🍣  🍱
```

or 
```shell
# on mac to copy into clipboard
./emojiterm beer sushi bento | pbcopy
```

or try to help yourself
```shell
emojiterm help                                                                  
NAME:
   emojiterm - text to emoji converter - Use underscore instead of space in emoji names - emojiterm "<space separeted list of tokens>" | Example: emojiterm "sun wine_glass beer)"
                                                                                          Names are based on this file: https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji.json

USAGE:
   emojiterm - text to emoji converter - Use underscore instead of space in emoji names [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
