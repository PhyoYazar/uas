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

	UNIQUE (name, semester, academic_year),
	UNIQUE (code, semester, academic_year),
	PRIMARY KEY (subject_id)
);

-- Version: 1.03
-- Description: Create table students
CREATE TABLE students (
	student_id       UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	email         	  TEXT 		  NOT NULL    UNIQUE,
	roll_number      INT  		  NOT NULL,
	phone_number	  TEXT  		  NOT NULL,
	year         	  TEXT 		  NOT NULL,
	academic_year    TEXT  		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (roll_number, year, academic_year),
	PRIMARY KEY (student_id)
);

-- Version: 1.04
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

-- Version: 1.05
-- Description: Create table course_outlines
CREATE TABLE course_outlines (
	co_id		        UUID        NOT NULL,
	subject_id		  UUID        NOT NULL,
	instance 		  INT			  NOT NULL,
	name          	  TEXT        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (co_id),
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);

-- Version: 1.06
-- Description: Create table graduate_attributes
CREATE TABLE graduate_attributes (
	ga_id		        UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	slug 				  TEXT 		  NOT NULL     UNIQUE,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (ga_id)
);

-- Version: 1.07
-- Description: Create table co_ga
CREATE TABLE co_ga (
	co_ga_id		     UUID        NOT NULL,
	co_id		        UUID        NOT NULL,
	ga_id		        UUID        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (co_id, ga_id),
	PRIMARY KEY (co_ga_id),
	FOREIGN KEY (co_id) REFERENCES course_outlines(co_id) ON DELETE CASCADE,
	FOREIGN KEY (ga_id) REFERENCES graduate_attributes(ga_id) ON DELETE CASCADE
);

-- Version: 1.08
-- Description: Create table attributes
CREATE TABLE attributes (
	attribute_id	  UUID        NOT NULL,
	name          	  TEXT        NOT NULL,
	type          	  TEXT        NOT NULL,
	instance         INT 		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	UNIQUE (name, instance, type),
	PRIMARY KEY (attribute_id)
);

-- Version: 1.09
-- Description: Create table marks
CREATE TABLE marks (
	mark_id	  		  UUID        NOT NULL,
	attribute_id	  UUID        NOT NULL,
	ga_id			  	  UUID        NOT NULL,
	subject_id		  UUID 		  NOT NULL,
	mark 				  INT 		  NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (mark_id),
	UNIQUE (attribute_id, subject_id, ga_id),
	FOREIGN KEY (attribute_id) REFERENCES attributes(attribute_id) ON DELETE CASCADE,
	FOREIGN KEY (ga_id) REFERENCES graduate_attributes(ga_id) ON DELETE CASCADE,
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);

-- Version: 1.10
-- Description: Create table co_attributes
CREATE TABLE co_attributes (
	co_attribute_id  UUID 		  NOT NULL,
	attribute_id	  UUID        NOT NULL,
	co_id			  	  UUID        NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (co_attribute_id),
	UNIQUE (attribute_id, co_id),
	FOREIGN KEY (attribute_id) REFERENCES attributes(attribute_id) ON DELETE CASCADE,
	FOREIGN KEY (co_id) REFERENCES course_outlines(co_id) ON DELETE CASCADE
);

-- Version: 1.11
-- Description: Update course_outlines table UNIQUE
ALTER TABLE course_outlines ADD UNIQUE (subject_id, instance, name);

-- Version: 1.12
-- Description: Add auto increment column in graduate attributes table
ALTER TABLE graduate_attributes ADD COLUMN incrementing_column SERIAL;

-- Version: 1.13
-- Description: Add unique (subject_id & instance) to the course outlines table
ALTER TABLE course_outlines ADD CONSTRAINT subject_id_instance_key UNIQUE (subject_id, instance);

-- Version: 1.14
-- Description: Add unique (subject_id & attribute_id) to the marks table
ALTER TABLE marks DROP CONSTRAINT marks_attribute_id_subject_id_ga_id_key;
ALTER TABLE marks ADD CONSTRAINT marks_attribute_id_subject_id_key UNIQUE (attribute_id, subject_id);

-- Version: 1.15
-- Description: roll back version 1.14 to original (attribute_id & subject_id & ga_id) to the marks table
ALTER TABLE marks DROP CONSTRAINT marks_attribute_id_subject_id_key;
ALTER TABLE marks ADD CONSTRAINT marks_attribute_id_subject_id_ga_id_key UNIQUE (attribute_id, subject_id, ga_id);

-- Version: 1.16
-- Description: Drop email and phone number columns from student table & create full_mark table
ALTER TABLE students DROP COLUMN email, DROP COLUMN phone_number;
CREATE TABLE full_marks (
	full_mark_id	  UUID        NOT NULL,
	attribute_id	  UUID        NOT NULL,
	subject_id		  UUID 		  NOT NULL,
	mark 				  INT 		  NOT NULL,
	date_created  	  TIMESTAMP   NOT NULL,
	date_updated  	  TIMESTAMP   NOT NULL,

	PRIMARY KEY (full_mark_id),
	UNIQUE (subject_id, attribute_id),
	FOREIGN KEY (attribute_id) REFERENCES attributes(attribute_id) ON DELETE CASCADE,
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);

-- Version: 1.17
-- Description: Drop student_subjects table & create student_marks table
DROP TABLE IF EXISTS student_subjects;
CREATE TABLE student_marks (
	student_mark_id	  	UUID        NOT NULL,
	attribute_id	  		UUID        NOT NULL,
	student_id		  		UUID        NOT NULL,
	subject_id		  		UUID 		   NOT NULL,
	mark 				  		INT 		   NULL,
	date_created  	  		TIMESTAMP   NOT NULL,
	date_updated  	  		TIMESTAMP   NOT NULL,

	PRIMARY KEY (student_mark_id),
	UNIQUE (subject_id, student_id, attribute_id),
	FOREIGN KEY (attribute_id) REFERENCES attributes(attribute_id) ON DELETE CASCADE,
	FOREIGN KEY (student_id) REFERENCES students(student_id) ON DELETE CASCADE,
	FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE CASCADE
);


-- Version: 1.18
-- Description: change student name column to student number column
ALTER TABLE students DROP COLUMN name;
ALTER TABLE students ADD COLUMN student_number INT NOT NULL;
