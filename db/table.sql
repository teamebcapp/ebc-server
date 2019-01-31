CREATE SCHEMA info;

CREATE SEQUENCE info.user_user_seq_seq;

CREATE SEQUENCE info.business_card_bc_seq_seq;
/**********************************/
/* Table Name: info.user */
/**********************************/
CREATE TABLE info.user(
		user_seq INT NOT NULL default nextval('info.user_user_seq_seq'::regclass),
		user_id VARCHAR(256) NOT NULL PRIMARY KEY,
		password VARCHAR(512) NOT NULL,
		name VARCHAR(128),
		company VARCHAR(128),
		position VARCHAR(128),
		duty VARCHAR(128),
		phone VARCHAR(16),
		email VARCHAR(256),
  CONSTRAINT IDX_user_1 UNIQUE (user_seq)
);

/**********************************/
/* Table Name: info.business_card */
/**********************************/
CREATE TABLE info.business_card(
		bc_seq INT NOT NULL PRIMARY KEY default nextval('info.business_card_bc_seq_seq'::regclass),
		user_Id VARCHAR(256) NOT NULL REFERENCES info.user (user_id),
		name VARCHAR(128),
		company VARCHAR(128),
		depart VARCHAR(256),
		team VARCHAR(128),
		position VARCHAR(128),
		duty VARCHAR(128),
		phone VARCHAR(16),
		tel VARCHAR(16),
		fax VARCHAR(32),
		adress VARCHAR(512),
		email VARCHAR(256),
		priority INT NOT NULL,
  CONSTRAINT IDX_business_card_1 UNIQUE (bc_seq)
);


CREATE INDEX IDX_user_2 ON info.user (user_id);

CREATE INDEX IDX_business_card_2 ON info.business_card (user_id);


