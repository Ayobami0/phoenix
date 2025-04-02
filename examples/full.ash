system:
  meta:
    name: Development Machine
    description: Linux dev environment setup
    version: 1.0.0
    author: Ayobami
    created: 2025-04-02
    target_os: ArchLinux
  
# Installation components
install:
  # Package management
  packages:
      - group: development-essentials
        manager: pacman
        manager_command: sudo pacman -Syu
        packages:
          - git
          - git-lfs
          - curl
          - wget
          - build-essential
          - cmake
      - group: editors
        manager: pacman
        manager_command: sudo pacman -Syu
        packages:
          - neovim
          - tmux
      - group: languages
        manager: pacman
        manager_command: sudo pacman -Syu
        packages:
          - python3
          - python3-pip
          - nodejs
          - npm
      - group: others
        manager: yay
        manager_command: yay -S
        packages:
          - brave
          - spotify

# Service configuration
services:
  enable:
    - ssh
    - docker
  disable:
    - cups
    - networking

# Environment configuration
environment:
  - name: PATH
    value: $HOME/.local/bin:$HOME/bin:$PATH
  - name: EDITOR
    value: nvim
  - name: VISUAL
    value: nvim
  - name: NODE_ENV
    value: development

# User configuration
user:
  - username: devuser
    shell: /bin/zsh
    groups:
      - docker
      - sudo
      - adm

# Filesystem setup
filesystem:
  directories:
   - path: $HOME/Projects
     mode: 755
   - path: $HOME/.config/nvim
     mode: 755
   - path: $HOME/.ssh
     mode: 700

  symlinks:
   - source: $HOME/dotfiles/.zshrc
     target: $HOME/.zshrc

# External resources
git:
  - source: https://github.com/username/dotfiles.git
    branch: main
    destination: $HOME/dotfiles

# Workflow configuration
workflow:
  pre_setup:
    - script: setup-network.sh
      args: []
    - script: install-yay.sh
      args: []
    - script: backup-existing-config.sh
      args: ["--force"]
  
  post_setup:
    - script: setup-dotfiles.sh
      args: []
    - script: configure-neovim.sh
      args: ["--plugins"]
