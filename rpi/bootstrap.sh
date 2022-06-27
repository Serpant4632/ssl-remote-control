#!/usr/bin/env bash

set -eu

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

SRC_VERSION=${1-}

function installService() {
    if [[ ! -f ~/.local/share/systemd/user/ssl-remote-control.service ]]; then
        mkdir -p ~/.local/share/systemd/user/
        cp "$SCRIPT_DIR/ssl-remote-control.service" ~/.local/share/systemd/user/ssl-remote-control.service
        systemctl --user enable ssl-remote-control.service
    fi
}

function installRemoteControl() {
    sudo apt-get install --no-install-recommends -y curl jq

    systemctl --user stop ssl-remote-control.service
    mkdir -p ~/.local/bin/
    if [[ -z "${SRC_VERSION}" ]]; then
        SRC_VERSION="$(curl -s https://api.github.com/repos/RoboCup-SSL/ssl-remote-control/releases/latest | jq -r '.tag_name')"
    fi
    echo "Installing version: ${SRC_VERSION}"
    wget "https://github.com/RoboCup-SSL/ssl-remote-control/releases/download/${SRC_VERSION}/ssl-remote-control_${SRC_VERSION}_linux_arm" -O ~/.local/bin/ssl-remote-control
    chmod +x ~/.local/bin/ssl-remote-control
    systemctl --user start ssl-remote-control.service
}

function installBrowser() {
    # https://blog.r0b.io/post/minimal-rpi-kiosk/
    sudo apt-get install --no-install-recommends -y \
        xserver-xorg-video-all xserver-xorg-input-all xserver-xorg-core xinit x11-xserver-utils \
        unclutter \
        xrandr \
        chromium-browser

    # Enable Auto-Login on console
    mkdir -p /etc/systemd/system/getty@tty1.service.d
    sudo cp "${SCRIPT_DIR}/autologin.conf" /etc/systemd/system/getty@tty1.service.d/autologin.conf

    # Configure to run browser when X starts
    cp "${SCRIPT_DIR}/.xinitrc" ~/.xinitrc

    # Configure to run X on tty1
    cp "${SCRIPT_DIR}/.bash_profile" ~/.bash_profile
}

function configurePi() {
    sudo cp "${SCRIPT_DIR}/config.txt" /boot/config.txt
}

sudo apt-get update
installService
installRemoteControl
installBrowser
configurePi

echo "You may need to restart the system to apply some settings. Reboot now? (y/n)"
read -r response
if [[ "${response}" == "y" ]]; then
    sudo reboot
fi
