#!/usr/bin/env bash

mkdir -p /etc/eigenlayer-rewards-updater || true
chown -R ubuntu:ubuntu /etc/eigenlayer-rewards-updater
cp ./cron/eigenlayer-rewards-updater.sh /usr/local/bin

cp ./cron/eigenlayer-rewards-updater.service /etc/systemd/system
cp ./cron/eigenlayer-rewards-updater.timer /etc/systemd/system

systemctl enable eigenlayer-rewards-updater.timer
systemctl start eigenlayer-rewards-updater.timer

cp config.yml.tpl ./config.yml


echo "Make sure to fill out config.yml"
