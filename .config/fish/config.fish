set fish_greeting
set TERM xterm-256color
set -gx EDITOR nvim
set -gx XDG_CONFIG_HOME $HOME/.config

if test -f $HOME/.config/fish/secrets.fish
    source $HOME/.config/fish/secrets.fish
end

fish_add_path -m $HOME/.bin $HOME/.local/bin /usr/local/bin $fish_user_paths
source $HOME/.config/fish/aliases.fish

function mkcd
    mkdir -p $argv
    and cd $argv
end

fzf --fish | source
mise activate fish | source
zoxide init fish | source
starship init fish | source
