DROP TABLE IF EXISTS card CASCADE;
DROP TABLE IF EXISTS wallet CASCADE;
DROP TABLE IF EXISTS person CASCADE;


CREATE TABLE person
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    username VARCHAR(120) UNIQUE,
    email VARCHAR(120) UNIQUE,
    password VARCHAR(120)
);


CREATE TABLE wallet
(
    id SERIAL PRIMARY KEY,
    real_limit DECIMAL NOT NULL,
    maximum_limit DECIMAL NOT NULL,
    person_id INTEGER  NOT NULL,
    FOREIGN KEY (person_id) REFERENCES person (id) ON DELETE CASCADE
);
   

CREATE TABLE card
(
    id SERIAL PRIMARY KEY,
    cc_number VARCHAR(16) NOT NULL UNIQUE,
    cc_due_date INTEGER NOT NULL,
    cc_expiration_month INTEGER NOT NULL,
    cc_expiration_year INTEGER NOT NULL,
    cc_cvv INTEGER NOT NULL,
    cc_real_limit DECIMAL NOT NULL,
    cc_current_limit DECIMAL NOT NULL,
    wallet_id INTEGER  NOT NULL,
    FOREIGN KEY (wallet_id) REFERENCES wallet (id) ON DELETE CASCADE
);

INSERT INTO person (name, username, email, password) VALUES ('Allison V.','allisonverdam','allison@g.com','$2a$14$GUpis37i8Z26V2GrtfuJie2jnHOjXUd/fMWMPAy7OEWTZ2xytKTuO');
INSERT INTO person (name, username, email, password) VALUES ('Jullyana C.','ju','jullyana@g.com','$2a$14$7P/mU6Z3Atlano2RmbQdRe5TPSzdkjcUelPAIK8iUjKCV0A0u8aAa');
INSERT INTO person (name, username, email, password) VALUES ('Beatriz B.','bia','beatriz@g.com','$2a$14$xN.rdPUvB.3GMoGHmt4bce/96XL5wjz72r0gb.6TmeNWfkMX6aTT.');

INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 1);
INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 2);
INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 3);

INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341230', 10, 06, 17, 123, 300, 180, 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341231', 10, 06, 17, 123, 400, 200, 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341232', 11, 08, 16, 123, 500, 450, 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341233', 12, 10, 14, 123, 600, 480, 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341234', 12, 01, 14, 123, 300, 120, 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341235', 13, 02, 15, 123, 400, 350, 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341236', 14, 03, 15, 123, 500, 480, 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341237', 15, 04, 16, 123, 600, 360, 3);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341238', 16, 05, 17, 123, 700, 670, 3);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_month, cc_expiration_year, cc_cvv, cc_real_limit, cc_current_limit, wallet_id) VALUES ('1234123412341239', 17, 06, 16, 123, 800, 600, 3);