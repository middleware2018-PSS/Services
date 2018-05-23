--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: back2school; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA back2school;


ALTER SCHEMA back2school OWNER TO postgres;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: receiver; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.receiver AS ENUM (
    'student',
    'parent',
    'teacher',
    'general'
);


ALTER TYPE public.receiver OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: accounts; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.accounts (
    "user" text NOT NULL,
    password text NOT NULL,
    kind text NOT NULL,
    id integer NOT NULL
);


ALTER TABLE back2school.accounts OWNER TO postgres;

--
-- Name: admins; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.admins (
    id integer NOT NULL,
    info text DEFAULT ' '::text,
    name text DEFAULT ' '::text,
    surname text DEFAULT ' '::text
);


ALTER TABLE back2school.admins OWNER TO postgres;

--
-- Name: admins_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.admins_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.admins_id_seq OWNER TO postgres;

--
-- Name: admins_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.admins_id_seq OWNED BY back2school.admins.id;


--
-- Name: appointments; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.appointments (
    student integer NOT NULL,
    teacher integer NOT NULL,
    location text DEFAULT ''::text,
    "time" timestamp without time zone NOT NULL,
    id integer NOT NULL
);


ALTER TABLE back2school.appointments OWNER TO postgres;

--
-- Name: appointments_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.appointments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.appointments_id_seq OWNER TO postgres;

--
-- Name: appointments_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.appointments_id_seq OWNED BY back2school.appointments.id;


--
-- Name: classes; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.classes (
    id integer NOT NULL,
    year integer,
    section text DEFAULT ' '::text,
    info text,
    grade integer
);


ALTER TABLE back2school.classes OWNER TO postgres;

--
-- Name: classes_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.classes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.classes_id_seq OWNER TO postgres;

--
-- Name: classes_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.classes_id_seq OWNED BY back2school.classes.id;


--
-- Name: enrolled; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.enrolled (
    student integer NOT NULL,
    class integer NOT NULL
);


ALTER TABLE back2school.enrolled OWNER TO postgres;

--
-- Name: grades; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.grades (
    student integer,
    grade integer,
    subject text DEFAULT ''::text,
    date timestamp without time zone NOT NULL,
    teacher integer NOT NULL,
    id integer NOT NULL
);


ALTER TABLE back2school.grades OWNER TO postgres;

--
-- Name: grades_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.grades_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.grades_id_seq OWNER TO postgres;

--
-- Name: grades_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.grades_id_seq OWNED BY back2school.grades.id;


--
-- Name: isparent; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.isparent (
    parent integer NOT NULL,
    student integer NOT NULL
);


ALTER TABLE back2school.isparent OWNER TO postgres;

--
-- Name: notification; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.notification (
    id integer NOT NULL,
    receiver integer,
    message text DEFAULT ''::text,
    "time" time without time zone,
    receiver_kind public.receiver DEFAULT 'general'::public.receiver NOT NULL
);


ALTER TABLE back2school.notification OWNER TO postgres;

--
-- Name: notification_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.notification_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.notification_id_seq OWNER TO postgres;

--
-- Name: notification_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.notification_id_seq OWNED BY back2school.notification.id;


--
-- Name: parents; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.parents (
    id integer NOT NULL,
    name text DEFAULT ''::text,
    surname text DEFAULT ''::text,
    mail text DEFAULT ''::text,
    info text DEFAULT ''::text
);


ALTER TABLE back2school.parents OWNER TO postgres;

--
-- Name: parents_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.parents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.parents_id_seq OWNER TO postgres;

--
-- Name: parents_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.parents_id_seq OWNED BY back2school.parents.id;


--
-- Name: payments; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.payments (
    id integer NOT NULL,
    amount integer,
    student integer,
    payed boolean DEFAULT false,
    reason text DEFAULT ''::text,
    emitted timestamp without time zone DEFAULT now()
);


ALTER TABLE back2school.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.payments_id_seq OWNED BY back2school.payments.id;


