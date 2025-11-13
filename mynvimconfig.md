# 📝 Мой Neovim — Полное руководство по конфигурации

> **Цель**: максимальная функциональность, скорость, минимализм.  
> **Среда**: Linux (Ubuntu), Alacritty, i3, Go + Lua разработка.

---

## 🔧 1. Базовые настройки (`init.lua`)

| Настройка | Значение | Пояснение |
|---------|--------|----------|
| `mapleader` / `maplocalleader` | `" "` (пробел) | Все пользовательские команды начинаются с **пробела**. |
| `scrolloff` | `10` | Всегда 10 строк "воздуха" сверху и снизу от курсора — удобно читать код. |
| `cursorline` | `true` | Подсвечивает **текущую строку**. |
| `number` + `relativenumber` | `true` | Абсолютный номер текущей строки + относительные номера остальных — идеально для навигации (`5j`, `10k`). |
| `tabstop` / `shiftwidth` | `4` | Отступы — **4 пробела** (стандарт для Go). |
| `expandtab` | `true` | Нажатие `Tab` → вставляет **пробелы**, а не символ табуляции. |
| `smartindent` | `true` | Умное автоформатирование отступов. |
| `wrap` | `false` | Длинные строки **не переносятся** — прокручиваются вправо. |
| `mouse` | `"a"` | Полная поддержка мыши (выделение, клики по окнам). |
| `swapfile` / `backup` | `false` | Не создаются временные файлы — чистота в рабочей директории. |
| `undofile` | `true` | История отмены сохраняется между сессиями. |
| `termguicolors` | `true` | Поддержка **24-битных цветов** (Alacritty — поддерживает). |
| `updatetime` | `300` | Быстрее обновление diagnostics, LSP, autosave-плагинов. |
| `signcolumn` | `"yes"` | Резервирует место слева под значки (ошибки, git) — интерфейс не прыгает. |
| **`clipboard`** | **`"unnamedplus"`** | ✅ **Ключевая настройка!**<br>Все операции копирования/вставки (`yy`, `p`, `v...y`) работают с **системным буфером** (`Ctrl+C` / `Ctrl+V`).<br>Можно копировать из Neovim → вставить в браузер и наоборот. |

---

## 🎨 2. Цветовая схема и подсветка

- Тема: **`gruber-darker.nvim`** (тёмная, без курсива)
- При смене темы автоматически переопределяются цвета:
  - **Номера строк**: `#6272a4` (мягкий серый)
  - **Номер текущей строки**: `#ff79c6` (**розовый**, **жирный**)
  - **Фон текущей строки**: `#282a36` (тёмно-серый)

---

## ⌨️ 3. Горячие клавиши (все `<leader>` = **Пробел**)

### 📁 Навигация и управление
| Клавиши | Действие |
|--------|--------|
| `<leader>t` | Открыть встроенный терминал |
| `<leader>e` | Включить/выключить файловый менеджер **Neo-tree** |
| `<leader>w` | Сохранить файл (`:w`) |
| `<leader>W` | Сохранить файл принудительно (`:w!`) |
| `<leader>Q` | Закрыть всё и выйти (`:qa`) |

### 🪟 Управление окнами
| Клавиши | Действие |
|--------|--------|
| `<leader>h` | Перейти в **левое** окно |
| `<leader>j` | Перейти в **нижнее** окно |
| `<leader>k` | Перейти в **верхнее** окно |
| `<leader>l` | Перейти в **правое** окно |

### 📄 Управление буферами
| Клавиши | Действие |
|--------|--------|
| `<M-h>` / `<M-l>` | Переключиться на **предыдущий / следующий буфер** |
| `<Space>q` | **Закрыть текущий буфер**, но не закрывать Neovim (используется `mini.bufremove`) |
| `<leader>1` – `<leader>9` | Переключиться на буфер с номером **1–9** |

### 🔍 Поиск и навигация по коду (LSP)
| Клавиши | Действие |
|--------|--------|
| `gd` | Перейти к **определению** |
| `gD` | Перейти к **объявлению** |
| `gi` | Найти **реализации** |
| `gr` | Найти **все использования** |
| `K` | Показать подсказку (**hover**) |
| `<leader>rn` | Переименовать символ во всём проекте |
| `<leader>ca` | Быстрые исправления (**code actions**) |
| `<leader>df` | Форматировать код (go fmt, etc.) |
| `<leader>ld` | Показать ошибку под курсором |
| `[d` / `]d` | Перейти к **предыдущей / следующей ошибке** |

### 🔎 Поиск через Telescope
| Клавиши | Действие |
|--------|--------|
| `<leader>ff` | Найти файлы (`find_files`) |
| `<leader>fg` | Поиск по коду (`live_grep`) |
| `<leader>fb` | Переключиться между буферами (`buffers`) |
| `<leader>fs` | Поиск в текущем файле (`current_buffer_fuzzy_find`) |

