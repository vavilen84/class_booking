CREATE TABLE timetable_item
(
    id       varchar(255) NOT NULL,
    class_id varchar(255) NOT NULL,
    date     date         not null unique,
    FOREIGN KEY (class_id) REFERENCES class (id),
    PRIMARY KEY (id)
)
