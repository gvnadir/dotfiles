return {
  'nvim-neo-tree/neo-tree.nvim',
  branch = 'v3.x',
  dependencies = {
    'plenary',
    'nvim-tree/nvim-web-devicons', -- not strictly required, but recommended
    'MunifTanjim/nui.nvim',
    -- "3rd/image.nvim", -- Optional image support in preview window: See `# Preview Mode` for more information
  },
  config = function()
    -- If you want icons for diagnostic errors, you'll need to define them somewhere:
    vim.fn.sign_define('DiagnosticSignError', { text = ' ', texthl = 'DiagnosticSignError' })
    vim.fn.sign_define('DiagnosticSignWarn', { text = ' ', texthl = 'DiagnosticSignWarn' })
    vim.fn.sign_define('DiagnosticSignInfo', { text = ' ', texthl = 'DiagnosticSignInfo' })
    vim.fn.sign_define('DiagnosticSignHint', { text = '󰌵', texthl = 'DiagnosticSignHint' })

    require('neo-tree').setup {
      close_if_last_window = true,
      sources = { 'filesystem', 'buffers', 'git_status', 'document_symbols' },
      default_component_configs = {
        filesystem = {
          hide_dotfiles = false,
          hide_by_name = {
            'node_modules',
          },
          never_show = {
            '.DS_Store',
          },
          use_libuv_file_watcher = true,
        },
      },
      event_handlers = {
        {
          event = 'file_opened',
          handler = function()
            -- auto close
            -- vimc.cmd("Neotree close")
            -- OR
            require('neo-tree.command').execute { action = 'close' }
          end,
        },
      },
    }

    vim.keymap.set('n', '<leader>e', function()
      require('neo-tree.command').execute { toggle = true }
    end, { desc = 'Toggle tree explorer' })

    vim.keymap.set('n', '<leader>ge', function()
      require('neo-tree.command').execute { source = 'git_status', toggle = true }
    end, { desc = 'Git Explorer' })

    vim.keymap.set('n', '<leader>be', function()
      require('neo-tree.command').execute { source = 'buffers', toggle = true }
    end, { desc = 'Buffers Explorer' })
  end,
}
