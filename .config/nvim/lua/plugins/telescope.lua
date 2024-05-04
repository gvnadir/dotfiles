-- NOTE: Plugins can specify dependencies.
--
-- The dependencies are proper plugin specifications as well - anything
-- you do for a plugin at the top level, you can do for a dependency.
--
-- Use the `dependencies` key to specify the dependencies of a particular plugin

return { -- Fuzzy Finder (files, lsp, etc)
  'nvim-telescope/telescope.nvim',
  name = 'telescope',
  event = 'VeryLazy',
  branch = '0.1.x',
  dependencies = {
    'plenary',
    { -- If encountering errors, see telescope-fzf-native README for install instructions
      'nvim-telescope/telescope-fzf-native.nvim',

      -- `build` is used to run some command when the plugin is installed/updated.
      -- This is only run then, not every time Neovim starts up.
      build = 'make',

      -- `cond` is a condition used to determine whether this plugin should be
      -- installed and loaded.
      cond = function()
        return vim.fn.executable 'make' == 1
      end,
    },
    { 'nvim-telescope/telescope-ui-select.nvim' },

    -- Useful for getting pretty icons, but requires special font.
    --  If you already have a Nerd Font, or terminal set up with fallback fonts
    --  you can enable this
    { 'nvim-tree/nvim-web-devicons' },

    { 'benfowler/telescope-luasnip.nvim' },

    { 'AckslD/nvim-neoclip.lua' },
  },
  config = function()
    -- Telescope is a fuzzy finder that comes with a lot of different things that
    -- it can fuzzy find! It's more than just a "file finder", it can search
    -- many different aspects of Neovim, your workspace, LSP, and more!
    --
    -- The easiest way to use telescope, is to start by doing something like:
    --  :Telescope help_tags
    --
    -- After running this command, a window will open up and you're able to
    -- type in the prompt window. You'll see a list of help_tags options and
    -- a corresponding preview of the help.
    --
    -- Two important keymaps to use while in telescope are:
    --  - Insert mode: <c-/>
    --  - Normal mode: ?
    --
    -- This opens a window that shows you all of the keymaps for the current
    -- telescope picker. This is really useful to discover what Telescope can
    -- do as well as how to actually do it!

    -- [[ Configure Telescope ]]
    -- See `:help telescope` and `:help telescope.setup()`
    require('telescope').setup {
      -- You can put your default mappings / updates / etc. in here
      --  All the info you're looking for is in `:help telescope.setup()`
      --
      -- defaults = {
      --   mappings = {
      --     i = { ['<c-enter>'] = 'to_fuzzy_refine' },
      --   },
      -- },
      -- pickers = {}
      extensions = {
        ['ui-select'] = {
          require('telescope.themes').get_dropdown(),
        },
      },
    }

    require('neoclip').setup()

    -- Enable telescope extensions, if they are installed
    pcall(require('telescope').load_extension, 'fzf')
    pcall(require('telescope').load_extension, 'ui-select')
    pcall(require('telescope').load_extension, 'luasnip')
    pcall(require('telescope').load_extension, 'notify')

    -- See `:help telescope.builtin`
    local builtin = require 'telescope.builtin'
    vim.keymap.set('n', '<leader>:', '<cmd>Telescope command_history<cr>', { desc = 'Command History' })
    vim.keymap.set('n', '<leader>fh', builtin.help_tags, { desc = 'Search Help' })
    vim.keymap.set('n', '<leader>fk', builtin.keymaps, { desc = 'Search Keymaps' })
    vim.keymap.set('n', '<leader>ff', function()
      builtin.find_files()
    end, { desc = 'Find Files' })
    vim.keymap.set('n', '<leader>fa', function()
      builtin.find_files { hidden = true }
    end, { desc = 'Find All Files (Including Hidden Ones)' })
    vim.keymap.set('n', '<leader><leader>', builtin.git_files, { desc = 'Search Git Files' })
    vim.keymap.set('n', '<leader>fs', builtin.builtin, { desc = 'Search Select Telescope' })
    -- Searches for all the occurrences of the word under the cursor
    vim.keymap.set('n', '<leader>fw', builtin.grep_string, { desc = 'Search current Word' })
    vim.keymap.set('n', '<leader>fp', builtin.live_grep, { desc = 'Search by Grep' })
    vim.keymap.set('n', '<leader>fd', builtin.diagnostics, { desc = 'Search Diagnostics' })
    vim.keymap.set('n', '<leader>fr', builtin.resume, { desc = 'Search Resume' })
    vim.keymap.set('n', '<leader>f.', builtin.oldfiles, { desc = 'Search Recent Files ("." for repeat)' })
    vim.keymap.set('n', '<leader>fb', builtin.buffers, { desc = 'Find existing buffers' })
    vim.keymap.set('n', '<leader>fi', function()
      require('telescope').extensions.luasnip.luasnip {}
    end, { desc = 'Search snippets' })
    vim.keymap.set('n', '<leader>fy', function()
      require('telescope').extensions.neoclip.default()
    end, { desc = 'Search yank history' })
    vim.keymap.set('n', '<leader>fn', function()
      require('telescope').extensions.notify.notify()
    end, { desc = 'Search notification history' })

    -- git
    vim.keymap.set('n', '<leader>gc', '<cmd>Telescope git_commits<cr>', { desc = 'Search Commits' })
    vim.keymap.set('n', '<leader>gs', '<cmd>Telescope git_status<cr>', { desc = 'Search Git Status' })

    -- Slightly advanced example of overriding default behavior and theme
    vim.keymap.set('n', '<leader>/', function()
      -- You can pass additional configuration to telescope to change theme, layout, etc.
      builtin.current_buffer_fuzzy_find(require('telescope.themes').get_dropdown {
        winblend = 10,
        previewer = false,
      })
    end, { desc = '/ Fuzzily search in current buffer' })

    -- Also possible to pass additional configuration options.
    --  See `:help telescope.builtin.live_grep()` for information about particular keys
    vim.keymap.set('n', '<leader>f/', function()
      builtin.live_grep {
        grep_open_files = true,
        prompt_title = 'Live Grep in Open Files',
      }
    end, { desc = 'Search / in Open Files' })

    -- Shortcut for searching your neovim configuration files
    vim.keymap.set('n', '<leader>fv', function()
      builtin.find_files { cwd = vim.fn.stdpath 'config' }
    end, { desc = 'Search Neovim files' })
  end,
}
