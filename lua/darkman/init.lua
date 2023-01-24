local M = {}

function M.setup()
  local projectPath = debug.getinfo(1,"S").source:match("^@?(.-)/lua/darkman/init.lua")
  local binPath = projectPath .. '/bin/darkman.nvim'

  local function require_darkman()
    return vim.fn.jobstart({binPath}, {rpc=true})
  end
  vim.fn['remote#host#Register']('darkman.nvim', '0', require_darkman)

  -- output from `./bin/darkman.nvim -manifest darkman.nvim`
  vim.cmd[[
  call remote#host#RegisterPlugin('darkman.nvim', '0', [
  \ {'type': 'function', 'name': 'Hello', 'sync': 1, 'opts': {}},
  \ ])
  ]]
end

return M
