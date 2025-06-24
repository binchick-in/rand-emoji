# Rand Emoji

_Print a random emoji to terminal_

# Why?
This tool is mainly for use with my [tmux session manager script](https://github.com/binchick-in/dotfiles/blob/master/bin/tmux-env) to help add some character to my tmux windows.

# Installation

```
go install github.com/binchick-in/rand-emoji@latest
```

# Usage

Print a random emoji from the global list
```
$> rand-emoji
🥟
```

Print a random animal emoji
```
$> rand-emoji -animals
🦒
```

Here are all the options available.
```
$> rand-emoji -help
Usage of ./rand-emoji:
  -animals
        Use only animal emojis
  -ape
        Use only ape emojis
  -extras
        Use only extras emojis
  -food
        Use only food emojis
  -list
        List all available emojis
  -version
        Show version
```
