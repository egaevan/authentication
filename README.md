# Authentication Mini Project (TDD Practice)

Dokumen ini berisi **Use Case dan Test Case** untuk mini project *Authentication* yang digunakan khusus untuk **latihan Test-Driven Development (TDD)**.

> ⚠️ Catatan penting
>
> * Project ini **dummy project**, bukan real production system
> * Semua test berfokus pada **domain logic**, bukan framework / database
> * Repository, session store, dan hashing **boleh in-memory atau mock**

---

## 🎯 Tujuan Project

* Melatih praktik **Red → Green → Refactor**
* Membiasakan menulis test **sebelum implementasi**
* Memahami edge case umum pada authentication system

---

## 🔐 Use Case 1 — Register User

### Deskripsi

User mendaftarkan akun baru menggunakan email dan password.

### Test Cases

#### 1.1 Register dengan data valid

**Given**

* Email belum terdaftar
* Password memenuhi aturan minimal

**When**

* Register dipanggil

**Then**

* User berhasil dibuat
* Password disimpan dalam bentuk hash
* Status user = active

```
Test: should create user when email and password are valid
```

---

#### 1.2 Register dengan email duplikat

**Given**

* Email sudah terdaftar

**When**

* Register dipanggil

**Then**

* Gagal
* Error: email already exists

```
Test: should fail when email already exists
```

---

#### 1.3 Register dengan format email tidak valid

**Given**

* Email tidak valid

**When**

* Register dipanggil

**Then**

* Gagal
* Error: invalid email format

```
Test: should reject invalid email format
```

---

#### 1.4 Register dengan password lemah

**Given**

* Password terlalu pendek / lemah

**When**

* Register dipanggil

**Then**

* Gagal
* Error: password too weak

```
Test: should reject weak password
```

---

## 🔑 Use Case 2 — Login User

### Deskripsi

User login menggunakan email dan password.

### Test Cases

#### 2.1 Login dengan credential benar

**Given**

* User terdaftar
* Password benar

**When**

* Login dipanggil

**Then**

* Login sukses
* Token/session dibuat

```
Test: should login successfully with correct credentials
```

---

#### 2.2 Login dengan password salah

**Given**

* Email terdaftar
* Password salah

**When**

* Login dipanggil

**Then**

* Gagal
* Error: invalid credentials

```
Test: should fail login with wrong password
```

---

#### 2.3 Login dengan email tidak terdaftar

**Given**

* Email tidak ditemukan

**When**

* Login dipanggil

**Then**

* Gagal
* Error: invalid credentials

```
Test: should fail login when email not found
```

---

#### 2.4 Login dengan user inactive

**Given**

* User status = inactive / locked

**When**

* Login dipanggil

**Then**

* Gagal
* Error: user inactive

```
Test: should prevent login for inactive user
```

---

## 🔁 Use Case 3 — Logout User

### Deskripsi

User logout dan token menjadi tidak valid.

### Test Cases

#### 3.1 Logout dengan token valid

**Given**

* User sudah login
* Token valid

**When**

* Logout dipanggil

**Then**

* Token dihapus / invalid

```
Test: should invalidate token on logout
```

---

#### 3.2 Logout dengan token tidak valid

**Given**

* Token tidak dikenal

**When**

* Logout dipanggil

**Then**

* Gagal
* Error: invalid token

```
Test: should fail logout with invalid token
```

---

## 🔄 Use Case 4 — Validate Token

### Deskripsi

Sistem memvalidasi token sebelum mengakses resource.

### Test Cases

#### 4.1 Validasi token aktif

**Given**

* Token aktif

**When**

* Validate token dipanggil

**Then**

* User ID dikembalikan

```
Test: should return user id for valid token
```

---

#### 4.2 Validasi token expired

**Given**

* Token sudah expired

**When**

* Validate token dipanggil

**Then**

* Gagal
* Error: token expired

```
Test: should reject expired token
```

---

#### 4.3 Validasi token revoked

**Given**

* Token sudah di-logout

**When**

* Validate token dipanggil

**Then**

* Gagal
* Error: token revoked

```
Test: should reject revoked token
```

---

## 🔐 Use Case 5 — Change Password

### Deskripsi

User mengganti password lama dengan password baru.

### Test Cases

#### 5.1 Ganti password dengan old password benar

**Given**

* User login
* Old password benar

**When**

* Change password dipanggil

**Then**

* Password baru tersimpan
* Password lama tidak valid

```
Test: should change password when old password is correct
```

---

#### 5.2 Ganti password dengan old password salah

**Given**

* Old password salah

**When**

* Change password dipanggil

**Then**

* Gagal
* Error: invalid old password

```
Test: should fail change password with wrong old password
```

---

## 🚫 Use Case 6 — Lock User After Failed Login (Optional)

### Deskripsi

User dikunci setelah gagal login berkali-kali.

### Test Cases

#### 6.1 Lock user setelah N kali gagal login

**Given**

* User gagal login sebanyak N kali

**When**

* Login kembali

**Then**

* User dikunci
* Error: account locked

```
Test: should lock account after multiple failed attempts
```

---

## 📌 Catatan TDD Penting

* Semua test harus:

    * Cepat
    * Deterministik
    * Mudah dibaca
* Test menguji **behavior**, bukan implementasi
* Dependency eksternal **wajib dimock**

---

## ✅ Rekomendasi Urutan Implementasi (TDD)

1. Register
2. Login (success & fail)
3. Validate token
4. Logout
5. Change password
6. Lock user (advanced)

---

## 🧪 Definition of Done (Dummy Project)

* Semua use case memiliki test
* Test hijau sebelum refactor
* Tidak ada database sungguhan
* Seluruh logic dapat dijalankan via unit test

---

> Dokumen ini dimaksudkan untuk dimasukkan langsung ke dalam `README.md` sebagai panduan TDD mini project Authentication.
