#!/bin/sh
sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config
service ssh restart

# Create training user.
useradd -d /home/training -m -s /bin/bash training

echo training:ogoh8xohQu | chpasswd

echo "training ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/training
