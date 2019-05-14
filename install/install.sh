#!/bin/sh

# Clone the repository and extract it.
wget -O ~/ddd.tar.gz https://github.com/luke-park/ddd/archive/v${VERSION}.tar.gz
tar -xzf ~/ddd.tar.gz -C ~/
mv ~/ddd-${VERSION} ~/ddd-repo

# Create the ddd directory and copy over stuff we need.
mv ~/ddd-repo/bin/ddds ~/ddds
chmod u+x ~/ddds
mkdir ~/ddd
mv ~/ddd-repo/install/ddds.toml ~/ddd/ddds.toml

# Add the service
sudo mv ~/ddd-repo/install/dddsd.service /etc/systemd/system/dddsd.service
sudo systemctl enable dddsd

# Cleanup
rm -R ~/ddd-repo
rm ~/ddd.tar.gz