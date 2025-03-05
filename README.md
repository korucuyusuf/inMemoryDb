
<img width="450" alt="image" src="https://github.com/user-attachments/assets/eef41b64-daaf-411a-b93e-11eacfb2f0df" /> <img width="450" alt="image" src="https://github.com/user-attachments/assets/b998d328-2167-479c-b017-70fba9917681" />



# Sample-Api

Bu proje, in-memory bir veritabanı kullanarak API'leri yönetmek ve verileri saklamak için tasarlanmıştır. API'ler aracılığıyla veri alabilir (GET) ve yeni veriler oluşturabilirsiniz (CREATE). Bu proje geliştiriciler ve API kullanıcıları için basit ve verimli bir çözüm sunmayı amaçlamaktadır.

## Özellikler

- **In-Memory Database**: Tüm veriler bellekte saklanır, veritabanı sürekli değildir TTL süresi ayarlanmalı.
- **API GET isteği**: Veritabanındaki mevcut verileri almanıza olanak tanır.
- **API CREATE isteği**: Yeni verileri veritabanına eklemenizi sağlar.

## Teknolojiler

- golang
- In-memory DB (buntDb)

## Kullanım

### Kurulum

1. Bu repo'yu bilgisayarınıza klonlayın:
    ```bash
    git clone https://github.com/korucuyusuf/inMemoryDb.git
    ```


   

Sunucu çalıştıktan sonra API'leri kullanmaya başlayabilirsiniz.

### API Endpoints

#### GET /api/items

Veritabanındaki tüm öğeleri listeleyen GET isteği.

**Curl İsteği:**
```bash
curl -X GET http://localhost:8080/get
```
### API Endpoints

#### Create /api/items

Kayıt oluşturma isteği.

**Curl İsteği:**
```bash
curl --location 'http://localhost:8080/create' \
--header 'Content-Type: application/json' \
--data '{"name":"KORUCU"}'
```
