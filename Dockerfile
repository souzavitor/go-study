FROM devtransition/golang-glide

RUN useradd -ms /bin/bash vitor

USER vitor

EXPOSE 8080