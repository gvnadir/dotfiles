return {
  { -- Adds git related signs to the gutter, as well as utilities for managing changes
    'lewis6991/gitsigns.nvim',
    event = 'VeryLazy',
    keys = {
      {
        '<leader>gp',
        ':Gitsigns preview_hunk<CR>',
        { desc = 'Show hunk change' },
      },
      {
        '<leader>gb',
        ':Gitsigns toggle_current_line_blame<CR>',
        { desc = 'Toggle Blame' },
      },
    },
    opts = {},
  },
}
