# 🌈 Adorable Go API Project ʕ •ᴥ•ʔ

Welcome to our love-filled Go API wonderland! This is a super cute RESTful API project using Gin framework and MongoDB ✨
Let's manage user data together and create a warm and happy backend service! (◕‿◕✿)

## 🎨 Special Features

- 🌟 Beautiful RESTful API design
- 🍃 MongoDB data storage
- 📚 Cute Swagger API documentation
- 🎯 Gentle error handling mechanism
- 🧪 Thoughtful unit testing
- 📝 Detailed logging
- 🔄 HATEOAS-compliant API responses for better discoverability

## 🎮 Preparation for Adventure

Make sure you have these companions:
- 🚀 Go 1.21+ magic tools
- 🗄️ MongoDB magical database
- 🐙 Git version control buddy

## 🌱 Plant the Project Seed

1. First, bring the project home:
```bash
git clone https://github.com/jebylinjbjob/go-api_for_main.git
cd go-api_for_main
```

2. Summon the required sprites:
```bash
go mod download
```

3. Make sure the MongoDB sprite is awake (it likes to stay at port 27017)

4. Let the project shine:
```bash
go run main.go
```

## 🎯 API Sprites

### 👥 User Management Squad
- `GET /api/v1/users` - Summon all users ✨
- `POST /api/v1/users` - Create new friends 🎉
- `GET /api/v1/users/:id` - Find specific friend 🔍
- `PUT /api/v1/users/:id` - Help friend change clothes 👕
- `DELETE /api/v1/users/:id` - Say goodbye (wave) 👋

### 🎪 System World
- `GET /ping` - Poke to see if we're awake 👉
- `GET /swagger/*any` - Browse our magic book 📖

### 🧪 Testing Environment

In testing, we change the magic path from v1 to test, let's play together!

Here are the test environment sprites:
- `GET /api/test/users` - Call test users to play 🎈
- `POST /api/test/users` - Invite new friends to play 🎪
- `GET /api/test/users/:id` - Find who's playing 🔮
- `PUT /api/test/users/:id` - Let friends change their style 🎭
- `DELETE /api/test/users/:id` - Play hide and seek (for now, say goodbye) 🎪


## 📚 Magic User Manual

Want to see more magic? After running the service, visit:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing Laboratory

Test all magic:
```bash
go test ./... -v
```

Test specific spells:
```bash
go test ./test -v
```

## 🏰 Project Castle Structure

```
.
├── controllers/     # 🎮 Control center
├── models/         # 📝 Data model house
│   ├── response.go  # 🔄 HATEOAS response structures
│   └── user.go      # 👤 User data model
├── routes/         # 🛣️ Route map
├── docs/          # 📚 Magic library
├── test/          # 🧪 Laboratory
├── main.go        # 🎯 Main entrance
└── README.md      # 📖 User manual
```

## 🔄 HATEOAS API Responses

Our API follows the HATEOAS (Hypermedia as the Engine of Application State) principle, allowing clients to navigate the API dynamically through hypermedia links:

- 🔗 All responses contain relevant links for possible actions
- 🧭 Clients can discover available operations through response links
- 🔍 No need to hardcode API endpoints in client applications
- 🎭 Supports API evolution with minimal client changes

Examples of our HATEOAS responses:
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  },
  "_links": [
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "self",
      "method": "GET",
      "title": "Get User Info"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "update",
      "method": "PUT",
      "title": "Update User"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "delete",
      "method": "DELETE",
      "title": "Delete User"
    }
  ]
}
```

## 🎨 Error Handling Helper

We use different expressions for different situations:

- 200: ✅ Success!
- 400: ❌ Oops, request has an issue
- 404: 🔍 Can't find what you want
- 500: 😱 Server sneezed
- 503: 🏥 Database sprite is resting

## 🎛️ Environment Setting Tools

### 🗄️ MongoDB Settings
- `MONGODB_URI`: MongoDB sprite's home (default is mongodb://localhost:27017)
- `MONGODB_DATABASE`: Database name (default is go_api_db)
- `MONGODB_USERNAME`: MongoDB authentication username (optional)
- `MONGODB_PASSWORD`: MongoDB authentication password (optional)
- `MONGODB_TIMEOUT`: MongoDB connection timeout in seconds (default is 10)

💫 **MongoDB Security Tips**:
1. Two ways to provide MongoDB authentication:
   - Using `MONGODB_USERNAME` and `MONGODB_PASSWORD` environment variables
   - Or directly in `MONGODB_URI` like: `mongodb://username:password@localhost:27017`
2. Recommended to store these secrets in a `.env` file
3. Never commit passwords to Git repository
4. Use a proper secrets management system in production

### 🖥️ Server Settings
- `PORT`: Service door location (default is port 8080)
- `GIN_MODE`: Server running mode (default is debug)

## 🚧 New Facilities Under Construction

- 🔐 User authentication and authorization system
- 🚦 Request rate limiter
- 💾 Cache memory space
- ✨ More data validation magic
- 📤 File upload portal

## 🌟 Join Our Adventure

Want to create magic together?

1. 🍴 Fork your own branch
2. 🌱 Create new features
3. ✨ Submit your magic
4. 🚀 Push to your branch
5. 🎉 Create Pull Request

## 📜 Magic License

MIT License (Love-filled open source license 💝)

---
Made with ❤️, hope you can feel the warmth too! (｡♥‿♥｡) 


[![image](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg))](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg)