--
-- Name: students; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.students (
    id integer NOT NULL,
    name text DEFAULT ''::text,
    surname text DEFAULT ''::text,
    mail text DEFAULT ''::text,
    info text DEFAULT ''::text
);


ALTER TABLE back2school.students OWNER TO postgres;

--
-- Name: students_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.students_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.students_id_seq OWNER TO postgres;

--
-- Name: students_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.students_id_seq OWNED BY back2school.students.id;


--
-- Name: subjects; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.subjects (
    id text DEFAULT ''::text NOT NULL
);


ALTER TABLE back2school.subjects OWNER TO postgres;

--
-- Name: teachers; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.teachers (
    id integer NOT NULL,
    name text DEFAULT ''::text,
    mail text DEFAULT ''::text,
    info text DEFAULT ''::text,
    surname text DEFAULT ''::text
);


ALTER TABLE back2school.teachers OWNER TO postgres;

--
-- Name: teachers_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.teachers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.teachers_id_seq OWNER TO postgres;

--
-- Name: teachers_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.teachers_id_seq OWNED BY back2school.teachers.id;


--
-- Name: teaches; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.teaches (
    teacher integer NOT NULL,
    subject text DEFAULT ''::text NOT NULL,
    class integer NOT NULL
);


ALTER TABLE back2school.teaches OWNER TO postgres;

--
-- Name: timetable; Type: TABLE; Schema: back2school; Owner: postgres
--

CREATE TABLE back2school.timetable (
    class integer NOT NULL,
    subject text DEFAULT ''::text NOT NULL,
    location text DEFAULT ''::text,
    start time without time zone,
    "end" time without time zone,
    info text DEFAULT ''::text,
    id integer NOT NULL
);


ALTER TABLE back2school.timetable OWNER TO postgres;

--
-- Name: timetable_id_seq; Type: SEQUENCE; Schema: back2school; Owner: postgres
--

CREATE SEQUENCE back2school.timetable_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE back2school.timetable_id_seq OWNER TO postgres;

--
-- Name: timetable_id_seq; Type: SEQUENCE OWNED BY; Schema: back2school; Owner: postgres
--

ALTER SEQUENCE back2school.timetable_id_seq OWNED BY back2school.timetable.id;


--
-- Name: admins id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.admins ALTER COLUMN id SET DEFAULT nextval('back2school.admins_id_seq'::regclass);


--
-- Name: appointments id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.appointments ALTER COLUMN id SET DEFAULT nextval('back2school.appointments_id_seq'::regclass);


--
-- Name: classes id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.classes ALTER COLUMN id SET DEFAULT nextval('back2school.classes_id_seq'::regclass);


--
-- Name: grades id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.grades ALTER COLUMN id SET DEFAULT nextval('back2school.grades_id_seq'::regclass);


--
-- Name: notification id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.notification ALTER COLUMN id SET DEFAULT nextval('back2school.notification_id_seq'::regclass);