### ✨ Автодополнение (nvim-cmp)
| Клавиши | Действие |
|--------|--------|
| `Ctrl+Space` | Открыть / закрыть меню автодополнения (**вручную!**) |
| `Tab` / `S-Tab` | Выбрать следующий / предыдущий вариант |
| `Enter` | Подтвердить выбор |
| `↑` / `↓` | Навигация по вариантам |

---

## 🧩 4. Плагины и их назначение

| Плагин | Назначение |
|-------|----------|
| **`mini.nvim`** (`bufremove`) | Безопасное закрытие буферов без закрытия окна или Neovim. |
| **`bufferline.nvim`** | Визуальная панель буферов (вкладки наверху). Иконки отключены. |
| **`nvim-cmp`** | Умное автодополнение (только по запросу — без всплытия). |
| **`mason.nvim` + `mason-lspconfig`** | Установка и настройка LSP-серверов (для Go, Lua и др.). |
| **`lualine.nvim`** | Современный статус-бар с веткой Git, ошибками, именем файла. |
| **`telescope.nvim`** | Мощный fuzzy-поиск по файлам, коду, буферам. |
| **`gruber-darker.nvim`** | Минималистичная тёмная тема без курсива. |
| **`nvim-treesitter`** | Улучшенная подсветка синтаксиса и навигация (Go, Lua, Bash, Markdown). |
| **`nvim-autopairs`** | Автоматическое закрытие скобок, кавычек, тегов. |
| **`neo-tree.nvim`** | Файловый менеджер с поддержкой Git и быстрой навигацией. |

---

## ⚙️ 5. Производительность

- Отключены встроенные плагины Neovim:  
  `gzip`, `matchit`, `netrw`, `tar`, `tohtml`, `tutor`, `zip` — **меньше мусора, быстрее запуск**.
- Используется **`lazy.nvim`** — современный менеджер плагинов с lazy-loading.

---

## ✅ 6. Особенности моей системы

- **Терминал**: Alacritty → поддерживает truecolor и корректно передаёт `Alt` как `<M-...>`.
- **Окружение**: i3 → навигация по окнам важна → настроены `<leader>h/j/k/l`.
- **Язык**: Go → отступы = 4 пробела, LSP, форматирование через `gofmt`.
- **Буфер обмена**: работает через `xclip` + `clipboard = "unnamedplus"` → **единый буфер с системой**.

---

