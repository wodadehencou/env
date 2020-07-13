export ENV=$HOME/.env

source $ENV/antigen/antigen.zsh

antigen bundle robbyrussell/oh-my-zsh lib/

antigen bundle git
antigen bundle docker
antigen bundle docker-compose
antigen bundle golang
antigen bundle brew
antigen bundle zsh-users/zsh-syntax-highlighting
antigen bundle zsh-users/zsh-completions
antigen bundle zsh-users/zsh-autosuggestions
antigen bundle zsh-users/zsh-history-substring-search
antigen bundle skywind3000/z.lua
antigen bundle $ENV/jam-zsh-plugin --no-local-clone
#antigen theme tonyseek/oh-my-zsh-seeker-theme seeker
antigen theme romkatv/powerlevel10k

antigen apply

