DROP TABLE IF EXISTS Item

CREATE TABLE Item (
	ID VARCHAR(20) NOT NULL PRIMARY KEY,
	ITEM_NAME VARCHAR(30),
	MIN_LEVEL INT,
	MAX_LEVEL INT,
	CONSUMABLE BIT
);

-- Add Default Items

INSERT INTO Item (ID, ITEM_NAME, MIN_LEVEL, MAX_LEVEL,CONSUMABLE)
VALUES ('yg79cgs8s7f2hi8pouhs','Health Potion',1,20,1),
('gifi0a60vaoezg99yjee','Ring of Health',1,5,0),
('072oi1xahcmh12y1llf7','Pathfinders Compass',15,20,0);

