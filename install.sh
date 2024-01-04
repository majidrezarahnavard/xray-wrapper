#!/bin/bash


install_dir=/root/xray-configuration
cd $install_dir

wget https://raw.githubusercontent.com/majidrezarahnavard/xray-wrapper/main/config.json
wget https://raw.githubusercontent.com/majidrezarahnavard/xray-wrapper/main/reinstall.sh
wget https://raw.githubusercontent.com/majidrezarahnavard/xray-wrapper/main/setting.json
wget https://raw.githubusercontent.com/majidrezarahnavard/xray-wrapper/main/make_subscribe.sh

#add permitions
sudo chmod +x $install_dir/reinstall.sh
sudo chmod +x $install_dir/make_subscribe.sh

rm -rf $install_dir/xray-telegram*
wget https://github.com/majidrezarahnavard/xray-wrapper/releases/download/v.1.0.0/xray-telegram
sudo chmod +x ./xray-telegram

rm -rf $install_dir/xray-wrapper*
wget https://github.com/majidrezarahnavard/xray-wrapper/releases/download/v.1.0.0/xray-wrapper
sudo chmod +x ./xray-wrapper

#instal monitoring
apt-get update
apt-get install nload
apt-get install htop
apt-get install iftop
apt-get install vnstat
apt-get install speedtest-cli
apt-get install net-tools
apt-get install git
apt-get install cron
apt-get install curl tar unzip jq -y
apt-get install -y jq


echo "net.ipv4.tcp_fastopen = 3" | sudo tee -a /etc/sysctl.conf
echo "net.core.default_qdisc = fq" | sudo tee -a /etc/sysctl.conf
echo "net.ipv4.tcp_congestion_control = bbr" | sudo tee -a /etc/sysctl.conf

sysctl -p


journalctl --vacuum-time=1d


timedatectl set-timezone UTC
timedatectl
echo "UTC" | sudo tee /etc/timezone
cat /etc/timezone


bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install --beta -u root



touch $install_dir/log.txt

cp $install_dir/config.json /usr/local/etc/xray/config.json



# Install apache2 and clone the website
apt-get install apache2

sudo sed -i 's/80/8090/' /etc/apache2/ports.conf
sudo sed -i 's/80/8090/' /etc/apache2/sites-enabled/000-default.conf #for your examples sake

service apache2 restart

cd /var/www/html/

git clone https://github.com/codingstella/vCard-personal-portfolio.git
cp -ar ./vCard-personal-portfolio/*  /var/www/html/
rm -rf ./vCard-personal-portfolio/

mkdir subs



Install cron job 
croncmd="cd $install_dir && $install_dir/xray-telegram > $install_dir/cronjob.log 2>&1"
cronjob="30 * * * * $croncmd"
( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -


