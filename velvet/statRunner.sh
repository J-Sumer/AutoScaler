#! /bin/bash
# printf "Memory\t\tDisk\t\tCPU\n"
end=$((SECONDS+3600))
URL=http://$1:$2/stat
echo $URL
while [ $SECONDS -lt $end ]; do
timestamp=$(date +%s)
MEMORY=$(free -m | awk 'NR==2{printf "%.2f%%", $3*100/$2 }')
# DISK=$(df -h | awk '$NF=="/"{printf "%s", $5}')
CPU=$[100-$(vmstat 1 2|tail -1|awk '{print $15}')]
# echo "$MEMORY$DISK$CPU"
echo "$timestamp: CPU: $CPU%  Memory: $MEMORY"
curl -X POST $URL -H "Content-Type: application/json" -d '{"Id": 79, "status": 3}'  
sleep 0.5
done