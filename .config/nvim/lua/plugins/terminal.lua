return {
  'akinsho/toggleterm.nvim',
  version = '*',
  config = function()
    local term = require 'toggleterm'
    term.setup()

    -- Toggle floating terminal and enter insert mode
    vim.keymap.set('n', '<leader>tf', function()
      term.toggle(1, 20, vim.fn.getcwd(), 'float')
      vim.cmd 'startinsert'
    end, { desc = 'Toggle floating terminal' })

    -- Toggle bottom terminal and enter insert mode
    vim.keymap.set('n', '<leader>tb', function()
      term.toggle(1, 20, vim.fn.getcwd(), 'horizontal')
      vim.cmd 'startinsert'
    end, { desc = 'Toggle bottom terminal' })

    function _G.set_terminal_keymaps()
      local opts = { buffer = 0 }
      vim.keymap.set('t', '<esc>', [[<C-\><C-n>]], opts)
      vim.keymap.set('t', '<C-x>', [[<C-\><C-n>ZZ]], opts)
      vim.keymap.set('t', '<C-h>', [[<Cmd>wincmd h<CR>]], opts)
      vim.keymap.set('t', '<C-j>', [[<Cmd>wincmd j<CR>]], opts)
      vim.keymap.set('t', '<C-k>', [[<Cmd>wincmd k<CR>]], opts)
      vim.keymap.set('t', '<C-l>', [[<Cmd>wincmd l<CR>]], opts)
    end

    vim.cmd 'autocmd! TermOpen term://* lua set_terminal_keymaps()'
  end,
}
