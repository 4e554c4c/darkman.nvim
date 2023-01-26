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

`setup` takes a dictionary of the following values (and defaults)
```lua
{
  change_background = true,
  send_user_event = false,
  colorscheme = nil, -- can be { dark = "x, light = "y" }
}
```

If `change_background` is true, `background` will be automatically set to
`light` or `dark`.
Please note that you can add extra functionality by listening to the `OptionSet
background` autocmd event.

If you would not like darkman.nvim to set `background`, you may set
`send_user_event=true`. In which case the the `User Darkmode` or `UserLightmode`
events will be triggered instead.

If the `colorscheme` option is set to a table with `dark` and `light` keys, the
colorschemes given will be set automatically.
