go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag

# @mangosteen package
wire gen ./internal/app
swag init --generalInfo ./bin/mangosteen/main.go --output ./docs/swagger

# @cmd mangosteen package
# native build
go build

# build for linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build

# deploy
# @cmd mangosteen package without cgo
scp mangosteen root@170.0.35.150:applications/z7z8
scp conf.toml root@170.0.35.150:applications/z7z8

# @cmd mangosteen src with cgo mangosteen.zip include vendor
git config --global core.autocrlf input
scp mangosteen.zip root@170.0.35.150:src
unzip mangosteen.zip -d mangosteen

cd mangosteen
export GO111MODULE=on
export GO111MODULE=off
/bin/cp -rf vendor/. /root/go/src
go build

# supervisor
pip install supervisor-XXX.wheel
mkdir /etc/supervisor
echo_supervisord_conf > /etc/supervisor/supervisord.conf

# modify at least 2 lines
[include]
files = /etc/supervisor/config.d/*.ini

supervisord -c /etc/supervisor/supervisord.conf

# supervisorctl

vim /etc/supervisor/config.d/mangosteen.ini
~~~~~~~~~~~~~~~~~~~~~
[program:mangosteen]
command=/root/applications/z7z8/mangosteen serve
stdout_logfile=/root/applications/z7z8/log/supervisor_mangosteen.out
autostart=true
autorestart=true
startsecs=5
priority=1
stopasgroup=true
killasgroup=true
~~~~~~~~~~~~~~~~~~~~

supervisorctl status
supervisorctl stop mangosteen
supervisorctl start mangosteen
supervisorctl restart mangosteen
supervisorctl reread
supervisorctl update

# systemctl
vim /etc/systemd/system/mangosteen.service
~~~~~~~~~~~~~~~~~~~~~~~~~~
[Unit]
Description=mangosteen server
[Service]
WorkingDirectory=/root/applications/z7z8
ExecStart=/root/applications/z7z8/mangosteen serve -c /root/applications/z7z8/conf.toml
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=on-failure
Type=notify
[Install]
Alias=mangosteen.service
~~~~~~~~~~~~~~~~~~~~~~~~~~

mangosteen.service
WorkingDirectory

systemctl status mangosteen -l
systemctl start mangosteen
systemctl restart mangosteen
systemctl stop mangosteen

go test -v ./internal/apps/unqualified/services -test.run TestUnqualSaveService_SaveUnqualHistoriesByDate

go test -v ./internal/apps/unqualified/services -test.run TestRollBreakStatService_GetRollBreakStatResult
go test -v ./internal/apps/fsp/dao -test.run TestLevel2AssuringDao_GetAssuringValueRecordByCoilId
go test -v ./internal/apps/fsp/dao -test.run TestCoilMainDao_getRecordsByDate

