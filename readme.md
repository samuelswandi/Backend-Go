# Samuel Christopher Swandi - Backend Developer - Technical Test Jobhun Internship 2023

## Assumptions

### Database
Terdapat banyak sekali asumsi yang ada untuk relasi yang terdapat pada database, dan untuk mempertahankan konsistensi, maka saya akan membuat relasi menjadi berhubungan seperti ini: <br>
1. Diadakannya tabel `mahasiswa_jurusan`, karena untuk mendapat detail mahasiswa, harus memiliki mapping ke tabel ini
2. Penghapusan mahasiswa tidak berpengaruh terhadap penghapusan hobi dan jurusan (hanya dihapus di mahasiswa_jurusan dan mahasiswa_hobi)

### Api request
1. API request untuk membuat mahasiswa baru akan membutuhkan nama, usia, gender, list jurusan, dan list hobi.
2. API updateMahasiswa hanya bisa update untuk nama, usia, dan gender. 
3. API mahasiswa akan mengembalikan nama, usia, gender, tanggal registrasi, jurusan, dan hobi

## List of API
| API | request | response | description
|-----|------|-----|----|
|`mahasiswa` | menggunakan parameter: <br> 1. nama (string) <br> 2. id (uint) <br> dan jika tidak ada parameter, maka akan mengembalikan seluruhnya|`{ "Message": "SUCCESS", "Data": [ { "Id": 1, "Nama": "Samuel Swandi", "Usia": 20, "Gender": 1, "TanggalRegistrasi": "2023-04-23T03:39:34Z", "Jurusan": "Ilmu Komputer, Teknik Informatika", "Hobi": "None" } ]}` | bertujuan untuk mendapatkan semua mahasiswa atau mahasiswa secara spesifik dengan query id, nama, jurusan 
`createMahasiswa` |``{ "Nama": "Samuel", "Usia": 20, "Gender": 1,"Jurusan": ["Teknik Informatika"],"Hobi": ["Membaca"]}` | `{ "Message": "SUCCESS", "Data": {   "Id": 6,   "Nama": "Samuel",   "Usia": 20,  "Gender": 1, "TanggalRegistrasi": "2023-04-23T03:04:30.806392+07:00"}}` | membuat mahasiswa baru
`updateMahasiswa` | `{"Id": 1,"Nama": "Samuel Swandi","Usia": 20,"Gender": 1 }` | `{ "Message": "SUCCESS", "Data": { "Id": 1, "Nama": "Samuel Swandi", "Usia": 20, "Gender": 1, "TanggalRegistrasi": "0001-01-01T00:00:00Z" }}` | update mahasiswa dengan id tertentu
`deleteMahasiswa` | hanya id di query param | success/tidak | menghapus mahasiswa dengan id tertentu

## Design pattern
Design pattern yang digunakan adalah DDD (Domain-Driven design) yang di mana terletak pemisahan terhadap handling request dan juga komunikasi dengan database. Hal tersebut ditujukan agar codebase ini dapat dikembangkan lebih lanjut seperti penambahan middleware, unit testing, bahkan microservices.


## Tools & Framework
- `Air` <br> 
used for hot reload
- `Docker` <br>
additional, used for deploy go & database concurrently
- `Echo` <br>
golang backend framework
- `Validator` <br>
go playground validator, to check if request existed or not


## How to start
1. Clone this repo

2. (SQL pre-requisite), add seeding located at .docker/init.sql then run `go run main.go`

3. (Docker pre-requisite) run `docker-compose up -d` 

## Author
Samuel Swandi <br>
https://www.linkedin.com/in/samuelswandi/