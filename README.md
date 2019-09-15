# FESPAY

## Run

```shell script
go build
cp .env.example .env
vi .env
vi fespay.service
sudo cp fespay.service /etc/systemd/system/
sudo service fespay start
```

