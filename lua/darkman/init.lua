local M = {}

local loaded = false

local defaults = {
  change_background = true,
  send_user_event = false,
  -- can be { dark = x, light = y }
  colorscheme = nil,
}

function M.setup(opts)
  opts = vim.tbl_deep_extend("force", defaults, opts or {})

  if loaded then
    return
  end
  loaded = true

  -- `@` is the first character since this is a file
  local projectPath = debug.getinfo(1,"S").source:match("^@(.-)/lua/darkman/init.lua")
  assert(projectPath, "Unable to find location of current lua file") 
  local binPath = projectPath .. '/bin/darkman.nvim'

  local function require_darkman()
    return vim.fn.jobstart({binPath}, {rpc=true})
  end
  vim.fn['remote#host#Register']('darkman.nvim', '0', require_darkman)

  -- output from `./bin/darkman.nvim -manifest darkman.nvim`
  vim.cmd[[
    call remote#host#RegisterPlugin('darkman.nvim', '0', [
    \ {'type': 'function', 'name': 'DarkmanGetMode', 'sync': 1, 'opts': {}},
    \ {'type': 'function', 'name': 'DarkmanSetup', 'sync': 1, 'opts': {}},
    \ ])
  ]]

  -- now setup darkman
  vim.fn.DarkmanSetup(opts.change_background, opts.send_user_event, opts.colorscheme)
end

return M
