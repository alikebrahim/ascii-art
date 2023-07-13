#!/bin/bash
release_file=/etc/os-release 

cat $release_file

if grep "Arch" $release_file
then
    sudo pacman -Syu
fi

if grep "debian" $release_file
then
    sudo apt update
fi