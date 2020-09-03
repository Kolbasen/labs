VERSION= "1.0.0"
PID_FILE = /tmp/labs.pid

start:
	@echo "Starting..."
	go run server.go & echo $$! > $(PID_FILE)

stop:
	@pkill -P `cat $(PID_FILE)` || true

restart: stop start
	@echo "Restarting..."
  
serve:	start	
	@echo "Serving..."
	fswatch -or --event Created --event Updated --event Renamed ./ | xargs -n1 -I {} make restart
  
.PHONY: start stop restart serve