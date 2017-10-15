DROP TABLE IF EXISTS card CASCADE;
DROP TABLE IF EXISTS wallet CASCADE;
DROP TABLE IF EXISTS person CASCADE;


CREATE TABLE person
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    username VARCHAR(120),
    email VARCHAR(120),
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
    cc_number VARCHAR(16),
    cc_due_date DATE,
    cc_expiration_date DATE,
    cc_cvv INTEGER,
    cc_limit DECIMAL NOT NULL,
    wallet_id INTEGER  NOT NULL,
    FOREIGN KEY (wallet_id) REFERENCES wallet (id) ON DELETE CASCADE
);

INSERT INTO person (name, username, email, password) VALUES ('Allison V.','allisonverdam','allison@g.com','$2a$14$GUpis37i8Z26V2GrtfuJie2jnHOjXUd/fMWMPAy7OEWTZ2xytKTuO');
INSERT INTO person (name, username, email, password) VALUES ('Jullyana C.','ju','jullyana@g.com','$2a$14$7P/mU6Z3Atlano2RmbQdRe5TPSzdkjcUelPAIK8iUjKCV0A0u8aAa');
INSERT INTO person (name, username, email, password) VALUES ('Beatriz B.','bia','beatriz@g.com','$2a$14$xN.rdPUvB.3GMoGHmt4bce/96XL5wjz72r0gb.6TmeNWfkMX6aTT.');

INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 1);
INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 2);
INSERT INTO wallet (real_limit, maximum_limit, person_id) VALUES (0, 0, 3);

INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341230', '2016/01/12', '2017/10/12', 123, '300', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341231', '2016/02/11', '2017/10/12', 123, '400', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341232', '2016/03/10', '2017/10/12', 123, '500', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341233', '2016/04/09', '2017/10/12', 123, '600', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341234', '2016/05/08', '2017/10/12', 123, '300', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341235', '2016/01/07', '2017/10/12', 123, '400', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341236', '2016/02/06', '2017/10/12', 123, '500', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341237', '2016/03/05', '2017/10/12', 123, '600', 3);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341238', '2016/04/04', '2017/10/12', 123, '700', 3);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, wallet_id) VALUES ('1234123412341239', '2016/05/03', '2017/10/12', 123, '800', 3);