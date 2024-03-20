# MCFunction LSP

Minecraft Function LSP for better development experience

<!-- vim-markdown-toc GFM -->

* [💻 Demonstration](#-demonstration)
* [🧩 Installation](#-installation)
    * [Neovim](#neovim)
    * [VSCode / VSCodium](#vscode--vscodium)
* [📦 Building](#-building)

<!-- vim-markdown-toc -->

# 💻 Demonstration

> \[INFO\]
> This project is still in-development. This section will appear later.

# 🧩 Installation

If your editor isn't listed here, refer to the editor's documentation regarding installation of custom LSP servers.

## Neovim

You will need to define `mcfunction` filetype:

```lua
vim.filetype.add {
  extension = {
    mcfunction = 'mcfunction',
  },
}
```

Then register the LSP:

```lua
local client = vim.lsp.start_client {
  name = 'mcfunction_lsp',
  cmd = { '/path/to/mcfunction-lsp' },
  root_dir = vim.fn.getcwd(),
}

if not client then
  vim.notify 'MCFunction LSP Client did not load!'
end

-- Attach the LSP client when opening an .mcfunction file
vim.api.nvim_create_autocmd('FileType', {
  pattern = 'mcfunction',
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

## VSCode / VSCodium

Install the [VSCode Extension](https://github.com/mend-lang/vscode-extension).

# 📦 Building

1. Clone the repository: `git clone https://github.com/mend-lang/mcfunction-lsp.git`.
1. Compile the project:
   2.1. Use `make dev` to compile a debug build.
   2.2. Use `make prod` to compile to every platform (Windows, MacOS, Linux)
1. Provide the path to binary in your editor's LSP config.
