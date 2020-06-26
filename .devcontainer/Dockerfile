FROM golang:1-alpine

RUN apk add --no-cache --virtual .base \
    sudo bash zsh zsh-vcs tmux \
    curl git vim openssh

# fix sudo (https://github.com/sudo-project/sudo/issues/42)
RUN echo "Set disable_coredump false" >> /etc/sudo.conf

ARG USER=docker
ENV HOME /home/$USER

# timezone
COPY configurations/localtime /etc/localtime

# add new user with zsh as the default shell
RUN adduser -D $USER -s /bin/zsh \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME

# oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" \
    && cd /tmp \
    && git clone https://github.com/powerline/fonts.git --depth=1 \
    && cd fonts && ./install.sh \
    && cd .. && rm -rf fonts \
    && git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions \
    && git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# user settings
COPY configurations/zsh.zshrc ${HOME}/.zshrc
COPY configurations/_tmux.conf ${HOME}/.tmux.conf

# key scan for well known git servers (e.g. github)
RUN mkdir -p ${HOME}/.ssh \
    && ssh-keyscan github.com >> ${HOME}/.ssh/known_hosts

# go tools
RUN go get github.com/mdempsky/gocode \
    && go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get github.com/ramya-rao-a/go-outline \
    && go get github.com/acroca/go-symbols \
    && go get golang.org/x/tools/cmd/guru \
    && go get golang.org/x/tools/cmd/gorename \
    && go get github.com/cweill/gotests/... \
    && go get github.com/fatih/gomodifytags \
    && go get github.com/josharian/impl \
    && go get github.com/davidrjenni/reftools/cmd/fillstruct \
    && go get github.com/haya14busa/goplay/cmd/goplay \
    && go get github.com/godoctor/godoctor \
    && go get github.com/go-delve/delve/cmd/dlv \
    && go get github.com/stamblerre/gocode \
    && go get github.com/rogpeppe/godef \
    && go get golang.org/x/tools/cmd/goimports \
    && go get golang.org/x/lint/golint \
    && go get golang.org/x/tools/gopls

ENTRYPOINT [ "/bin/zsh" ] 
