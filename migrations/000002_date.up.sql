ALTER TABLE songs
ALTER COLUMN release_date TYPE DATE
USING TO_DATE(release_date, 'DD.MM.YYYY');