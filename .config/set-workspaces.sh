#!/bin/bash

# Execute this file as a startup program from "Startup applications preferences"

# Disable Super+0..9 for launching dock applications
for i in {1..9}; do
  gsettings set org.gnome.shell.keybindings switch-to-application-$i "[]"
done
gsettings set org.gnome.shell.keybindings switch-to-application-0 "[]"

# Map Super+1..9 to workspaces 1..9
for i in {1..9}; do
  gsettings set org.gnome.desktop.wm.keybindings switch-to-workspace-$i "['<Super>$i']"
done

# Map Super+0 to workspace 10
gsettings set org.gnome.desktop.wm.keybindings switch-to-workspace-10 "['<Super>0']"

# Map Super+Shift+1..0 to move windows to workspaces 1..9
for i in {1..9}; do
  gsettings set org.gnome.desktop.wm.keybindings move-to-workspace-$i "['<Super><Shift>$i']"
done
#
# Map Super+Shift+0 to move windows to workspace 10
gsettings set org.gnome.desktop.wm.keybindings move-to-workspace-10 "['<Super><Shift>0']"

# (Optional) Lock workspaces to 10 total instead of dynamic workspaces
# gsettings set org.gnome.mutter dynamic-workspaces false
# gsettings set org.gnome.desktop.wm.preferences num-workspaces 10
