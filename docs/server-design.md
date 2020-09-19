# Server application design

## DB structure

### class

- id varchar(255) PK
- name varchar(255) not null
- capacity int(10) not null

### timetable_item

- id varchar(255) PK
- class_id varchar(255) REFERENCES class not null
- date date not null unique

### visitor

- id varchar(255) PK
- email varchar(255) not null unique

### visitor_to_timetable_item

- visitor_id varchar(255) not null
- timetable_item_id varchar(255) not null
PK (visitor_id, timetable_item_id)


## Booking flow

- Register class
- Create class timetable
- Register user
- User is able to book a class for a given date
