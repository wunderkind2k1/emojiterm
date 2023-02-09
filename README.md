# emojiterm
A simple cli tool that encodes words into emojis and echos them to the terminal. use this in your scripts.

![using emojiterm](vhs/emojiterm-running.gif)


## Installation

### With homebrew

```shell
brew tap wunderkind2k1/tap
brew install emojiterm
```

![installing emojiterm](vhs/emojiterm-installation.gif)

### Go standard: Through _go install_

```shell
go install github.com/wunderkind2k1/emojiterm@latest
```

## Usage

```shell
emojiterm beer sushi bento
 🍺  🍣  🍱
```

__or (since v.0.0.2)__
```shell
# to copy into clipboard. Needs pbcopy or xclip installed though
emojiterm -c beer sushi bento
```

Of course pipes are going to work like before
```shell
# use pipes and don't eat sushi - note: -c is optional
emojiterm -c beer bento sushi | cut -c 1,4
> 🍺🍱
# your clipboard still has the sushi
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
   --clipboard, -c  write emojis to clipboard
   --help, -h  show help
```
