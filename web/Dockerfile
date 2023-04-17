FROM golang:1.20.3-bullseye
ARG USERNAME=user-name-goes-here
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN mkdir /app
# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
RUN chown $USER_UID:$USER_GID /app
USER $USERNAME
WORKDIR /app
COPY . /app/
RUN go get ./...
