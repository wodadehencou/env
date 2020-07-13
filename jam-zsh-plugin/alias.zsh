
cdls(){chdir $1;ls}

alias work="cd $HOME/Workspace"
alias q="exit"
alias b="cd .."
alias view="vim -R"
alias vi="vim"
alias cd='cdls'
alias ssh='ssh -o serveraliveinterval=360 -X'

alias cp='cp -i'
alias mv='mv -i'
alias rm='rm -r -i'
# alias ls='ls -F --color=auto'
alias ll='ls -l'
alias la='ls -a'
alias grep='grep --color=auto'

alias format_json='python -mjson.tool'

#GOROOT=~/programs/go
#GOROOT=~/Programs/go1.10
#GOPATH=~/Programs/gopath
#PATH=$GOROOT/bin:$GOPATH/bin:$PATH

#export GOROOT
#export GOPATH
#export PATH

#function initgo () {
#	export GOPATH=`pwd`
#}


function setgopath() {
	DIR=`pwd`
	echo $DIR
	#echo $HOME

	while [ $DIR != $HOME ]; do
		#echo $DIR
		if [ -d ${DIR}/src ]; then
			break
		else
			DIR=`dirname ${DIR}`
		fi
	done

	if [ $DIR = $HOME ]; then
		echo "Can't find a valid GOPATH"
	else
		export GOPATH=${DIR}
		echo "Set env GOPATH as ${DIR}"
		export PATH=${DIR}/bin:${PATH}
		echo "Add PATH: ${DIR}/bin"
	fi
}

function addgopath () {

	DIR=`pwd`
	echo $DIR
	#echo $HOME

	while [ $DIR != $HOME ]; do
		#echo $DIR
		if [ -d ${DIR}/src ]; then
			break
		else
			DIR=`dirname ${DIR}`
		fi
	done

	if [ $DIR = $HOME ]; then
		echo "Can't find a valid GOPATH"
	else
		export GOPATH=${DIR}:${GOPATH}
		echo "Set env GOPATH as ${DIR}"
		export PATH=${DIR}/bin:${PATH}
		echo "Add PATH: ${DIR}/bin"
	fi
}

function gfw () {
	# export http_proxy=socks5://127.0.0.1:1080
	# export https_proxy=socks5://127.0.0.1:1080
	export http_proxy=http://127.0.0.1:8118
	export https_proxy=http://127.0.0.1:8118
}
#export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64


# setterm -term linux -regtabs 4

# keyboard speed
# xset r rate 300 40

unsetopt share_history

#PATH=~/.local/bin:$PATH
PATH=$ENV/bin:$PATH
#PATH=~/Programs/bin:$PATH
export PATH

