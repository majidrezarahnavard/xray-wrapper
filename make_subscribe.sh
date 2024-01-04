#!/bin/bash

install_dir=/root/xray-configuration

cp $install_dir/config.json /usr/local/etc/xray/config.json


rm -rf /var/www/html/subs/s.*
cp  $install_dir/subs/s.* /var/www/html/


# Restart xray service
systemctl restart xray