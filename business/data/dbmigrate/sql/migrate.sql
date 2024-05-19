-- Version: 1.01
-- Description: Create table users
CREATE TABLE users (
	user_id       UUID        NOT NULL,
	name          TEXT        NOT NULL,
	email         TEXT UNIQUE NOT NULL,
	roles         TEXT[]      NOT NULL,
	password_hash TEXT        NOT NULL,
   department    TEXT        NULL,
   enabled       BOOLEAN     NOT NULL,
	date_created  TIMESTAMP   NOT NULL,
	date_updated  TIMESTAMP   NOT NULL,

	PRIMARY KEY (user_id)
);

-- Version: 1.02
-- Description: Create table products
CREATE TABLE products (
	product_id   UUID           NOT NULL,
   user_id      UUID           NOT NULL,
	name         TEXT           NOT NULL,
	cost         NUMERIC(10, 2) NOT NULL,
	quantity     INT            NOT NULL,
	date_created TIMESTAMP      NOT NULL,
	date_updated TIMESTAMP      NOT NULL,

	PRIMARY KEY (product_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Add user_summary view.
CREATE OR REPLACE VIEW user_summary AS
SELECT
    u.user_id   AS user_id,
	 u.name      AS user_name,
    COUNT(p.*)  AS total_count,
    SUM(p.cost) AS total_cost
FROM
    users AS u
JOIN
    products AS p ON p.user_id = u.user_id
GROUP BY
    u.user_id

-- Version: 1.04
-- Description: Create table subjects
CREATE TABLE subjects (
	subject_id       UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	code         	  TEXT  		  NOT NULL,
	year         	  TEXT 		  NOT NULL,
	academic_year    TEXT  		  NOT NULL,
	semester         TEXT        NOT NULL,
	instructor	     TEXT  		  NOT NULL,
	exam			     INT  		  NOT NULL,
	practical		  INT  		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (name, semester, year),
	UNIQUE (code, semester, year),
	PRIMARY KEY (subject_id)
);

-- Version: 1.05
-- Description: Create table students
CREATE TABLE students (
	student_id       UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	email         	  TEXT 		  NOT NULL    UNIQUE,
	roll_number      TEXT  		  NOT NULL,
	phone_number	  TEXT  		  NOT NULL,
	year         	  TEXT 		  NOT NULL,
	academic_year    TEXT  		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (roll_number, year, academic_year),
	PRIMARY KEY (student_id)
);

-- Version: 1.06
-- Description: Create table student_subjects
CREATE TABLE student_subjects (
	student_subject_id       UUID        NOT NULL,
	subject_id					 UUID        NOT NULL,
	student_id					 UUID        NOT NULL,
	mark          	  			 INT         NULL,
	date_created  	 			 TIMESTAMP   NOT NULL,
	date_updated  	 			 TIMESTAMP   NOT NULL,

	UNIQUE (subject_id, student_id),
	PRIMARY KEY (student_subject_id),
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE,
	FOREIGN KEY (student_id) REFERENCES students(student_id) ON DELETE CASCADE
);

-- Version: 1.07
-- Description: Create table course_outlines
CREATE TABLE course_outlines (
	co_id		        UUID        NOT NULL,
	subject_id		  UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (co_id),
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);

-- Version: 1.08
-- Description: Create table graduate_attributes
CREATE TABLE graduate_attributes (
	ga_id		        UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (ga_id)
);

-- Version: 1.09
-- Description: Create table co_ga
CREATE TABLE co_ga (
	co_ga_id		     UUID        NOT NULL,
	co_id		        UUID        NOT NULL,
	ga_id		        UUID        NOT NULL,
	mark          	  INT         NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (co_id, ga_id),
	PRIMARY KEY (co_ga_id),
	FOREIGN KEY (co_id) REFERENCES course_outlines(co_id) ON DELETE CASCADE,
	FOREIGN KEY (ga_id) REFERENCES graduate_attributes(ga_id) ON DELETE CASCADE
);

-- Version: 1.10
-- Description: Create table marks
CREATE TABLE marks (
	mark_id		  	   UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	type          	  TEXT        NOT NULL,
	instance         INT 		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (mark_id)
);

-- Version: 1.11
-- Description: Create table co_mark
CREATE TABLE co_marks (
	co_mark_id	  	  UUID        NOT NULL,
	mark_id		  	  UUID        NOT NULL,
	co_id			  	  UUID        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (co_mark_id),
	FOREIGN KEY (mark_id) REFERENCES marks(mark_id) ON DELETE CASCADE,
	FOREIGN KEY (co_id) REFERENCES course_outlines(co_id) ON DELETE CASCADE
);

-- Version: 1.12
-- Description: Update students table roll_number TEXT to INT
ALTER TABLE students DROP COLUMN roll_number;
ALTER TABLE students ADD COLUMN roll_number INT;

-- Version: 1.13
-- Description: Update ga table slug TEXT
ALTER TABLE graduate_attributes ADD COLUMN slug CHAR(3) UNIQUE NOT NULL;

-- Version: 1.14
-- Description: Update mark table UNIQUE
ALTER TABLE marks ADD UNIQUE (name, instance);

-- Version: 1.15
-- Description: Update subjects table UNIQUE
ALTER TABLE subjects DROP CONSTRAINT subjects_name_semester_year_key;
ALTER TABLE subjects DROP CONSTRAINT subjects_code_semester_year_key;
ALTER TABLE subjects ADD CONSTRAINT subjects_name_semester_academic_year_key UNIQUE (name, semester, academic_year);
ALTER TABLE subjects ADD CONSTRAINT subjects_code_semester_academic_year_key UNIQUE (code, semester, academic_year);

-- Version: 1.16
-- Description: Update co_marks table UNIQUE
ALTER TABLE co_marks ADD UNIQUE (mark_id, co_id);