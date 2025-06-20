# 🌈 โปรเจค API Go แสนน่ารัก ʕ •ᴥ•ʔ

ยินดีต้อนรับสู่โลกแห่งความรักของ Go API! นี่คือโปรเจค RESTful API สุดน่ารักที่ใช้ Gin framework และ MongoDB ✨
มาจัดการข้อมูลผู้ใช้ด้วยกันและสร้างบริการแบ็คเอนด์ที่อบอุ่นและมีความสุขกันเถอะ! (◕‿◕✿)

## 🎨 คุณสมบัติพิเศษ

- 🌟 การออกแบบ RESTful API ที่สวยงาม
- 🍃 การจัดเก็บข้อมูลด้วย MongoDB
- 📚 เอกสาร Swagger API น่ารัก
- 🎯 กลไกการจัดการข้อผิดพลาดที่นุ่มนวล
- 🧪 การทดสอบหน่วยที่ใส่ใจ
- 📝 การบันทึกรายละเอียดที่ครบถ้วน
- 🔄 การตอบสนอง API ที่สอดคล้องกับหลักการ HATEOAS เพื่อการค้นพบที่ดีขึ้น

## 🎮 การเตรียมตัวก่อนการผจญภัย

ตรวจสอบให้แน่ใจว่าคุณมีเพื่อนร่วมทางเหล่านี้:
- 🚀 เครื่องมือวิเศษ Go 1.21+
- 🗄️ ฐานข้อมูลวิเศษ MongoDB
- 🐙 เพื่อนควบคุมเวอร์ชัน Git

## 🌱 ปลูกเมล็ดพันธุ์โปรเจค

1. ก่อนอื่น นำโปรเจคกลับบ้าน:
```bash
git clone https://github.com/jebylinjbjob/go-api_for_main.git
cd go-api_for_main
```

2. เรียกวิญญาณที่จำเป็น:
```bash
go mod download
```

3. ตรวจสอบให้แน่ใจว่าวิญญาณ MongoDB ตื่นแล้ว (มันชอบอยู่ที่พอร์ต 27017)

4. ให้โปรเจคเปล่งประกาย:
```bash
go run main.go
```

## 🎯 วิญญาณ API

### 👥 ทีมจัดการผู้ใช้
- `GET /api/v1/users` - เรียกผู้ใช้ทั้งหมด ✨
- `POST /api/v1/users` - สร้างเพื่อนใหม่ 🎉
- `GET /api/v1/users/:id` - ค้นหาเพื่อนเฉพาะ 🔍
- `PUT /api/v1/users/:id` - ช่วยเพื่อนเปลี่ยนเสื้อผ้า 👕
- `DELETE /api/v1/users/:id` - บอกลา (โบกมือ) 👋

### 🎪 โลกของระบบ
- `GET /ping` - แตะเพื่อดูว่าเรายังตื่นอยู่ไหม 👉
- `GET /swagger/*any` - เปิดดูหนังสือเวทมนตร์ของเรา 📖

### 🧪 สภาพแวดล้อมการทดสอบ

ในการทดสอบ เราเปลี่ยนทางเวทมนตร์จาก v1 เป็น test มาร่วมกันเล่นกันเถอะ!

นี่คือวิญญาณที่ทดสอบทั้งหมด:
- `GET /api/test/users` - เรียกผู้ใช้ทดสอบเพื่อเล่น 🎈
- `POST /api/test/users` - เชิญเพื่อนใหม่เพื่อเล่น 🎪
- `GET /api/test/users/:id` - ค้นหาเพื่อนที่เล่น 🔮
- `PUT /api/test/users/:id` - ให้เพื่อนเปลี่ยนรูปลักษณ์ 🎭
- `DELETE /api/test/users/:id` - เล่นหลบหนี (ตอนนี้ บอกลา) 🎪

## 📚 คู่มือการใช้เวทมนตร์

อยากเห็นเวทมนตร์เพิ่มเติมไหม? หลังจากรันบริการ เข้าไปที่:
```
http://localhost:8080/swagger/index.html
```

## 🧪 ห้องทดลอง

ทดสอบเวทมนตร์ทั้งหมด:
```bash
go test ./... -v
```

ทดสอบคาถาเฉพาะ:
```bash
go test ./test -v
```

## 🏰 โครงสร้างปราสาทโปรเจค

```
.
├── controllers/     # 🎮 ศูนย์ควบคุม
├── models/         # 📝 บ้านโมเดลข้อมูล
│   ├── response.go  # 🔄 โครงสร้างการตอบสนอง HATEOAS
│   └── user.go      # 👤 โมเดลข้อมูลผู้ใช้
├── routes/         # 🛣️ แผนที่เส้นทาง
├── docs/          # 📚 ห้องสมุดเวทมนตร์
├── test/          # 🧪 ห้องทดลอง
├── main.go        # 🎯 ทางเข้าหลัก
└── README.md      # 📖 คู่มือการใช้งาน
```

## 🔄 การตอบสนอง API แบบ HATEOAS

API ของเราปฏิบัติตามหลักการ HATEOAS (Hypermedia as the Engine of Application State) ซึ่งช่วยให้ไคลเอนต์สามารถนำทาง API แบบไดนามิกผ่านลิงก์ไฮเปอร์มีเดีย:

- 🔗 การตอบสนองทั้งหมดจะมีลิงก์ที่เกี่ยวข้องสำหรับการดำเนินการที่เป็นไปได้
- 🧭 ไคลเอนต์สามารถค้นพบการดำเนินการที่มีอยู่ผ่านลิงก์ในการตอบสนอง
- 🔍 ไม่จำเป็นต้องเขียนรหัส API endpoints อย่างตายตัวในแอปพลิเคชันไคลเอนต์
- 🎭 รองรับการพัฒนา API โดยที่มีการเปลี่ยนแปลงน้อยที่สุดสำหรับไคลเอนต์

ตัวอย่างการตอบสนองแบบ HATEOAS:
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "จอห์น โด",
    "email": "john@example.com",
    "age": 30
  },
  "_links": [
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "self",
      "method": "GET",
      "title": "ดูข้อมูลผู้ใช้"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "update",
      "method": "PUT",
      "title": "อัปเดตผู้ใช้"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "delete",
      "method": "DELETE",
      "title": "ลบผู้ใช้"
    }
  ]
}
```

## 🎨 ผู้ช่วยจัดการข้อผิดพลาด

เราใช้การแสดงออกที่แตกต่างกันสำหรับสถานการณ์ต่างๆ:

- 200: ✅ สำเร็จ!
- 400: ❌ อุ๊ปส์ คำขอมีปัญหา
- 404: 🔍 หาสิ่งที่ต้องการไม่พบ
- 500: 😱 เซิร์ฟเวอร์จาม
- 503: 🏥 วิญญาณฐานข้อมูลกำลังพักผ่อน

## 🎛️ เครื่องมือตั้งค่าสภาพแวดล้อม

สามารถกำหนดค่าต่างๆ ผ่านตัวแปรสภาพแวดล้อมดังนี้:

### 🖥️ การตั้งค่าเซิร์ฟเวอร์
- `PORT`: พอร์ตที่ใช้รันเซิร์ฟเวอร์ (ค่าเริ่มต้น: 8080)
- `GIN_MODE`: โหมดการทำงานของ Gin (ค่าเริ่มต้น: debug, ค่าที่เป็นไปได้: debug, release)

### 🗄️ การตั้งค่า MongoDB
- `MONGODB_URI`: URI สำหรับเชื่อมต่อ MongoDB (ค่าเริ่มต้น: mongodb://localhost:27017)
- `MONGODB_DATABASE`: ชื่อฐานข้อมูล (ค่าเริ่มต้น: go_api_db)
- `MONGODB_USERNAME`: ชื่อผู้ใช้สำหรับการยืนยันตัวตน MongoDB (ไม่บังคับ)
- `MONGODB_PASSWORD`: รหัสผ่านสำหรับการยืนยันตัวตน MongoDB (ไม่บังคับ)
- `MONGODB_TIMEOUT`: เวลาหมดเวลาการเชื่อมต่อเป็นวินาที (ค่าเริ่มต้น: 10)

💡 **หมายเหตุเกี่ยวกับความปลอดภัย MongoDB**:
1. สามารถกำหนดข้อมูลการยืนยันตัวตนได้ 2 วิธี:
   - ผ่านตัวแปร `MONGODB_USERNAME` และ `MONGODB_PASSWORD`
   - หรือรวมในตัวแปร `MONGODB_URI` เช่น: `mongodb://username:password@localhost:27017`
2. แนะนำให้ใช้ไฟล์ `.env` สำหรับเก็บข้อมูลที่ละเอียดอ่อน
3. อย่าเก็บรหัสผ่านใน Git repository
4. ในการผลิตจริง ควรใช้ระบบจัดการความลับ (secrets management)

### 🔐 การตั้งค่า JWT
- `JWT_SECRET_KEY`: คีย์ลับสำหรับการเข้ารหัส JWT
- `JWT_EXPIRATION_HOURS`: อายุของ token เป็นชั่วโมง (ค่าเริ่มต้น: 24)

### 📝 การตั้งค่าการบันทึก
- `LOG_LEVEL`: ระดับการบันทึก (ค่าเริ่มต้น: debug, ค่าที่เป็นไปได้: debug, info, warn, error)
- `LOG_FILE`: ที่อยู่ไฟล์บันทึก (ค่าเริ่มต้น: ./logs/app.log)

### 🚦 การตั้งค่าการจำกัดอัตรา
- `RATE_LIMIT_REQUESTS`: จำนวนคำขอที่อนุญาตต่อหน้าต่างเวลา (ค่าเริ่มต้น: 100)
- `RATE_LIMIT_WINDOW`: ขนาดหน้าต่างเวลาเป็นนาที (ค่าเริ่มต้น: 1)

### 🌐 การตั้งค่า CORS
- `ALLOWED_ORIGINS`: รายการ origin ที่อนุญาต คั่นด้วยเครื่องหมายจุลภาค (ค่าเริ่มต้น: http://localhost:3000,http://localhost:8080)

## 🌟 เข้าร่วมการผจญภัยของเรา

อยากสร้างเวทมนตร์ด้วยกันไหม?

1. 🍴 Fork สาขาของคุณเอง
2. 🌱 สร้างฟีเจอร์ใหม่
3. ✨ ส่งเวทมนตร์ของคุณ
4. 🚀 Push ไปยังสาขาของคุณ
5. 🎉 สร้าง Pull Request

## 📜 ใบอนุญาตเวทมนตร์

MIT License (ใบอนุญาตโอเพนซอร์สที่เต็มไปด้วยความรัก 💝)

---
สร้างด้วย ❤️ หวังว่าคุณจะรู้สึกถึงความอบอุ่นด้วย! (｡♥‿♥｡) 


[![image](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg))](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg)
