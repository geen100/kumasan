CREATE DATABASE IF NOT EXISTS bear_sighting_db;
USE bear_sighting_db;

CREATE TABLE IF NOT EXISTS sightings (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    latitude DECIMAL(9,6) NOT NULL,  
    longitude DECIMAL(9,6) NOT NULL, 
    city VARCHAR(255) NOT NULL,      -- 市町村名
    area VARCHAR(255),               -- 目撃地区の詳細（例: 山の名前、公園名など）
    date DATE NOT NULL,              -- 目撃日
    time TIME NOT NULL,              -- 目撃時刻
    situation TEXT NOT NULL,         -- 目撃の状況（例: クマの動作や様子）
    details TEXT,                    -- 追加の詳細情報
    classification ENUM('目撃', '痕跡', '噂') NOT NULL, -- 区分
    user_id BIGINT UNSIGNED,         -- 目撃情報を登録したユーザーID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,      -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新日時
    PRIMARY KEY (id)
);

CREATE INDEX idx_latitude_longitude ON sightings (latitude, longitude);  -- 緯度・経度のインデックス
CREATE INDEX idx_date ON sightings (date);  -- 日付のインデックス
CREATE INDEX idx_city ON sightings (city);  -- 市町村のインデックス
