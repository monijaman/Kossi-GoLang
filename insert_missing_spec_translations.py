import os
import psycopg2
from dotenv import load_dotenv

# Load .env.production
env_path = os.path.join(os.path.dirname(__file__), '.env.production')
load_dotenv(dotenv_path=env_path)

DATABASE_URL = os.getenv('DATABASE_URL')

conn = psycopg2.connect(DATABASE_URL)
cur = conn.cursor()

cur.execute("SELECT id, value FROM specifications")
specs = cur.fetchall()

for spec_id, value in specs:
    cur.execute("""
        SELECT 1 FROM specification_translations
        WHERE specification_id = %s AND language = 'bn'
    """, (spec_id,))
    if not cur.fetchone():
        # Insert translation (replace with actual Bangla translation if available)
        bangla_value = f"[Bangla] {value}"
        cur.execute("""
            INSERT INTO specification_translations (specification_id, language, value)
            VALUES (%s, 'bn', %s)
        """, (spec_id, bangla_value))
        print(f"Inserted Bangla translation for spec id {spec_id}")

conn.commit()
cur.close()
conn.close()
