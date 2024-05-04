return {
  'windwp/nvim-autopairs',
  event = 'InsertEnter',
  config = true,
  -- use opts = {} for passing setup options
  -- this is equalent to setup({}) function
  keys = {
    {
      '<leader>up',
      function()
        vim.g.autopairs_disabled = not vim.g.autopairs_disabled
        if vim.g.autopairs_disabled then
          require('nvim-autopairs').disable()
          vim.notify('Disabled auto pairs', 'info')
        else
          require('nvim-autopairs').enable()
          vim.notify('Enabled auto pairs', 'info')
        end
      end,
    },
  },
}
