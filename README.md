# FESPAY

페스페이는 학교 축제용 간편결제시스템입니다. 학생의 학생증의 바코드를 인식해 결제합니다. Vue.js로 짜여져 있습니다.

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
