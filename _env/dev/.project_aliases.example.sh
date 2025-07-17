#!/bin/bash

# =================================================================
# Docker & Application Aliases
# =================================================================
alias g="docker exec -it chronos-api"
alias gl="docker compose logs -f chronos-api"
alias gd="docker compose up -d"
alias gb="g sh"
alias gds="docker compose stats"

# =================================================================
# Go Tooling Aliases
# =================================================================
alias gg="g go"
alias gr="gg run"
alias gtest="gg test"

# =================================================================
# Goose Migration Aliases
# =================================================================
alias gs="g goose"
alias gup="gs up"
alias gdown="gs down"
alias greset="gs reset"
alias gcreate="gs create"
