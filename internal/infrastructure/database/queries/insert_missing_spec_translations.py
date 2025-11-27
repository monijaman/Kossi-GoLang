import os
import psycopg2
from dotenv import load_dotenv

# Always load .env.production from gocrit_server root
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
ENV_PATH = os.path.abspath(os.path.join(SCRIPT_DIR, '..', '..', '..', '.env.production'))
load_dotenv(dotenv_path=ENV_PATH)

DATABASE_URL = os.getenv('DATABASE_URL')
if not DATABASE_URL:
    raise RuntimeError(f"DATABASE_URL not found in {ENV_PATH}")

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