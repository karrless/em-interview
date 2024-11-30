ALTER TABLE songs
ALTER COLUMN release_date TYPE TEXT
USING TO_CHAR(release_date, 'DD.MM.YYYY');