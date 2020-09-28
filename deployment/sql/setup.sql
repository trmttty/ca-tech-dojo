DROP TABLE IF EXISTS users;
CREATE TABLE users (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	token VARCHAR(255) NOT NULL,
	PRIMARY KEY (id) 
);

INSERT INTO users (name, token) VALUES ("tatsuya", "12345678910");
INSERT INTO users (name, token) VALUES ("toru", "12345678910");
INSERT INTO users (name, token) VALUES ("satoshi", "12345678910");
