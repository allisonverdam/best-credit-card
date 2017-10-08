DROP TABLE IF EXISTS card CASCADE;
DROP TABLE IF EXISTS person CASCADE;


CREATE TABLE person
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    username VARCHAR(120),
    email VARCHAR(120),
    password VARCHAR(120)
);
   

CREATE TABLE card
(
    id SERIAL PRIMARY KEY,
    cc_number VARCHAR(16),
    cc_due_date DATE,
    cc_expiration_date DATE,
    cc_cvv INTEGER,
    cc_limit DECIMAL NOT NULL,
    person_id INTEGER  NOT NULL,
    FOREIGN KEY (person_id) REFERENCES person (id) ON DELETE CASCADE
);

INSERT INTO person (name, username, email, password) VALUES ('Allison V.','allisonverdam','allison@g.com','1234');
INSERT INTO person (name, username, email, password) VALUES ('Jullyana C.','ju','jullyana@g.com','12345678');


INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341230', '2016/10/12', '2017/10/12', 123, '1000', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341231', '2016/10/11', '2017/10/12', 123, '1000', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341232', '2016/10/10', '2017/10/12', 123, '1000', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341233', '2016/10/09', '2017/10/12', 123, '1000', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341234', '2016/10/08', '2017/10/12', 123, '1000', 1);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341235', '2016/10/07', '2017/10/12', 123, '1000', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341236', '2016/10/06', '2017/10/12', 123, '1000', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341237', '2016/10/05', '2017/10/12', 123, '1000', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341238', '2016/10/04', '2017/10/12', 123, '1000', 2);
INSERT INTO card (cc_number, cc_due_date, cc_expiration_date, cc_cvv, cc_limit, person_id) VALUES ('1234123412341239', '2016/10/03', '2017/10/12', 123, '1000', 2);