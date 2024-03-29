maksud dari len(strings)-1 itu apa ? kenapa kita ambil panjang strings yang diterima namun di - 1

ChatGPT
Penggunaan len(strings)-1 dalam loop for merupakan cara umum untuk menghindari akses indeks yang melebihi panjang array atau slice. Dalam konteks ini, len(strings) adalah panjang array strings. Karena indeks dimulai dari 0, kita ingin memastikan bahwa saat kita mengakses elemen array menggunakan indeks dalam loop, kita tidak mencoba mengakses elemen di luar batas array, yang akan menyebabkan kesalahan (out of range).

Misalnya, jika panjang array strings adalah 4, maka indeksnya akan berkisar dari 0 hingga 3. Dengan menggunakan len(strings)-1 sebagai batas loop, kita memastikan bahwa kita tidak akan mencoba mengakses elemen di indeks 4 (yang sebenarnya tidak ada dalam array tersebut).

Jadi, len(strings)-1 digunakan agar loop berjalan hingga indeks terakhir yang valid pada array, sehingga menghindari kesalahan indeks di luar batas array.