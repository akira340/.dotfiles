local status_ok, configs = pcall(require, "nvim-treesitter.configs")
if not status_ok then
  return
end

configs.setup {
  ensure_installed = "all",
  sync_install = true,
  ignore_install = { "" },
  highlight = {
    enable = true,
    disable = { "" },
    additional_vim_regex_highlighting = true,
  },
  indent = {
    -- enable = true,
    disable = { "yaml" }
  },
  context_commentstring = {
    enable = true,
    enable_autocmd = false,
  },
}
