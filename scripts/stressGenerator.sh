
#!/bin/bash

# START=$(date +%s)

# while [ $(( $(date +%s) - $START )) -lt 3600 ]
# do  
# # Generate a random number between 10 and 30 (in minutes)
# random_time=$(shuf -i 10-20 -n 1)

# echo "Sleeping for $random_time minutes..."
# sleep "$((random_time * 60))"

# random_stress=$(shuf -i 1-3 -n 1)

# # Run your command here
# echo "Executing your command now!"
# pumba stress -d 1m rubis1

# done

pumba stress -d 1m rubis3

sleep 15m
pumba stress -d 2m rubis3

sleep 25m
pumba stress -d 1m rubis3

