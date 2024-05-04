return { -- Collection of various small independent plugins/modules
  'echasnovski/mini.nvim',
  event = 'VeryLazy',
  config = function()
    -- Better Around/Inside textobjects
    --
    -- Examples:
    --  - va)  - [V]isually select [A]round [)]paren
    --  - yinq - [Y]ank [I]nside [N]ext [']quote
    --  - ci'  - [C]hange [I]nside [']quote
    require('mini.ai').setup { n_lines = 500 }

    -- Add/delete/replace surroundings (brackets, quotes, etc.)
    --
    -- - saiw) - [S]urround [A]dd [I]nner [W]ord [)]Paren
    -- - sd'   - [S]urround [D]elete [']quotes
    -- - sr)'  - [S]urround [R]eplace [)] [']
    require('mini.surround').setup()

    require('mini.comment').setup {
      mappings = {
        -- comment_line = 'g;',
      },
      options = {
        custom_commentstring = function()
          return require('ts_context_commentstring.internal').calculate_commentstring() or vim.bo.commentstring
        end,
      },
    }

    vim.keymap.set('n', '<leader>up', function()
      vim.g.minipairs_disable = not vim.g.minipairs_disable
      if vim.g.minipairs_disable then
        vim.notify('Disabled auto pairs', 'info')
      else
        vim.notify('Enabled auto pairs', 'info')
      end
    end, { desc = 'Toggle Autopairs' })

    -- ... and there is more!
    --  Check out: https://github.com/echasnovski/mini.nvim
  end,
}
