alias reload='source $HOME/.config/fish/config.fish'
alias la='ls -la'

if type -q lsd
    alias ls='lsd'
end

if type -q bat
    alias cat='bat'
else if type -q batcat
    alias cat='batcat'
end
