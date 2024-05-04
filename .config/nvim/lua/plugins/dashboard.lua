return {
  'nvimdev/dashboard-nvim',
  event = 'VimEnter',
  opts = function()
    local logo = [[                                                    
       ███╗   ██╗███████╗ ██████╗ ██╗   ██╗██╗███╗   ███╗ 
       ████╗  ██║██╔════╝██╔═══██╗██║   ██║██║████╗ ████║
       ██╔██╗ ██║█████╗  ██║   ██║██║   ██║██║██╔████╔██║
       ██║╚██╗██║██╔══╝  ██║   ██║╚██╗ ██╔╝██║██║╚██╔╝██║
       ██║ ╚████║███████╗╚██████╔╝ ╚████╔╝ ██║██║ ╚═╝ ██║
       ╚═╝  ╚═══╝╚══════╝ ╚═════╝   ╚═══╝  ╚═╝╚═╝     ╚═╝
       ]]

    logo = string.rep('\n', 8) .. logo .. '\n\n'

    return {
      theme = 'doom',
      config = {
        header = vim.split(logo, '\n'),
        center = {
          { action = 'Telescope find_files', desc = ' Find File', icon = ' ', key = 'f' },
          { action = 'ene | startinsert', desc = ' New File', icon = ' ', key = 'n' },
          { action = 'Telescope oldfiles', desc = ' Recent Files', icon = ' ', key = 'r' },
          { action = 'Telescope live_grep', desc = ' Find Text', icon = ' ', key = 'g' },
          { action = [[lua require('telescope.builtin').find_files { cwd = vim.fn.stdpath 'config' }]], desc = ' Config', icon = ' ', key = 'c' },
          { action = 'Lazy', desc = ' Lazy', icon = '󰒲 ', key = 'l' },
          { action = 'qa', desc = ' Quit', icon = ' ', key = 'q' },
        },
      },
    }
  end,
  dependencies = { { 'nvim-tree/nvim-web-devicons' } },
}
