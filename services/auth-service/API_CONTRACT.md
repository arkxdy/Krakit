# 🌐 API Contract (v1)

## Base URL

```
/api/v1
```

---

# 📦 Standard Response Format

## ✅ Success

```json
{
  "success": true,
  "data": {},
  "message": "optional"
}
```

## ❌ Error

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message"
  }
}
```

---

# 🔐 AUTH APIs

## 1. Signup

**POST** `/auth/signup`

### Request

```json
{
  "email": "user@email.com",
  "password": "StrongPassword123",
  "first_name": "Aditya",
  "last_name": "Kr"
}
```

### Response

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@email.com",
      "full_name": "Aditya Kr",
      "role": "candidate",
      "plan": "free"
    },
    "access_token": "jwt",
    "expires_in": 900
  }
}
```

> 🔐 Refresh token is stored in **HttpOnly cookie**

---

## 2. Login

**POST** `/auth/login`

### Request

```json
{
  "email": "user@email.com",
  "password": "password"
}
```

### Response

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@email.com",
      "full_name": "Aditya Kr",
      "role": "admin",
      "plan": "premium"
    },
    "access_token": "jwt",
    "expires_in": 900
  }
}
```

---

## 3. Refresh Token

**POST** `/auth/refresh`

### Request

```json
{}
```

> Uses refresh token from cookie

### Response

```json
{
  "success": true,
  "data": {
    "access_token": "new_jwt",
    "expires_in": 900
  }
}
```

---

## 4. Logout (Current Session)

**POST** `/auth/logout`

### Response

```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

---

## 5. Logout All Devices

**POST** `/auth/logout-all`

### Headers

```
Authorization: Bearer <access_token>
```

### Response

```json
{
  "success": true,
  "message": "All sessions revoked"
}
```

---

# 👤 USER APIs

## 1. Get Current User

**GET** `/users/me`

### Headers

```
Authorization: Bearer <access_token>
```

### Response

```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "email": "user@email.com",
    "first_name": "Aditya",
    "last_name": "Kr",
    "full_name": "Aditya Kr",
    "role": "admin",
    "plan": "premium",
    "last_login_at": "timestamp"
  }
}
```

---

## 2. Update Profile

**PUT** `/users/me`

### Request

```json
{
  "first_name": "Aditya",
  "last_name": "Kumar"
}
```

### Response

```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "full_name": "Aditya Kumar"
  }
}
```

---

## 3. Change Password

**POST** `/users/change-password`

### Request

```json
{
  "old_password": "old_password",
  "new_password": "NewStrongPassword123"
}
```

### Response

```json
{
  "success": true,
  "message": "Password updated"
}
```

---

# 🔄 SESSION APIs

## 1. List Sessions

**GET** `/sessions`

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "platform": "web",
      "created_at": "timestamp",
      "expires_at": "timestamp",
      "is_current": true
    }
  ]
}
```

---

## 2. Revoke Session

**DELETE** `/sessions/{session_id}`

### Response

```json
{
  "success": true,
  "message": "Session revoked"
}
```

---

# 🔐 RBAC APIs (Admin)

## 1. List Permissions

**GET** `/permissions`

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "name": "user.read",
      "description": "Read users"
    }
  ]
}
```

---

## 2. Assign Permission to Role

**POST** `/roles/{role}/permissions`

### Request

```json
{
  "permission_id": "uuid"
}
```

---

## 3. Remove Permission

**DELETE** `/roles/{role}/permissions/{permission_id}`

---

# 🧠 JWT Contract

All services must rely on this payload:

```json
{
  "sub": "user_id",
  "email": "user@email.com",
  "role": "admin",
  "plan": "premium",
  "permissions": [
    "user.read",
    "exam.create"
  ],
  "jti": "session_id",
  "exp": 1710000000,
  "iat": 1700000000
}
```

---

# 🔌 Service Responsibilities

Each service must:

1. Verify JWT using **public key (RS256)**
2. Extract claims
3. Enforce permissions

---

# 🔥 HTTP Status Codes

| Case         | Code |
| ------------ | ---- |
| Success      | 200  |
| Created      | 201  |
| Bad Request  | 400  |
| Unauthorized | 401  |
| Forbidden    | 403  |
| Not Found    | 404  |
| Conflict     | 409  |

---

# 🚀 Notes

* Access token → short-lived (~15 min)
* Refresh token → HttpOnly cookie
* Use refresh token rotation
* Never expose password or refresh token in API response
