# pyenv
export PYENV_ROOT=$HOME/source/pyenv
export PATH=$PYENV_ROOT/bin:$PATH
if command -v pyenv 1>/dev/null 2>&1; then
  eval "$(pyenv init --path)"
fi

# pyenv-virtualenv
eval "$(pyenv virtualenv-init -)"

# nodenv
export NODENV_ROOT=$HOME/source/nodenv
export PATH=$NODENV_ROOT/bin:$PATH
if command -v nodenv 1>/dev/null 2>&1; then
  eval "$(nodenv init -)"
fi

# FZF
export FZF_DEFAULT_COMMAND="fd --type f --hidden --follow --exclude .git"

# bun
export BUN_INSTALL="$HOME/.bun"
export PATH="$BUN_INSTALL/bin:$PATH"
# bun completions
[ -s "$HOME/.bun/_bun" ] && source "$HOME/.bun/_bun"


# poetry
export PATH="$HOME/.local/bin:$PATH"

# neovim
export PATH="$HOME/neovim/bin:$PATH"
