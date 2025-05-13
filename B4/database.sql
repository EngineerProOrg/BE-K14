create database BEK14;


use BEK14;
create table if not exists Course(
    course_id int not null auto_increment,
    course_name varchar(255),
    PRIMARY KEY (course_id)
);

create table if not exists Professor(
    prof_id int not null auto_increment,
    prof_lname varchar(50),
    prof_fname varchar(50),
    PRIMARY KEY (prof_id)
);

create table if not exists Student(
    stud_id int not null auto_increment,
    stud_lname varchar(50),
    stud_fname varchar(50),
    stud_street varchar(255),
    stud_city varchar(50),
    stud_zip varchar(10),
    PRIMARY KEY (stud_id)
);

create table if not exists Class(
    class_id int not null auto_increment,
    class_name varchar(255),
    prof_id int,
    course_id int,
    room_id int unique,
    PRIMARY KEY (class_id),
    FOREIGN KEY (prof_id) REFERENCES Professor(prof_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

create table if not exists Room(
    room_id int not null auto_increment,
    room_loc varchar(50),
    room_cap varchar(50),
    class_id int unique,
    PRIMARY KEY (room_id),
    FOREIGN KEY (class_id) REFERENCES Class(class_id)
);

create table if not exists Enroll(
    stud_id int not null,
    class_id int not null,
    grade varchar(3),
    CONSTRAINT PK_enroll PRIMARY KEY (stud_id,class_id),
    FOREIGN KEY (stud_id) REFERENCES Student(stud_id),
    FOREIGN KEY (class_id) REFERENCES Class(class_id)
);



-- INSERT--
insert into Course (course_name) values 
    ("JAVA"), 
    ("GO"), 
    ("C#"), 
    ("HIV");
insert into Professor(prof_lname, prof_fname) values 
    ("TOET", "Discord"), 
    ("VIET", "Discord"), 
    ("TOET", "Zalo"), 
    ("VIET", "Zalo");
insert into Room(room_loc, room_cap) values 
    ("Room 1", "100"), 
    ("Room 2", "101"),
    ("Room 3", "102");
insert into Class(class_name, prof_id, course_id, room_id) values
    ("BK14", 1, 2, 1),
    ("BK14", 2, 2, 2),
    ("HIV-AIDS", 3, 4, 3);

update Room set class_id = 1 where room_id = 1;
update Room set class_id = 2 where room_id = 2;
update Room set class_id = 3 where room_id = 3;

insert into Student(stud_lname, stud_fname, stud_street, stud_city, stud_zip) values
    ("An", "Le", "Street 1", "Ha Noi", "Z111"),
    ("Zip", "Po", "Street 2", "Ha Tay", "Z112"),
    ("Trump", "Donal", "Street New York", "Cali", "Z113"),
    ("Tap", "Can Binh", "Street Bac Kinh", "China", "Z114");

insert into Enroll(stud_id, class_id, grade) values
    (1, 1, "C1"),
    (1, 2, "C2"),
    (2, 3, "C3"),
    (1, 3, "C3"),
    (2, 2, "C2"),
    (3, 2, "C2"),
    (4, 2, "C2");
----------

-- 1--những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
select CONCAT(prof_lname, " ", prof_fname) as ProfessorName,
    CONCAT(s.stud_lname, " ", s.stud_fname) as StudentName,
    c.class_name as ClassName
from Professor p
inner join Class c on p.prof_id = c.prof_id
inner join Enroll e on c.class_id = e.class_id
inner join Student s on e.stud_id = s.stud_id;
-----

-- 2--những course (distinct) mà 1 professor cụ thể đang dạy
select distinct(cs.course_name) as CourseName
from Course cs
inner join Class c on cs.course_id = c.course_id;
-----

-- 3--những course (distinct) mà 1 student cụ thể đang học
select distinct(cs.course_name) as CourseName
from Course cs
inner join Class c on cs.course_id = c.course_id
inner join Enroll e on c.class_id = e.class_id;
-----

-- 4---điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
select (
	case 
		when e.grade = 'A' then 10
        when e.grade = 'B' then 8
        when e.grade = 'C' then 6
        when e.grade = 'D' then 4
        when e.grade = 'E' then 2
        else 0
    end
) as Grade from Enroll e;
------

-- 5---điểm số trung bình của 1 học sinh cụ thể 
-- (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
select CONCAT(s.stud_lname, " ", s.stud_fname) as StudentName,
	(case when temp.AverageGrade < 5 then "weak"
          when temp.AverageGrade >= 5 and temp.AverageGrade < 8 then "average"
          else "good"
     end) as Ability
from Enroll e
inner join Student s on e.stud_id = s.stud_id
inner join (
	select e1.stud_id, avg(
		case 
			when e1.grade = 'A' then 10
			when e1.grade = 'B' then 8
			when e1.grade = 'C' then 6
			when e1.grade = 'D' then 4
			when e1.grade = 'E' then 2
			else 0
		end) as AverageGrade
	from Enroll e1
    group by e1.stud_id) temp on e.stud_id = temp.stud_id
group by e.stud_id;
------

-- 6--- điểm số trung bình của các class (quy ra lại theo chữ cái)
select c.class_name, (case when temp.AverageGrade < 5 then "weak"
          when temp.AverageGrade >= 5 and temp.AverageGrade < 8 then "average"
          else "good"
     end) as Grade
from Class c 
inner join Enroll e on c.class_id = e.class_id
inner join (
	select e1.class_id, avg(
		case 
			when e1.grade = 'A' then 10
			when e1.grade = 'B' then 8
			when e1.grade = 'C' then 6
			when e1.grade = 'D' then 4
			when e1.grade = 'E' then 2
			else 0
		end) as AverageGrade
	from Enroll e1
    group by e1.class_id) temp on e.class_id = temp.class_id
group by c.class_id;
-------

-- 7--- điểm số trung bình của các course (quy ra lại theo chữ cái)
select temp.course_name,
	 (case 
		when temp.Grade >= 8 then "good"
        when temp.Grade < 8 and temp.Grade >= 5 then "average"
        else "weak"
     end) as AverageGrade
from (SELECT 
  cs.course_name,
  avg(case 
		when e.grade = 'A' then 10
		when e.grade = 'B' then 8
		when e.grade = 'C' then 6
		when e.grade = 'D' then 4
		when e.grade = 'E' then 2
		else 0
	 end) as Grade
FROM Course cs
inner join Class c on cs.course_id = c.course_id
inner join Enroll e on c.class_id = e.class_id
group by cs.course_id) temp;