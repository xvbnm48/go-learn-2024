# Defer golang

Defer adalah salah satu fitur di golang yang digunakan untuk menunda eksekusi suatu fungsi sampai blok fungsi selesai dieksekusi. Defer biasanya digunakan untuk menutup file, koneksi database, atau melakukan sesuatu yang perlu dilakukan setelah blok fungsi selesai dieksekusi.

singkatnya defer ini tuh function yang bisa kita jadwalkan untuk dieksekusi setelah function lain selesai dieksekusi.

namun defer akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi.

# panic

Panic adalah function yang dapat digunakan untuk memberhentikan eksekusi program. Panic biasanya digunakan untuk menghentikan eksekusi program ketika terjadi error yang tidak bisa diatasi, dan defer akan tetap di panggil.

# recover 
recover ini tuh function yang bisa digunakan untuk menangkap data panic, jadi dengan recover ini proses panic akan terhenti, dan program tetep jalan