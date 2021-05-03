-- Database: DungeonFetcher
-- DROP TABLE [IF EXISTS] Item;


-- DROP DATABASE "DungeonFetcher";

CREATE DATABASE "DungeonFetcher"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_Canada.1252'
    LC_CTYPE = 'English_Canada.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
	
CREATE TABLE IF NOT EXISTS Item (
   ID VARCHAR(20) PRIMARY KEY,
   ITEM_NAME VARCHAR(30) NOT NULL,
   MIN_LEVEL INTEGER,
   MAX_LEVEL INTEGER,
   CONSUMABLE BOOLEAN
);
	

	
