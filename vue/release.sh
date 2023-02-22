#!/bin/bash -e

rm -rf ./dist/*
npm run build

sudo rm -rf /var/www/html/*
sudo mv ./dist/* /var/www/html/

sudo systemctl restart nginx
