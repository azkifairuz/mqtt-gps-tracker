# 📡 MQTT GPS Tracker API Documentation

Dokumentasi sementara untuk API HTTP yang digunakan dalam proyek `mqtt-gps-tracker`.

---

## 🌐 Base URL
[text](http://api-gps-tracker.azkidev.my.id)


---

## 📍 Endpoints

### 1. `GET /ping`

**Deskripsi:**  
Endpoint untuk memastikan bahwa API aktif.

**Response:**

```json
{
  "message": "pong"
}

### 1. `GET /ping`

**Deskripsi:**  
Endpoint untuk memastikan bahwa API aktif.

**Response:**

```json
{
  "message": "pong"
}
```
### 2. `POST /device/power-status`

**Deskripsi:**  
Digunakan untuk mengirim status power dari device (ON/OFF atau lainnya).

**Request Body:**

```json
{
  "deviceId": "string",
  "status": "string"
}
```
**Response:**

```json
{
  "message": "status received",
  "device": "string",
  "status": "string"
}

```

### 3. `POST /device/position`

**Deskripsi:**  
Digunakan untuk mengirim posisi device berupa latitude dan longitude.


**Request Body:**

```json
{
  "deviceId": "string",
  "lat": 0.0,
  "long": 0.0
}

```
**Response:**

```json
{
  "message": "status received",
  "device": "string",
  "lat": 0.0,
  "long": 0.0
}
```
### 4. `POST /device/status`

**Deskripsi:**  
Digunakan untuk mengirim status device berupa nilai voltase dan kWh.

**Request:**

```json
{
  "deviceId": "string",
  "kwh": "string",
  "volt": "string"
}
```
**Response:**

```json
{
  "message": "status received",
  "device": "string",
  "kwh": "string",
  "volt": "string"
}

```