--
-- Name: parents id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.parents ALTER COLUMN id SET DEFAULT nextval('back2school.parents_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.payments ALTER COLUMN id SET DEFAULT nextval('back2school.payments_id_seq'::regclass);


--
-- Name: students id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.students ALTER COLUMN id SET DEFAULT nextval('back2school.students_id_seq'::regclass);


--
-- Name: teachers id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teachers ALTER COLUMN id SET DEFAULT nextval('back2school.teachers_id_seq'::regclass);


--
-- Name: timetable id; Type: DEFAULT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.timetable ALTER COLUMN id SET DEFAULT nextval('back2school.timetable_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.accounts ("user", password, kind, id) FROM stdin;
pippo	pluto	Parent	3
a	p	Admin	2
admin	password	Admin	1
\.


--
-- Data for Name: admins; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.admins (id, info, name, surname) FROM stdin;
\.


--
-- Data for Name: appointments; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.appointments (student, teacher, location, "time", id) FROM stdin;
1	1	\N	2018-05-06 12:17:24.289	2
\.


--
-- Data for Name: classes; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.classes (id, year, section, info, grade) FROM stdin;
2	1994	a	1	5
3	1999	a	1	4
\.


--
-- Data for Name: enrolled; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.enrolled (student, class) FROM stdin;
1	2
\.


--
-- Data for Name: grades; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.grades (student, grade, subject, date, teacher, id) FROM stdin;
1	10	science	2019-05-06 11:17:55.784	1	1
\.


--
-- Data for Name: isparent; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.isparent (parent, student) FROM stdin;
3	1
\.


--
-- Data for Name: notification; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.notification (id, receiver, message, "time", receiver_kind) FROM stdin;
1	1	prova	\N	student
2	\N	prova	\N	general
\.


--
-- Data for Name: parents; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.parents (id, name, surname, mail, info) FROM stdin;
4	pi	pi	\N	\N
3	pippa	pippo		
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.payments (id, amount, student, payed, reason, emitted) FROM stdin;
1	100	1	f		2018-05-06 17:07:46.55
2	100	2	t		2018-05-06 17:07:48.518
3	\N	\N	f	\N	2018-05-06 17:09:32.253582
\.


--
-- Data for Name: students; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.students (id, name, surname, mail, info) FROM stdin;
1	philippe	scorsolini	mail@mail.com	\N
2	lorenzo	petrangeli	mail@mail.it	\N
\.


--
-- Data for Name: subjects; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.subjects (id) FROM stdin;
science
\.


--
-- Data for Name: teachers; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.teachers (id, name, mail, info, surname) FROM stdin;
1	mantesso	mail		mail
\.


--
-- Data for Name: teaches; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.teaches (teacher, subject, class) FROM stdin;
1	science	2
1	science	3
\.


--
-- Data for Name: timetable; Type: TABLE DATA; Schema: back2school; Owner: postgres
--

COPY back2school.timetable (class, subject, location, start, "end", info, id) FROM stdin;
2	science		\N	\N		2
\.


--
-- Name: admins_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.admins_id_seq', 1, false);


--
-- Name: appointments_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.appointments_id_seq', 1, false);


--
-- Name: classes_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.classes_id_seq', 1, false);


--
-- Name: grades_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.grades_id_seq', 1, true);


--
-- Name: notification_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.notification_id_seq', 1, false);


--
-- Name: parents_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.parents_id_seq', 1, false);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.payments_id_seq', 1, false);


--
-- Name: students_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.students_id_seq', 1, false);


--
-- Name: teachers_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.teachers_id_seq', 1, false);


--
-- Name: timetable_id_seq; Type: SEQUENCE SET; Schema: back2school; Owner: postgres
--

SELECT pg_catalog.setval('back2school.timetable_id_seq', 2, true);


--
-- Name: accounts accounts_user_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.accounts
    ADD CONSTRAINT accounts_user_pk PRIMARY KEY ("user");


--
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- Name: appointments appointments_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.appointments
    ADD CONSTRAINT appointments_pkey PRIMARY KEY (id);


--
-- Name: classes classes_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.classes
    ADD CONSTRAINT classes_pkey PRIMARY KEY (id);


--
-- Name: enrolled enrolled_student_class_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.enrolled
    ADD CONSTRAINT enrolled_student_class_pk PRIMARY KEY (student, class);


--
-- Name: grades grades_id_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.grades
    ADD CONSTRAINT grades_id_pk PRIMARY KEY (id);


--
-- Name: isparent isparent_parent_student_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.isparent
    ADD CONSTRAINT isparent_parent_student_pk PRIMARY KEY (parent, student);


--
-- Name: notification notification_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.notification
    ADD CONSTRAINT notification_pkey PRIMARY KEY (id);


--
-- Name: parents parents_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.parents
    ADD CONSTRAINT parents_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: subjects subjects_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.subjects
    ADD CONSTRAINT subjects_pkey PRIMARY KEY (id);


--
-- Name: teachers teachers_pkey; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teachers
    ADD CONSTRAINT teachers_pkey PRIMARY KEY (id);


--
-- Name: teaches teaches_teacher_subject_class_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teaches
    ADD CONSTRAINT teaches_teacher_subject_class_pk PRIMARY KEY (teacher, subject, class);


--
-- Name: timetable timetable_id_pk; Type: CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.timetable
    ADD CONSTRAINT timetable_id_pk PRIMARY KEY (id);


--
-- Name: accounts_password_uindex; Type: INDEX; Schema: back2school; Owner: postgres
--

CREATE UNIQUE INDEX accounts_password_uindex ON back2school.accounts USING btree (password);


--
-- Name: accounts_user_uindex; Type: INDEX; Schema: back2school; Owner: postgres
--

CREATE UNIQUE INDEX accounts_user_uindex ON back2school.accounts USING btree ("user");


--
-- Name: grades_student_subject_date_pk; Type: INDEX; Schema: back2school; Owner: postgres
--

CREATE UNIQUE INDEX grades_student_subject_date_pk ON back2school.grades USING btree (student, subject, date);


--
-- Name: appointments appointments_student_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.appointments
    ADD CONSTRAINT appointments_student_id_fk FOREIGN KEY (student) REFERENCES back2school.students(id);


--
-- Name: appointments appointments_teacher_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.appointments
    ADD CONSTRAINT appointments_teacher_id_fk FOREIGN KEY (teacher) REFERENCES back2school.teachers(id);


--
-- Name: enrolled enrolled_classes_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.enrolled
    ADD CONSTRAINT enrolled_classes_id_fk FOREIGN KEY (class) REFERENCES back2school.classes(id);


--
-- Name: enrolled enrolled_students_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.enrolled
    ADD CONSTRAINT enrolled_students_id_fk FOREIGN KEY (student) REFERENCES back2school.students(id);


--
-- Name: grades grades_students_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.grades
    ADD CONSTRAINT grades_students_id_fk FOREIGN KEY (student) REFERENCES back2school.students(id);


--
-- Name: grades grades_teachers_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.grades
    ADD CONSTRAINT grades_teachers_id_fk FOREIGN KEY (teacher) REFERENCES back2school.teachers(id);


--
-- Name: isparent isparent_parents_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.isparent
    ADD CONSTRAINT isparent_parents_id_fk FOREIGN KEY (parent) REFERENCES back2school.parents(id);


--
-- Name: isparent isparent_students_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.isparent
    ADD CONSTRAINT isparent_students_id_fk FOREIGN KEY (student) REFERENCES back2school.students(id);


--
-- Name: payments payments_students_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.payments
    ADD CONSTRAINT payments_students_id_fk FOREIGN KEY (student) REFERENCES back2school.students(id);


--
-- Name: teaches teaches_classes_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teaches
    ADD CONSTRAINT teaches_classes_id_fk FOREIGN KEY (class) REFERENCES back2school.classes(id);


--
-- Name: teaches teaches_subjects_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teaches
    ADD CONSTRAINT teaches_subjects_id_fk FOREIGN KEY (subject) REFERENCES back2school.subjects(id);


--
-- Name: teaches teaches_teachers_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.teaches
    ADD CONSTRAINT teaches_teachers_id_fk FOREIGN KEY (teacher) REFERENCES back2school.teachers(id);


--
-- Name: timetable timetable_classes_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.timetable
    ADD CONSTRAINT timetable_classes_id_fk FOREIGN KEY (class) REFERENCES back2school.classes(id);


--
-- Name: timetable timetable_subjects_id_fk; Type: FK CONSTRAINT; Schema: back2school; Owner: postgres
--

ALTER TABLE ONLY back2school.timetable
    ADD CONSTRAINT timetable_subjects_id_fk FOREIGN KEY (subject) REFERENCES back2school.subjects(id);


--
-- PostgreSQL database dump complete
--

