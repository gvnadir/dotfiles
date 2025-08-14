return {
  "nvim-lualine/lualine.nvim",
  opts = {
    options = {
      theme = {
        normal = {
          a = { fg = "#ffffff", bg = "#181825" },
          b = { fg = "#ffffff", bg = "#181825" },
          c = { fg = "#ffffff", bg = "#181825" },
        },
        insert = {
          a = { fg = "#ffffff", bg = "#181825" },
          b = { fg = "#ffffff", bg = "#181825" },
          c = { fg = "#ffffff", bg = "#181825" },
        },
        visual = {
          a = { fg = "#ffffff", bg = "#181825" },
          b = { fg = "#ffffff", bg = "#181825" },
          c = { fg = "#ffffff", bg = "#181825" },
        },
        replace = {
          a = { fg = "#ffffff", bg = "#181825" },
          b = { fg = "#ffffff", bg = "#181825" },
          c = { fg = "#ffffff", bg = "#181825" },
        },
        command = {
          a = { fg = "#ffffff", bg = "#181825" },
          b = { fg = "#ffffff", bg = "#181825" },
          c = { fg = "#ffffff", bg = "#181825" },
        },
        inactive = {
          a = { fg = "#cccccc", bg = "#181825" },
          b = { fg = "#cccccc", bg = "#181825" },
          c = { fg = "#cccccc", bg = "#181825" },
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
