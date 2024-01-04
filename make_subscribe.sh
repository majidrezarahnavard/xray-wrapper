#!/bin/bash

install_dir=/root/xray-configuration

cp $install_dir/config.json /usr/local/etc/xray/config.json


rm -rf /var/www/html/subs/s.*
cp  $install_dir/s.* /var/www/html/subs/


# Restart xray service
systemctl restart xray