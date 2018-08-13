
DROP TABLE IF EXISTS "TEAMJJ";
CREATE TABLE public."TEAMJJ"
(
"ID" serial,
"NAME" character varying(10) UNIQUECOLLATE pg_catalog."default" NOT NULL,
"TEAMNAME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"AGE" character varying(10)  NOT NULL,
"EMAIL" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"DETAIL" character varying(100)  NOT NULL,
"PROJECTSID" integer REFERENCES products (PROJECTSID),
CONSTRAINT "TEAMJJ_pkey" PRIMARY KEY ("ID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;

DROP TABLE IF EXISTS "PROJECTS";
CREATE TABLE public."PROJECTS"
(
"PROJECTSID" serial,
"NAME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"URL" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"DETAIL" character varying(100)  NOT NULL
CONSTRAINT "PROJECTS_pkey" PRIMARY KEY ("ID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;


GRANT ALL PRIVILEGES ON DATABASE "teamjj" to admin;
ALTER USER admin WITH SUPERUSER;
