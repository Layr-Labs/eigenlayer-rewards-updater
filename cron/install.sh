#!/usr/bin/env bash

mkdir -p /etc/payment-updater || true
cp payment-updater.sh /usr/local/bin

cp payment-updater.service /etc/systemd/system
cp payment-updater.timer /etc/systemd/system

systemctl enable payment-updater.timer
systemctl start payment-updater.timer

cp config.yml.tpl /etc/payment-updater/config.yml

echo "Make sure to fill out config.yml"
