USE bear_sighting_db;

-- sightingsテーブルが空の場合のみデータを挿入
INSERT INTO sightings (latitude, longitude, city, area, date, time, situation, details)
SELECT * FROM (
    SELECT
        35.6895 AS latitude, 139.6917 AS longitude, 'Tokyo' AS city, 'Shinjuku' AS area,
        '2024-11-28' AS date, '09:30:00' AS time,
        'A bear was seen crossing the hiking trail.' AS situation,
        'The bear appeared to be an adult, approximately 200kg.' AS details
    UNION ALL
    SELECT
        35.0116, 135.7681, 'Kyoto', 'Gion',
        '2024-11-27', '14:15:00',
        'Bear tracks were discovered near the riverbank.',
        'Tracks seemed fresh, possibly from the morning.'
    UNION ALL
    SELECT
        43.0642, 141.3469, 'Sapporo', 'Odori Park',
        '2024-11-26', '18:45:00',
        'A resident reported hearing bear sounds during the night.',
        NULL
    UNION ALL
    SELECT
        34.6937, 135.5023, 'Osaka', 'Namba',
        '2024-11-25', '21:00:00',
        'A bear was sighted near a garbage disposal area.',
        'Bear was eating leftover food.'
    UNION ALL
    SELECT
        36.2048, 138.2529, 'Nagano', 'Matsumoto',
        '2024-11-24', '10:30:00',
        'Claw marks were found on trees near the hiking trail.',
        'Claw marks measured approximately 15cm.'
) AS tmp
WHERE NOT EXISTS (SELECT 1 FROM sightings);
