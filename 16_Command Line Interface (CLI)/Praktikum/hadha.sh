#!/bin/bash

hadha="$1"
hadha7873="$2"
hadha-k-66059b266="$3"

mkdir "$hadha at $(date)"

mkdir -p "$hadha at $(date)/about_me/personal"
mkdir -p "$hadha at $(date)/about_me/professional"

echo "https://www.facebook.com/$2" > "$hadha at $(date)/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$hadha at $(date)/about_me/professional/linkedin.txt"

mkdir "$hadha at $(date)/my_friends"

curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$1 at $(date)/my_friends/list_of_my_friends.txt"

mkdir "$hadha at $(date)/my_system_info"

echo "User: $(whoami)" >> "$hadha at $(date)/my_system_info/about_this_laptop.txt"
echo "$(uname -a)" >> "$hadha at $(date)/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$hadha at $(date)/my_system_info/internet_connection.txt"

cmd //c "tree /F"


