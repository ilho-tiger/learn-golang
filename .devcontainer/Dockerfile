FROM golang:1.14-alpine

RUN apk add --no-cache --virtual .base \
    sudo bash zsh zsh-vcs tmux \
    curl git vim openssh

# fix sudo (https://github.com/sudo-project/sudo/issues/42)
RUN echo "Set disable_coredump false" >> /etc/sudo.conf

ARG USER=isong
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

RUN ssh-keyscan -H github.com ${HOME}/.ssh/known_hosts

ENTRYPOINT [ "/bin/zsh" ] 
