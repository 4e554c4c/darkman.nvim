# Darkman.nvim

**Darkman.nvim** is a neovim plugin written in Golang, designed to interface
with the Freedesktop Dark-mode standard. It is named after and integrates with
[darkman](https://darkman.whynothugo.nl/), a daemon which implements this
standard. However, the plugin should work on all compliant desktop environments
(including Gnome and KDE).


## Maintenance status

darkman.nvim is relatively stable, so updates do not come regularly. There is no
need to assume that the repository is unmaintained if an update hasn't occurred
for several months. This section will change if the maintenance status for
darkman.nvim changes.

## Requirements

To compile darkman.nvim a go compiler is required due to
4e554c4c/darkman.nvim#1. We are working on changing this, but it is currently
impossible due to the architecture of common neovim package managers.

## Installation

Using [lazy.nvim](https://github.com/folke/lazy.nvim)
```lua
{
  '4e554c4c/darkman.nvim',
  event = 'VimEnter',
  build = 'go build -o bin/darkman.nvim',
  opts = {
    -- configuration here
  },
}
```


## Configuration

`setup` takes a dictionary of the following values (and defaults)
```lua
{
  change_background = true,
  send_user_event = false,
  colorscheme = nil, -- can be { dark = "x", light = "y" }
}
```

If `change_background` is true, `background` will be automatically set to
`light` or `dark`.
Please note that you can add extra functionality by listening to the `OptionSet
background` autocmd event.

If you would not like darkman.nvim to set `background`, you may set
`send_user_event=true`. In which case the `User DarkMode` or `User LightMode`
events will be triggered instead.

If the `colorscheme` option is set to a table with `dark` and `light` keys, the
colorschemes given will be set automatically.
