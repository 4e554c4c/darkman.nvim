## 🚧 Under Construction 🚧

This plugin is not yet complete. Please use with caution.

# Darkman.nvim

**Darkman.nvim** is a neovim plugin written in Golang, designed to interface
with the Freedesktop Dark-mode standard. It is named after and integrates with
[darkman](https://darkman.whynothugo.nl/), a daemon which implements this
standard. However, the plugin should work on all compliant desktop environments
(including Gnome and KDE).

## Requirements

This plugin is currently in development, and no versions have been released.
Thus, to compile darkman.nvim a go compiler is required.

## Installation

Using [packer.nvim](https://github.com/wbthomason/packer.nvim)
```lua
use {
  '4e554c4c/darkman.nvim',
  run = 'go build -o bin/darkman.nvim',
  config = function()
    require 'darkman'.setup()
  end,
}
```

## Configuration

TODO lol
