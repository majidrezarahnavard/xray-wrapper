#!/bin/bash

apt update 

echo "Uninstalling..."


install_dir=/root/xray-configuration

# Remove files
rm $install_dir/config.json
rm $install_dir/s.*
rm /var/www/html/subs/s.*


vnstat > $install_dir/log.txt


ls -laht /var/log
echo "" > /var/log/syslog
echo "" > /var/log/syslog.1

journalctl --vacuum-time=1d



echo "Unistall DONE!"



bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install --beta -u root