```lua
-- === 1. Leader и базовые настройки ===
vim.g.mapleader = " "       -- Leader = пробел
vim.g.maplocalleader = " "  -- Local leader = пробел
vim.opt.scrolloff = 10      -- Всегда 10 строк "воздуха" сверху/снизу

-- Надёжный способ: применять после каждой смены темы
vim.api.nvim_create_autocmd("ColorScheme", {
  pattern = "*",
  callback = function()
    vim.cmd([[hi LineNr       guifg=#6272a4 gui=NONE]])
    vim.cmd([[hi CursorLineNr guifg=#ff79c6 gui=bold]])
  end,
})

vim.opt.cursorline = true  -- включить подсветку строки
vim.cmd([[hi CursorLine ctermbg=236 guibg=#282a36]])  -- фон строки (тёмно-серый)

vim.opt.clipboard = "unnamedplus"  -- копировать в системный буфер (Ctrl+C / Ctrl+V)

-- Номера строк
vim.opt.number = true
vim.opt.relativenumber = true

-- Отступы
vim.opt.tabstop = 4
vim.opt.shiftwidth = 4
vim.opt.expandtab = true
vim.opt.smartindent = true

-- Поведение
vim.opt.wrap = false
vim.opt.mouse = ""
vim.opt.swapfile = false
vim.opt.backup = false
vim.opt.undofile = true
vim.opt.mouse = "a"

-- Цвета и интерфейс
vim.opt.termguicolors = true
vim.opt.updatetime = 300
vim.opt.signcolumn = "yes"




-- === 3. Остальные бинды ===
vim.keymap.set("n", "<leader>t", "<cmd>terminal<cr>", { desc = "Open terminal" })
vim.keymap.set("n", "<leader>e", "<cmd>Neotree toggle<cr>", { desc = "Toggle file tree" })
vim.keymap.set("n", "<leader>w", "<cmd>w<cr>", { desc = "Save file" })
vim.keymap.set("n", "<leader>W", "<cmd>w!<cr>", { desc = "Force save" })

vim.keymap.set("n", "<leader>Q", "<cmd>qa<cr>", { desc = "Quit all" })
vim.keymap.set("n", "<leader>h", "<C-w>h", { desc = "Go to left window" })
vim.keymap.set("n", "<leader>l", "<C-w>l", { desc = "Go to right window" })
vim.keymap.set("n", "<leader>j", "<C-w>j", { desc = "Go to window below" })
vim.keymap.set("n", "<leader>k", "<C-w>k", { desc = "Go to window above" })

-- === 4. Установка lazy.nvim (ВАЖНО: URL без пробелов!) ===
local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git", "clone", "--filter=blob:none", "--branch=stable",
    "https://github.com/folke/lazy.nvim.git",  -- ИСПРАВЛЕНО: нет пробелов!
    lazypath,
  })
end
vim.opt.rtp:prepend(lazypath)
package.path = package.path .. ";" .. lazypath .. "/?.lua"


-- === LSP бинды ===
vim.keymap.set("n", "gd", "<cmd>lua vim.lsp.buf.definition()<CR>", { desc = "Go to definition" }) --Перейти к определению
vim.keymap.set("n", "gD", "<cmd>lua vim.lsp.buf.declaration()<CR>", { desc = "Go to declaration" }) --К объявлению
vim.keymap.set("n", "gi", "<cmd>lua vim.lsp.buf.implementation()<CR>", { desc = "Go to implementation" }) -- Реализации
vim.keymap.set("n", "gr", "<cmd>lua vim.lsp.buf.references()<CR>", { desc = "Find references" }) --Где ещё используется эта функция/переменная?
vim.keymap.set("n", "K", "<cmd>lua vim.lsp.buf.hover()<CR>", { desc = "Show hover info" }) --Подсказка
vim.keymap.set("n", "<leader>rn", "<cmd>lua vim.lsp.buf.rename()<CR>", { desc = "Rename symbol" }) --Переименовать in al project
vim.keymap.set("n", "<leader>ca", "<cmd>lua vim.lsp.buf.code_action()<CR>", { desc = "Code action (quickfix)" }) --Code action
vim.keymap.set("n", "<leader>df", "<cmd>lua vim.lsp.buf.format({ async = true })<CR>", { desc = "Format buffer" }) --Форматирование
vim.keymap.set("n", "<leader>ld", "<cmd>lua vim.diagnostic.open_float()<CR>", { desc = "Show diagnostic under cursor" }) --Показать ошибку
vim.keymap.set("n", "[d", "<cmd>lua vim.diagnostic.goto_prev()<CR>", { desc = "Go to previous diagnostic" }) --Переключение по ошибкам
vim.keymap.set("n", "]d", "<cmd>lua vim.diagnostic.goto_next()<CR>", { desc = "Go to next diagnostic" }) --Переключение по ошибкам

-- Закрыть буфер, но не закрывать Neovim (рабочая версия)
vim.keymap.set('n', '<Space>q', function()
  require('mini.bufremove').delete(0)  -- 0 = текущий буфер
end, { silent = true, desc = "Close current buffer safely" })
-- буфер перемещение
vim.keymap.set('n', '<M-h>', ':bprev<CR>', { silent = true })
vim.keymap.set('n', '<M-l>', ':bnext<CR>', { silent = true })

-- Переключение на буферы по номеру: [Space] + 1..9
for i = 1, 9 do
  vim.keymap.set("n", "<leader>" .. i, ":buffer " .. i .. "<CR>", {
    silent = true,
    desc = "Switch to buffer " .. i
  })
end

-- === 5. Плагины ===
require("lazy").setup({
{
  "echasnovski/mini.nvim",
  version = "*",
  config = function()
    require('mini.bufremove').setup()
  end,
},

 --buffers
{
  "akinsho/bufferline.nvim",
  version = "*",
  dependencies = "nvim-tree/nvim-web-devicons", -- иконки (если хочешь)
  opts = {
    options = {
      mode = "buffers",
      separator_style = "slant",
      show_buffer_close_icons = false,
      show_close_icon = false,
      -- Чтобы колёсико закрывало вкладки:
      close_command = "bdelete! %d",
    }
  }
},


-- Автодополнение (вручную вызываемое по Ctrl+j)
{
  "hrsh7th/nvim-cmp",
  dependencies = {
    "hrsh7th/cmp-nvim-lsp",
    "L3MON4D3/LuaSnip",
  },
  config = function()
    local cmp = require("cmp")

    cmp.setup({
      completion = { autocomplete = false }, -- ❌ отключаем автопоявление
      mapping = {
        ["<C-Space>"] = cmp.mapping(function()
          if cmp.visible() then
            cmp.close()
          else
            cmp.complete() -- открыть меню
          end
        end, { "i", "c" }),

        ["<Tab>"] = cmp.mapping.select_next_item({ behavior = cmp.SelectBehavior.Select }),
        ["<S-Tab>"] = cmp.mapping.select_prev_item({ behavior = cmp.SelectBehavior.Select }),
        ["<CR>"] = cmp.mapping.confirm({ select = true }), -- Enter = выбрать

        -- Для удобства — стрелки ↑ ↓ тоже работают
        ["<Down>"] = cmp.mapping.select_next_item({ behavior = cmp.SelectBehavior.Select }),
        ["<Up>"] = cmp.mapping.select_prev_item({ behavior = cmp.SelectBehavior.Select }),
      },

      sources = {
        { name = "nvim_lsp" },
        { name = "buffer" },
      },
    })
  end,
},



{
    "mason-org/mason-lspconfig.nvim",
    opts = {},
    dependencies = {
        { "mason-org/mason.nvim", opts = {} },
        "neovim/nvim-lspconfig",
    },
},


  -- Статус-бар
  {
    "nvim-lualine/lualine.nvim",
    dependencies = { "nvim-tree/nvim-web-devicons" },
    opts = {
      options = {
      --  theme = "gruber_darker",
        section_separators = { left = "", right = "" },
        component_separators = { left = "", right = "" },
        icons_enabled = true,
      },
      sections = {
        lualine_a = { "mode" },
        lualine_b = { "branch", "diff", "diagnostics" },
        lualine_c = { "filename" },
        lualine_x = { "encoding", "fileformat", "filetype" },
        lualine_y = { "progress" },
        lualine_z = { "location" },
      },
      inactive_sections = {
        lualine_a = { "filename" },
        lualine_c = {},
        lualine_x = {},
        lualine_y = {},
        lualine_z = { "location" },
      },
    },
  },

  -- Поиск (Telescope)
  {
    "nvim-telescope/telescope.nvim",
    dependencies = {
      "nvim-lua/plenary.nvim",
      { "nvim-telescope/telescope-fzf-native.nvim", build = "make" },
    },
    keys = {
      { "<leader>ff", "<cmd>Telescope find_files<cr>", desc = "Find files" },
      { "<leader>fg", "<cmd>Telescope live_grep<cr>", desc = "Live grep" },
      { "<leader>fb", "<cmd>Telescope buffers<cr>", desc = "Switch buffers" },
      { "<leader>fs", "<cmd>Telescope current_buffer_fuzzy_find<cr>", desc = "Search in file" },
    },
    config = function()
      require("telescope").setup({
        defaults = {
          vimgrep_arguments = {
            "rg", "--color=never", "--no-heading", "--with-filename",
            "--line-number", "--column", "--smart-case"
          },
          prompt_prefix = "🔍 ",
          selection_caret = "➤ ",
          path_display = { "truncate" },
          layout_config = { horizontal = { width = 0.9, height = 0.8 } },
        },
      })
      pcall(require("telescope").load_extension, "fzf")
    end,
  },

  -- Тема: gruber-darker без курсива
  {
    "blazkowolf/gruber-darker.nvim",
    priority = 1000,
    opts = {
      italic = {
        strings = false, comments = false, operators = false,
        folds = false, keywords = false, functions = false, variables = false,
      },
      bold = false,
    },
    config = function(_, opts)
      require("gruber-darker").setup(opts)
      vim.cmd("colorscheme gruber-darker")
    end,
  },

  -- Tree-sitter
  {
    "nvim-treesitter/nvim-treesitter",
    build = ":TSUpdate",
    opts = {
      highlight = { enable = true },
      indent = { enable = true },
      ensure_installed = { "lua", "go", "gomod", "markdown", "bash" },
      auto_install = true,
    },
    config = function(_, opts)
      require("nvim-treesitter.configs").setup(opts)
    end,
  },

  -- Автозакрытие скобок
  {
    "windwp/nvim-autopairs",
    event = "InsertEnter",
    config = function()
      require("nvim-autopairs").setup({})
    end,
  },

  -- Файловый менеджер
  {
    "nvim-neo-tree/neo-tree.nvim",
    branch = "v3.x",
    dependencies = {
      "nvim-lua/plenary.nvim",
      "MunifTanjim/nui.nvim",
      "nvim-tree/nvim-web-devicons",
    },
    keys = { { "<leader>e", "<cmd>Neotree toggle<cr>", desc = "Toggle file tree" } },
    opts = {
      close_if_last_window = true,
      filesystem = {
        filtered_items = {
          visible = false,
          hide_dotfiles = false,
          hide_gitignored = false,
          hide_hidden = false,
        },
      },
      window = {
        width = 30,
        mappings = { ["<space>"] = "none" },
      },
    },
  },
}, {
  rocks = { enabled = false },
  performance = {
    rtp = {
      disabled_plugins = {
        "gzip", "matchit", "netrwPlugin", "tarPlugin", "tohtml", "tutor", "zipPlugin"
      },
    },
  },
})

```
