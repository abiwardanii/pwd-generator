import random
import string

def generate_password(length=12):
    """Generate a secure password."""
    # Menggabungkan karakter huruf besar, huruf kecil, angka, dan simbol
    characters = string.ascii_letters + string.digits + string.punctuation

    # Mengacak urutan karakter
    secure_password = ''.join(random.choice(characters) for _ in range(length))
    return secure_password

# Menghasilkan password dengan panjang 12 karakter
password = generate_password()
print("Kata sandi baru:", password)
