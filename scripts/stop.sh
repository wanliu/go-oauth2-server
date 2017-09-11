OAUTH=$(pgrep oauth2-server-linux)
if [ -n $OAUTH ]; then
	. scripts/start.sh
else
	kill -2 $OAUTH
fi
service nginx stop
