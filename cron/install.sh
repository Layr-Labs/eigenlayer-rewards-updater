#!/usr/bin/env bash

mkdir -p /etc/payment-updater || true
chown -R ubuntu:ubuntu /etc/payment-updater
cp ./cron/payment-updater.sh /usr/local/bin

cp ./cron/payment-updater.service /etc/systemd/system
cp ./cron/payment-updater.timer /etc/systemd/system

systemctl enable payment-updater.timer
systemctl start payment-updater.timer

cp config.yml.tpl ./config.yml


echo "Make sure to fill out config.yml"
