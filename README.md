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

## 고민거리

- `PATCH /register`는 `cardCode`로 체크하나 아님 `tokenID`로 체크하나?
  - 일단 개발 속도를 위해 `cardCode`로 체크
- [ ] limit, page 구현
- [x] 동결, 정지 구현 
- [ ] 어드민 로그