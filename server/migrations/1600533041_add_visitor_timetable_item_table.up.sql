CREATE TABLE visitor_timetable_item
(
    id                varchar(255) NOT NULL,
    visitor_id        varchar(255) NOT NULL,
    timetable_item_id varchar(255) NOT NULL,
    FOREIGN KEY (visitor_id) REFERENCES visitor (id),
    FOREIGN KEY (timetable_item_id) REFERENCES timetable_item (id),
    PRIMARY KEY (id),
    UNIQUE (visitor_id, timetable_item_id)
)
