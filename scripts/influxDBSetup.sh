#! /bin/bash

# This script installs influxDB on the server including the InfluxDB CLI

wget -q https://repos.influxdata.com/influxdata-archive_compat.key

echo "-------------------Signing Key-------------------"
echo '393e8779c89ac8d958f81f942f9ad7fb82a25e133faddaf92e15b16e6ac9ce4c influxdata-archive_compat.key' | sha256sum -c && cat influxdata-archive_compat.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/influxdata-archive_compat.gpg > /dev/null
echo 'deb [signed-by=/etc/apt/trusted.gpg.d/influxdata-archive_compat.gpg] https://repos.influxdata.com/debian stable main' | sudo tee /etc/apt/sources.list.d/influxdata.list

echo "-------------------Updating apt repository-------------------"
sudo apt-get update -y

echo "-------------------Install InfluxDB-------------------"
sudo apt-get install influxdb2 -y

echo "-------------------Start influx DB-------------------"
sudo service influxdb start

echo "-------------------Check status of influxDB-------------------"
systemctl status influxdb

#influx setup -u USERNAME -p PASSWORD -t TOKEN -o ORGANIZATION_NAME -b BUCKET_NAME -f

echo "-------------------Setup influx DB with credentials-------------------"
influx setup -u admin -p password -t TOKEN -o NCSU -b ADS -f