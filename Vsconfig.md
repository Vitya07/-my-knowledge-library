# Vim в VS Code — Минималистичная настройка

Все бинды работают в **normal mode** (нажми `Esc`, чтобы выйти из ввода).

> **Leader** = `Space`  
> Все команды — **последовательно**: сначала `Space`, отпускаешь, потом нужная клавиша.

---

## 🔑 Основные бинды

| Бинд            | Действие                                      |
|-----------------|-----------------------------------------------|
| `gd`            | Перейти к определению                         |
| `K` (`Shift+k`) | Показать подсказку (hover)                    |
| `<Space>f`      | Форматировать файл                            |
| `<Space>r`      | Переименовать символ (умно, через LSP)        |
| `<Space>e`      | Закрыть/открыть боковую панель                |
| `<Space>E`      | Открыть проводник (Explorer)                  |
| `<Space>q`      | Закрыть текущую вкладку                       |
| `<Space>/`      | Поиск по всем файлам проекта                  |
| `<Space>s`      | Замена по всем файлам                         |
| `<Space>b`      | Быстрое открытие файла (fuzzy search)         |
| `<Space>w`      | Сохранить все открытые файлы                  |

---

## 💡 Автодополнение (только по запросу)

- **Автоматическое автодополнение отключено**.
- Чтобы вызвать подсказки: нажми **`Ctrl+Space`**.
- Навигация: **стрелки `↑` / `↓`**.
- Вставить: **`Enter`** или **`Tab`**.

---

## 🔁 Работа со скобками (встроено в Vim)

| Команда | Действие                              |
|--------|----------------------------------------|
| `ci)`  | Изменить содержимое внутри `( ... )`   |
| `ci}`  | Изменить содержимое внутри `{ ... }`   |
| `da"`  | Удалить всё включая кавычки `"..."`    |
| `yi'`  | Скопировать содержимое внутри `'...'`  |
| `vi[`  | Выделить содержимое внутри `[ ... ]`   |
| `%`    | Перейти к парной скобке                |

---

## 🧪 Примеры

- **Изменить строку**: курсор на `"Hello"` → `ci"` → ввести новое → `Esc`
- **Посмотреть описание функции**: курсор на `Println` → нажать `K`
- **Найти переменную**: `<Space>/` → ввести `user.ID`
- **Переименовать**: курсор на переменной → `<Space>r` → ввести новое имя
- **Открыть файл**: `<Space>b` → ввести часть имени → выбрать

---

## ⚠️ Важно

- Убедись, что установлено расширение для твоего языка (например, **Go** от golang.go).
- `<Space>E` = `Space` → отпустить → `Shift+E` (не одновременно!).
- Все бинды работают **только в normal mode** (`Esc` перед использованием).


```json
{
  "editor.mouseWheelZoom": true,
  "files.autoSave": "afterDelay",
  "editor.smoothScrolling": true,
  "workbench.list.smoothScrolling": true,
  "terminal.integrated.smoothScrolling": true,
  "editor.fontWeight": "600",
  "workbench.iconTheme": "catppuccin-mocha",
  "explorer.confirmDelete": false,
  "editor.fontFamily": "JetBrainsMono Nerd Font Mono",
  "editor.fontLigatures": "'ss01', 'ss02', 'ss03', 'ss06', 'ss07', 'ss08', 'calt', 'liga'",
  "editor.cursorBlinking": "blink",
  "workbench.editor.empty.hint": "hidden",
  "path-intellisense.showHiddenFiles": true,
  "editor.hideCursorInOverviewRuler": true,
  "workbench.view.showQuietly": {
    "workbench.panel.output": false
  },
  "python.createEnvironment.trigger": "off",
  "explorer.confirmDragAndDrop": false,
  "editor.guides.bracketPairs": "active",
  "gopls": {
    "ui.semanticTokens": true
  },
  "editor.cursorStyle": "line",
  "editor.cursorWidth": 2,
  "editor.letterSpacing": 0.4,
  "editor.lineHeight": 1.4,

  // ───────────────────────────────
  // УБИВАЕМ ITALIC
  // ───────────────────────────────
  "editor.tokenColorCustomizations": {
    "textMateRules": [
      {
        "scope": [
          "comment",
          "comment.block",
          "comment.line",
          "string",
          "keyword.control",
          "storage.type",
          "entity.name.type",
          "support.type",
          "variable.language",
          "meta.var.expr"
        ],
        "settings": {
          "fontStyle": ""
        }
      }
    ]
  },

  "editor.fontVariations": false,
  "editor.italicFontStyle": false,
  "files.autoSaveDelay": 100,
  "editor.formatOnSave": true,

  // terminal cursor
  "terminal.integrated.cursorBlinking": true,
  "terminal.integrated.cursorStyle": "line",
  "workbench.startupEditor": "none",
  "workbench.colorTheme": "Gruber Darker",

  // ───────────────────────────────
  // VIM — 100% рабочая настройка
  // ───────────────────────────────
  "vim.useSystemClipboard": true,
  "vim.useCtrlKeys": true,
  "vim.leader": " ",
  "vim.handleKeys": {
    "<C-c>": false,
    "<C-v>": false,
    "<C-x>": false,
    "<C-z>": false,
    "<C-y>": false,
    "<C-a>": false
  },
  "editor.lineNumbers": "relative",
  "editor.cursorSurroundingLines": 10,

  // === ПОЛНОЕ ОТКЛЮЧЕНИЕ АВТОДОПОЛНЕНИЯ ===
  "editor.inlineSuggest.enabled": false,
  "editor.suggestOnTriggerCharacters": false,
  "editor.quickSuggestions": {
    "other": false,
    "comments": false,
    "strings": false
  },
  "editor.parameterHints.enabled": false,
  "editor.suggest.showIcons": true,
  "editor.tabCompletion": "off",

  // === VIM БИНДЫ ===
  "vim.normalModeKeyBindingsNonRecursive": [
    // Go to definition
    {
      "before": ["g", "d"],
      "commands": ["editor.action.revealDefinition"]
    },
    // HOVER по K — ЭТО РАБОТАЕТ!
    {
      "before": ["K"],
      "commands": ["editor.action.showHover"]
    },
    // Форматировать
    {
      "before": ["<leader>", "f"],
      "commands": ["editor.action.formatDocument"]
    },
    // Переименовать
    {
      "before": ["<leader>", "r"],
      "commands": ["editor.action.rename"]
    },
    // Закрыть проводник (свернуть боковую панель)
    {
      "before": ["<leader>", "e"],
      "commands": ["workbench.action.toggleSidebarVisibility"]
    },
    // Открыть проводник (гарантированно)
    {
      "before": ["<leader>", "E"],
      "commands": ["workbench.view.explorer"]
    },
    // Закрыть вкладку
    {
      "before": ["<leader>", "q"],
      "commands": ["workbench.action.closeActiveEditor"]
    },
    // Поиск по проекту
    {
      "before": ["<leader>", "/"],
      "commands": ["workbench.action.findInFiles"]
    },
    // Замена по проекту
    {
      "before": ["<leader>", "s"],
      "commands": ["workbench.action.replaceInFiles"]
    },
    // Быстрое открытие файла
    {
      "before": ["<leader>", "b"],
      "commands": ["workbench.action.quickOpen"]
    },
    // Сохранить всё
    {
      "before": ["<leader>", "w"],
      "commands": ["workbench.action.files.saveAll"]
    }
  ]
}
```







