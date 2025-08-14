return {
  "nvim-lualine/lualine.nvim",
  opts = {
    options = {
      theme = {
        normal = {
          a = { fg = "#ffffff", bg = "#11111b" },
          b = { fg = "#ffffff", bg = "#11111b" },
          c = { fg = "#ffffff", bg = "#11111b" },
        },
        insert = {
          a = { fg = "#ffffff", bg = "#11111b" },
          b = { fg = "#ffffff", bg = "#11111b" },
          c = { fg = "#ffffff", bg = "#11111b" },
        },
        visual = {
          a = { fg = "#ffffff", bg = "#11111b" },
          b = { fg = "#ffffff", bg = "#11111b" },
          c = { fg = "#ffffff", bg = "#11111b" },
        },
        replace = {
          a = { fg = "#ffffff", bg = "#11111b" },
          b = { fg = "#ffffff", bg = "#11111b" },
          c = { fg = "#ffffff", bg = "#11111b" },
        },
        command = {
          a = { fg = "#ffffff", bg = "#11111b" },
          b = { fg = "#ffffff", bg = "#11111b" },
          c = { fg = "#ffffff", bg = "#11111b" },
        },
        inactive = {
          a = { fg = "#cccccc", bg = "#11111b" },
          b = { fg = "#cccccc", bg = "#11111b" },
          c = { fg = "#cccccc", bg = "#11111b" },
        },
      },
      section_separators = "",
      component_separators = "",
    },
    sections = {
      lualine_a = { "mode" },
      lualine_b = { "branch" },
      lualine_c = { "filename" },
      lualine_x = { "progress" },
      lualine_y = {},
      lualine_z = { { "datetime", style = "%H:%M" } },
    },
  },
}
