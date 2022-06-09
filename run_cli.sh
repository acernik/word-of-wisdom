LOCAL_IP=$(ifconfig | grep 'inet ' | grep -Fv 127.0.0.1 | awk '{print $2}')
echo $LOCAL_IP
export CLIENT_ADDRESS=$LOCAL_IP