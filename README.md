# FESPAY (페스페이)

<div align="center">
  <img src="./generate-card/logo2-cropped.png" alt="FESPAY Logo" width="600"/>
  
  <strong>Digital Payment System for the 2019 festival of [Suncheon Wangun Middle School](https://gwangcheol.hs.jne.kr/).</strong>
</div>

- Built in Go and Vue.js.
- Operated at Dec. 23th, 2019.
- Accomplished US$2000+ transactions (KRW 2,088,600).

## Motivation

In the festivals of SWMS, selling foods or products in booths was allowed but using real bills was prohibited(perhaps teachers thought students were too young to manage their bills). So the festival committee sold 'Wangun Coupons' (business card sized paper, 1 Wangun Coupon = 500 won) before the festival started until the 2018 festival. Those coupons were only valid for that year's festival, so they would often end up on the floor when the festival was over.

Ordering lots of paper coupons and throwing away them looked like environmental desctruction to me. Moreover, it was hard to calculate and locate usage of coupons. A new solution was needed.

## How It Worked: Preparation

The main idea was simple: let students use their ID cards as prepaid cards. 

Every students of Suncheon Wangun Middle School(SWMS) had a student ID card. Each student ID had a unique barcode that was used to check out books. Thanks to the cooperation of the school library, I could get a PDF of the students' bar-code list. Converted them into a csv file and imported to the database(in this case it was PostgreSQL). Constructed an API server that could process payments and built a web app that could scan bar-codes.

In order to keep a refill lane during the festival from being crowded, I asked each class president to conduct prepaid applications before the festival started.

| Prepaid Application for Students |
|---|
| <img width="600" alt="Prepaid Application for Students" src="https://github.com/user-attachments/assets/473f2ccd-cfb3-44b1-a6f1-4703b6054aff"> |

Members of the booths were granted to acess the web app. 

## How It Worked: During The Festival

| Main | Scan & Pay |
|---|---|
| <img src="https://github.com/user-attachments/assets/5c184c6c-f0ea-4b69-8b3d-a7909f454cfb" width="300" /> | <video autoplay loop muted src="https://github.com/user-attachments/assets/cbedd07f-e752-48f1-a025-36eb44a0848c" width="300" /> |

When a customer student wanted to buy a product, a seller student opened the app and entered an amount. The customer showed the seller their student ID, the seller scaned it, and the checkout was complete.

If the balance of the customer was insufficent or the account of the customer was suspended(this wasn't used actually), the checkout was blocked and the error message was shown up.

## Run

```bash
go build
cp .env.example .env
vi .env
vi fespay.service
sudo cp fespay.service /etc/systemd/system/
sudo service fespay start
```
