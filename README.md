# .dotfiles

## Prerequisite
- Alacritty
- oh-my-zsh
- Nerd Fonts
- tmux
- Neovim

### Install omz
Make sure your default terminal is zsh
```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### Install nerd fonts
#### Mac
psst: You need brew  

all nerd fonts
```
brew tap homebrew/cask-fonts
brew search '/font-.*-nerd-font/' | awk '{ print $1 }' | xargs -I{} brew install --cask {} || true
```

### Neovimrc
https://github.com/asashakira/neovimrc


## Install

```
alias dotfiles='/usr/bin/git --git-dir=$HOME/.dotfiles/ --work-tree=$HOME'
echo ".dotfiles" >> .gitignore
git clone --bare https://github.com/asashakira/.dotfiles.git $HOME/.dotfiles
dotfiles checkout
dotfiles config --local status.showUntrackedFiles no
```
