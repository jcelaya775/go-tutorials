#!/usr/bin/env bash
tmux new-window
tmux select-window -t 0
tmux rename-window "tour-of-go"
nvim .
